package user

import (
	"fmt"
	"yournovel/yournovel/conf"
)

func TestUser() {

	fmt.Println(conf.Config["host"])
}
