package server

import (
	"encoding/json"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/tony24681379/inu-golang-backend/esclient"
)

type amiast struct {
	Vtt       []string
	Customer0 []string
	Agent0    []string
}

//Server init
func Server(elasticSearchIP, elasticSearchPort string) {
	ec, err := esclient.CreateESClient(elasticSearchIP, elasticSearchPort)
	if err != nil {
		print(err)
	}
	// Search Using Raw json String
	// searchJSON := `{
	// 	"from": 0,
	// 	"size": 200,
	// 	"query" :{
	// 		"match": {
	// 			"vtt": "不能"
	// 		}
	// 	},
	// 	"_source":[
	// 		"vtt",
	// 		"agent*",
	// 		"customer*"
	// 	]
	// }`
	/*	out, err := ec.Search("logs-2016.05.20", "amiast", nil, searchJSON)
		if err != nil {
			fmt.Printf("error to search")
		}
		if len(out.Hits.Hits) > 0 {
			o, err := out.Hits.Hits[0].Source.MarshalJSON()
			if err != nil {
				fmt.Printf("error to get hit")
			}
			fmt.Printf("%v", string(o))
		}
	*/
	base, err := ec.Get("logs-2016.05.11", "amiast", "B95F18EFB979D565AB84-A71D30C0210399FB-ED6A41971D1-A2EDE7FD00000CF31FC00AD7", nil)
	if err != nil {
		fmt.Printf("error to get index")
	}
	if base.Found {
		o, err := base.Source.MarshalJSON()
		if err != nil {
			logrus.Error("error to get base")
		}
		var vtt amiast
		json.Unmarshal(o, &vtt)
		fmt.Print(vtt)
	}

	//	ch := make(chan int)
	//	<-ch
}
