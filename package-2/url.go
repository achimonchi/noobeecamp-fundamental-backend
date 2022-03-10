package main

import (
	"fmt"
	"net/url"
	"strings"
)

func main() {
	urlstring := "https://noobee.id:443/search?title=&categories=1-2-3&types=1-3-2&tags=1-3-2&kind=Artikel"

	dataUrl, err := url.Parse(urlstring)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Printf("Host :%s\nPath :%s\nPort :%s\n", dataUrl.Host, dataUrl.Path, dataUrl.Port())
	fmt.Printf("Raw Query :%s\n", dataUrl.RawQuery)
	fmt.Printf("Query :%v\n", dataUrl.Query())

	qs := dataUrl.Query()
	cat := qs.Get("categories")
	fmt.Println(strings.Split(cat, "-"))
}
