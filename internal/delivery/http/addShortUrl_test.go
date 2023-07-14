package http

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	. "ozonTech/muhtarov/internal/delivery/mock"
	"ozonTech/muhtarov/internal/models"
	"testing"
)

func TestHandleAddShortUrl(t *testing.T) {
	type mockBehavior func(shorty *MockURLShorty)
	testTable := []struct {
		name               string
		testId             int
		inputLongUrl       string
		mockBehavior       mockBehavior
		expectedStatusCode int
		expectedShortURL   string
	}{
		{
			name:         "OK",
			testId:       1,
			inputLongUrl: "www.google.com",
			mockBehavior: func(shorty *MockURLShorty) {
				shorty.EXPECT().AddUrl("www.google.com").Return(models.UrlStruct{
					LongUrl:  "www.google.com",
					ShortUrl: "1wyI4Mkj7L",
					Id:       26413135951547755,
				}, nil)
			},
			expectedStatusCode: 200,
			expectedShortURL:   `1wyI4Mkj7L`,
		},
		{
			name:               "empty longURL",
			testId:             2,
			inputLongUrl:       "",
			mockBehavior:       func(shorty *MockURLShorty) {},
			expectedStatusCode: 400,
			expectedShortURL:   models.NotValidUrlErr.Error() + "\n",
		},
		{
			name:               "already short url have length 10",
			testId:             3,
			inputLongUrl:       "somepop",
			mockBehavior:       func(shorty *MockURLShorty) {},
			expectedStatusCode: 400,
			expectedShortURL:   models.NotValidUrlErr.Error() + "\n",
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			urlShorty := NewMockURLShorty(ctrl)
			testCase.mockBehavior(urlShorty)

			handler := NewHandler(urlShorty)
			mux := http.NewServeMux()
			mux.HandleFunc("/make-short-url", handler.makeShortUrl)

			req := httptest.NewRequest(http.MethodPost, "/make-short-url?url="+testCase.inputLongUrl, nil)
			recorder := httptest.NewRecorder()

			mux.ServeHTTP(recorder, req)
			assert.Equal(t, testCase.expectedStatusCode, recorder.Code)
			assert.Equal(t, testCase.expectedShortURL, recorder.Body.String())
		})
	}
}
