package server

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/Sirupsen/logrus"
	"github.com/tony24681379/inu-golang-backend/esclient"
	elastic "gopkg.in/olivere/elastic.v5"
)

type amiast struct {
	Vtt       []string
	Customer0 []string
	Agent0    []string
}

//Server init
func Server2(elasticSearchIP, elasticSearchPort string) {
	ec, err := esclient.CreateESClient(elasticSearchIP, elasticSearchPort)
	if err != nil {
		logrus.Error("Can not connect to ElasticSearch")
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

//Server init
func Server(elasticSearchIP, elasticSearchPort string) {
	client, err := elastic.NewClient(elastic.SetSniff(false), elastic.SetURL("http://"+elasticSearchIP+":"+elasticSearchPort))
	if err != nil {
		fmt.Println(err)
		return
	}
	/*	esversion, err := client.ElasticsearchVersion("http://" + elasticSearchIP + ":" + elasticSearchPort)
		if err != nil {
			panic(err)
		}
		fmt.Printf("Elasticsearch version %s", esversion)
	*/
	// Ping the Elasticsearch server to get e.g. the version number
	info, code, err := client.Ping("http://" + elasticSearchIP + ":" + elasticSearchPort).Do(context.Background())
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Elasticsearch returned with code %d and version %s\n", code, info.Version.Number)

	get1, err := client.Get().
		Index("logs-2016.05.11").
		Type("amiast").
		Id("B95F18EFB979D565AB84-A71D30C0210399FB-ED6A41971D1-A2EDE7FD00000CF31FC00AD7").
		Do(context.Background())
	if err != nil {
		fmt.Println(err)
	} else if get1.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
		o, err := get1.Source.MarshalJSON()
		if err != nil {
			logrus.Error("error to get base")
		}
		var vtt amiast
		json.Unmarshal(o, &vtt)
		fmt.Print(vtt)
	}
}
