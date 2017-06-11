package main

import (
    "github.com/julienschmidt/httprouter"
    "net/http"
    "encoding/json"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "time"
    "os"
    "io"
    "io/ioutil"
)

const (
    // username:password@protocol(address)/dbname?param=value 
    cConnectString = "root:123456@tcp(localhost:3306)/kytt?charset=utf8"
    cMaxConnectionCount = 128
    cLoggerFile = "server.log"
)

var Logger *log.Logger

type Server struct {
    mRouter             *httprouter.Router
    mDB                 *sql.DB
    mLogFile         *os.File
}

type RecommendData struct {
    Id              int                 `json:"id"`
    UserId          int                 `json:"userId"`
    UserNickname    string              `json:"userNickname"`
    Title           string              `json:"title"`
    PostDate        string              `json:"postDate"`
    LikeCount       int                 `json:"likeCount"`
    CommentCount    int                 `json:"commentCount"`
    ForwardCount    int                 `json:"forwardCount"`
    ViewCount       int                 `json:"ViewCount"`
    Tag             string              `json:"tag"`
    Summary         string              `json:"summary"`
    QuestionId      int                 `json:"questionId"`
}

type Recommend struct {
    Type            string              `json:"type"`
    IsOfficial      bool                `json:"isOfficial"`
    Data            RecommendData       `json:"data"`
}

type Content struct {
    Content         string              `json:"content"`
}

type BaseHeadline struct {
    Id              string              `json:"id"`
    UserId          string              `json:"userId"`
    UserNickname    string              `json:"userNickname"`
    Title           string              `json:"title"`
    Content         string              `json:"content"`
    PostDate        string              `json:"postDate"`
    LikeCount       int                 `json:"likeCount"`
    CommentCount    int                 `json:"commentCount"`
    ForwardCount    int                 `json:"forwardCount"`
    ViewCount       int                 `json:"viewCount"`
    Tag             string              `json:"tag"`
    TitleImage      string              `json:"titleImage"`
    Summary         string              `json:"summary"`
}

type Headline struct {
    BaseHeadline
    IsOfficial      bool                `json:"isOfficial"`
}

func (this *Server) StmtPrepare(operation string) *sql.Stmt {
    stmt, err := this.mDB.Prepare(operation)
    if err != nil {
        Logger.Panic("[StmtPrepare] err = ", err)
    }
    return stmt
}

func GetCurrentTime() string {
    return time.Now().Format("2006-01-02 15:04:05")
}

func (this *Server) InitRouter() {
    this.mRouter.GET("/v1/home/recommend/:userId", this.GetUserRecommend)
    this.mRouter.GET("/v1/headlines/content/:headlineId", this.GetHeadlineContent)
    this.mRouter.GET("/v1/answer/content/:answerId", this.GetAnswerContent)
    this.mRouter.POST("/v1/headlines", this.PostHeadlines)
    //this.mRouter.POST("/v1/questions", this.PostQuestions)
    //this.mRouter.POST("/v1/answers", this.PostAnswers)


    this.mRouter.POST("/v1/signup", this.PostSignup)
    this.mRouter.POST("/v1/signin", this.PostSignin)

    this.mRouter.GET("/v1/headlines", this.GetHeadlines)
//    this.mRouter.GET("/v1/headlines/:headlineId", this.GetHeadline)
//    this.mRouter.POST("/v1/headlines/:headlineId/comments", this.PostHeadlineComments)
//    this.mRouter.POST("/v1/headlines/:headlineId/likes", this.PostHeadlineLikes)
    
    this.mRouter.POST("/v1/comment/:commentId/comments", this.PostCommentComments)
    this.mRouter.POST("/v1/comment/:commentId/likes", this.PostCommentLikes)

    this.mRouter.GET("/v1/users/:userId/followings", this.PostUserFollowings)
 //   this.mRouter.GET("/v1/headlines/:headlineId/followings", this.PostHeadlineFollowings)
}

