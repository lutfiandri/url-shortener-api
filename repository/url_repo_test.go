package repository_test

import (
	"os"
	"testing"

	"url-shortener-api/config"
	"url-shortener-api/entity"
	"url-shortener-api/repository"

	"github.com/lithammer/shortuuid/v3"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

const (
	tmpLocation = "../.tmp"
	dbLocation  = tmpLocation + "/url-shortener.db"
)

var database *gorm.DB

var url entity.Url = entity.Url{
	Id:      shortuuid.New(),
	LongUrl: "https://google.com",
	Title:   "Google",
}

func TestMain(m *testing.M) {
	// before
	os.Mkdir(tmpLocation, os.ModePerm)
	database = config.NewSqliteDatabase(dbLocation)

	// run the tests
	m.Run()

	// after
	os.Remove("../.tmp/url-shortener.db")
}

func TestFindAll(t *testing.T) {
	urlRepository := repository.NewUrlRepository(database)

	// DataPoint
	urls := []entity.Url{
		{
			Id:      shortuuid.New(),
			LongUrl: "https://google.com",
			Title:   "Google",
		},
		{
			Id:      shortuuid.New(),
			LongUrl: "https://facebook.com",
			Title:   "Facebook",
		},
		{
			Id:      shortuuid.New(),
			LongUrl: "https://manjaro.org",
			Title:   "Manjaro",
		},
	}

	// Insert to DB
	for _, url := range urls {
		database.Create(&url)
	}

	// Remove from DB at the end
	defer func() {
		for _, url := range urls {
			database.Delete(&url)
		}
	}()

	// Test the testcases
	results, err := urlRepository.FindAll()

	assert.NoError(t, err)
	assert.ElementsMatch(t, urls, results)
}

func TestFindById_Positive(t *testing.T) {
	urlRepository := repository.NewUrlRepository(database)

	database.Create(&url)

	result, err := urlRepository.FindById(url.Id)

	assert.NoError(t, err)
	assert.Equal(t, url, result)

	database.Delete(&url)
}

func TestFindById_Negative(t *testing.T) {
	urlRepository := repository.NewUrlRepository(database)

	result, err := urlRepository.FindById(url.Id)

	assert.Error(t, err)
	assert.Empty(t, result)
}

func TestCreate(t *testing.T) {
	urlRepository := repository.NewUrlRepository(database)

	err := urlRepository.Create(url)

	var result entity.Url
	database.First(&result, "id = ?", url.Id)

	assert.NoError(t, err)
	assert.Equal(t, url, result)

	database.Delete(&url)
}

func TestDeleteById_Positive(t *testing.T) {
	urlRepository := repository.NewUrlRepository(database)

	database.Create(&url)

	err := urlRepository.DeleteById(url.Id)

	var result entity.Url
	database.First(&result, "id = ?", url.Id)

	assert.NoError(t, err)
	assert.Empty(t, result)
}

func TestDeleteById_Negative(t *testing.T) {
	urlRepository := repository.NewUrlRepository(database)

	err := urlRepository.DeleteById(url.Id)

	assert.Error(t, err)
}
