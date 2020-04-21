package controllers

import (
	"fmt"
	"net/http"

	"github.com/theguptaji/bookstore_items-api/domain/items"
	"github.com/theguptaji/bookstore_items-api/services"
	"github.com/theguptaji/bookstore_oauth-go/oauth"
)

var (
	ItemsController itemsControllerInterface = &itemsController{}
)

type itemsControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Get(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (i *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if oauth.AuthenticateRequest(r) != nil {
		// TODO: Return error JSON to the user
	}

	item := items.Item{
		Seller: oauth.GetCallerId(r),
	}

	result, err := services.ItemsService.Create(item)
	if err != nil {
		// TODO: Return error JSON to user
	}
	fmt.Println(result)
	// TODO: Return created item as JSON with HTTP status 201 - created
}

func (i *itemsController) Get(w http.ResponseWriter, r *http.Request) {

}
