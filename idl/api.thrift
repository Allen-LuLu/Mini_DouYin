namespace go api

struct RegisterReq{
    1:string username (api.query="username")
    2:string password (api.query="password")
}

struct RegisterResp{
    1:i64 status_code
    2:string status_msg
    3:i64 user_id
    4:string token
}

struct LoginReq{
    1:string username (api.query="username")
    2:string password (api.query="password")
}

struct LoginResp{
    1:i64 status_code
    2:string status_msg
    3:i64 user_id
    4:string token
}

struct User{
    1:i64 id
    2:string name
    3:i64 follow_count
    4:i64 follower_count
    5:bool is_follow
}

struct UserInfoReq{
    1:i64 user_id (api.query="user_id")
    2:string token (api.query="token")
}

struct UserInfoResp{
    1:i64 status_code
    2:string status_msg
    3:User user
}


service UserApiService{
    RegisterResp UserRegister(1:RegisterReq req) (api.post="/douyin/user/register/")
    LoginResp UserLogin(1:LoginReq req) (api.post="/douyin/user/login/")
    UserInfoResp GetUserInfo(1:UserInfoReq req) (api.get="/douyin/user/")
}