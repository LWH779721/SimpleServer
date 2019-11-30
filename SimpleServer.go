package main

import (
    "github.com/gin-gonic/gin"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jinzhu/gorm"
    "net/http"
    "os"
    "io"
    "fmt"
)

func index(c *gin.Context){
    c.String(200, "Hello, liaowenhu")
    //c.JSON(200, gin.H{"name":"liaowenhua","age":"27"})
}

func snapshotHandler(c *gin.Context){
    file, header, err := c.Request.FormFile("image")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"result":"Bad request"})
        return
    }
    
    filename := header.Filename
    if filename == "" {
        c.JSON(http.StatusBadRequest, gin.H{"result":"Bad request"})
        return
    }
 
    fmt.Println(filename)
    out, err := os.Create("upload/"+filename)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"result":"Bad request"})
        return
    }
    defer out.Close()
    _, err = io.Copy(out, file)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"result":"Bad request"})
        return
    }
    
    c.String(http.StatusCreated, "upload successful")
}

type User struct{
    Sno  string 
    Name string
}

var db *gorm.DB
var err error

func main(){
    db, err = gorm.Open("mysql", "lwh:123456@/db1?charset=utf8&parseTime=True&loc=Local")
    if err != nil{
        fmt.Println(err)
        return
    }
    
    defer db.Close()
    db.AutoMigrate(&User{})
    
    r := gin.Default()
    r.GET("/", index)
    r.POST("/snapshot", snapshotHandler)
    
    r.Run()
}