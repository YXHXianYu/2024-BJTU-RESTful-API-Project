package test

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"restfulapi/internal/utils"
)

func ConvertToJson(data interface{}) string {
	json_str, _ := json.Marshal(data)
	return string(json_str)
}

func ConvertToBytes(data interface{}) io.Reader {
	json_str, _ := json.Marshal(data)
	return bytes.NewBuffer(json_str)
}

func ConvertToMap(jsonStr string) map[string]interface{} {
	var data map[string]interface{}
	json.Unmarshal([]byte(jsonStr), &data)
	return data
}

func TestUser(t *testing.T) {
	r := utils.Setup()

	// === User ===

	{ // get user (unauthorized)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/16", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	}

	var token string

	{ // login
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/sessions", ConvertToBytes(
			struct {
				Username string `json:"username"`
				Password string `json:"password"`
			} {
				Username: "admin",
				Password: "20021012",
			},
		))
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		data := ConvertToMap(w.Body.String())
		token = data["data"].(map[string]interface{})["token"].(string)
		t.Log("token: " + token)
	}

	{ // get user
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/16", nil)
		req.Header.Set("Authorization", token)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		
		data := ConvertToMap(w.Body.String())
		assert.Equal(t, "admin", data["data"].(map[string]interface{})["username"])
		assert.Equal(t, "admin", data["data"].(map[string]interface{})["role"])
		assert.Equal(t, float64(16), data["data"].(map[string]interface{})["id"])
	}
	{ // get user (unauthorized)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/users/16", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	}
	{ // get all users
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/users", nil)
		req.Header.Set("Authorization", token)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	}
	{ // create user (bad request)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", ConvertToBytes(
			struct {
				Username string `json:"username"`
				Password string `json:"password"`
				Email string `json:"email"`
				Telephone string `json:"telephone"`
			} {
				Username: "admin",
				Password: "20021012",
				Email: "2943003@qq.com",
				Telephone: "1234567890",
			},
		))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	}
	{ // create user (bad request)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/users", ConvertToBytes(
			struct {
				Username string `json:"username"`
				Password string `json:"password"`
				Email string `json:"email"`
				Telephone string `json:"telephone"`
			} {
				Username: "fields_not_complete",
				Password: "20021012",
			},
		))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	}

	// === Blog ===

	{ // create blog (unauthorized)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/blogs", ConvertToBytes(
			struct {
				Title string `json:"title"`
				Content string `json:"content"`
			} {
				Title: "title",
				Content: "content",
			},
		))
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	}
	{ // create blog (bad request)
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/blogs", ConvertToBytes(
			struct {
				Title string `json:"title"`
				Content string `json:"content"`
			} {
				Title: "title",
				Content: "content",
			},
		))
		req.Header.Set("Authorization", token)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	}
	{ // create blog (bad request (fields not complete))
		req, _ := http.NewRequest(http.MethodPost, "/api/v1/blogs", ConvertToBytes(
			struct {
				Title string `json:"title"`
				Content string `json:"content"`
			} {
				Title: "title",
			},
		))
		req.Header.Set("Authorization", token)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
	}
	{ // get blog (unauthorized)
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/blogs/1", nil)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	}
	{ // get blog
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/blogs/1", nil)
		req.Header.Set("Authorization", token)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	}
	{ // get all blogs
		req, _ := http.NewRequest(http.MethodGet, "/api/v1/blogs", nil)
		req.Header.Set("Authorization", token)
		w := httptest.NewRecorder()
		r.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	}
}