namespace go publish
struct Video{//to define videos
1: required i64 id;
2: required User author;
3: required string play_url;
4: required string cover_url;
5: required i64 favorite_count;
6: required i64 comment_count;
7: required bool is_favorite;
8: required string title;
}
struct User{
1: required i64 id;
2: required string name;
3: optional i64 follow_count;
4: optional i64 fillower_count;
5: required bool is_follow;
}