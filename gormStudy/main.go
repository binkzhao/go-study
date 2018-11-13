package gormStudy

import (
	"github.com/fvbock/endless"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"syscall"
)

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func main() {
	router := gin.Default()
	router.POST("/login-json", func(ctx *gin.Context) {
		var json Login
		if err := ctx.BindJSON(&json); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "params error"})
		} else {
			if json.User == "admin" && json.Password == "admin" {
				ctx.JSON(http.StatusOK, gin.H{"message": "you have login success"})
			} else {
				ctx.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			}
		}
	})

	server := endless.NewServer("127.0.0.1:8008", router)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d\n", syscall.Getpid())
		log.Printf("Service runing addr: %s\n", add)
	}
	err := server.ListenAndServe()

	if err != nil {
		log.Println("Service run error: ", err)
	}

	log.Println("All servers stopped.Exiting.")
	os.Exit(0)
}
