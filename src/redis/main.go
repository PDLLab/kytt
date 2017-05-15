package main
import (
    "github.com/go-redis/redis"
    "fmt"
    "sync"
    "log"
    "os"
    "io"
)    

var Logger *log.Logger

type RedisClient struct {
    Client  *redis.Client
}

// 连接客户端
func NewRedisClient() *redis.Client {
    Logger.Println("[NewRedisClient Begin]")
    client := redis.NewClient(&redis.Options {
        Addr: "127.0.0.1:6379",
        Password: "",
        DB: 0,
        PoolSize: 1024, // 连接池大小
    })
    _, err := client.Ping().Result()
    if err != nil {
        Logger.Panicln("[NewRedisClient] err = ", err)
        return nil
    }
    Logger.Println("[NewRedisClient End]")
    return client
}

func SetKeyValue(key string, value string, client *redis.Client) {
    Logger.Println("[SetKeyValue Begin] key = ", key, " value = ", value)
    client.SAdd(key, value)
    Logger.Println("[SetKeyValue End]")
}

func GetKeyValue(key string, client *redis.Client) *redis.StringSliceCmd {
    Logger.Println("[GetKeyValue Begin] key = ", key)
    return client.SMembers(key)
}

func (RC *RedisClient) SetUserFollower(userID string, followerID string) {
    Logger.Println("[SetUserFollower Begin] userID = ", userID, " followerID = ", followerID)
    prefix := "user:follower:"
    key := prefix + userID
    value := followerID
    SetKeyValue(key, value, RC.Client) 
    Logger.Println("[SetUserFollower End]")
}

func (RC *RedisClient) GetUserFollower(userID string) []string {
    Logger.Println("[GetUserFollower Begin] userID = ", userID)
    prefix := "user:follower:"
    key := prefix + userID
    ret := GetKeyValue(key, RC.Client)
    all, err := ret.Result()
    if err != nil {
        Logger.Panicln("[GetUserFollower] err = ", err)
    }

    Logger.Println("[GetUserFollower End] all = ", all)
    return all 
}

func (RC *RedisClient) SetHeadlineFollower(headlineID, followerID string) {
    Logger.Println("[SetHeadlineFollower Begin] headlineID = ", headlineID, " followerID = ", followerID)
    prefix := "headline:follower:"
    key := prefix + headlineID
    value := followerID
    SetKeyValue(key, value, RC.Client) 
    Logger.Println("[SetHeadlineFollower End]")
}
    
func (RC *RedisClient) GetHeadlineFollower(headlineID string) []string {
    Logger.Println("[GetHeadlineFollower Begin] headlineID = ", headlineID)
    prefix := "headline:follower:"
    key := prefix + headlineID
    ret := GetKeyValue(key, RC.Client)
    all, err := ret.Result()
    if err != nil {
        Logger.Panicln("[GetHeadlineFollower] err = ", err)
    }

    Logger.Println("[GetHeadlineFollower End] all = ", all)
    return all 
}

func (RC *RedisClient) SetHeadlineCommnet(headlineID, commentUserID string) {
    Logger.Println("[SetHeadlineComment Begin] headlineID = ", headlineID, " commentUserID = ", commentUserID)
    prefix := "headline:comment:"
    key := prefix + headlineID
    value := commentUserID
    SetKeyValue(key, value, RC.Client) 
    Logger.Println("[SetHeadlineComment End]")
}
    
func (RC *RedisClient) GetHeadlineComment(headlineID string) []string {
    Logger.Println("[GetHeadlineComment Begin] headlineID = ", headlineID)
    prefix := "headline:comment:"
    key := prefix + headlineID
    ret := GetKeyValue(key, RC.Client)
    all, err := ret.Result()
    if err != nil {
        Logger.Panicln("[GetHeadlineComment] err = ", err)
    }

    Logger.Println("[GetHeadlineComment End] all = ", all)
    return all 
}

