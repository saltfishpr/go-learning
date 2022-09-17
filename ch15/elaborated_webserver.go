// @file: elaborated_webserver.go
// @description: 精巧的多功能网页服务器
// @author: SaltFish
// @date: 2020/09/09

// Package ch15 is chapter 15
package ch15

import (
	"bytes"
	"expvar"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

var helloRequests = expvar.NewInt("hello-requests")
var webroot = flag.String("root", "/home/user", "web root directory")
var booleanflag = flag.Bool("boolean", true, "another flag for testing")

type Counter struct {
	n int
}

type intChan chan int

func ElaboWeb() {
	flag.Parse()
	http.Handle("/", http.HandlerFunc(logger))
	http.Handle("/go/hello", http.HandlerFunc(helloServer))
	ctr := new(Counter)
	expvar.Publish("counter", ctr)
	http.Handle("/counter", ctr)
	http.Handle("/go/", http.StripPrefix("/go/", http.FileServer(http.Dir(*webroot))))
	http.Handle("/flags", http.HandlerFunc(flagServer))
	http.Handle("/args", http.HandlerFunc(argServer))
	http.Handle("/chan", chanCreate())
	http.Handle("/date", http.HandlerFunc(dateServer))
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panicln("ListenAndServe: ", err)
	}
}

func logger(w http.ResponseWriter, r *http.Request) {
	log.Print(r.URL.String())
	w.WriteHeader(http.StatusNotFound)
	w.Write([]byte("oops"))
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	helloRequests.Add(1)
	io.WriteString(w, "hello, world\n")
}

func (ctr *Counter) String() string {
	return fmt.Sprintf("%d", ctr.n)
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		ctr.n++
	case "POST":
		buf := new(bytes.Buffer)
		io.Copy(buf, r.Body)
		body := buf.String()
		if n, err := strconv.Atoi(body); err != nil {
			fmt.Fprintf(w, "bad POST: %v\nbody: [%v]\n", err, body)
		} else {
			ctr.n = n
			fmt.Fprint(w, "counter reset\n")
		}
	}
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

func flagServer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	fmt.Fprint(w, "Flags:\n")
	flag.VisitAll(
		func(f *flag.Flag) {
			if f.Value.String() != f.DefValue {
				fmt.Fprintf(w, "%s = %s [default = %s]\n", f.Name, f.Value.String(), f.DefValue)
			} else {
				fmt.Fprintf(w, "%s = %s\n", f.Name, f.Value.String())
			}
		},
	)
}

func argServer(w http.ResponseWriter, r *http.Request) {
	for _, s := range os.Args {
		fmt.Fprint(w, s, "")
	}
}

func chanCreate() intChan {
	c := make(intChan)
	go func(c intChan) {
		for x := 0; ; x++ {
			c <- x
		}
	}(c)
	return c
}

func (ch intChan) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, fmt.Sprintf("channel send #%d\n", <-ch))
}

func dateServer(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
	r, w, err := os.Pipe()
	if err != nil {
		fmt.Fprintf(rw, "pipe: %s\n", err)
		return
	}
	p, err := os.StartProcess(
		"/bin/date",
		[]string{"date"},
		&os.ProcAttr{Files: []*os.File{nil, w, w}},
	)
	defer r.Close()
	w.Close()
	if err != nil {
		fmt.Fprintf(rw, "fork/exec: %s\n", err)
		return
	}
	defer p.Release()
	io.Copy(rw, r)
	wait, err := p.Wait()
	if err != nil {
		fmt.Fprintf(rw, "wait: %s\n", err)
		return
	}
	if !wait.Exited() {
		fmt.Fprintf(rw, "date: %v\n", wait)
		return
	}
}
