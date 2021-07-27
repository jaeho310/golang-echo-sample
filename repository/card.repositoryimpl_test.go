package repository_test

import (
	"fmt"
	"platform-sample/infrastructure/database"
	"platform-sample/infrastructure/server"
	"platform-sample/repository"
	"testing"
)

func initCardRepository() *repository.CardRepositoryImpl {
	mockDb := database.SqlStore{}.GetMockDb()
	mockServer := server.Server{MainDb: mockDb}
	return mockServer.InjectCardRepository()
}

func Test_GetCards(t *testing.T) {
	cardRepository := initCardRepository()
	cards, err := cardRepository.GetCards()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(cards)
}
