package login

import (
	"due-learn/utils"
	"fmt"

	"github.com/dobyte/due/component/http/v2"
	"github.com/dobyte/due/registry/etcd/v2"
	"github.com/dobyte/due/v2"
	"github.com/dobyte/due/v2/codes"
	"github.com/dobyte/due/v2/log"
	"github.com/dobyte/due/v2/utils/xtime"
)

func Init() {
	utils.SetEtcConfigurator("./etc/login")

	// 创建容器
	container := due.NewContainer()

	// 创建HTTP组件
	component := http.NewServer(
		http.WithRegistry(etcd.NewRegistry()),
	)
	// 初始化应用
	initApp(component.Proxy())
	// 添加网格组件
	container.Add(component)
	// 启动容器
	container.Serve()
}

// 初始化应用
func initApp(proxy *http.Proxy) {
	// 路由器
	router := proxy.Router()
	// 注册路由
	router.Get("/greet", greetHandler)
}

// 请求
type greetReq struct {
	Message string `json:"message"`
}

// 响应
type greetRes struct {
	Message string `json:"message"`
}

// 测试接口
// @Summary 测试接口
// @Tags 测试
// @Schemes
// @Accept json
// @Produce json
// @Param request body greetReq true "请求参数"
// @Response 200 {object} http.Resp{Data=greetRes} "响应参数"
// @Router /greet [get]
func greetHandler(ctx http.Context) error {
	req := &greetReq{}

	if err := ctx.Bind().JSON(req); err != nil {
		return ctx.Failure(codes.InvalidArgument)
	}

	log.Info(req.Message)

	return ctx.Success(&greetRes{
		Message: fmt.Sprintf("I'm login server, and the current time is: %s", xtime.Now().Format(xtime.DateTime)),
	})
}
