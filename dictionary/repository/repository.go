package repository

import (
	"dictionary/model"
	"errors"
	"fmt"
	"github.com/arriqaaq/flashdb"
	"log"
	"time"
)

type Repository struct {
	Database  map[string]model.Entry
	Database2 *flashdb.FlashDB
}

func (d *Repository) Init() {
	config := &flashdb.Config{Path: "/db", EvictionInterval: 10}
	db, err := flashdb.New(config)
	if err != nil {
		log.Fatal(err)
	}
	//defer db.Close()
	d.Database2 = db
}

func (d *Repository) Add(word string, definition string) {
	//d.Database[word] = model.Entry{Value: definition, Date: time.Now()}
	err := d.Database2.Update(func(tx *flashdb.Tx) error {
		err := tx.Set(word, definition)
		return err
	})
	if err != nil {
		panic(err)
	}
}

func (d *Repository) Get(word string) (model.Entry, error) {
	var value string
	err := d.Database2.View(func(tx *flashdb.Tx) error {
		data, err := tx.Get(word)
		if data != "" {
			value = data
		}
		return err
	})
	if err != nil {
		return model.Entry{}, errors.New("The key " + word + " is not in our dictionary")
	}
	return model.Entry{
		Value: value,
		Date:  time.Now(),
	}, nil
	/*value, contains := d.Database[word]
	if contains {
		return value, nil
	} else {
		return model.Entry{}, errors.New("The key " + word + " is not in our dictionary")
	}*/
}

func (d *Repository) Remove(word string) {
	//delete(d.Database, word)
	err := d.Database2.Update(func(tx *flashdb.Tx) error {
		err := tx.Delete(word)
		return err
	})
	if err != nil {
		panic(err)
	}
	fmt.Println("Removed value for key", word, "successfully!")
}

func (d *Repository) List() {
	entries := make(map[string]string, 0)
	d.Database2.View(func(tx *flashdb.Tx) error {
		keys := tx.HKeys("")
		for _, key := range keys {
			entries[key], _ = tx.Get(key)
		}
		return nil
	})
	fmt.Println("Word	| Definition			| Date")
	for key, element := range entries {
		fmt.Println(key, "	| ", element, "		| "+time.Now().Format(time.DateTime))
	}
}

func (d *Repository) Size() int {
	return len(d.Database)
}
