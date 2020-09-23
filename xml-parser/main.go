package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

type siteMapIndex struct {
	Locations []Locations `xml:"sitemap"`
}

type Locations struct {
	Loc string `xml:"loc"`
}

func (l Locations) String() string {
	return fmt.Sprintf(l.Loc)
}

func main() {

	resp, _ := http.Get("https://www.washingtonpost.com/news-sitemaps/index.xml")
	bytes, _ := ioutil.ReadAll(resp.Body)
	resp.Body.Close()

	var siteMap siteMapIndex
	xml.Unmarshal(bytes, &siteMap)
	fmt.Println(siteMap.Locations)

}
