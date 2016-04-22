package main // "import.moetang.info/go/tool/go-import-server"

import (
	"log"

	"github.com/gin-gonic/gin"
)

import (
	"import.moetang.info/go/lib/gin-startup"
)

func main() {
	entries, err := ReadEntries()
	if err != nil {
		log.Println(err)
		return
	}

	g := gin_startup.NewGinStartup()
	g.EnableFastCgi("tcp://127.0.0.1:15001")
	g.EnableHttp("tcp://127.0.0.1:15002")
	g.Custom(func(r *gin.Engine) {
		r.Use(gin.Recovery(), gin.Logger())

		for _, v := range entries {
			RegisterAction(v, r)
		}
	})
	g.Start()
	c := make(chan bool)
	<-c
}