func (this *Server) GetUserRecommend(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    userId := ps.ByName("userId")
    Logger.Println("[GetUserRecommend] userId = ", userId)
    var recommend []Recommend
    var mockRecommend Recommend
    mockRecommend.Type = "headline"
    mockRecommend.IsOfficial = true
    mockRecommend.Data.Id = 12832401734
    mockRecommend.Data.UserId = 238978232
    mockRecommend.Data.UserNickname = "迷糊的小蘑菇"
    mockRecommend.Data.Title = "以梦为马——选择比努力重要（跨专业没你想象的那么难）"
    mockRecommend.Data.PostDate = "2017-06-07 09:27:23"
    mockRecommend.Data.LikeCount = 306
    mockRecommend.Data.CommentCount = 46
    mockRecommend.Data.ForwardCount = 2
    mockRecommend.Data.ViewCount = 283232
    mockRecommend.Data.Tag = "考研,会计,经验"
    mockRecommend.Data.Summary = "　这篇是写给所有想跨专业但在犹豫的考生的强心剂，也是想考现当代文学的考生的经验帖，其实去年就想把自己艰难的考研之路记录下来，但觉得全是教训，负能量满满，不想提笔正视当时的自己，但如今已经确认考上自己喜欢的学校，喜欢的专业，并且以初试400+的成绩跨入自己只用不到七个月时间准备的专业，所以真心想用自己的亲身经历给想跨专业的战友们打打鸡血，也让希望走上研究生道路的小汪们避免一些弯路。"

    recommend = append(recommend, mockRecommend)
    content, err := json.Marshal(recommend)
    if err != nil {
        Logger.Println("[GetUserRecommend] = err", err)
    }
    
    w.Write(content)
    w.WriteHeader(http.StatusOK)
}

