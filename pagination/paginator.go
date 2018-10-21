package pagination

import (
	"bytes"
	"html/template"
	"log"
	"math"

	"github.com/go-kira/kon"

	"github.com/Lafriakh/kira/helpers"
)

const defaultTemplate = "default"

// Pagination ...
type Pagination struct {
	total, perPage, currentPage int
	url                         string

	numPages       int
	maxPagesToShow int
	lastPage       float64
	template       string

	configs *kon.Kon
}

// Page ...
type Page struct {
	Number    int
	URL       string
	IsCurrent bool
}

func New(configs *kon.Kon, total, perPage, currentPage int, url string) *Pagination {
	pagination := &Pagination{
		total:          total,
		perPage:        perPage,
		currentPage:    currentPage,
		url:            url,
		maxPagesToShow: 5,
		lastPage:       math.Max(float64(int(math.Ceil(float64(total)/float64(perPage)))), float64(1)),
		template:       defaultTemplate,
		configs:        configs,
	}

	pagination.updateNumPages()

	return pagination
}

func (p *Pagination) updateNumPages() {
	p.numPages = int(math.Ceil(float64(p.total) / float64(p.perPage)))
}

// SetTemplate - set pagination template.
func (p *Pagination) SetTemplate(name string) {
	p.template = name
}

// Pages - calculate pagination pages.
func (p *Pagination) Pages() []Page {
	var pages []Page
	var slidingStart int
	var slidingEnd int

	if p.numPages <= 1 {
		return make([]Page, 0)
	}

	if p.numPages <= p.maxPagesToShow {
		for i := 1; i <= p.numPages; i++ {
			pages = append(pages, p.createPage(i, i == p.currentPage))
		}
	} else {
		numAdjacents := int(math.Floor(float64((p.maxPagesToShow - 3) / 2)))
		if (p.currentPage + numAdjacents) > p.numPages {
			// $slidingStart = $this->numPages - $this->maxPagesToShow + 2;
			slidingStart = p.numPages - p.maxPagesToShow + 2
		} else {
			slidingStart = p.currentPage - numAdjacents
		}
		if slidingStart < 2 {
			slidingStart = 2
		}

		slidingEnd = slidingStart + p.maxPagesToShow - 3
		if slidingEnd >= p.numPages {
			slidingEnd = p.numPages - 1
		}

		// Build the list of pages.
		pages = append(pages, p.createPage(1, p.currentPage == 1))
		if slidingStart > 2 {
			pages = append(pages, p.createPageEllipsis())
		}

		for i := slidingStart; i <= slidingEnd; i++ {
			pages = append(pages, p.createPage(i, i == p.currentPage))
		}

		if slidingEnd < p.numPages-1 {
			pages = append(pages, p.createPageEllipsis())
		}

		pages = append(pages, p.createPage(p.numPages, p.currentPage == p.numPages))
	}

	return pages
}

func (p *Pagination) createPage(number int, isCurrent bool) Page {
	return Page{
		Number:    number,
		URL:       p.url + helpers.ConvertToString(number),
		IsCurrent: isCurrent,
	}
}

func (p *Pagination) createPageEllipsis() Page {
	return Page{
		Number:    0,
		URL:       "",
		IsCurrent: false,
	}
}

// OnFirstPage - check if we in first page or not.
func (p *Pagination) OnFirstPage() bool {
	return p.currentPage <= 1
}

// HasMorePages - check if we have more pages.
func (p *Pagination) HasMorePages() bool {
	return p.currentPage < int(p.lastPage)
}

// PrevPage - return prev page number.
func (p *Pagination) PrevPage() int {
	if p.currentPage > 1 {
		return p.currentPage - 1
	}
	return 0
}

// PrevURL - return prev page url.
func (p *Pagination) PrevURL() string {
	if p.PrevPage() == 0 {
		return ""
	}

	return p.url + helpers.ConvertToString(p.PrevPage())
}

// NextPage - return nrxt page number.
func (p *Pagination) NextPage() int {
	if p.currentPage < p.numPages {
		return p.currentPage + 1
	}
	return 0
}

// NextURL - return next page url.
func (p *Pagination) NextURL() string {
	if p.NextPage() == 0 {
		return ""
	}

	return p.url + helpers.ConvertToString(p.NextPage())
}

func getTemplate(template string) string {
	if template == "" {
		return defaultTemplate
	}

	return template
}

// Render - render the pagination.
func (p *Pagination) Render() template.HTML {
	var out bytes.Buffer
	tmpl := template.Must(template.ParseFiles("./storage/framework/views/pagination/" + getTemplate(p.configs.GetString("PAGINATION_TEMPLATE")) + ".go.html"))

	err := tmpl.Execute(&out, p)
	if err != nil {
		log.Fatalf("template execution: %s", err)
	}

	return template.HTML(out.String())

}
