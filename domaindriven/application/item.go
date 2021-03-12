package application

import "magical.dev/tesing/basic"

type ItemApplicationService struct{
	itemList []*basic.Item
}

func (a *ItemApplicationService) CreateItem() {
	item := basic.NewItem(100)
	a.itemList = append(a.itemList, item)
}

func (a *ItemApplicationService) GetAllItems() []*basic.Item {
	return a.itemList
}
