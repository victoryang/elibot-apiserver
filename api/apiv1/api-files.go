package apiv1

import (
	"os"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)

var RootPath string

func SetRootPath(path string) {
	RootPath = path
}

func handleUploadFile(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(32 << 20)
	src, handler, err := r.FormFile("fileupload")
	if err!=nil {
		Log.Error("Fail to get reader: ", err)
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}
	defer src.Close()

	des, err := os.OpenFile(RootPath + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
    	Log.Error("Fail to open file: ", err)
        WriteInternalServerErrorResponse(w, ERRREQUESTFAIL)
		return
    }
    defer des.Close()

    if encoding := r.Header.Get("Content-Encoding"); encoding!="" {
    	Log.Print("content-Encoding set to gzip")
    } else if {
    	if _, err := io.Copy(des, src); err!=nil {
	    	Log.Error("Fail to copy data into file: ", err)
	    	WriteInternalServerErrorResponse(w, ERRREQUESTFAIL)
			return
	    }
    }

    WriteSuccessResponse(w, "succeed in uploading file " + handler.Filename)
}

func getJBIList(w http.ResponseWriter, r *http.Request) {
	files, err := ioutil.ReadDir(RootPath)
    if err != nil {
        WriteInternalServerErrorResponse(w, ERRREQUESTFAIL)
		return
    }

    var list []string
    for _, f := range files {
    	if strings.Contains(f.Name(), ".jbi") {
    		list = append(list, f.Name())
    	}
    }

    WriteSuccessResponse(w, list)
}