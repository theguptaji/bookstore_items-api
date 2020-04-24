package items

import (
	"errors"

	"github.com/theguptaji/bookstore_items-api/clients/elasticsearch"
	"github.com/theguptaji/bookstore_utils-go/rest_errors"
)

const (
	IndexItems = "items"
)

func (i *Item) Save() rest_errors.RestErr {
	result, err := elasticsearch.Client.Index(IndexItems, i)
	if err != nil {
		return rest_errors.NewInternalServerError("error when trying to save index", errors.New("database error"))
	}
	i.Id = result.Id
	return nil
}
