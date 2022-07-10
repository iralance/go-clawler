package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCarList(t *testing.T) {
	contents, err := ioutil.ReadFile("carlist_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCarList(contents)

	const resultSize = 230
	expectedUrls := []string{
		"https://newcar.xcar.com.cn/car/0-0-0-0-1-0-0-0-0-0-0-0/",
		"https://newcar.xcar.com.cn/car/0-0-0-0-56-0-0-0-0-0-0-0/",
		"https://newcar.xcar.com.cn/car/0-0-0-0-78-0-0-0-0-0-0-0/",
	}

	//for _, r := range result.Requests {
	//	log.Println(r.Url)
	//}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests; but had %d", resultSize, len(result.Requests))
	}

	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s; but was %s", i, url, result.Requests[i].Url)
		}
	}

	if len(result.Items) != resultSize {
		t.Errorf("result should have %d items; but had %d", resultSize, len(result.Items))
	}
}
