package main

import(
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type RSS struct{
	XMLName xml.Name
	Channel *Channel
}

type Channel struct{
	Title string
	ItemList []Item
}

type Item struct{
	Title string
	Link string
	Traffic string
	NewsItems []News
}

type News struct{
	Headline string
	HeadlineLink string
}

func main() {
	readGoogleTrends
}

func readGoogleTrends() {
	getGoogleTrends
}

func getGoogleTrends() {

}
