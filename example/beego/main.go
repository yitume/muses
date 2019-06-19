package main

import (
	"fmt"
	"github.com/yitume/muses"
	"github.com/yitume/muses/pkg/tmpl/beego"
)

var cfg = `
[muses.tmpl.beego]
    debug = true
`

func main() {
	if err := muses.Container(
		[]byte(cfg),
		beego.Register,
	); err != nil {
		panic(err)
	}
	obj, err := beego.Caller("index")
	if err != nil {
		fmt.Println("err------>", err)
		return
	}
	obj.Data["hello"] = "hello yitu"
	output, err := obj.RenderBytes()
	if err != nil {
		fmt.Println("err------>", err)
		return
	}

	fmt.Println(string(output))
}
