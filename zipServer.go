package main

import (
	"./vfs/"
	"./vfs/zipfs"
	"archive/zip"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

var fs vfs.FileSystem

func handler(w http.ResponseWriter, r *http.Request) {
	fi, err := fs.Open("/" + r.URL.Path[1:])
	if err != nil {
		http.NotFound(w, r)
		return
	}
	defer fi.Close()
	byt, err := ioutil.ReadAll(fi)
	if err != nil {
		fmt.Println("site 2: ", err)
	}
	fmt.Fprintf(w, "%s", string(byt))
}

func main() {
	zipfilepath := os.Args[1]
	rc, _ := zip.OpenReader(zipfilepath)
	defer rc.Close()
	fs = zipfs.New(rc, zipfilepath)
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
