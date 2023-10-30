package tests

import (
	"context"
	"github.com/HunCoding/meu-primeiro-crud-go/src/controller"
	"github.com/HunCoding/meu-primeiro-crud-go/src/model/repository"
	"github.com/HunCoding/meu-primeiro-crud-go/src/model/service"
	"github.com/HunCoding/meu-primeiro-crud-go/src/tests/connection"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"io"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"
)

var (
	UserController controller.UserControllerInterface
	Database       *mongo.Database
)

func TestMain(m *testing.M) {
	err := os.Setenv("MONGODB_USER_DB", "test_user")
	if err != nil {
		return
	}

	closeConnection := func() {}
	Database, closeConnection = connection.OpenConnection()

	repo := repository.NewUserRepository(Database)
	userService := service.NewUserDomainService(repo)
	UserController = controller.NewUserControllerInterface(userService)

	defer func() {
		os.Clearenv()
		closeConnection()
	}()

	os.Exit(m.Run())
}

func TestFindUserByEmail(t *testing.T) {

	t.Run("user_not_found_with_this__email", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}

		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserController.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("user__found_with_specified_email", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		_, err := Database.
			Collection("test_user").
			InsertOne(context.Background(), bson.M{"_id": id, "name": t.Name(), "email": "test@test.com"})
		if err != nil {
			t.Fatal(err)
			return
		}

		param := []gin.Param{
			{
				Key:   "userEmail",
				Value: "test@test.com",
			},
		}

		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserController.FindUserByEmail(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func TestFindUserById(t *testing.T) {

	t.Run("user_not_found_with_this_id", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		id := primitive.NewObjectID().Hex()

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id,
			},
		}

		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserController.FindUserByID(ctx)

		assert.EqualValues(t, http.StatusNotFound, recorder.Code)
	})

	t.Run("user_found_with_specified_id", func(t *testing.T) {
		recorder := httptest.NewRecorder()
		ctx := GetTestGinContext(recorder)
		id := primitive.NewObjectID()

		_, err := Database.
			Collection("test_user").
			InsertOne(context.Background(), bson.M{"_id": id, "name": t.Name(), "email": "test@test.com"})
		if err != nil {
			t.Fatal(err)
			return
		}

		param := []gin.Param{
			{
				Key:   "userId",
				Value: id.Hex(),
			},
		}

		MakeRequest(ctx, param, url.Values{}, "GET", nil)
		UserController.FindUserByID(ctx)

		assert.EqualValues(t, http.StatusOK, recorder.Code)
	})
}

func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}

	return ctx
}

func MakeRequest(
	c *gin.Context,
	param gin.Params,
	u url.Values,
	method string,
	body io.ReadCloser) {
	c.Request.Method = method
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param

	c.Request.Body = body
	c.Request.URL.RawQuery = u.Encode()
}
