package data

import (
	"context"
	"crawler/engine"
	"errors"
	"log"

	"gopkg.in/olivere/elastic.v5"
)

func ItemSaver() (chan engine.Item, error) {
	client, err := elastic.NewClient(
		elastic.SetURL("http://myubuntu:9200"),
		elastic.SetSniff(false))
	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item  #%d: %v", itemCount, item)
			itemCount++

			err := save(client, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()
	return out, nil
}

func save(client *elastic.Client, item engine.Item) (err error) {
	if item.Type == "" {
		return errors.New("must supply Type")
	}
	indexService := client.Index().
		Index("dating_profile").
		Type(item.Type).
		BodyJson(item)
	if item.Id != "" {
		indexService.Id(item.Id)
	}
	_, err = indexService.
		Do(context.Background())

	if err != nil {
		return err
	}
	return nil
}
