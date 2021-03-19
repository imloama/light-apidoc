package main

import (
	_ "github.com/imloama/light-apidoc/boot"
	_ "github.com/imloama/light-apidoc/router"

	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
