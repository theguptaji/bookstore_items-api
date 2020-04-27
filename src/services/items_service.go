package services

import (
	"github.com/theguptaji/bookstore_items-api/src/domain/items"
	"github.com/theguptaji/bookstore_items-api/src/domain/queries"
	"github.com/theguptaji/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(item items.Item) (*items.Item, rest_errors.RestErr)
	Get(string) (*items.Item, rest_errors.RestErr)
	Search(queries.EsQuery)([]items.Item, rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, rest_errors.RestErr) {
	item := items.Item{Id: id}
	if err := item.Get(); err!=nil{
		return nil, err
	}
	return &item, nil
}

func(s *itemsService) Search(query queries.EsQuery)([]items.Item, rest_errors.RestErr){
	dao := items.Item{}
	return dao.Search(query)
}
