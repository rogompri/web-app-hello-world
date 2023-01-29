package main

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"regexp"
	"testing"

	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
)

func TestTitle(t *testing.T) {
	// Hacer una petición HTTP para obtener el HTML
	res, err := http.Get("http://localhost:8080")
	if err != nil {
		t.Fatal(err)
	}
	defer res.Body.Close()

	// Leer el contenido del HTML
	html, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	// Utilizar goquery para analizar el HTML
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(html))
	if err != nil {
		t.Fatal(err)
	}

	// Extraer el título del body
	title := doc.Find("h1").Text()

	// Comprobar que el título del body cumple la expresión regular
	match, _ := regexp.MatchString("(?i)^Hello World from.*", title)
	assert.True(t, match)
}
