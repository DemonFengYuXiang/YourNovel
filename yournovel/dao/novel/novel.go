package novel

import (
	"fmt"
	"yournovel/yournovel/conf"
)

func TestNovel() {
	fmt.Println(conf.Config["host"])
	conf.Config["host"] = "10.1.1.1"
}
