package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.RawQuery)
	fmt.Fprintf(w, `Hello World.`)
}

func startHttpServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func startRedis() {
	var ctx = context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "172.17.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	err := rdb.Set(ctx, "user:100", "lile", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "user:001").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("user:100 =>", val)
}

func main() {
	fmt.Print(`
	##         .
## ## ##        ==
## ## ## ## ##    ===
/"""""""""""""""""\___/ ===
{                       /  ===-
\______ O           __/
\    \         __/
\____\_______/

Hello from Docker!
`)
	startRedis()
	startHttpServer()
}
