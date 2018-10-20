package blog

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"path"
	"strconv"
	"strings"
)

func Auth(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if AUTH != request.Header.Get("Authorization") {
			writer.WriteHeader(http.StatusForbidden)
			fmt.Fprintf(writer, "Forbidden")
			return
		}
		f(writer, request)
	}
}

func Index(writer http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(writer, "%v", StorageMap.IndexHtml())
}

func Upload(writer http.ResponseWriter, request *http.Request) {
	var buf bytes.Buffer
	if request.Method != "POST" {
		writer.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(writer, "Not Found")
	}
	file, header, err := request.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	rawFilename := SafeFilename(header.Filename)
	filenameArr := strings.Split(rawFilename, ".")
	filename := filenameArr[0]
	var ext string
	if filenameArr[1] == "md" {
		ext = filenameArr[1]
	} else {
		ext = "wi" // with images
	}
	io.Copy(&buf, file)
	md := MarkdownStorage{
		Id:    0,
		Title: filename,
		Ext:   ext,
	}
	nextId := StorageMap.Append(md)
	md.Id = nextId
	if ext == "md" {
		SaveUploadFile(buf.Bytes(), md.FilePath())
	} else {
		rawFilePath := path.Join(COMPRESS_FILE_PATH, rawFilename)
		SaveUploadFile(buf.Bytes(), rawFilePath)
		UnCompressFile(rawFilePath, md.FilePath())
	}
	fmt.Fprintf(writer, "Success")
}

func ReadPost(writer http.ResponseWriter, request *http.Request) {
	url := request.URL.Path
	urlList := strings.Split(url, "/")
	if len(urlList) == 3 {
		id, err := strconv.Atoi(urlList[2])
		if err == nil {
			html, ok := StorageMap[id]
			if ok {
				fmt.Fprintf(writer, "%s", html.ReadHtml())
				return
			}
		}
	}
	writer.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(writer, "Not Found")
}