func (RC *RedisClient) SetHeadlineLike(headlineID, likeUserID string) {
    Logger.Println("[SetHeadlineLike Begin] headlineID = ", headlineID, " likeUserID = ", likeUserID)
    prefix := "headline:like:"
    key := prefix + headlineID
    value := likeUserID
    SetKeyValue(key, value, RC.Client) 
    Logger.Println("[SetHeadlineLike End]")
}
    
func (RC *RedisClient) GetHeadlineLike(headlineID string) []string {
    Logger.Println("[GetHeadlineLike Begin] headlineID = ", headlineID)
    prefix := "headline:like:"
    key := prefix + headlineID
    ret := GetKeyValue(key, RC.Client)
    all, err := ret.Result()
    if err != nil {
        Logger.Panicln("[GetHeadlineLike] err = ", err)
    }

    Logger.Println("[GetHeadlineLike End] all = ", all)
    return all 
}

func (RC *RedisClient) SetCommentCommnet(commentID, commentUserID string) {
    Logger.Println("[SetCommentComment Begin] commentID = ", commentID, " commentUserID = ", commentUserID)
    prefix := "comment:comment:"
    key := prefix + commentID
    value := commentUserID
    SetKeyValue(key, value, RC.Client) 
    Logger.Println("[SetCommentComment End]")
}
    
func (RC *RedisClient) GetCommentComment(commentID string) []string {
    Logger.Println("[GetCommentComment Begin] commentID = ", commentID)
    prefix := "comment:comment:"
    key := prefix + commentID
    ret := GetKeyValue(key, RC.Client)
    all, err := ret.Result()
    if err != nil {
        Logger.Panicln("[GetCommentComment] err = ", err)
    }

    Logger.Println("[GetCommentComment End] all = ", all)
    return all 
}

func (RC *RedisClient) SetCommentLike(commentID, likeUserID string) {
    Logger.Println("[SetCommentLike Begin] commentID = ", commentID, " likeUserID = ", likeUserID)
    prefix := "comment:like:"
    key := prefix + commentID
    value := likeUserID
    SetKeyValue(key, value, RC.Client) 
    Logger.Println("[SetCommentLike End]")
}
    
func (RC *RedisClient) GetCommentLike(commentID string) []string {
    Logger.Println("[GetCommentLike Begin] commentID = ", commentID)
    prefix := "comment:like:"
    key := prefix + commentID
    ret := GetKeyValue(key, RC.Client)
    all, err := ret.Result()
    if err != nil {
        Logger.Panicln("[GetCommentLike] err = ", err)
    }

    Logger.Println("[GetCommentLike End] all = ", all)
    return all 
}

// benchmark
func TestingPool(client *redis.Client) {
    wg := sync.WaitGroup{}
    wg.Add(1000)

    for i := 0; i < 1000; i++ {
        go func() {
            defer wg.Done()

            for j := 0; j < 100; j++ {
                client.Set(fmt.Sprintf("name%d", j), fmt.Sprintf("xys%d", j), 0).Err()
                client.Get(fmt.Sprintf("name%d", j)).Result()
            }

            fmt.Printf("PoolStats, TotalConns: %d, FreeConns: %d\n", client.PoolStats().TotalConns, client.PoolStats().FreeConns);
        }()
    }

    wg.Wait()
}

func main() {
    logFile, err := os.Create("redis_debug.log")
    defer logFile.Close()
    if err != nil {
        log.Fatalln("open file error.")
    }
    
    writers := []io.Writer {
        logFile,
        os.Stdout,
    }
    fileAndStdoutWriter := io.MultiWriter(writers...)
    Logger = log.New(fileAndStdoutWriter, "[Debug]", log.LstdFlags | log.Lshortfile)

    var redisClient RedisClient
    redisClient.Client = NewRedisClient()
    redisClient.SetUserFollower("001", "002")
    redisClient.SetUserFollower("001", "003")
    redisClient.GetUserFollower("001")

    redisClient.SetHeadlineFollower("123123", "123456")
    redisClient.GetHeadlineFollower("123123")
    redisClient.GetHeadlineFollower("1234")

    Logger.Println("[main End]")
}


