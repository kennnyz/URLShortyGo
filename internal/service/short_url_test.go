package service

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"math/rand"
	"ozonTech/muhtarov/internal/models"
	mock_repository "ozonTech/muhtarov/internal/repository/mock"
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
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mock_repository.NewMockURLShortyRepository(ctrl)

	// Ожидаем вызов метода AddUrl с определенным аргументом и возвращаем значения
	expectedUrlStruct := models.UrlStruct{
		LongUrl:  "www.example.com",
		ShortUrl: "78HxvnMug9",
		Id:       96569598279802005,
	}
	mockRepo.EXPECT().AddUrl(models.UrlStruct{
		LongUrl:  "www.example.com",
		ShortUrl: "78HxvnMug9",
		Id:       96569598279802005,
	}).Return(expectedUrlStruct, nil)

	// Создаем экземпляр сервиса с моком репозитория
	s := NewURLShortyService(mockRepo)

	// Вызываем метод AddUrl
	//urlStruct, err := s.AddUrl("www.example.com")

	url, err := s.repo.AddUrl(models.UrlStruct{
		LongUrl:  "www.example.com",
		ShortUrl: "78HxvnMug9",
		Id:       96569598279802005,
	})
	if err != nil {
		t.Error(err)
	}

	// Проверяем, что нет ошибок и значения соответствуют ожиданиям
	assert.NoError(t, err)
	assert.Equal(t, expectedUrlStruct, url)
}
