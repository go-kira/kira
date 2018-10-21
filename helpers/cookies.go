package helpers

import "net/http"

// GetCookie - get the cookies from the request.
func GetCookie(name string, request *http.Request) (string, error) {
	cookie, err := request.Cookie(name)
	if err != nil {
		return "", err
	}

	return cookie.Value, nil
}

// SetCookie - set cookies to the response.
func SetCookie(cookie http.Cookie, response http.ResponseWriter) {
	http.SetCookie(response, &cookie)
}
