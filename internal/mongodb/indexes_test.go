package mongodb

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreateIndex(t *testing.T) {
	t.Parallel()

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))

	mt.Run("success", func(mt *mtest.T) {
		dbw := DatabaseWrapper{mt.DB}

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		err := dbw.CreateIndex("test", "test", false)
		assert.Nil(t, err)
	})
}
