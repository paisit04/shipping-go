package rest_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/paisit04/shipping-go/handlers/rest"
)

func TestTranslateAPI(t *testing.T) {
	tt := []struct {
		Endpoint            string
		StatusCode          int
		ExpectedLanguage    string
		ExpectedTranslation string
	}{
		{
			Endpoint:            "/hello",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "english",
			ExpectedTranslation: "hello",
		},
		{
			Endpoint:            "/hello?language=german",
			StatusCode:          http.StatusOK,
			ExpectedLanguage:    "german",
			ExpectedTranslation: "hallo",
		},
		{
			Endpoint:            "/hello?language=dutch",
			StatusCode:          http.StatusNotFound,
			ExpectedLanguage:    "",
			ExpectedTranslation: "",
		},
	}

	handler := http.HandlerFunc(rest.TranslateHandler)

	for _, test := range tt {
		rr := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", test.Endpoint, nil)

		handler.ServeHTTP(rr, req)

		if rr.Code != test.StatusCode {
			t.Errorf("Expected status %d, but got %d",
				test.StatusCode, rr.Code)
		}

		var resp rest.Resp
		_ = json.Unmarshal(rr.Body.Bytes(), &resp)

		if resp.Language != test.ExpectedLanguage {
			t.Errorf(`Expected language "%s", but got %s`,
				test.ExpectedLanguage, resp.Language)
		}
		if resp.Translation != test.ExpectedTranslation {
			t.Errorf(`Expected translation "%s", but got %s`,
				test.ExpectedTranslation, resp.Translation)
		}
	}

}
