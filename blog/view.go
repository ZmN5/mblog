package blog

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func Auth(f http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if AUTH != request.Header.Get("Authenticate") {
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
	url := request.URL.Path
	fmt.Fprintf(writer, "%s, upload", url)
}

func ReadPost(writer http.ResponseWriter, request *http.Request) {
	url := request.URL.Path
	urlList := strings.Split(url, "/")
	if len(urlList) == 3 {
		id, err := strconv.Atoi(urlList[2])
		if err == nil {
			html, ok := StorageMap[id]
			if ok {
				fmt.Fprintf(writer, "<html>%s</html>", html.ReadHtml())
				return
			}
		}
	}
	writer.WriteHeader(http.StatusNotFound)
	fmt.Fprintf(writer, "Not Found")
}
