package esclient

import (
	"log"

	elastigo "github.com/mattbaird/elastigo/lib"
)

//CreateESClient create ElasticSearch client
func CreateESClient(esAddress, esPort string) (*elastigo.Conn, error) {
	ec := elastigo.NewConn()
	ec.Domain = esAddress
	ec.Port = esPort
	ec.RequestTracer = func(method, url, body string) {
		log.Printf("Requesting %s %s", method, url)
		log.Printf("Request body: %s", body)
	}
	_, err := ec.Health()
	if err != nil {
		return nil, err
	}

	return ec, nil
}
