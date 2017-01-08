package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/k0kubun/pp"
	"github.com/mholt/binding"
)

// We expect a multipart/form-data upload with
// a file field named 'data'
type MultipartForm struct {
	Data *multipart.FileHeader `json:"data"`
}

func (f *MultipartForm) FieldMap(req *http.Request) binding.FieldMap {
	return binding.FieldMap{
		&f.Data: "data",
	}
}

// Handlers are still clean and simple
func handler(resp http.ResponseWriter, req *http.Request) {
	multipartForm := new(MultipartForm)
	errs := binding.Bind(req, multipartForm)
	if errs.Handle(resp) {
		return
	}

	pp.Println(req)

	// To access the file data you need to Open the file
	// handler and read the bytes out.
	var fh io.ReadCloser
	var err error
	if fh, err = multipartForm.Data.Open(); err != nil {
		http.Error(resp,
			fmt.Sprint("Error opening Mime::Data %+v", err),
			http.StatusInternalServerError)
		return
	}
	defer fh.Close()
	dataBytes := bytes.Buffer{}
	var size int64
	if size, err = dataBytes.ReadFrom(fh); err != nil {
		http.Error(resp,
			fmt.Sprint("Error reading Mime::Data %+v", err),
			http.StatusInternalServerError)
		return
	}
	// Now you have the attachment in databytes.
	// Maximum size is default is 10MB.
	log.Printf("Read %v bytes with filename %s",
		size, multipartForm.Data.Filename)
}

func main() {
	http.HandleFunc("/upload", handler)
	pp.Println("listening 3000")
	http.ListenAndServe(":3000", nil)
}
