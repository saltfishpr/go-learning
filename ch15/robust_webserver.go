// @file: robust_webserver.go
// @description: 闭包错误处理增强网页健壮性
// @author: SaltFish
// @date: 2020/09/09

// Package ch15 is chapter 15
package ch15

import (
	"io"
	"log"
	"net/http"
)

const form1 = `<html><body><form action="#" method="post" name="bar">
		<input type="text" name="in"/>
		<input type="submit" value="Submit"/>
	</form></html></body>`

type HandleFnc func(http.ResponseWriter, *http.Request)

/* handle a simple get request */
func SimpleServer1(w http.ResponseWriter, request *http.Request) {
	io.WriteString(w, "<h1>hello, world</h1>")
}

/* handle a form, both the GET which displays the form
   and the POST which processes it.*/
func FormServer(w http.ResponseWriter, request *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	switch request.Method {
	case "GET":
		/* display the form to the user */
		io.WriteString(w, form1)
	case "POST":
		/* handle the form data, note that ParseForm must
		   be called before we can extract form data*/
		// request.ParseForm();
		// io.WriteString(w, request.Form["in"][0])
		io.WriteString(w, request.FormValue("in"))
	}
}

func RobustWeb() {
	http.HandleFunc("/test1", logPanics(SimpleServer1))
	http.HandleFunc("/test2", logPanics(FormServer))
	if err := http.ListenAndServe(":8088", nil); err != nil {
		panic(err)
	}
}

func logPanics(function HandleFnc) HandleFnc {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if x := recover(); x != nil {
				log.Printf("[%v] caught panic: %v", request.RemoteAddr, x)
			}
		}()
		function(writer, request)
	}
}
