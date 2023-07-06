package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"math/rand"
	"ozonTech/muhtarov/internal/models"
	mock_service "ozonTech/muhtarov/internal/service/mock"
	"testing"
)

func TestEncodeBase62(t *testing.T) {
	testTable := []struct {
		num         int64
		expectedAns string
	}{
		{0, "0"},
		{1, "1"},
		{10, "A"},
		{61, "z"},
		{123, "1z"},
		{998, "G6"},
		{1000, "G8"},
		{8829381, "b2vN"},
		{623, "A3"},
		{593, "9Z"},
		{1233, "Jt"},
		{29873492783, "Wbi88l"},
		{-992, "-G0"},
		{67931256112695545, "517nrE0PEn"},
		{45383646439471391, "3LrAXqPqWt"},
	}

	for _, testCase := range testTable {
		result := encodeBase62(testCase.num)
		assert.Equal(t, testCase.expectedAns, result, "Unexpected result for input: %d", testCase.num)
	}
}

func TestGenerateID(t *testing.T) {
	ids := make(map[int64]bool) // Мапа для отслеживания уникальности идентификаторов

	for i := 0; i < 1000; i++ {
		longURL := fmt.Sprintf("https://%dexample.com/%d", i, i+23)
		id := generateID(longURL)

		assert.False(t, ids[id], "Duplicate ID generated: %d", id)
		assert.True(t, id >= 0, "Negative ID generated: %d", id)
		ids[id] = true
	}
}

func TestMakeShortURL(t *testing.T) {
	s := &URLShortyService{}
	for i := 0; i < 1000; i++ {
		longURL := fmt.Sprintf("https://%dexample.com/%d.ru/test/somestring/___99999kjdnv", i, rand.Int())
		urlStruct := s.makeShortURL(longURL)

		assert.NotEmpty(t, urlStruct.ShortUrl, "ShortUrl is empty")
		assert.True(t, len(urlStruct.ShortUrl) == 10, "Invalid len short url: %d", urlStruct.Id, len(urlStruct.ShortUrl), urlStruct.ShortUrl, urlStruct.LongUrl)
	}
}

func TestURLShortyService(t *testing.T) {
	type mockBehavior func(s *mock_service.MockURLShortyRepository, urlStruct models.UrlStruct)
	expectedUrlStruct := models.UrlStruct{
		LongUrl:  "www.example.com",
		ShortUrl: "78HxvnMug9",
		Id:       96569598279802005,
	}

	testTable := []struct {
		name              string
		inputUrlStruct    models.UrlStruct
		mockBehavior      mockBehavior
		expectedUrlStruct models.UrlStruct
		expectedErr       error
	}{
		{
			name:           "ok",
			inputUrlStruct: expectedUrlStruct,
			mockBehavior: func(s *mock_service.MockURLShortyRepository, urlStruct models.UrlStruct) {
				s.EXPECT().AddUrl(expectedUrlStruct).Return(expectedUrlStruct, nil)
			},
			expectedUrlStruct: expectedUrlStruct,
		},
		{
			name:           "invalid url",
			inputUrlStruct: expectedUrlStruct,
			mockBehavior: func(s *mock_service.MockURLShortyRepository, urlStruct models.UrlStruct) {
				s.EXPECT().AddUrl(expectedUrlStruct).Return(models.UrlStruct{}, models.NotValidUrlErr)
			},
			expectedUrlStruct: models.UrlStruct{},
			expectedErr:       models.NotValidUrlErr,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// init deps
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockRepo := mock_service.NewMockURLShortyRepository(ctrl)
			testCase.mockBehavior(mockRepo, testCase.inputUrlStruct)

			s := NewURLShortyService(mockRepo)
			url, err := s.repo.AddUrl(expectedUrlStruct)

			assert.Equal(t, err, testCase.expectedErr)
			assert.Equal(t, testCase.expectedUrlStruct, url)
		})
	}
}
