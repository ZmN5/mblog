package blog

import "os"

const MARKDOWN_PATH = "/data/blog/markdown"
const HTML_PATH = "/data/blog/html"

var AUTH = os.Getenv("AUTH")
