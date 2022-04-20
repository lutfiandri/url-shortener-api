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
	UserId:  shortuuid.New(),
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
