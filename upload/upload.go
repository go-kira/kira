package upload

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Lafriakh/kira"
	"github.com/Lafriakh/kira/helpers"
	"github.com/go-kira/kog"
	"github.com/go-kira/kon"
)

var (
	errNotImage           = errors.New("the file not an image")
	errWhileWriteFile     = errors.New("error while write a file")
	errImageToNotFound    = errors.New("set an image format to convert to it")
	errFileTypeNotAllowed = "the file type: %s not allowed"
	errFileSize           = "the file %s size: %s too big"
	errFileRequired       = "the field %s is required"
)

// MB - one MB.
const (
	KB = 1 << 10
	MB = 1 << 20
	GB = 1 << 30
)

var isImage = []string{
	Mimes["jpeg"],
	Mimes["png"],
	Mimes["gif"],
	Mimes["bmp"],
	Mimes["svg"],
}

// Upload - upload struct.
type Upload struct {
	Response         http.ResponseWriter
	Request          *http.Request
	name, path, form string
	isImage          bool
	imageTo          string
	ext              []string
	size             int64
	required         bool
}

// Uploaded - is the finale file with returned to the user.
type Uploaded struct {
	Name string
	Mime string
	Size int64
	Path string
}

// New - return upload instance.
func New(config *kon.Kon, w http.ResponseWriter, r *http.Request) *Upload {

	// return upload instance.
	return &Upload{
		Response: w,
		Request:  r,
		size:     config.GetInt64("UPLOAD_MAX_SIZE") * MB, // this the default max upload size from config.
		required: true,
	}
}

// Form - set the form.
func (u *Upload) Form(form string) *Upload {
	u.form = form

	return u
}

// Name - set the name.
func (u *Upload) Name(name string) *Upload {
	u.name = name

	return u
}

// Path - set the path.
func (u *Upload) Path(path string) *Upload {
	u.path = path

	return u
}

// IsImage - set the upload file as an image.
func (u *Upload) IsImage() *Upload {
	u.isImage = true

	return u
}

// Ext - check if the file that we upload is one of this types.
func (u *Upload) Ext(types []string) *Upload {
	u.ext = types

	return u
}

// Size - to validate the file size.
func (u *Upload) Size(size int64) *Upload {
	u.size = size

	return u
}

// ImageTo - when this set convert the image to this type.
func (u *Upload) ImageTo(itype string) *Upload {
	u.imageTo = itype

	return u
}

// NotRequired - make the form not required.
func (u *Upload) NotRequired() *Upload {
	u.required = false

	return u
}

// Upload - to upload one file.
// dst: ./storage/{dst}
func (u *Upload) Upload() (Uploaded, error) {
	var uploaded os.FileInfo
	var path string

	// file
	file, header, err := u.Request.FormFile(u.form)
	if err != nil {
		// if the form not required, return nil error.
		if !u.required && err == http.ErrMissingFile {
			return Uploaded{}, nil
		}
		return Uploaded{}, err
	}
	defer file.Close()

	// if the file not required return nil error if file nil and size is 0.
	if !u.required && file == nil || !u.required && header.Size == 0 {
		return Uploaded{}, nil
	}

	// if the file required and the size 0.
	if u.required && header.Size == 0 {
		return Uploaded{}, fmt.Errorf(errFileRequired, u.form)
	}

	// file header
	fileHeaderMime, err := u.getFileType(file)
	if err != nil {
		return Uploaded{}, err
	}

	// check if the file type equal to one of the types the user sets.
	if len(u.ext) > 0 {
		var isOK bool
		for _, ext := range u.ext {
			if fileHeaderMime == Mimes[ext] {
				isOK = true
				break
			}
		}
		if !isOK {
			// file mime.
			typeWithoutDot := filepath.Ext(header.Filename)[1:len(filepath.Ext(header.Filename))]
			// return error if the type of the file not allowed.
			return Uploaded{}, fmt.Errorf(errFileTypeNotAllowed, typeWithoutDot)
		}
	}

	// check the file size
	if header.Size > u.size {
		return Uploaded{}, fmt.Errorf(errFileSize, header.Filename, helpers.BytesFormat(float64(header.Size), 1))
	}

	// upload file
	if !u.isImage {
		// do upload
		path = u.getFilePath(header.Filename)
		uploaded, err = writeFile(file, path)
		if err != nil {
			return Uploaded{}, err
		}
	} else {
		// upload the image
		// check if the uploaded file is an image.
		if !helpers.Contains(isImage, fileHeaderMime) {
			return Uploaded{}, errNotImage
		}
		if u.imageTo != "" {
			// upload image and convert.
			path = filepath.Join(u.path, u.name+"."+u.imageTo)
			uploaded, err = writeImage(file, path, u.imageTo)
			if err != nil {
				return Uploaded{}, err
			}
		} else {
			// do upload
			path = u.getFilePath(header.Filename)
			uploaded, err = writeFile(file, path)
			if err != nil {
				return Uploaded{}, err
			}
		}
	}

	return Uploaded{
		Name: uploaded.Name(),
		Mime: fileHeaderMime,
		Size: uploaded.Size(),
		Path: path,
	}, nil
}

func (u *Upload) getFileType(file multipart.File) (string, error) {
	// Create a buffer to store the header of the file in
	fileHeaderBuffer := make([]byte, 512)
	// Copy the headers into the FileHeader buffer
	if _, err := file.Read(fileHeaderBuffer); err != nil {
		return "", err
	}
	// set position back to start.
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}

	return http.DetectContentType(fileHeaderBuffer), nil
}

func (u *Upload) getFilePath(fname string) string {
	return filepath.Join(u.path, u.name+filepath.Ext(fname))
}

func writeFile(file multipart.File, dst string) (os.FileInfo, error) {
	location := kira.PathApp + dst

	f, err := os.OpenFile(location, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = io.Copy(f, file)
	if err != nil {
		return nil, errWhileWriteFile
	}

	// return file location.
	fileInfo, err := f.Stat()
	if err != nil {
		kog.Panic(err)
	}

	return fileInfo, nil
}

func writeImage(file multipart.File, dst string, utype string) (os.FileInfo, error) {
	location := kira.PathApp + dst

	f, err := os.OpenFile(location, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	if utype == "png" {
		err = convertToPNG(f, file)
		if err != nil {
			return nil, err
		}
	} else if utype == "jpeg" {
		err = convertToJPEG(f, file)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errImageToNotFound
	}

	// return file location.
	fileInfo, err := f.Stat()
	if err != nil {
		kog.Panic(err)
	}

	return fileInfo, nil
}

// ParseForm - parse the request form.
func ParseForm(r *http.Request) {
	// parse the request before upload.
	err := r.ParseMultipartForm(32 << 20)
	if err != nil {
		kog.Panic(err)
	}

}
