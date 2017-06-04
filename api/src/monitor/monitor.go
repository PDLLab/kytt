package main
import (
    "os/exec"
    "fmt"
    "strings"
    "time"
//    "io/ioutil"
    "encoding/json"
    "net/http"
    "bytes"
) 

type WhoInfo struct {
    mLogin          map[string]bool
    mScanInterval   int
    mExitChan       chan int
}

type DingContent struct {
    Content         string      `json:"content"`
}

type DingAtMobiles struct {
    AtMobiles       []string    `json:"atMobiles"`
}

type DingAt struct {
    DingAtMobiles
    IsAtAll         bool        `json:"isAtAll"`
}

type DingMessage struct {
    Msgtype         string      `json:"msgtype"`
    Text            DingContent `json:"text"`
    At              DingAt      `json:"at"`
}

func RunCommand(name string, args []string) (error, []byte) {
    cmd := exec.Command(name, args...)
    output, err := cmd.CombinedOutput()
    if err != nil {
        return err, output
    }
    return nil, output
}

func (this *WhoInfo) MonitorRun() {
    for {
        cmd := "who"
        args := []string{"-s"}
        err, output := RunCommand(cmd, args)
        if err != nil {
        } else {
            info := strings.Split(string(output), "\n")
            login := make(map[string]bool)

            for _, i := range info {
                if i != "" {
                    login[i] = true
                    _, ok := this.mLogin[i]
                    if !ok {
                        this.Notify("有新用户登录服务器：" +  i)
                        //fmt.Println("有新用户登录服务器：", i)
                        // send message
                    }
                }
            }
            this.mLogin = login
            fmt.Println(this.mLogin)
        }
        time.Sleep(time.Duration(this.mScanInterval) * time.Second)
    }
    this.mExitChan <- 1
}    

func (this *WhoInfo) Notify(content string) {
    var dingMessage DingMessage
    dingMessage.Msgtype = "text"
    dingMessage.Text.Content = content
    dingMessage.At.AtMobiles = []string{}
    dingMessage.At.IsAtAll = true
    message, _ := json.Marshal(dingMessage)
    url := "https://oapi.dingtalk.com/robot/send?access_token=766333e4a0e47de7a22f2c79dac9254154bcbda9741a990ff62bdc4367bc957c"
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
    if err != nil {
        fmt.Println("err = ", err)
    }
    req.Header.Set("Content-Type", "application/json")
    client := &http.Client{}
    _, err = client.Do(req)
    if err != nil {
        fmt.Println("err = ", err)
    }
/*    defer resp.Body.Close()
    fmt.Println(resp.Status)
    fmt.Println(resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println(string(body))*/
}
 
func (this *WhoInfo) Init() {
    this.mLogin = make(map[string]bool)
    this.mScanInterval = 2
    this.mExitChan = make(chan int)
}

func main() {
    var whoInfo WhoInfo
    whoInfo.Init()
    go whoInfo.MonitorRun()
    <- whoInfo.mExitChan
}

//USER     TTY      FROM             LOGIN@   IDLE   JCPU   PCPU WHAT
