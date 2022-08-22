package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/qwark97/go-images/internal/mocks"
	"github.com/qwark97/go-images/storage"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestShouldCreateNewHandlersObject(t *testing.T) {
	// given
	var h *Handlers
	storage := new(mocks.Storage)

	// when
	h = NewHandlers(storage)

	// then
	assert.NotNil(t, h)
}

func TestShouldCheckPostHandler(t *testing.T) {
	// given
	expectedStatus := http.StatusNoContent
	expectedResponse := storage.CreateResp{
		Status: expectedStatus,
		Msg:    "ok",
	}
	var fetchedResponse storage.CreateResp

	s := new(mocks.Storage)
	s.On("Create", mock.AnythingOfType("storage.CreateData")).Return(storage.CreateResp{
		Status: expectedStatus,
		Msg:    "ok",
	}, nil)
	h := NewHandlers(s)

	body := bytes.NewReader(preparePostBody())
	request := httptest.NewRequest("POST", "/post", body)
	writer := httptest.NewRecorder()

	// when
	h.Post(writer, request)

	// then
	json.NewDecoder(writer.Result().Body).Decode(&fetchedResponse)
	assert.Equal(t, expectedStatus, writer.Result().StatusCode)
	assert.Equal(t, expectedResponse, fetchedResponse)
}

func preparePostBody() []byte {
	body := storage.CreateData{
		User:        "test-user",
		Description: "test-description",
	}
	res, _ := json.Marshal(body)
	return res
}

func TestShouldCheckGetHandler(t *testing.T) {
	// given
	userID := primitive.ObjectID([12]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12})
	now := time.Now()

	expectedStatus := http.StatusOK
	expectedResponse := []storage.ReadResp{
		{
			ID:          userID,
			User:        "test-user",
			Description: "test description",
			CreatedAt:   primitive.NewDateTimeFromTime(now),
		},
	}
	var fetchedResponse []storage.ReadResp

	s := new(mocks.Storage)
	s.On("Read").Return([]storage.ReadResp{
		{
			ID:          userID,
			User:        "test-user",
			Description: "test description",
			CreatedAt:   primitive.NewDateTimeFromTime(now),
		},
	}, nil)
	h := NewHandlers(s)

	body := bytes.NewReader([]byte{})
	request := httptest.NewRequest("GET", "/get", body)
	writer := httptest.NewRecorder()

	// when
	h.Get(writer, request)

	// then
	json.NewDecoder(writer.Result().Body).Decode(&fetchedResponse)
	assert.Equal(t, expectedStatus, writer.Result().StatusCode)
	assert.Equal(t, expectedResponse, fetchedResponse)
}
