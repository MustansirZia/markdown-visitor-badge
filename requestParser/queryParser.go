package requestParser

import (
	"errors"
	"net/http"

	"golang.org/x/exp/slices"
)

// RequestParams - Parameters contaioned in the request.
type RequestParams struct {
	// Label - Label for the visitors.
	Label string
	// LabelBgColor - Background color for the visitors label area.
	LabelBgColor string
	// LabelColor - Color for the visitors label.
	LabelColor string
	// CountBgColor - Background color for the count area.
	CountBgColor string
	// CountColor - Color for the count.
	CountColor string
	// Key - Key for the badge.
	Key string
}

// RequestParser - Interface for the request parser.
type RequestParser interface {
	// Parse - Parses the given HTTP request and returns the badge parameters.
	Parse(r *http.Request) (RequestParams, error)
}

// NewParser - Constructs and returns a new RequestParser.
func NewParser() RequestParser {
	return queryParser{}
}

var validKeys = []string{"github_visitors"}

// ErrInvalidKey - If an invalid value is passed for the key parameter.
var ErrInvalidKey = errors.New("this key is not supported")

type queryParser struct{}

func (p queryParser) Parse(r *http.Request) (RequestParams, error) {
	queryParams := r.URL.Query()
	var label string
	var labelColor string
	var labelBgColor string
	var countColor string
	var countBgColor string
	var key string
	if queryParams.Has("label") {
		label = queryParams.Get("label")
	} else {
		label = "VISITORS"
	}
	if queryParams.Has("labelColor") {
		labelColor = queryParams.Get("labelColor")
	} else {
		labelColor = "#fff"
	}
	if queryParams.Has("labelBgColor") {
		labelBgColor = queryParams.Get("labelBgColor")
	} else {
		labelBgColor = "#555"
	}
	if queryParams.Has("countColor") {
		countColor = queryParams.Get("countColor")
	} else {
		countColor = "#000"
	}
	if queryParams.Has("countBgColor") {
		countBgColor = queryParams.Get("countBgColor")
	} else {
		countBgColor = "#f47373"
	}
	if queryParams.Has("key") {
		key = queryParams.Get("key")
	} else {
		key = "github_visitors"
	}
	if !slices.Contains(validKeys, key) {
		return RequestParams{}, ErrInvalidKey
	}
	requestParams := RequestParams{Label: label, LabelColor: labelColor, LabelBgColor: labelBgColor, CountColor: countColor, CountBgColor: countBgColor, Key: key}
	return requestParams, nil
}
