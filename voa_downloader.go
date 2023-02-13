package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
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

func DownloadFile(fileUrl string) {
	// Create a download directory if not existing
	downloadPath := "downloads"
	_ = os.Mkdir(downloadPath, os.ModePerm)

	// Build fileName from fileUrl
	fileURL, err := url.Parse(fileUrl)
	if err != nil {
		log.Fatal(err)
	}
	path := fileURL.Path
	segments := strings.Split(path, "/")
	fileName := downloadPath + "/" + segments[len(segments)-1]

	// Create blank file
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatal(err)
	}
	client := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	// Put content on file
	resp, err := client.Get(fileUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	size, err := io.Copy(file, resp.Body)

	defer file.Close()

	fmt.Printf("Downloaded a file %s with size %d \n", fileName, size)
}

func main() {
	if len(os.Args) != 2 {
		log.Println("Usage: app <number of files to download>")
		return
	}
	numberOfFiles, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatalln("Entered number was not valid")
	}

	if err := os.RemoveAll("./downloads"); err != nil {
		log.Fatal(err)
	}

	rss := GetRssFeed()
	mp3List := GetMp3List(rss, numberOfFiles)

	var wg sync.WaitGroup
	for _, mp3Url := range mp3List {
		wg.Add(1)
		go func(url string) {
			DownloadFile(url)
			wg.Done()
		}(mp3Url)
	}
	wg.Wait()
}
