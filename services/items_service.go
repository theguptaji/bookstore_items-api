package services

import (
	"net/http"

	"github.com/theguptaji/bookstore_items-api/domain/items"
	"github.com/theguptaji/bookstore_utils-go/rest_errors"
)

var (
	ItemsService itemsServiceInterface = &itemsService{}
)

type itemsServiceInterface interface {
	Create(item items.Item) (*items.Item, *rest_errors.RestErr)
	Get(string) (*items.Item, *rest_errors.RestErr)
}

type itemsService struct{}

func (s *itemsService) Create(item items.Item) (*items.Item, *rest_errors.RestErr) {
	if err := item.Save(); err != nil {
		return nil, err
	}
	return &item, nil
}

func (s *itemsService) Get(id string) (*items.Item, *rest_errors.RestErr) {
	return nil, rest_errors.NewRestError("implement me senpai!!", http.StatusNotImplemented, "not implemented", nil)
}
