package main

import (
	"io/ioutil"
	"reflect"
	"testing"
)

func TestGetMp3List(t *testing.T) {
	buf, err := ioutil.ReadFile("testdata/rss_sample.txt")
	if err != nil {
		t.Errorf("rss sample file read failed")
	}

	got := GetMp3List(string(buf), 2)
	want := []string{"https://av.voanews.com/clips/VEN/2022/04/21/20220421-220500-VEN060-program.mp3", "https://av.voanews.com/clips/VEN/2022/04/20/20220420-220500-VEN060-program.mp3"}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
