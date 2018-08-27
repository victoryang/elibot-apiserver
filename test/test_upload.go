package main

import (
        "net/http"
        "fmt"
        "bytes"
        "os"
        "mime/multipart"
        "io"
        "path/filepath"
)

func newfileUploadRequest(uri string, params map[string]string, paramName, path string) (*http.Request, error) {
        file, err := os.Open(path)
        if err != nil {
        return nil, err
        }
        defer file.Close()

        body := &bytes.Buffer{}
        writer := multipart.NewWriter(body)
        part, err := writer.CreateFormFile(paramName, filepath.Base(path))
        if err != nil {
              return nil, err
        }
        _, err = io.Copy(part, file)

        for key, val := range params {
              _ = writer.WriteField(key, val)
        }
        err = writer.Close()
        if err != nil {
              return nil, err
        }

        req, err := http.NewRequest("POST", uri, body)
        req.Header.Set("Content-Type", writer.FormDataContentType())
        return req, err
}

func main() {

        extraParams := map[string]string{
        "title":       "My Document",
        "author":      "Matt Aimonetti",
        "description": "A document with all the Go programming language secrets",
        }
        request, err := newfileUploadRequest("http://192.168.1.253:9000/v1/files/upload", extraParams, "fileupload", "/tmp/testfile")
        if err != nil {
                fmt.Println(err)
        }
        client := &http.Client{}
        resp, err := client.Do(request)
        if err != nil {
                fmt.Println(err)
        } else {
                body := &bytes.Buffer{}
                _, err := body.ReadFrom(resp.Body)
                if err != nil {
                        fmt.Println(err)
                }
                resp.Body.Close()
                fmt.Println(resp.StatusCode)
                fmt.Println(resp.Header)
                fmt.Println(body)
        }
}