func (this *Server) GetHeadlineContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    headlineId := ps.ByName("healineId")
    Logger.Println("[GetHeadlineContent] headlineId = ", headlineId)
    var content Content
    content.Content = `<p>(要转载可以私信说，不打招呼就拿那叫偷)</p><p>今年4月，李安在NAB年会上首次发布了新片《半场无战事》的11分钟片段，前段时间我们也得以见到这部电影的首支预告片（<a href="http://link.zhihu.com/?target=https%3A//movie.douban.com/trailer/196573/%23content" class=" wrap external" target="_blank" rel="nofollow noreferrer">《半场无战事》中文预告片<i class="icon-external"></i></a>），之后所有的讨论都是围绕这三个数字——3D、4K、120fps——展开的，3D大家都比较了解，4K分辨率的播放设备也逐渐开始应用，iphone6S已经能拍摄4K视频，最后这个120fps才是李安这次冒险的重头戏，他这是把电影拍摄的帧率一下子提高了四倍，如果这部电影最终呈现的效果令人满意，李安带来的技术革新意义堪比当年的阿凡达在3D电影上的技术革新。<br></p><br>120帧/秒，或许将和VR一样成为电影的下一个关键词。从24到120，电影拍摄帧率是怎么一步步进化的？<br><br>其实早期的电影拍摄并没有一个标准，因为那时候采用的是手摇式摄影装置，拍摄的时候难免会有拍摄速度不统一，而且帧率相差的非常大，比如说爱迪生的电影是以40帧/秒拍摄的，而卢米埃兄弟的电影则是以16帧/秒拍摄的。<br><br>后来发展到有声电影，引入了同期录音技术，可将声画同时录制到胶片上，24帧/秒能保持最高的声音清晰度，如果再低的话，音轨上就会有太多的表面杂音，于是24帧/秒就成为了电影拍摄帧率的标准。<br><br>超过24帧/秒拍摄的电影就叫“超高帧”电影，有很多电影尝试以超高帧拍摄：<br><br>1950年诞生了“全景电影”（也叫西尼拉玛），这种电影使用3台摄影机分别拍摄3条方形素材，然后由3台放映机无缝拼接，投影到银幕上，以产生3倍宽高比的影像，这种电影是以26帧/秒拍摄的，这种电影主要是拍摄风光电影，后来由于放映设备昂贵逐渐被市场淘汰。<br><noscript><img src="https://pic1.zhimg.com/5b3d635b8678a49c80188ff1f6ddf084_b.jpg" class="content_image">（第一部“西尼拉玛”电影《这就是西尼拉玛》，可以看到屏幕是环形的）</noscript><img src="//zhstatic.zhihu.com/assets/zhihu/ztext/whitedot.jpg" class="content_image lazy" data-actualsrc="https://pic1.zhimg.com/5b3d635b8678a49c80188ff1f6ddf084_b.jpg">（第一部“西尼拉玛”电影《这就是西尼拉玛》，可以看到屏幕是环形的）<br><br>1956年的《环游世界80天》是以30帧/秒拍摄的。<br><noscript><img src="https://pic2.zhimg.com/e31b4a39d69752270af51347b6e9b3b9_b.jpg" data-rawwidth="528" data-rawheight="768" class="origin_image zh-lightbox-thumb" width="528" data-original="https://pic2.zhimg.com/e31b4a39d69752270af51347b6e9b3b9_r.jpg">在1992年的塞尔维亚世界博览会上，第一部使用IMAX HD技术拍摄的电影《Momentum》是以48帧/秒拍摄的。</noscript><img src="//zhstatic.zhihu.com/assets/zhihu/ztext/whitedot.jpg" data-rawwidth="528" data-rawheight="768" class="origin_image zh-lightbox-thumb lazy" width="528" data-original="https://pic2.zhimg.com/e31b4a39d69752270af51347b6e9b3b9_r.jpg" data-actualsrc="https://pic2.zhimg.com/e31b4a39d69752270af51347b6e9b3b9_b.jpg">在1992年的塞尔维亚世界博览会上，第一部使用IMAX HD技术拍摄的电影《Momentum》是以48帧/秒拍摄的。<br><noscript><img src="https://pic2.zhimg.com/0994b10d8297e839765e1d703de9d9c5_b.jpg" data-rawwidth="1294" data-rawheight="1674" class="origin_image zh-lightbox-thumb" width="1294" data-original="https://pic2.zhimg.com/0994b10d8297e839765e1d703de9d9c5_r.jpg">1999年诞生了Maxivision 48技术，是结合了48帧/秒和35mm胶片的底片格式，可以以48帧/秒的速度放映影片，受到了业内人士的推广，但是大多数影院依然采用24帧/秒的放映设备，所以Maxivision仍没有得到广泛的投入使用。</noscript><img src="//zhstatic.zhihu.com/assets/zhihu/ztext/whitedot.jpg" data-rawwidth="1294" data-rawheight="1674" class="origin_image zh-lightbox-thumb lazy" width="1294" data-original="https://pic2.zhimg.com/0994b10d8297e839765e1d703de9d9c5_r.jpg" data-actualsrc="https://pic2.zhimg.com/0994b10d8297e839765e1d703de9d9c5_b.jpg">1999年诞生了Maxivision 48技术，是结合了48帧/秒和35mm胶片的底片格式，可以以48帧/秒的速度放映影片，受到了业内人士的推广，但是大多数影院依然采用24帧/秒的放映设备，所以Maxivision仍没有得到广泛的投入使用。<p>道格拉斯·特鲁姆布（《2001太空漫步》、《星球大战》、《银翼杀手》的特效导演）旗下的Showscan视效工作室，在着力推广高帧率（60帧/秒）结合70mm拍摄技术（能科学上网的可以看一下这个视频介绍 <a href="http://link.zhihu.com/?target=https%3A//www.youtube.com/watch%3Fv%3DNkWLZy7gbLg" class=" wrap external" target="_blank" rel="nofollow noreferrer">Showscan Digital from Douglas Trumbull<i class="icon-external"></i></a>），但是到目前为止，Showscan也依然是纸上谈兵，除了几部试验性的短片，他们还没有尝试长片拍摄。</p><p>道格拉斯之所以推崇高帧率拍摄，是因为：</p><blockquote>“理论上，这样的极致帧率会使动态画面更流畅，解决摄影机摆动过程中产生的频闪或晃动问题，带来更舒适的3D和无可比拟的敏锐与真实感。但这仍然只是理论，因为还从没有人看过用这种技术拍摄的电影。即便这些都被完美地实现了，这个世界上也几乎没有地方能让观众欣赏到它。”</blockquote><p>2011年，卡梅隆宣布他的《阿凡达2》将和Showscan合作，采用他们的60帧/秒技术拍摄。</p><p>2012年，彼得·杰克逊的《霍比特人：意外之旅》先行一步，采用了48帧/秒的速度拍摄，后继两部霍比特人也是以48帧/秒拍摄。当时业内毁誉参半，批评者认为画面太过清晰反而像游戏或油画的质感，缺乏了传统电影的模糊和频闪效果，也就失去了电影的美感（并非所有影院都能以48fps放映，所以放映时候有24fps和48fps两个版本，如果你觉得没有这种体验，很可能是因为当时的影院没有使用48规格放映）。</p><br><p>对此，彼得大帝回应说，“（你们这群没见识的凡人）等你们习惯了就好了”。</p><p><noscript><img src="https://pic2.zhimg.com/651c727b3556b1e82781aa07060f0119_b.jpg" data-rawwidth="1888" data-rawheight="814" class="origin_image zh-lightbox-thumb" width="1888" data-original="https://pic2.zhimg.com/651c727b3556b1e82781aa07060f0119_r.jpg">（霍比特人的画面被批评太像油画质感）</noscript><img src="//zhstatic.zhihu.com/assets/zhihu/ztext/whitedot.jpg" data-rawwidth="1888" data-rawheight="814" class="origin_image zh-lightbox-thumb lazy" width="1888" data-original="https://pic2.zhimg.com/651c727b3556b1e82781aa07060f0119_r.jpg" data-actualsrc="https://pic2.zhimg.com/651c727b3556b1e82781aa07060f0119_b.jpg">（霍比特人的画面被批评太像油画质感）</p><p>李安所采用的影像格式被索尼官方称为“<b>Immersive Digital</b>”(还未有官方的名称，网友译版是“沉浸数字式”)，是第一次将3D、4K、120帧/秒这几中规格结合到一起。</p><p>120帧/秒面对的最大问题就是观影习惯，比如说3D技术刚开始推广的时候，也有诸如画面太暗、会晕的抱怨，李安要解决的问题比彼得·杰克逊还要大，48帧/秒已经给人一种不适应感，120帧/秒带来的观影感受肯定是颠覆性的。</p><p>但是李安并不是受彼得·杰克逊的启发，而是受了卡梅隆的启发，他在拍《少年派》的时候在3D技术上就参考过卡梅隆等人的意见，所有人都告诉他，用24帧/秒拍3D会加大本来就有的频闪（就是看3D的时候会觉得晕的原因），但是当时没有更好的技术解决办法，只能妥协。</p><noscript><img src="https://pic4.zhimg.com/44d3fc1e0c11aa6ba8c5892995277123_b.jpg" data-rawwidth="5200" data-rawheight="2925" class="origin_image zh-lightbox-thumb" width="5200" data-original="https://pic4.zhimg.com/44d3fc1e0c11aa6ba8c5892995277123_r.jpg"></noscript><img src="//zhstatic.zhihu.com/assets/zhihu/ztext/whitedot.jpg" data-rawwidth="5200" data-rawheight="2925" class="origin_image zh-lightbox-thumb lazy" width="5200" data-original="https://pic4.zhimg.com/44d3fc1e0c11aa6ba8c5892995277123_r.jpg" data-actualsrc="https://pic4.zhimg.com/44d3fc1e0c11aa6ba8c5892995277123_b.jpg"><p>后来卡梅隆宣布使用高帧率拍摄《阿凡达2》，给了李安启发，但是他还是觉得不够，于是就想到用4K技术来解决清晰率的问题，最终选择120帧/秒，其实是为了省钱——120正好同时是24帧/秒（电影通用）和30帧/秒（电视通用）的整倍数，就可以用同一种格式制作电视和电影。而李安认为《半场无战事》是一个最好的机会，因为——</p><blockquote>“这部电影的重点就是“感受”，不仅仅是讲述故事，而是想办法让观众体验，给他们一种全新的经历。”</blockquote><p>李安接受3D技术的初衷并不是为了炫技，他和很多使用3D技术的导演和制作公司的意见可以说是完全相反的，就连我们普通观众也认为像阿凡达、漫威的大场面动作电影才应该用3D，而建国大业那种叙事性电影用3D简直是莫名其妙，而李安反而认为：</p><blockquote>“2D和3D最大的区别在于面部拍摄，而不是在于动作或大场面，我觉得2D更适合动作片，而3D更适合拍剧情片，因为3D可以营造更亲密的感受。”</blockquote><p>所以他对于伊拉克战场和中场表演两中不同的主要场景运用了不同的技术处理，战场场景会尽可能地呈现“真实”镜头，也就是用120帧/秒拍摄，预告片里可以看到，在演员跑动的过程中，镜头几乎没有一点摇晃的感觉。</p><p><noscript><img src="https://pic2.zhimg.com/d852d0889935d0a7049d5339b4a74d9d_b.png" data-rawwidth="639" data-rawheight="347" class="origin_image zh-lightbox-thumb" width="639" data-original="https://pic2.zhimg.com/d852d0889935d0a7049d5339b4a74d9d_r.png">（原谅我的渣截图水平）</noscript><img src="//zhstatic.zhihu.com/assets/zhihu/ztext/whitedot.jpg" data-rawwidth="639" data-rawheight="347" class="origin_image zh-lightbox-thumb lazy" width="639" data-original="https://pic2.zhimg.com/d852d0889935d0a7049d5339b4a74d9d_r.png" data-actualsrc="https://pic2.zhimg.com/d852d0889935d0a7049d5339b4a74d9d_b.png">（原谅我的渣截图水平）</p><p>而中场表演场景是在体育馆内，有大量的舞台灯光照明，为了避免大家都担心的背景过于虚化的效果，只有部分镜头（比如无光照的观众席）使用了120帧/秒拍摄。从预告片中可以看到，对于男主角比利的面部特写镜头非常多，着重刻画他的表情和情绪，因为背景虚化本来就很适合大头特写，而中场表演的镜头则大多是全景。</p><p><noscript><img src="https://pic3.zhimg.com/dfc7d68a5720b3a801a38794f559409a_b.jpg" data-rawwidth="3000" data-rawheight="2000" class="origin_image zh-lightbox-thumb" width="3000" data-original="https://pic3.zhimg.com/dfc7d68a5720b3a801a38794f559409a_r.jpg">（目前为止爆出的唯一一张剧照，还看不太出4K、120帧/秒的门道）</noscript><img src="//zhstatic.zhihu.com/assets/zhihu/ztext/whitedot.jpg" data-rawwidth="3000" data-rawheight="2000" class="origin_image zh-lightbox-thumb lazy" width="3000" data-original="https://pic3.zhimg.com/dfc7d68a5720b3a801a38794f559409a_r.jpg" data-actualsrc="https://pic3.zhimg.com/dfc7d68a5720b3a801a38794f559409a_b.jpg">（目前为止爆出的唯一一张剧照，还看不太出4K、120帧/秒的门道）</p><p>在看了11分钟的片段之后，（前文提到的Showscan）道格拉斯说：<br></p><blockquote>“有种身临其境的感觉，看起来他（李安）会给我们带来一次全所未有的观影体验，我简直激动地颤抖了。”</blockquote><p>卡梅隆看过之后更是觉得3D、4K、120帧/秒将成为行业内的最高规格标准，有观众评论“这段11分钟的电影让我觉得镜头、幕布消失了，我看到了最真实的画面。”</p><p>其实120帧/秒更大的问题是放映，就连李安自己在拍摄过程中都没法看到最终效果，因为技术负担太大了，他在片场拍摄的时候只能用3D、2K、60帧/秒的格式播放。</p><p>之前在LA的11分钟试映，是采用了科视的双4K激光投影机和赢康公司的7sense Delta Infinity III服务器，配备杜比3D眼镜，而且当时他们面向的是NAB的行业内技术人员，没有人比他们更懂行，所以获得的大多是惊叹和赞誉。</p><p>但是如果想面向普通大众，李安就要有准备面对整个电影行业的颠覆性改变。普通的影院现在还没有设备可以同时满足3D、4K、120fps的播放规格，普通观众的反应和接受程度也无法预料，至今为止索尼还没有宣布这部电影将如何大范围上映，为了陪李安冒险，索尼在尽可能地寻找方式，甚至考虑了先期放映，但是最终为了大范围上映，很有可能他们只能把标准起码降低到60帧/秒以及用2K代替4K。<br></p><br><p>李安和卡梅隆、道格拉斯等人一样，都站在了先驱者的行列，希望能通过这种革命让电影院看电影重新变成一种朝圣般的体验，一种在电脑手机上无法比拟的体验，如果可以让人们重新走入电影院，那么即使是要让全世界所有的电影院都更换新的放映设备也是值得的。<br></p><br><blockquote> “改变人们的习惯和颠覆一种文化是非常难的，我的好奇心的确是有点旺盛，但是我已经不年轻了，我不愿再等。”</blockquote><p>当他拍《断背山》的时候，很多人劝他的演员，拍这部戏就是葬送自己的演员生涯，结果他的四位演员获得了三个奥斯卡演技提名，当他拍《少年派的奇幻漂流》的时候，所有人都说这部小说太难银幕化了，这是自讨苦吃，然后他拿到了第二座小金人。</p><p>现在，几乎没有人对李安说“不可能”三个字。</p><p>对于《半场无战事》，想借用一部电影的名字《好戏还在后面》：“Vous n'avez encore rien vu 你们见到的还不算什么”，我们拭目以待。</p><p>——————————————————————————————————————————</p><br><p>打假揭幕，还原事实，偶尔来点福利——如果喜欢我们的内容，欢迎关注<a href="http://zhuanlan.zhihu.com/bigertech" class="internal">笔戈科技</a></p>`
    c, err := json.Marshal(content)
    if err != nil {
        Logger.Println("[GetHeadlineContent] err = ", err)
    }
    w.Write(c)
    w.WriteHeader(http.StatusOK)
}

