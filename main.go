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

	for i, element := range objList {
		log.Println(i)
		var test util.ApiTest
		bytes, err := json.Marshal(element)
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

}
