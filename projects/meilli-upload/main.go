package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/meilisearch/meilisearch-go"
)

func main() {

	var client = meilisearch.NewClient(meilisearch.Config{
		Host: "http://localhost:7700",
	})

	_, err := client.Indexes().Create(meilisearch.CreateIndexRequest{
		UID: "documents",
	})

	if err != nil {
		panic(err)
	}

	jsonFile, _ := os.Open("csvjson.json")
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var documents []map[string]interface{}
	json.Unmarshal(byteValue, &documents)

	fmt.Println(len(documents))
	i := 0
	for i = 0; i < 5000-1; i++ {
		var temp []map[string]interface{}
		temp = append(temp, documents[i])
		_, errDB := client.Documents("documents").AddOrUpdate(temp)
		if errDB != nil {
			panic(errDB)
		}
		temp = nil
	}

}
