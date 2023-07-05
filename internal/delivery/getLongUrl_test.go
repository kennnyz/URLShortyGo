package delivery

import (
	"go.uber.org/mock/gomock"
	"ozonTech/muhtarov/internal/models"
	mock_repository "ozonTech/muhtarov/internal/repository/mock"
	"testing"
)

func TestHandler_getLongUrl(t *testing.T) {
	ctrl := gomock.NewController(t)

	// Assert that Bar() is invoked.
	defer ctrl.Finish()

	m := mock_repository.NewMockURLShortyRepository(ctrl)
	m.EXPECT().GetFullUrlByShort("www.google.com").Return(models.UrlStruct{
		LongUrl:  "www.google.com",
		ShortUrl: "7Nqd5nZyf4",
		Id:       23,
	})
}
