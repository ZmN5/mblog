package blog

import (
	"os"
)

// const WORKDIR = "/data/blog/"
const WORKDIR = "/Users/cangyufu/Documents/articles/"
const MARKDOWN_PATH = WORKDIR + "markdown"
const COMPRESS_FILE_PATH = WORKDIR + "compress"
const CERTS = WORKDIR + "certs"

var AUTH = os.Getenv("AUTH")
var MODE = os.Getenv("MODE")
var DOMAIN = os.Getenv("DOMAIN")

func init() {

	if MODE == "" {
		MODE = "HTTP"
	}
}
