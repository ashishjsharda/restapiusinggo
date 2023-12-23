package model

import "errors"

var items []Item

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

func GetAllItems() []Item {
	return items
}

func AddItem(item Item) {
	items = append(items, item)
}

func UpdateItem(id string, newItem Item) (Item, error) {
	for i, item := range items {
		if item.ID == id {
			items[i] = newItem
			return newItem, nil
		}
	}
	return Item{}, errors.New("item not found")
}

func DeleteItem(id string) error {
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			return nil
		}
	}
	return errors.New("item not found")
}
