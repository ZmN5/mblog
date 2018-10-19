package blog

import "os"

const WORKDIR = "/data/blog/"
const MARKDOWN_PATH = WORKDIR + "markdown"
const HTML_PATH = WORKDIR + "html"
const CERTS = WORKDIR + "certs"

var AUTH = os.Getenv("AUTH")
var DOMAIN = os.Getenv("DOMAIN")
var MODE = os.Getenv("MODE")
