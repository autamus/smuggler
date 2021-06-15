package web

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/autamus/smuggler/config"
)

// Upload handles POST requests from http clients and saves the uploaded files
// into the save path.
// Report Syntax: http://server/upload
// Report Example: curl -F "upload=@FILENAME" http://server/upload
func Upload(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	uploadFile, handler, err := r.FormFile("upload")
	if err != nil {
		log.Println("Error Retrieving the File")
		log.Println(err)
		return
	}
	defer uploadFile.Close()

	err = os.MkdirAll(config.Global.Save.Path, os.ModePerm)
	if err != nil {
		reportError(w, r, "web.Upload", err)
		return
	}
	file, err := os.Create(filepath.Join(config.Global.Save.Path, handler.Filename))
	if err != nil {
		reportError(w, r, "web.Upload", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, uploadFile)
	if err != nil {
		reportError(w, r, "web.Upload", err)
		return
	}
	fmt.Println("Successfully Smuggled File :P")
}
