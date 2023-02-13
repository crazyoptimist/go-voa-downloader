package main

import (
	"reflect"
	"testing"
)

func TestGetMp3List(t *testing.T) {
	// buf, err := ioutil.ReadFile("testdata/rss_sample.txt")
	// if err != nil {
	// 	t.Errorf("rss sample file read failed")
	// }

	got := GetMp3List()
	want := []string{""}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
