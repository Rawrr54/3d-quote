package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/quote", quote)
	http.HandleFunc("/reports", reports)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func quote(w http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()
		ext := strings.Split(fh.Filename, ".")[1]
		fname := strings.Split(fh.Filename, ".")[0] + "." + ext
		//create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "3dfiles", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		mf.Seek(0, 0)
		io.Copy(nf, mf)
	}
	tpl.ExecuteTemplate(w, "quote.gohtml", nil)
}

func index(w http.ResponseWriter, req *http.Request) {

	if req.Method == http.MethodPost {
		mf, fh, err := req.FormFile("nf")
		if err != nil {
			fmt.Println(err)
		}
		defer mf.Close()
		ext := strings.Split(fh.Filename, ".")[1]
		fname := strings.Split(fh.Filename, ".")[0] + "." + ext
		//create new file
		wd, err := os.Getwd()
		if err != nil {
			fmt.Println(err)
		}
		path := filepath.Join(wd, "public", "3dfiles", fname)
		nf, err := os.Create(path)
		if err != nil {
			fmt.Println(err)
		}
		defer nf.Close()

		mf.Seek(0, 0)
		io.Copy(nf, mf)
	}
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}

func reports(w http.ResponseWriter, req *http.Request) {

}
