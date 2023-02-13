namespace go relation

struct BaseResp {
    1:i64 status_code
    2:string status_message
    3:i64 service_time
}

struct ActionRelationRequest {
    1:i64 user_id
    2:i64 follow_id
    3:i32 action_type
}

struct ActionRelationResponse {
    1:BaseResp base_resp
}

struct ListAllOperateRequest{
    1:i64 user_id
    2:i32 operate_type
}

struct ListAllOperateResponse {
    1:BaseResp base_resp
    2:list<i64> user_ids
}


service RelationService {
    ActionRelationResponse ActionRelation(1:ActionRelationRequest req)
    ListAllOperateResponse ListAllOperate(1:ListAllOperateRequest req)
}

