package elasticsearch

import (
	"bytes"
	"context"
	// "crypto/tls"
	"encoding/json"
	"fmt"
	// "net/http"
	"strings"

	"github.com/elastic/go-elasticsearch/v8"
)

type Document struct {
	Wname   string `json:"wname"`
	Data string `json:"content"`
}

func StartElastic(data string) {
	// Create a custom HTTP transport to skip certificate verification
	// transport := &http.Transport{
	// 	TLSClientConfig: &tls.Config{
	// 		InsecureSkipVerify: true, // Skip SSL certificate verification
	// 	},
	// }

	conf := elasticsearch.Config{
		Addresses: []string{
			"https://18.163.38.96:9200", // URL on which elastic search is running
		},
		Username:  "dev_account",
		Password:  "PMh4IqZnuQ2I34kbjXNceq8H908zm6",
		APIKey: "",
	}

	client, err := elasticsearch.NewClient(conf)
	if err != nil {
		panic(fmt.Sprintf("Error creating the Elasticsearch client: %s", err))
	}

	AddDocument(client, data)
	GetDocument(client)
}

func AddDocument(client *elasticsearch.Client, data string) {
	indexName := "fiction_novel"

	doc := Document{
		Wname: "bqgda",
		Data: data,
	}

	docJSON, err := json.Marshal(doc)
	if err != nil {
		panic(fmt.Sprintf("Error marshaling the document: %s", err))
	}

	var dataBody strings.Builder

	// First operation: Index operation metadata (add an index action)
	dataBody.WriteString(`{ "index": { "_index": "` + indexName + `", "_id": "2" } }` + "\n")
	// First document: Actual data to index
	// dataBody.WriteString(`{ "title": "bqgda", "content": "` + string(docJSON) + `" }` + "\n")
	dataBody.WriteString(string(docJSON) + "\n")

	bulkBodyString := dataBody.String()
	if !strings.HasSuffix(bulkBodyString, "\n") {
		bulkBodyString += "\n"
	}

	buf := bytes.NewBufferString(bulkBodyString)

	ingestResult, err := client.Bulk(
		bytes.NewReader(buf.Bytes()),
		client.Bulk.WithIndex("fiction_novel"),
		client.Bulk.WithPipeline("ent-search-generic-ingestion"),
	  )
	  
	  fmt.Println(ingestResult, err)
}

func GetDocument(client *elasticsearch.Client) {
	searchResp, err := client.Search(
		client.Search.WithContext(context.Background()),
		client.Search.WithIndex("fiction_novel"),
		client.Search.WithQuery("title"),
		client.Search.WithTrackTotalHits(true),
		client.Search.WithPretty(),
	  )
	  
	  fmt.Println(searchResp, err)
}