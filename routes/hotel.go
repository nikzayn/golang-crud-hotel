package routes

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/nikzayn/golang-crud-hotel/config"
	"github.com/nikzayn/golang-crud-hotel/models"

	"github.com/nikzayn/golang-crud-hotel/routes/response"
)

func listHotel(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	appCtx := ctx.Value(AppContextKey).(*config.AppContext)
	appCtx.Log.Info("A request came to get the list of hotels")

	hotels, err := models.GetHotels(appCtx)
	if err != nil {
		//error while getting the list of hotels from db
		appCtx.Log.Info("error while getting the list of hotels from the database", err)
		response.WriteError(res, response.Error{Err: "error while getting the list of hotels from database"}, http.StatusInternalServerError)
		return
	}

	response.Write(res, response.Message{Message: "Fetched the list of hotels", Data: hotels})
}

func createHotel(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	appCtx := ctx.Value(AppContextKey).(*config.AppContext)
	appCtx.Log.Info("A request came to create the an hotel")

	h := &models.Hotel{}
	err := json.NewDecoder(req.Body).Decode(h)
	if err != nil {
		//error while parsing the request
		appCtx.Log.Info("error while parsing the request", err)
		response.WriteError(res, response.Error{Err: "error while parsing the request"}, http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	err = h.Create(appCtx)
	if err != nil {
		//error while saving the hotel in the db
		appCtx.Log.Info("error while saving the hotel in the db", err)
		response.WriteError(res, response.Error{Err: "error while saving the hotel in the database"}, http.StatusInternalServerError)
		return
	}

	response.Write(res, response.Message{Message: "Successfully created the hotel", Data: h})
}

func updateHotel(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	appCtx := ctx.Value(AppContextKey).(*config.AppContext)
	appCtx.Log.Info("A request came to update the hotel")

	h := &models.Hotel{}
	err := json.NewDecoder(req.Body).Decode(h)
	if err != nil {
		//error while parsing the request
		appCtx.Log.Info("error while parsing the request", err)
		response.WriteError(res, response.Error{Err: "error while parsing the request"}, http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	err = h.Update(appCtx)
	if err != nil {
		//error while updating the hotel in the db
		appCtx.Log.Info("error while updating the hotel in the db", err)
		response.WriteError(res, response.Error{Err: "error while updating the hotel in the database"}, http.StatusInternalServerError)
		return
	}

	response.Write(res, response.Message{Message: "Successfully updated the hotel", Data: h})
}

func deleteHotel(ctx context.Context, res http.ResponseWriter, req *http.Request) {
	appCtx := ctx.Value(AppContextKey).(*config.AppContext)
	appCtx.Log.Info("A request came to delete the hotel")

	h := &models.Hotel{}
	err := json.NewDecoder(req.Body).Decode(h)
	if err != nil {
		//error while parsing the request
		appCtx.Log.Info("error while parsing the request", err)
		response.WriteError(res, response.Error{Err: "error while parsing the request"}, http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	err = h.Delete(appCtx)
	if err != nil {
		//error while deleting the hotel in the db
		appCtx.Log.Info("error while saving the deleting the db", err)
		response.WriteError(res, response.Error{Err: "error while deleting the hotel in the database"}, http.StatusInternalServerError)
		return
	}

	response.Write(res, response.Message{Message: "Successfully deleted the hotel", Data: h})
}

func init() {
	AddRoutes(
		Route{
			Version:     "v1",
			HandlerFunc: listHotel,
			Pattern:     "/hotel",
		},
		Route{
			Version:     "v1",
			HandlerFunc: createHotel,
			Pattern:     "/hotel/create",
		},
		Route{
			Version:     "v1",
			HandlerFunc: updateHotel,
			Pattern:     "/hotel/update",
		},
		Route{
			Version:     "v1",
			HandlerFunc: deleteHotel,
			Pattern:     "/hotel/delete",
		},
	)
}
