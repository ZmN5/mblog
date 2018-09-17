package blog

import (
	"fmt"
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

func ParseMarkdownName(filename string) (int, string, error) {
	ss := strings.Split(filename, "--")
	Id, err := strconv.Atoi(ss[0])
	if err != nil {
		return -1, "", err
	}
	return Id, strings.TrimSuffix(ss[1], ".md"), nil
}

func parseHtml(title string, body string) string {
	metaHtml := "<meta charset=\"UTF-8\">"
	titleHtml := fmt.Sprintf("<title>%s</title>", title)
	styleHtml := fmt.Sprintf("<style>#wrapper{width: 960px;margin: 0 auto;border:0px solid;}</style>")

	headHtml := fmt.Sprintf("<head>%s%s%s</head>", metaHtml, titleHtml, styleHtml)
	bodyHtml := fmt.Sprintf("<body><div id=\"wrapper\" align=\"left\">%s</div></body>", body)

	return fmt.Sprintf("<html>%s%s</html>", headHtml, bodyHtml)

}
