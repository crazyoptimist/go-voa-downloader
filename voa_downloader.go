package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

const RSS_URL string = "https://www.voanews.com/podcast/?zoneId=5082"

func GetRssFeed() string {
	resp, err := http.Get(RSS_URL)
	if err != nil {
		log.Fatalln(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	sb := string(body)

	return sb

}

func GetMp3List(rssFeed string, numberOfFiles int) (mp3List []string) {
	r := regexp.MustCompile(`(url)\=\"https:\/\/(www.)?(.*?)\.(mp3)`)
	matches := r.FindAllString(rssFeed, -1)
	for i := 0; i < numberOfFiles; i++ {
		mp3Url := matches[i][5:]
		mp3List = append(mp3List, mp3Url)
	}

	return
}

func main() {
	rss := GetRssFeed()
	fmt.Println(rss)
}
