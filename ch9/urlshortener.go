// @file: urlshortener.go
// @description:
// @author: SaltFish
// @date: 2020/08/10

// Package ch9 is chapter 9
package ch9

import (
	"fmt"
	"net/http"
	"text/template"

	"google.golang.org/api/urlshortener/v1"
)

// MyURLShortener is fun
func MyURLShortener() {
	http.HandleFunc("/", root)
	http.HandleFunc("/short", short)
	http.HandleFunc("/long", long)

	http.ListenAndServe("localhost:8080", nil)
}

// the template used to show the forms and the results web page to the user
var rootHTMLTmpl = template.Must(
	template.New("rootHtml").Parse(
		`
<html><body>
<h1>URL SHORTENER</h1>
{{if .}}{{.}}<br /><br />{{end}}
<form action="/short" type="POST">
Shorten this: <input type="text" name="longUrl" />
<input type="submit" value="Give me the short URL" />
</form>
<br />
<form action="/long" type="POST">
Expand this: http://goo.gl/<input type="text" name="shortUrl" />
<input type="submit" value="Give me the long URL" />
</form>
</body></html>
`,
	),
)

func root(w http.ResponseWriter, r *http.Request) {
	rootHTMLTmpl.Execute(w, nil)
}
func short(w http.ResponseWriter, r *http.Request) {
	longURL := r.FormValue("longUrl")
	urlshortenerSvc, _ := urlshortener.New(http.DefaultClient)
	url, _ := urlshortenerSvc.Url.Insert(&urlshortener.Url{LongUrl: longURL}).Do()
	rootHTMLTmpl.Execute(
		w, fmt.Sprintf(
			"Shortened version of %s is : %s",
			longURL, url.Id,
		),
	)
}

func long(w http.ResponseWriter, r *http.Request) {
	shortURL := "http://goo.gl/" + r.FormValue("shortUrl")
	urlshortenerSvc, _ := urlshortener.New(http.DefaultClient)
	url, err := urlshortenerSvc.Url.Get(shortURL).Do()
	if err != nil {
		fmt.Println("error: ", err)
		return

	}
	rootHTMLTmpl.Execute(
		w, fmt.Sprintf(
			"Longer version of %s is : %s",
			shortURL, url.LongUrl,
		),
	)
}
