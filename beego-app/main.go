package main

import (
	_ "beego-app/routers"
	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	beego.Run()
}
