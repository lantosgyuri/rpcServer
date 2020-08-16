package main

type Item struct {
	Title string
	Body  string
}

type RpcCrudStore interface {
	GetItemByName(title string, reply *Item) error
	AddItem(item Item, reply *Item) error
	EditItem(edit Item, reply *Item) error
	DeleteItem(item Item, reply *Item) error
}

type InMemoryCrudStore struct {
	database []Item
}

func NewInMemoryCrudStore() *InMemoryCrudStore {
	return &InMemoryCrudStore{[]Item{}}
}

func isSameItem(titleOne, titleTwo string) bool {
	return titleOne == titleTwo
}

func (i *InMemoryCrudStore) GetItemByName(title string, reply *Item) error {
	var _item Item
	for _, val := range i.database {
		if isSameItem(val.Title, title) {
			_item = val
		}
	}
	*reply = _item
	return nil
}

func (i *InMemoryCrudStore) AddItem(item Item, reply *Item) error {
	i.database = append(i.database, item)
	*reply = item
	return nil
}

func (i *InMemoryCrudStore) EditItem(editedItem Item, reply *Item) error {
	var changedItem Item
	for idx, val := range i.database {
		if isSameItem(val.Title, editedItem.Title) {
			i.database[idx] = editedItem
			changedItem = editedItem
		}
	}
	*reply = changedItem
	return nil
}

func (i *InMemoryCrudStore) DeleteItem(item Item, reply *Item) error {
	var del Item
	for idx, val := range i.database {
		if isSameItem(val.Title, item.Title) && val.Body == item.Body {
			i.database = append(i.database[:idx], i.database[idx+1:]...)
			del = item
			break
		}
	}
	*reply = del
	return nil
}
