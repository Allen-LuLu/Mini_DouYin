namespace go user

struct BaseResp{
    1:i64 code
    2:string errmsg
}

struct RegisterReq{
    1:string username
    2:string password
}

struct RegisterResp{
    1:i64 userID
    2:string token
    3:BaseResp base
}

struct LoginReq{
    1:string username
    2:string password
}

struct LoginResp{
    1:i64 userID
    2:string token
    3:BaseResp base
}

struct User{
    1:i64 id
    2:string name
    3:i64 follow_count
    4:i64 follower_count
}

struct UserInfoReq{
    1:i64 userID
    2:string token
}

struct UserInfoResp{
    1:User user
    2:BaseResp base
}

service UserService{
    RegisterResp Register(1:RegisterReq req)  // 用户注册接口
    LoginResp Login(1:LoginReq req)  //用户登录接口
    UserInfoResp UserInfo(1:UserInfoReq req) //获取用户信息接口
}