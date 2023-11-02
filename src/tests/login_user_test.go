package tests

import (
	"encoding/json"
	"fmt"
	"github.com/HunCoding/meu-primeiro-crud-go/src/controller/model/request"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"math/rand"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestLoginUser(t *testing.T) {

	t.Run("user_and_password_is_not__valid", func(t *testing.T) {
		recorderCreateUser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateUser)

		recorderLoginUser := httptest.NewRecorder()
		ctxLoginUser := GetTestGinContext(recorderLoginUser)

		email := fmt.Sprintf("%d@testfds.com", rand.Int())
		password := fmt.Sprintf("%d$#@$#@", rand.Int())

		userCreateRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "Test User",
			Age:      10,
		}

		bCreate, _ := json.Marshal(userCreateRequest)
		stringReaderCreate := io.NopCloser(strings.NewReader(string(bCreate)))

		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReaderCreate)
		UserController.CreateUser(ctxCreateUser)

		userLoginRequest := request.UserLogin{
			Email:    "test@testnotpassing.com",
			Password: "fdsfdsr43$#@$#@$@#",
		}

		bLogin, _ := json.Marshal(userLoginRequest)
		stringReaderLogin := io.NopCloser(strings.NewReader(string(bLogin)))

		MakeRequest(ctxLoginUser, []gin.Param{}, url.Values{}, "POST", stringReaderLogin)
		UserController.LoginUser(ctxLoginUser)

		assert.EqualValues(t, http.StatusForbidden, recorderLoginUser.Result().StatusCode)
	})

	t.Run("user_and_password_is_valid", func(t *testing.T) {
		recorderCreateUser := httptest.NewRecorder()
		ctxCreateUser := GetTestGinContext(recorderCreateUser)

		recorderLoginUser := httptest.NewRecorder()
		ctxLoginUser := GetTestGinContext(recorderLoginUser)

		email := fmt.Sprintf("%d@test.com", rand.Int())
		password := fmt.Sprintf("%d$#@$#@", rand.Int())

		userCreateRequest := request.UserRequest{
			Email:    email,
			Password: password,
			Name:     "Test User",
			Age:      10,
		}

		bCreate, _ := json.Marshal(userCreateRequest)
		stringReaderCreate := io.NopCloser(strings.NewReader(string(bCreate)))

		MakeRequest(ctxCreateUser, []gin.Param{}, url.Values{}, "POST", stringReaderCreate)
		UserController.CreateUser(ctxCreateUser)

		userLoginRequest := request.UserLogin{
			Email:    email,
			Password: password,
		}

		bLogin, _ := json.Marshal(userLoginRequest)
		stringReaderLogin := io.NopCloser(strings.NewReader(string(bLogin)))

		MakeRequest(ctxLoginUser, []gin.Param{}, url.Values{}, "POST", stringReaderLogin)
		UserController.LoginUser(ctxLoginUser)

		assert.EqualValues(t, http.StatusOK, recorderLoginUser.Result().StatusCode)
		assert.NotEmpty(t, recorderLoginUser.Result().Header.Get("Authorization"))
	})
}
