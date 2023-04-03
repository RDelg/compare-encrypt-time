package routers_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/RDelg/compare-encrypt-time/src/models"
	"github.com/RDelg/compare-encrypt-time/src/routers"
	setting "github.com/RDelg/compare-encrypt-time/src/settings"
	"github.com/go-playground/assert/v2"
)

func TestInitRouter(t *testing.T) {

	os.Setenv("CYPHER_KEY", "HJkPmTz+uY7wd0p1+w//DABgbvPq9/230RwEG2sJ9mo=")
	os.Setenv("CYPHER_IV", "AAAAAAAAAAAAAAAAAAAAAA==")
	setting.Setup()
	defer os.Unsetenv("CYPHER_KEY")
	defer os.Unsetenv("CYPHER_IV")

	r := routers.InitRouter()

	t.Run("RemoteFunctionAdapter route encrypts", func(t *testing.T) {
		requestBody, _ := json.Marshal(models.Message{
			RequestId: "test",
			Caller:    "test",
			Session:   "test",
			Context: map[string]string{
				"action": "encrypt",
			},
			Calls: [][]string{
				{"test1"},
			},
		})
		request, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("RemoteFunctionAdapter route decrypts", func(t *testing.T) {
		requestBody, _ := json.Marshal(models.Message{
			RequestId: "test",
			Caller:    "test",
			Session:   "test",
			Context: map[string]string{
				"action": "decrypt",
			},
			Calls: [][]string{
				{"nt7TPnfBIZa8MSFmLura1Q=="},
			},
		})
		request, _ := http.NewRequest("POST", "/", bytes.NewBuffer(requestBody))
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)

		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Encrypt route", func(t *testing.T) {
		requestBody, _ := json.Marshal(models.Message{
			RequestId: "test",
			Caller:    "test",
			Session:   "test",
			Context: map[string]string{
				"test": "test",
			},
			Calls: [][]string{
				{"test1"}, {"test2"},
			},
		})
		request, _ := http.NewRequest("POST", "/encrypt", bytes.NewBuffer(requestBody))
		response := httptest.NewRecorder()
		fmt.Println(response)
		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("Decrypt route", func(t *testing.T) {
		requestBody, _ := json.Marshal(models.Message{
			RequestId: "test",
			Caller:    "test",
			Session:   "test",
			Context: map[string]string{
				"test": "test",
			},
			Calls: [][]string{
				{"nt7TPnfBIZa8MSFmLura1Q=="}, {"QhwCxkKGgYUrVu8/2lzzYw=="},
			},
		})
		request, _ := http.NewRequest("POST", "/decrypt", bytes.NewBuffer(requestBody))
		response := httptest.NewRecorder()

		r.ServeHTTP(response, request)
		assert.Equal(t, http.StatusOK, response.Code)
	})
}
