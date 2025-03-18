package main

import (
	"./s"
	"github.com/gin-gonic/gin"
)

func main() {
	s.ListExcludeHost()
	//fmt.Println(s.ExcludeHost)
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.POST("/post", func(ctx *gin.Context) {
		s.Analyzing(ctx.PostForm("Post"))
		//log.Println("s.HostPath:", s.HostPath)
	})

	err := r.Run(":1010")
	if err != nil {
		panic("启动失败")
	}
}
