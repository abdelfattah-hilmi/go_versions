package main

import (
	r "example/go_versions/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()

	// defining endpoints

	router.POST("/vms/installed-packages", r.AddVm)
	router.GET("/vms/installed-packages", r.GetVms)

	router.Run("localhost:8000")
}