func (this *Server) GetAnswerContent(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    answerId := ps.ByName("answerId")
    Logger.Println("[GetAnswerContent] answerId = ", answerId)
}

func (this *Server) PostQuestion(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

}

func (this *Server) PostSignup(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostSignin(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostHeadlines(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    content, err := ioutil.ReadAll(r.Body)
    if err != nil {
        Logger.Println("[PostHeadlines] err = ", err)
    }
    var headline Headline
    err = json.Unmarshal(content, &headline)
    if err != nil {
        Logger.Println("[PostHeadlines] err = ", err)
    }
    Logger.Println("headline = ", headline)
    w.WriteHeader(http.StatusCreated)
}

func (this *Server) GetHeadlines(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) GetHeadline(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostHeadlineComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostHeadlineLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostCommentComments(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostCommentLikes(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostUserFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) PostHeadlineFollowings(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
}

func (this *Server) InitLogger() {
    logFile, err := os.Create(cLoggerFile)
    this.mLogFile = logFile 
    if err != nil {
        os.Exit(1)
    }
    writers := []io.Writer {
        logFile,
        os.Stdout,
    }
    fileAndStdoutWriter := io.MultiWriter(writers...)
    Logger = log.New(fileAndStdoutWriter, "[Debug] ", log.LstdFlags | log.Lshortfile)
}

func (this *Server) InitDB() {
    db, err := sql.Open("mysql", cConnectString)
    if err != nil {
        Logger.Panic("[InitDB] err = ", err)
    }
    this.mDB = db
    err = this.mDB.Ping()
    if err != nil {
        Logger.Panic("[InitDB] err = ", err)
    }
}

func (this *Server) Init() {
    this.InitLogger()
    Logger.Println("[Init Begin]")
    router := httprouter.New()
    this.mRouter = router
    this.InitRouter()
    Logger.Println("[Init End]")
}

func (this *Server) Start() {
    http.ListenAndServe(":80", this.mRouter)
}
   
func (this *Server) Stop() {
    this.mLogFile.Close()
    this.mDB.Close()
}
func main() {
    var server Server
    server.Init()
    server.Start()
    defer server.Stop()
}
