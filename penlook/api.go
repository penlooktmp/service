package penlook

import (
	"github.com/penlook/gin"
	lib "github.com/penlook/service/penlook/library"
)

func Api() {
	doc := lib.doc
	doc.Print()

}
