package mw

import (
	"Mini_DouYin/cmd/api/biz/dao"
	"Mini_DouYin/common/consts/errmsg"
	"context"
	"errors"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/go-redis/redis"
	"net/http"
	"time"

	"Mini_DouYin/cmd/api/biz/model/api"
	"Mini_DouYin/cmd/api/biz/rpc"
	"Mini_DouYin/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

var JwtMiddleware *jwt.HertzJWTMiddleware
const SecretKey = "secret"
const IdentityKey = "userID"
const Expire = 3600

func InitJWT() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key:           []byte(SecretKey),
		TokenLookup:   "header: Authorization, query: token, cookie: jwt", // 该属性决定路由从前端请求的head中的什么字段提取token
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		IdentityKey:   IdentityKey,
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} { // 该函数用于jwt中间件鉴权时，会从token中提取相关信息，返回值将会传给校验函数Authorizator,Authorizator在下面定义
			claims := jwt.ExtractClaims(ctx, c)
			userid, _:= claims[IdentityKey].(float64) //从token中解析出userid,注意这里有个坑，我们记录的userid是int64的，但是jwt底层帮我们存储的是float64的，需要我们自己转类型
			return int64(userid)
		},
		PayloadFunc: func(data interface{}) jwt.MapClaims { // 登录验证成功后，会调用该函数，该函数负责将token保存的信息放到MapClaims中，后续jwt鉴权时调用IdentityHandler函数就能解析出相关信息
			// 该函数的传入参数由登录验证函数Authenticator的返回值提供，该函数在后面定义
			if v, ok := data.(int64); ok { // 将登录验证函数传出的userID保存到token中
				return jwt.MapClaims{
					IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) { // 登录验证函数，登录时会调用该函数，该函数的返回值将传入PayloadFunc函数
			var err error
			var req api.LoginReq
			if err = c.BindAndValidate(&req); err != nil { // 获取用户名和密码
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Username) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			resp,err := rpc.CheckUser(context.Background(), &user.CheckUserReq{ // 调用rpc服务，验证用户名和密码是否正确
				Username: req.Username,
				Password: req.Password,
			})
			if err != nil {
				return "", jwt.ErrFailedAuthentication
			}
			c.Set(IdentityKey,resp.UserID) // 记录当前登录的userID，便于返回时到redis中查询token
			return resp.UserID,nil
		},
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) { //登录成功后，会调用该函数，该函数负责响应客户端
			resp := new(api.LoginResp)

			id,_ := c.Get(IdentityKey)
			userid := id.(int64) //获取当前登录用户的userID

			t,err := dao.QueryToken(userid) // 到redis中查询是否存在token(注册时的token可能过期了)
			if err != nil{
				if errors.Is(err,redis.Nil){ // 如果查不到token,则将当前token存储到redis
					dao.StoreToken(userid,token,Expire*time.Second) //过期时间默认1h=3600s
					t = token
				}else{
					resp.StatusCode = consts.StatusInternalServerError
					resp.StatusMsg = errmsg.QUERYTOKEN
					c.JSON(http.StatusOK,resp)
					return
				}
			}

			token = t // 如果能查到，则用之前保存在redis中的token
			resp.UserID = userid
			resp.Token = token
			c.JSON(http.StatusOK, resp)
		},
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) { // 使用jwt中间件鉴权失败时，调用该函数
			c.JSON(http.StatusOK, utils.H{
				"status_code":    consts.StatusUnauthorized,
				"status_message": message,
			})
		},
		Authorizator: func(data interface{}, ctx context.Context, c *app.RequestContext) bool { //使用jwt中间件鉴权时，调用IdentityHandler后，调用该函数，IdentityHandler的返回值作为该函数的传入参数
			var req api.UserInfoReq
			var err error

			if err = c.BindAndValidate(&req); err != nil {
				return false
			}

			if v, ok := data.(int64); ok { // 将token中解析出来的userID与当前调用该服务的userID进行比较，如果不匹配说明用户A拿了用户B的token对路由进行访问，这可能是恶意访问，需要拒绝访问
				if v == req.UserID { //如果匹配则鉴权通过
					return true
				}
			}
			return false
		},
	})
}

func GenToken(userID int64)(string,time.Time,error){
	return JwtMiddleware.TokenGenerator(userID)
}
