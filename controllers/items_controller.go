package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/theguptaji/bookstore_items-api/domain/queries"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/theguptaji/bookstore_utils-go/rest_errors"

	"github.com/theguptaji/bookstore_items-api/utils/http_utils"

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
	Search(http.ResponseWriter, *http.Request)
}

type itemsController struct{}

func (i *itemsController) Create(w http.ResponseWriter, r *http.Request) {
	if err := oauth.AuthenticateRequest(r); err != nil {
		// TODO : Fix the respond error function, sending no response
		http_utils.RespondError(w, err)
		return
	}

	sellerId := oauth.GetCallerId(r)
	if sellerId == 0 {
		respErr := rest_errors.NewUnauthorizedError("invalid access token")
		http_utils.RespondError(w, respErr)
		return
	}

	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		respErr := rest_errors.NewBadRequestError("invalid request body")
		http_utils.RespondError(w, respErr)
		return
	}
	defer r.Body.Close()

	var itemRequest items.Item
	if err := json.Unmarshal(requestBody, &itemRequest); err != nil {
		respErr := rest_errors.NewBadRequestError("invalid item json body")
		http_utils.RespondError(w, respErr)
		return
	}

	itemRequest.Seller = sellerId

	result, createErr := services.ItemsService.Create(itemRequest)
	if createErr != nil {
		// TODO : Fix the respond error function, sending no response
		http_utils.RespondError(w, createErr)
		return
	}

	http_utils.RespondJson(w, http.StatusCreated, result)
}

func (i *itemsController) Get(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	itemId := strings.TrimSpace(vars["id"])

	item, err := services.ItemsService.Get(itemId)
	if err!=nil{
		http_utils.RespondError(w, err)
		return
	}
	http_utils.RespondJson(w, http.StatusOK, item)

}

func (i *itemsController) Search(w http.ResponseWriter, r *http.Request) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w,apiErr)
		return
	}
	defer r.Body.Close()

	var query queries.EsQuery
	if err := json.Unmarshal(bytes, &query); err!=nil{
		apiErr := rest_errors.NewBadRequestError("invalid json body")
		http_utils.RespondError(w,apiErr)
		return
	}
	items, searchErr:=  services.ItemsService.Search(query)
	if searchErr!=nil{
		http_utils.RespondError(w,searchErr)
		return
	}
	http_utils.RespondJson(w, http.StatusFound, items)
}
