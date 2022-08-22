package integrationtests_test

import (
	"context"
	"net/http"
	"testing"

	"github.com/qwark97/go-images/storage"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestShouldAddDataIntoDB(t *testing.T) {
	// given
	const uri = "mongodb://localhost:27017/"
	s := storage.NewStorage()
	defer s.Disconnect()
	user := storage.CreateData{
		User:        "test-user",
		Description: "test description",
	}
	var fetchedUser storage.ReadResp

	// when
	resp, err := s.Create(user)

	// then
	assert.Nil(t, err)
	assert.Equal(t, "ok", resp.Msg)
	assert.Equal(t, http.StatusNoContent, resp.Status)

	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	db := client.Database("exampleDB")
	coll := db.Collection("exampleColl")
	res := coll.FindOne(context.TODO(), bson.D{{"user", "test-user"}})
	err = res.Decode(&fetchedUser)
	assert.Nil(t, err)

	assert.NotNil(t, fetchedUser.CreatedAt)
	assert.Equal(t, user.User, fetchedUser.User)
	assert.Equal(t, user.Description, fetchedUser.Description)
}
