package blog

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	// "os"
	"path"
	"sort"
	"strings"
)

var Cache = make(map[int]string)

var StorageMap = make(MarkdownStorageMap)

func init() {
	StorageMap.Init()
}

type MarkdownStorage struct {
	Id    int
	Title string
	Ext   string
}

func (md MarkdownStorage) String() string {
	return fmt.Sprintf("%d--%s.%s", md.Id, md.Title, md.Ext)
}

func (md MarkdownStorage) FilePath() string {
	return path.Join(MARKDOWN_PATH, md.String())
}

//func (md MarkdownStorage) Delete() {
//	deletedpath := strings.Join(md.RawFilePath, "_del")
//	os.Rename(md.RawFilePath, deletedpath)
//}

func (md MarkdownStorage) ReadHtml() string {
	if md.Ext == "md" {
		return md.ReadHtmlFromFile()
	} else {
		return md.ReadHtmlFromDir()
	}
}

func (md MarkdownStorage) ReadHtmlFromDir() string {
	files, err := ioutil.ReadDir(md.FilePath())

	if err != nil {
		return ""
	}
	for _, file := range files {
		if file.IsDir() {
			continue
		}
		filename := file.Name()
		if !strings.HasSuffix(filename, "md") {
			continue
		}

		body, err := ioutil.ReadFile(path.Join(md.FilePath(), filename))
		if err != nil {
			return ""
		}
		html := string(blackfriday.Run(body))
		return parseHtml(md.Title, html)
	}
	return ""
}

func (md MarkdownStorage) ReadHtmlFromFile() string {
	var body []byte
	var err error
	body, err = ioutil.ReadFile(md.FilePath())
	if err != nil {
		return ""
	}
	html := string(blackfriday.Run(body))
	return parseHtml(md.Title, html)
}

type MarkdownStorageMap map[int]MarkdownStorage

func (mds MarkdownStorageMap) Count() int {
	return len(mds)
}

func (mds MarkdownStorageMap) Init() error {
	files, err := ioutil.ReadDir(MARKDOWN_PATH)
	if err != nil {
		return err
	}
	for _, file := range files {
		filename := file.Name()
		id, title, ext, err := ParseMarkdownName(filename)
		if err != nil {
			return err
		}
		md := MarkdownStorage{
			Id:    id,
			Title: title,
			Ext:   ext,
		}
		mds[id] = md
	}
	return nil
}

func (mds MarkdownStorageMap) Append(md MarkdownStorage) int {
	maxId := 0
	for id, _ := range mds {
		if id > maxId {
			maxId = id
		}
	}
	maxId++
	mds[maxId] = md
	return maxId
}

// func (mds MarkdownStorageMap) Delete(id int) {
// 	mds[id].Delete()
// 	delete(mds, id)
// }

// func (mds MarkdownStorageMap) Update(id int, md MarkdownStorage) error {
// 	if md.Markdown == nil {
// 		return fmt.Errorf("The markdown is empty!")
// 	}
// 	_, ok := mds[id]
// 	if !ok {
// 		return fmt.Errorf("The markdown id doesn't exists!")
// 	}
// 	mds.Delete(id)
// 	md.SaveMarkDown()
// 	mds[id] = md
// 	delete(mds, id)
// 	return nil
// }

func (mds MarkdownStorageMap) SortList() []MarkdownStorage {
	var allId []int
	for id := range mds {
		allId = append(allId, id)
	}
	sort.Ints(allId)
	var sortedMds []MarkdownStorage
	for _, id := range allId {
		sortedMds = append(sortedMds, mds[id])
	}
	return sortedMds
}

func (mds MarkdownStorageMap) IndexMarkdown() string {
	header := "# GoPy\n---\n\n"
	var titles []string
	for _, md := range mds.SortList() {
		title := fmt.Sprintf("### %d. [%s](post/%d)", md.Id, md.Title, md.Id)
		titles = append(titles, title)
	}
	titleStr := strings.Join(titles, "\n")
	return header + titleStr
}

func (mds MarkdownStorageMap) IndexHtml() string {
	body := []byte(mds.IndexMarkdown())
	html := string(blackfriday.Run(body))
	return parseHtml("FCY", html)
}
