package test1

import (
	"embed"
	_ "embed"
	"fmt"
)

//go:embed build
var Build embed.FS

//go:embed template
var Template embed.FS

//go:embed txt/hello.txt
var info string

func TestEmbedTxtStr() {
	fmt.Println(info)
}
