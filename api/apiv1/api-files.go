package apiv1

import (
	"os"
	"io"
	"net/http"
)

var StorePath string = "/rbctrl/"

func handleUploadFile(w http.ResponseWriter, r *http.Request) {
	src, handler, err := r.FormFile("fileupload")
	if err!=nil {
		WriteInternalServerErrorResponse(w, ERRINVALIDBODY)
		return
	}
	defer src.Close()

	des, err := os.OpenFile(StorePath + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        WriteInternalServerErrorResponse(w, ERRREQUESTFAIL)
		return
    }
    defer des.Close()

    if _, err := io.Copy(des, src); err!=nil {
    	WriteInternalServerErrorResponse(w, ERRREQUESTFAIL)
		return
    }

    WriteSuccessResponse(w, "succeed in uploading file " + handler.Filename)
}