package main

import (
	"cinema/internal/database"
	"fmt"

	"github.com/gin-gonic/gin"
)

func main(){
    r := gin.Default()

    db := database.GetDBInstance()

    fmt.Printf("db: %v\n", db)
    
    r.Run(":8080")
}