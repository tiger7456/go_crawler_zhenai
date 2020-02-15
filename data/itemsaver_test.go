package data

import (
	"context"
	"crawler/engine"
	"crawler/model"
	"encoding/json"
	"testing"

	"gopkg.in/olivere/elastic.v5"
)

func TestSave(t *testing.T) {
	profile := engine.Item{
		Url:  "test",
		Id:   "test",
		Type: "zhenai",
		Payload: model.Profile{
			Name:       "test",
			Gender:     "test",
			Age:        0,
			Height:     0,
			Income:     "1000000",
			Marriage:   "yes",
			Education:  "test",
			Occupation: "test",
			Hukou:      "test",
			Xinzuo:     "test",
		},
	}

	// save profile
	err := save(profile)
	if err != nil {
		panic(err)
	}
	// TODO: 这里要启动elasticsearch, 可以采用docker go client 解决
	client, err := elastic.NewClient(
		elastic.SetURL("http://myubuntu:9200"),
		elastic.SetSniff(false))
	resp, err := client.Get().Index("dating_profile").Type(profile.Type).Id(profile.Id).Do(context.Background())
	if err != nil {
		panic(err)
	}
	t.Logf("%s", resp.Source)
	var actual engine.Item
	err = json.Unmarshal([]byte(*resp.Source), &actual)
	if err != nil {
		panic(err)
	}

	actualProfile, _ := model.FromJsonObj(actual.Payload)
	actual.Payload = actualProfile

	// verify
	if actual != profile {
		t.Errorf("got %v; expected %v", actual, profile)
	}
}
