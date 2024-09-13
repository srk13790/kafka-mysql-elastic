package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/elastic/go-elasticsearch/v8/esapi"
	
)

type Document2 struct {
	Wname   string `json:"wname"`
	Data string `json:"content"`
}

func StartElastic2(data string) {
	client, err := elasticsearch.NewDefaultClient()
	if err != nil {
		panic(fmt.Sprintf("Error creating the Elasticsearch client: %s", err))
	}

	AddDocument2(client, data)
	GetDocument2(client)
}

func AddDocument2(client *elasticsearch.Client, data string) {
	doc := Document2{
		Wname: "bqgda",
		Data: data,
	}

	docJoson, err := json.Marshal(doc)
	if err != nil {
		panic(fmt.Sprintf("Error marshaling the document: %s", err))
	}

	req := esapi.IndexRequest{
		Index: "fiction_novel",
		DocumentID: "1",
		Body: strings.NewReader(string(docJoson)),
		Refresh: "true",
	}

	res, err := req.Do(context.Background(), client)
	if err != nil {
		panic(fmt.Sprintf("Error indexing the document: %s", err))
	}
	defer res.Body.Close()

	if res.IsError() {
		panic(fmt.Sprintf("Error indexing document: %s", res.String()))
	} else {
		fmt.Println("Document Successfully Indexed")
	}
}

func GetDocument2(client *elasticsearch.Client) {
	req:= esapi.GetRequest{
		Index: "fiction_novel",
		DocumentID: "1",
	}

	res, err := req.Do(context.Background(), client)
	if err != nil {
		panic(fmt.Sprintf("Error indexing the document: %s", err))
	}
	defer res.Body.Close()

	if res.IsError() {
		panic(fmt.Sprintf("Error indexing document: %s", res.String()))
	} else {
		buf := new(bytes.Buffer)
		buf.ReadFrom(res.Body)
		fmt.Printf("Retrieved document: %s\n", buf.String())
	}
}