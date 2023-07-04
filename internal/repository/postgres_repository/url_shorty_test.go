package postgres_repository

import (
	"fmt"
	"log"
	"ozonTech/muhtarov/internal/models"
	"ozonTech/muhtarov/pkg/database/postgres"
	"testing"
)

func TestEmailRepo_AddUser(t *testing.T) {
	db, err := postgres.NewClient("host=localhost port=5432 user=postgres password=password dbname=email_users sslmode=disable timezone=UTC connect_timeout=5")
	if err != nil {
		log.Println(err)
		return
	}

	repo := NewEmailRepo(db)

	r, err := repo.AddUrl(models.UrlStruct{})
	if err != nil {
		log.Println(err)
		//return
	}

	_, err = repo.GetFullUrlByShort("497bd685-a4a7-4da6-8926-a71ae956d5b0")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("497bd685-a4a7-4da6-8926-a71ae956d5b0")
	}

	fmt.Print(r)

}
