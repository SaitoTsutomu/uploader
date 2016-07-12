package main

import (
	. "fmt"
	"io/ioutil"
	. "net/http"
)

func home(w ResponseWriter, r *Request) {
	Fprintf(w,
		`<html><head><title>Uploader</title></head>
		 <body><form action="/upload" method="POST" enctype="multipart/form-data">
		 <input type="file" id="file" name="file">
		 <input type="submit" name="upload" value="upload">
		 </form></body></html>`)
}
func upload(w ResponseWriter, r *Request) {
	if r.Method == "GET" {
		Redirect(w, r, "/", 301)
	}
	_, h, _ := r.FormFile("file")
	pth := Sprintf("files/%s", h.Filename)
	f, _ := h.Open()
	buf, _ := ioutil.ReadAll(f)
	ioutil.WriteFile(pth, buf, 0644)
	Redirect(w, r, "/"+pth, 301)
}
func main() {
	HandleFunc("/", home)
	HandleFunc("/upload", upload)
	svr := StripPrefix("/files/", FileServer(Dir("files/")))
	Handle("/files/", svr)
	Println("Open http://localhost:8888/")
	panic(ListenAndServe(":8888", nil))
}
