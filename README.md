# RESTful API Golang + MySQL 

 kali ini saya membagikan bagaiamna cara membuat restfull api dengan  *Golang**, **Gin**, dan **MySQL** dengan struktur minimal cukup memakai **1 file main.go** + beberapa folder dasar.

 cara ini saya dapatkan dari kelas hacktiv belajar go free

---

## Struktur Folder Sederhana

```
project/
│  main.go
├─ models/
│   └─ post.go
│   └─ setup.go
├─ controllers/
│   └─ postController.go
```

---

## Install Dependensi

```
go mod init project

go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm gorm.io/driver/mysql
```

---

## File: models/database.go

```go
package models

import (
    "fmt"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
    dsn := "root:password@tcp(127.0.0.1:3306)/testdb?charset=utf8mb4&parseTime=True&loc=Local"

    database, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("Failed to connect to database!")
    }

    DB = database
    fmt.Println("Database connected...")
}
```

---

## File: controllers/postController.go

```go
package controllers

import (
    "net/http"
    "github.com/gin-gonic/gin"
)

func GetPost(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "title": "Hello From Controller",
        "content": "This is a simple post API",
    })
}
```

---

##  File: main.go

```go
package main

import (
    "fmt"
    "project/controllers"
    "project/models"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.GET("/ping", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "pong"})
    })

    for _, r := range router.Routes() {
        fmt.Println(r.Method, r.Path)
    }

    models.ConnectDatabase()

    router.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "hai broe",
        })
    })

    router.GET("/api/post", controllers.GetPost)

    router.Run(":3000")
}
```

---

## Cara Running

```
go run main.go
```

Server berjalan di:
```
http://localhost:3000
```

---

## Endpoint yang ada

| Method | Endpoint      | Deskripsi |
|--------|----------------|-----------|
| GET    | `/ping`        | Cek server |
| GET    | `/`            | Home |
| GET    | `/api/post`    | Ambil data post sederhana |

---

## Impoertant
- Struktur ini **sengaja dibuat sederhana**, supaya mudah paham.
- Boleh di tambahkan module lain nya:
  - CRUD lengkap
  - JWT Auth
  - Middleware
  - Router Group

 ## let's GO🐧

