package persist

import (
	"context"
	"encoding/json"
	"github.com/iralance/go-clawler/engine"
	"github.com/iralance/go-clawler/model"
	"github.com/olivere/elastic/v7"
	"testing"
)

func TestSaver(t *testing.T) {
	expected := engine.Item{
		Url:  "https://newcar.xcar.com.cn/151/",
		Type: "xcar",
		Id:   "151",
		Payload: model.Car{
			Name:         "安静的雪",
			Price:        0,
			ImageURL:     "",
			Size:         "",
			Fuel:         0,
			Transmission: "",
			Engine:       "",
			Displacement: 0,
			MaxSpeed:     0,
			Acceleration: 0,
		},
	}

	// TODO: Try to start up elastic search
	// here using docker go client.
	client, err := elastic.NewClient(
		elastic.SetSniff(false))

	if err != nil {
		panic(err)
	}

	const index = "dating_test"
	// Save expected item
	err = Save(client, index, expected)

	if err != nil {
		panic(err)
	}

	// Fetch saved item
	resp, err := client.Get().
		Index(index).
		Type(expected.Type).
		Id(expected.Id).
		Do(context.Background())

	if err != nil {
		panic(err)
	}

	t.Logf("%s", resp.Source)

	var actual engine.Item
	json.Unmarshal(resp.Source, &actual)

	actualProfile, _ := model.FromJsonObj(
		actual.Payload)
	actual.Payload = actualProfile

	// Verify result
	if actual != expected {
		t.Errorf("got %v; expected %v",
			actual, expected)
	}
}
