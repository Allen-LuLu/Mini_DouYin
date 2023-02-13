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
    3:BaseResp base
}

struct CheckUserReq{
    1:string username
    2:string password
}

struct CheckUserResp{
    1:i64 userID
}

struct User{
    1:i64 id
    2:string name
    3:i64 follow_count
    4:i64 follower_count
}

struct UserInfoReq{
    1:list<i64> userIDS
}

struct UserInfoResp{
    1:list<User> users
    2:BaseResp base
}

service UserService{
    RegisterResp Register(1:RegisterReq req)
    CheckUserResp CheckUser(1:CheckUserReq req)
    UserInfoResp UserInfo(1:UserInfoReq req)
}