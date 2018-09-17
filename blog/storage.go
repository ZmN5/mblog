package blog

import (
	"fmt"
	"github.com/russross/blackfriday"
	"io/ioutil"
	"os"
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
	Id       int
	Title    string
	Markdown []byte
}

func (md MarkdownStorage) getFilePath() string {
	return path.Join(MARKDOWN_PATH, fmt.Sprintf("%d--%s.md", md.Id, md.Title))

}

func (md MarkdownStorage) SaveMarkDown() error {
	data := []byte(md.Markdown)
	fd, err := TouchFile(md.getFilePath())
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

func (md MarkdownStorage) Delete() {
	oldpath := md.getFilePath()
	deletedpath := strings.Replace(oldpath, ".md", ".old", 1)
	os.Rename(oldpath, deletedpath)
}

func (md MarkdownStorage) ReadHtml() string {
	var body []byte
	var err error
	if md.Markdown == nil {
		body, err = ioutil.ReadFile(md.getFilePath())
		if err != nil {
			return ""
		}
		md.Markdown = body
	} else {
		body = md.Markdown
	}
	html, ok := Cache[md.Id]
	if !ok {
		html = string(blackfriday.Run(body))
		if md.Id > 0 {
			Cache[md.Id] = html
		}
	}
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
		isDir := file.IsDir()
		if !isDir && strings.HasSuffix(filename, ".md") {
			id, title, err := ParseMarkdownName(filename)
			if err == nil {
				md := MarkdownStorage{
					Id:       id,
					Title:    title,
					Markdown: nil,
				}
				mds[id] = md
			}
		}
	}
	return nil
}

func (mds MarkdownStorageMap) Append(md MarkdownStorage) error {
	if md.Markdown == nil {
		return fmt.Errorf("The markdown is empty!")
	}
	maxId := 0
	for id, _ := range mds {
		if id > maxId {
			maxId = id
		}
	}
	maxId++
	md.Id = maxId
	md.SaveMarkDown()
	mds[maxId] = md
	return nil
}

func (mds MarkdownStorageMap) Delete(id int) {
	mds[id].Delete()
	delete(mds, id)
}

func (mds MarkdownStorageMap) Update(id int, md MarkdownStorage) error {
	if md.Markdown == nil {
		return fmt.Errorf("The markdown is empty!")
	}
	_, ok := mds[id]
	if !ok {
		return fmt.Errorf("The markdown id doesn't exists!")
	}
	mds.Delete(id)
	md.SaveMarkDown()
	mds[id] = md
	delete(mds, id)
	return nil
}

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
	b := []byte(mds.IndexMarkdown())
	md := MarkdownStorage{Id: -1, Title: "Index", Markdown: b}
	return md.ReadHtml()
}
