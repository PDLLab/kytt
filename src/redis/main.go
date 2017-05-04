package main
import (
    "github.com/go-redis/redis"
    "fmt"
)    

func NewRedisClient() *redis.Client {
    client := redis.NewClient(&redis.Options {
        Addr: "127.0.0.1:6379",
        Password: "",
        DB: 0,
    })
    pong, err := client.Ping().Result()
    fmt.Println(pong, err)
    return client
}

func main() {
    cli := NewRedisClient()
    err := cli.Set("name", "zcf", 0).Err()
    if err != nil {
        panic(err)
    }

    val, err := cli.Get("name").Result()
    if err != nil {
        panic(err)
    }
    fmt.Println("name:", val)
}


