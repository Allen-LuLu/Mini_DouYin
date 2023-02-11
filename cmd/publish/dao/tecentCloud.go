package dao

import (
	"Mini_DouYin/cmd/publish/conf"
	"context"
	"encoding/base64"
	"github.com/tencentyun/cos-go-sdk-v5"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var tenClient *cos.Client

func InitTencentCloud() {
	tenCfg := conf.Cfg.TencentCfg
	//"https://examplebucket-1250000000.cos.COS_REGION.myqcloud.com"
	tencentOSSURL := "https://" + tenCfg.BucketName + "-" + tenCfg.BucketName + ".cos." + tenCfg.CosRegion + ".myqcloud.com"
	//"https://cos.COS_REGION.myqcloud.com"
	tencentOSSURLGet := "https://cos." + tenCfg.CosRegion + ".mycloud.com"

	u, _ := url.Parse(tencentOSSURL)
	su, _ := url.Parse(tencentOSSURLGet) //For get service
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	tenClient = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  os.Getenv(tenCfg.SecretId),
			SecretKey: os.Getenv(tenCfg.SecretKey),
		},
	})

}

// UploadObj Upload video by Tencent interface and return Video URL and Cover URL.
func UploadObj(title string) (string, string) {
	vidKey, covKey := GenKey(title)
	log.Println("Key generated:", vidKey)
	filePath := "../tmp/" + title + ".mp4"
	uploadRes, _, err := tenClient.Object.Upload(
		context.Background(), vidKey, filePath, nil,
	)
	if err != nil {
		panic(err)
	}

	//to make URL
	tenCfg := conf.Cfg.TencentCfg
	baseURL := "https://" + tenCfg.BucketName + "-" + tenCfg.BucketName + ".cos." + tenCfg.CosRegion + ".myqcloud.com"
	playURL, covURL := baseURL+uploadRes.Key, baseURL+covKey
	log.Printf("Video upload success:\nTitle:\t%v\nPlayURL:\t%v\nCoverURL:\t%v", title, playURL, covURL)
	return baseURL, covURL
}

// GenKey generate key for COS:base64(title+time.now(Unix); return mp4 and jpg key
func GenKey(title string) (string, string) {
	return base64.StdEncoding.EncodeToString([]byte(title+strconv.FormatInt(time.Now().Unix(), 10))) + ".mp4", base64.StdEncoding.EncodeToString([]byte(title+strconv.FormatInt(time.Now().Unix(), 10))) + ".jpg"
}

// IfURLExisted to find if the key existed
func IfURLExisted(key string) bool {
	ok, err := tenClient.Object.IsExist(context.Background(), key)
	if err == nil && ok {
		log.Printf("object exists\n")
		return true
	} else if err != nil {
		log.Printf("head object failed: %v\n", err)
	} else {
		log.Printf("object does not exist\n")
	}
	return false
}
