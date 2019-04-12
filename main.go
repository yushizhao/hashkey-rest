package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/yushizhao/hashkey-rest/util"
)

func main() {
	var TestPath = flag.String("test", "test.json", "the file")
	var configPath = flag.String("path", "config.json", "the file")

	flag.Parse()

	err := util.Init(configPath)
	if err != nil {
		log.Fatal(err)
	}

	var rawFile interface{}
	// Open our jsonFile
	jsonFile, err := os.Open(*TestPath)
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Successfully Opened %s\n", *TestPath)
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValue, &rawFile)
	objList := rawFile.([]interface{})

	objListLen := len(objList)
	for i := 0; i < objListLen; i++ {
		log.Println(i)

		if util.Conf.Debug {
			var debugInput string
			fmt.Scanln(&debugInput)
			switch debugInput {
			case "back":
				if i == 0 {
					i = -1
				} else {
					i = i - 2
				}
				continue

			case "skip":
				continue

			case "end":
				i = objListLen // -1
				continue

			case "view":
				log.Println(objList[i])
				i = i - 1
				continue

			case "check":
				log.Println("orderLocalID:")
				fmt.Scanln(&debugInput)
				err := util.CheckOrder(debugInput)
				if err != nil {
					log.Println(err)
				}
				i = i - 1
				continue
			}

		}

		var test util.ApiTest
		bytes, err := json.Marshal(objList[i])
		if err != nil {
			log.Println(err)
			continue
		}
		err = json.Unmarshal(bytes, &test)
		if err != nil {
			log.Println(err)
			continue
		}
		_, err = test.Send()
		if err != nil {
			log.Println(err)
			continue
		}
	}

	log.Println("Cancel All")
	if util.Conf.Debug {
		var debugInput string
		fmt.Scanln(&debugInput)
	}
	util.CancelAll()

}
