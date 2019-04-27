package api

import (
	option "bbssh/options/sshrsa"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
)

// StartServer 启动web服务接口
func StartServer(bindAddr string) {
	router := gin.Default()
	router.GET("/pkey/:name", func(c *gin.Context) {
		name := c.Param("name")
		fname := option.KeyDir + "/" + name + "_rsa.pri"
		pri, err := ioutil.ReadFile(fname)
		if err == nil {
			c.JSON(200, map[string]string{"pkey": string(pri)})
		} else {
			log.Fatalf("%v", err)
			c.JSON(200, map[string]string{"pkey": ""})
		}
	})
	router.POST("/validate", func(c *gin.Context) {
		c.JSON(200, map[string]bool{"success": true})
	})
	router.Run(bindAddr)
}
