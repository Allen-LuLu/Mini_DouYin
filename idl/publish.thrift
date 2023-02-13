namespace go publish
include "basic_def.thrift"
struct ActionReq{
    //1: file data;文件上传需要弄清楚
    1: required string token;
    2: required string title;
    3: required byte data;// video uploaded
    }
struct ActionResp{
1: required i16 status;
2: optional string msg
}
struct ListReq{
1: required i64 user_id;
2: required string token;
}

struct ListResp{
1: required i32 status;
2: optional string status_msg;
3: optional list <basic_def.Video> video_list;

}

service Publish{
ActionResp Action(1:ActionReq req)
ListResp List(1:ListReq req)
}