package conf

//Read config and pass to dao

import (
	"Mini_DouYin/cmd/publish/model"
	"encoding/json"
	"log"
	"os"
)

var Cfg = new(model.Cfg)

const confPath = "../../common/conf/conf.json"

func Init() {
	byteValue, err := os.ReadFile(confPath)
	if err != nil {
		log.Panicln(err)
	}

	json.Unmarshal([]byte(byteValue), &Cfg)

	log.Println(Cfg)

}
