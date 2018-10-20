package blog

import (
	"archive/zip"
	"fmt"
	"io"
	"os"
	"path"
	"strconv"
	"strings"
)

func PathExists(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err

}

func Mkdir(filepath string) error {
	exists, err := PathExists(filepath)
	if err == nil {
		if exists {
			return nil
		} else {
			err := os.MkdirAll(filepath, 0777)
			return err
		}
	}
	return err
}

func SafeFilename(filename string) string {
	trimName := strings.Trim(filename, ".")
	return strings.Replace(trimName, "/", "_", -1)
}

func TouchFile(filename string) (*os.File, error) {
	dir := path.Dir(filename)
	dirErr := Mkdir(dir)
	if dirErr == nil {
		fd, fileErr := os.Create(filename)
		if fileErr == nil {
			return fd, fileErr
		} else {
			return nil, fileErr
		}
	}
	return nil, dirErr
}

func SaveUploadFile(data []byte, filePath string) error {
	fd, err := TouchFile(filePath)
	defer fd.Close()
	if err == nil {
		_, err := fd.Write(data)
		if err == nil {
			return nil
		} else {
			return err
		}
	}
	return err
}

func UnCompressFile(zipFilePath, unzipFilePath string) error {
	r, err := zip.OpenReader(zipFilePath)
	if err != nil {
		return err
	}
	for _, k := range r.Reader.File {
		parsedName := path.Join(strings.Split(k.Name, "/")[1:]...)
		filename := path.Join(unzipFilePath, parsedName)
		if k.FileInfo().IsDir() {
			Mkdir(filename)
			continue
		}
		r, err := k.Open()
		if err != nil {
			continue
		}
		defer r.Close()
		NewFile, err := TouchFile(filename)
		if err != nil {
			fmt.Println(err)
			continue
		}
		io.Copy(NewFile, r)
		NewFile.Close()
	}
	return nil

}

func ParseMarkdownName(filename string) (int, string, string, error) {
	ss := strings.Split(filename, "--")
	Id, err := strconv.Atoi(ss[0])
	if err != nil {
		return -1, "", "", err
	}
	var ext string
	var title string
	if strings.HasSuffix(filename, ".md") {
		ext = "md"
		title = strings.TrimSuffix(ss[1], ".md")
	} else if strings.HasSuffix(filename, ".wi") {
		ext = "wi"
		title = strings.TrimSuffix(ss[1], ".wi")
	}

	return Id, title, ext, nil
}

func parseHtml(title string, body string) string {
	metaHtml := "<meta charset=\"UTF-8\">"
	titleHtml := fmt.Sprintf("<title>%s</title>", title)
	styleHtml := fmt.Sprintf("<style>#wrapper{width: 960px;margin: 0 auto;border:0px solid;}</style>")

	headHtml := fmt.Sprintf("<head>%s%s%s</head>", metaHtml, titleHtml, styleHtml)
	bodyHtml := fmt.Sprintf("<body><div id=\"wrapper\" align=\"left\">%s</div></body>", body)

	return fmt.Sprintf("<html>%s%s</html>", headHtml, bodyHtml)

}
