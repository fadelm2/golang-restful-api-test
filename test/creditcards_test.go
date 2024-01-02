package test

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"golang-restful-api-technical-test/internal/entity"
	"golang-restful-api-technical-test/internal/model"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateContact(t *testing.T) {
	TestLogin(t)

	user := new(entity.User)
	err := db.Where("id = ?", "Fadel").First(user).Error
	assert.Nil(t, err)

	requestBody := model.CreateCreditcardRequest{
		Type:    "Visa",
		Name:    "fadel card",
		Number:  "1223232",
		Expired: "12-12-2030",
		Cvv:     "122",
		UserId:  user.ID,
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/creditcards", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CreditcardResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.Type, responseBody.Data.Type)
	assert.Equal(t, requestBody.Name, responseBody.Data.Name)
	assert.Equal(t, requestBody.Number, responseBody.Data.Number)
	assert.Equal(t, requestBody.Expired, responseBody.Data.Expired)
	assert.Equal(t, requestBody.Cvv, responseBody.Data.Cvv)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
}

func TestCreateContactFailed(t *testing.T) {
	TestLogin(t)

	user := new(entity.User)
	err := db.Where("id = ?", "Fadel").First(user).Error
	assert.Nil(t, err)

	requestBody := model.CreateCreditcardRequest{
		Type:   "",
		Name:   "",
		Number: "",
		Cvv:    "",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPost, "/api/creditcards", strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CreditcardResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)
	assert.NotNil(t, responseBody.Errors)

}
func TestGetConnect(t *testing.T) {
	TestCreateContact(t)

	user := new(entity.User)
	err := db.Where("id = ?", "Fadel").First(user).Error
	assert.Nil(t, err)

	creditcard := new(entity.Creditcard)
	err = db.Where("user_id = ?", user.ID).First(creditcard).Error
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodGet, "/api/creditcards/"+creditcard.ID, nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CreditcardResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, creditcard.ID, responseBody.Data.ID)
	assert.Equal(t, creditcard.Type, responseBody.Data.Type)
	assert.Equal(t, creditcard.Name, responseBody.Data.Name)
	assert.Equal(t, creditcard.Number, responseBody.Data.Number)
	assert.Equal(t, creditcard.Cvv, responseBody.Data.Cvv)
	assert.Equal(t, creditcard.CreatedAt, responseBody.Data.CreatedAt)
	assert.Equal(t, creditcard.UpdatedAt, responseBody.Data.UpdatedAt)
}

func TestGetContactFailed(t *testing.T) {
	TestCreateContact(t)

	user := new(entity.User)
	err := db.Where("id = ?", "Fadel").First(user).Error
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodGet, "/api/creditcards/"+uuid.NewString(), nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CreditcardResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestUpdateCreditcard(t *testing.T) {
	TestCreateContact(t)

	user := new(entity.User)
	err := db.Where("id = ?", "Fadel").First(user).Error
	assert.Nil(t, err)

	creditcards := new(entity.Creditcard)
	err = db.Where("user_id = ?", user.ID).First(creditcards).Error
	assert.Nil(t, err)

	requestBody := model.UpdateCreditcardRequest{
		Type:    "Master",
		Name:    "Udin",
		Number:  "12445566",
		Expired: "2013-12",
		Cvv:     "223",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/creditcards/"+creditcards.ID, strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CreditcardResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, requestBody.Type, responseBody.Data.Type)
	assert.Equal(t, requestBody.Name, responseBody.Data.Name)
	assert.Equal(t, requestBody.Number, responseBody.Data.Number)
	assert.Equal(t, requestBody.Cvv, responseBody.Data.Cvv)
	assert.NotNil(t, responseBody.Data.ID)
	assert.NotNil(t, responseBody.Data.CreatedAt)
	assert.NotNil(t, responseBody.Data.UpdatedAt)
}

func TestUpdateCreditcardFailed(t *testing.T) {
	TestCreateContact(t)

	user := new(entity.User)
	err := db.Where("id = ?", "Fadel").First(user).Error
	assert.Nil(t, err)

	creditcards := new(entity.Creditcard)
	err = db.Where("user_id = ?", user.ID).First(creditcards).Error
	assert.Nil(t, err)

	requestBody := model.UpdateCreditcardRequest{
		Name: "Udin",
	}

	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/creditcards/"+creditcards.ID, strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CreditcardResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusBadRequest, response.StatusCode)

}

func TestUpdateCreditcardNotFound(t *testing.T) {
	TestCreateContact(t)

	user := new(entity.User)
	err := db.Where("id = ?", "Fadel").First(user).Error
	assert.Nil(t, err)

	requestBody := model.UpdateCreditcardRequest{
		Name: "Udin",
	}
	bodyJson, err := json.Marshal(requestBody)
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodPut, "/api/contacts/"+uuid.NewString(), strings.NewReader(string(bodyJson)))
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[model.CreditcardResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestDeleteCreditcard(t *testing.T) {
	TestCreateContact(t)

	user := new(entity.User)
	err := db.Where("id = ?", "Fadel").First(user).Error
	assert.Nil(t, err)

	creditcards := new(entity.Creditcard)
	err = db.Where("user_id = ?", user.ID).First(creditcards).Error
	assert.Nil(t, err)

	request := httptest.NewRequest(http.MethodDelete, "/api/creditcards/"+creditcards.ID, nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[bool])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, true, responseBody.Data)

}

func TestDeleteCreditcardFailed(t *testing.T) {
	TestCreateContact(t)

	user := new(entity.User)
	err := db.Where("id = ?", "Fadel").First(user).Error
	assert.Nil(t, err)
	request := httptest.NewRequest(http.MethodDelete, "/api/contacts/"+uuid.NewString(), nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[bool])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusNotFound, response.StatusCode)
}

func TestSearchCreditcard(t *testing.T) {
	TestLogin(t)

	user := new(entity.User)
	err := db.Where("id = ?", "Fadel").First(user).Error
	assert.Nil(t, err)

	CreateCreditCards(user, 20)

	request := httptest.NewRequest(http.MethodGet, "/api/creditcards", nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.CreditcardResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 10, len(responseBody.Data))
	assert.Equal(t, int64(20), responseBody.Paging.TotalItem)
	assert.Equal(t, int64(2), responseBody.Paging.TotalPage)
	assert.Equal(t, 1, responseBody.Paging.Page)
	assert.Equal(t, 10, responseBody.Paging.Size)
}

func TestSearchCreditcardWithFilter(t *testing.T) {
	TestLogin(t)

	user := new(entity.User)
	err := db.Where("id = ?", "Fadel").First(user).Error
	assert.Nil(t, err)

	CreateCreditCards(user, 20)

	request := httptest.NewRequest(http.MethodGet, "/api/creditcards?size=10&page=1&name=vis&number=112", nil)
	request.Header.Set("Accept", "application/json")
	request.Header.Set("Authorization", user.Token)

	response, err := app.Test(request)
	assert.Nil(t, err)

	bytes, err := io.ReadAll(response.Body)
	assert.Nil(t, err)

	responseBody := new(model.WebResponse[[]model.CreditcardResponse])
	err = json.Unmarshal(bytes, responseBody)
	assert.Nil(t, err)

	assert.Equal(t, http.StatusOK, response.StatusCode)
	assert.Equal(t, 10, len(responseBody.Data))
	assert.Equal(t, int64(20), responseBody.Paging.TotalItem)
	assert.Equal(t, int64(2), responseBody.Paging.TotalPage)
	assert.Equal(t, 1, responseBody.Paging.Page)
	assert.Equal(t, 10, responseBody.Paging.Size)
}
