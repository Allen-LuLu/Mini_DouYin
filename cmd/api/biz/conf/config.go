package conf

import (
	"Mini_DouYin/cmd/api/biz/model"
	"gopkg.in/ini.v1"
)

var Cfg =  new(model.Cfg)
const confPath = "../../common/conf/config.ini"

func Init(){
	err := ini.MapTo(Cfg,confPath)
	if err != nil {
		panic(any(err))
	}
}
