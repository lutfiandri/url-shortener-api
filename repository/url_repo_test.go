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

	// tests table (@datapoint)
	tests := []struct {
		name     string
		inputId  string
		expected entity.Url
		urls     []entity.Url
	}{
		{
			name:    "TestFindById_Positive-001",
			inputId: "id002",
			expected: entity.Url{
				Id:      "id002",
				LongUrl: "https://facebook.com",
				Title:   "Facebook",
			},
			urls: []entity.Url{
				{
					Id:      "id001",
					LongUrl: "https://google.com",
					Title:   "Google",
				},
				{
					Id:      "id002",
					LongUrl: "https://facebook.com",
					Title:   "Facebook",
				},
				{
					Id:      "id003",
					LongUrl: "https://manjaro.org",
					Title:   "Manjaro",
				},
			},
		},
		{
			name:    "TestFindById_Positive-002",
			inputId: "id002",
			expected: entity.Url{
				Id:      "id002",
				LongUrl: "https://facebook.com",
				Title:   "Facebook",
			},
			urls: []entity.Url{
				{
					Id:      "id001",
					LongUrl: "https://google.com",
					Title:   "Google",
				},
				{
					Id:      "id002",
					LongUrl: "https://facebook.com",
					Title:   "Facebook",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup data in db (using gorm)
			for _, url := range test.urls {
				database.Create(&url)
			}

			// delete data after usage in db (using gorm)
			defer func() {
				for _, url := range test.urls {
					database.Delete(&url)
				}
			}()

			// test (@test)
			actual, err := urlRepository.FindById(test.inputId)
			assert.NoErrorf(t, err, "error on %s", test.name)
			assert.Equalf(t, test.expected, actual, "not equal on %s", test.name)
		})
	}
}

func TestFindById_Negative(t *testing.T) {
	urlRepository := repository.NewUrlRepository(database)

	// tests table (@datapoint)
	tests := []struct {
		name    string
		inputId string
		urls    []entity.Url
	}{
		{
			name:    "TestFindById_Negative-001",
			inputId: "id004",
			urls: []entity.Url{
				{
					Id:      "id001",
					LongUrl: "https://google.com",
					Title:   "Google",
				},
				{
					Id:      "id002",
					LongUrl: "https://facebook.com",
					Title:   "Facebook",
				},
				{
					Id:      "id003",
					LongUrl: "https://manjaro.org",
					Title:   "Manjaro",
				},
			},
		},
		{
			name:    "TestFindById_Negative-002",
			inputId: "id005",
			urls: []entity.Url{
				{
					Id:      "id001",
					LongUrl: "https://google.com",
					Title:   "Google",
				},
				{
					Id:      "id002",
					LongUrl: "https://facebook.com",
					Title:   "Facebook",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup data in db (using gorm)
			for _, url := range test.urls {
				database.Create(&url)
			}

			// delete data after usage in db (using gorm)
			defer func() {
				for _, url := range test.urls {
					database.Delete(&url)
				}
			}()

			// test (@test)
			actual, err := urlRepository.FindById(test.inputId)
			assert.Errorf(t, err, "not error on %s", test.name)
			assert.Emptyf(t, actual, "not empty on %s", test.name)
		})
	}
}

func TestCreate_Positive(t *testing.T) {
	urlRepository := repository.NewUrlRepository(database)

	// tests table (@datapoint)
	tests := []struct {
		name     string
		inputUrl entity.Url
		urls     []entity.Url
	}{
		{
			name: "TestCreate_Positive-001",
			inputUrl: entity.Url{
				Id:      "id005",
				LongUrl: "https://google.com",
				Title:   "Google",
			},
			urls: []entity.Url{
				{
					Id:      "id001",
					LongUrl: "https://google.com",
					Title:   "Google",
				},
				{
					Id:      "id002",
					LongUrl: "https://facebook.com",
					Title:   "Facebook",
				},
				{
					Id:      "id003",
					LongUrl: "https://manjaro.org",
					Title:   "Manjaro",
				},
			},
		},
		{
			name: "TestCreate_Positive-002",
			inputUrl: entity.Url{
				Id:      "id006",
				LongUrl: "https://something2.com",
				Title:   "something2",
			},
			urls: []entity.Url{
				{
					Id:      "id001",
					LongUrl: "https://google.com",
					Title:   "Google",
				},
				{
					Id:      "id002",
					LongUrl: "https://facebook.com",
					Title:   "Facebook",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup data in db (using gorm)
			for _, url := range test.urls {
				database.Create(&url)
			}

			// delete data after usage in db (using gorm)
			defer func() {
				for _, url := range test.urls {
					database.Delete(&url)
				}
			}()

			// test (@test)
			err := urlRepository.Create(test.inputUrl)
			assert.NoErrorf(t, err, "error on %s", test.name)

			// verify the data is inserted
			var result entity.Url
			database.First(&result, "id = ?", test.inputUrl.Id)
			assert.NotEmptyf(t, result, "empty on %s", test.name)
		})
	}
}

func TestCreate_Negative(t *testing.T) {
	urlRepository := repository.NewUrlRepository(database)

	// tests table (@datapoint)
	tests := []struct {
		name     string
		inputUrl entity.Url
		urls     []entity.Url
	}{
		{
			name: "TestCreate_Negative-001",
			inputUrl: entity.Url{
				Id:      "id002",
				LongUrl: "https://google.com",
				Title:   "Google",
			},
			urls: []entity.Url{
				{
					Id:      "id001",
					LongUrl: "https://google.com",
					Title:   "Google",
				},
				{
					Id:      "id002",
					LongUrl: "https://facebook.com",
					Title:   "Facebook",
				},
				{
					Id:      "id003",
					LongUrl: "https://manjaro.org",
					Title:   "Manjaro",
				},
			},
		},
		{
			name: "TestCreate_Negative-002",
			inputUrl: entity.Url{
				Id:      "id001",
				LongUrl: "https://something2.com",
				Title:   "something2",
			},
			urls: []entity.Url{
				{
					Id:      "id001",
					LongUrl: "https://google.com",
					Title:   "Google",
				},
				{
					Id:      "id002",
					LongUrl: "https://facebook.com",
					Title:   "Facebook",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup data in db (using gorm)
			for _, url := range test.urls {
				database.Create(&url)
			}

			// delete data after usage in db (using gorm)
			defer func() {
				for _, url := range test.urls {
					database.Delete(&url)
				}
			}()

			// test (@test)
			err := urlRepository.Create(test.inputUrl)
			assert.Errorf(t, err, "error on %s", test.name)
		})
	}
}

func TestDeleteById_Positive(t *testing.T) {
	urlRepository := repository.NewUrlRepository(database)

	// tests table (@datapoint)
	tests := []struct {
		name    string
		inputId string
		urls    []entity.Url
	}{
		{
			name:    "TestDeleteById_Positive-001",
			inputId: "id002",
			urls: []entity.Url{
				{
					Id:      "id001",
					LongUrl: "https://google.com",
					Title:   "Google",
				},
				{
					Id:      "id002",
					LongUrl: "https://facebook.com",
					Title:   "Facebook",
				},
				{
					Id:      "id003",
					LongUrl: "https://manjaro.org",
					Title:   "Manjaro",
				},
			},
		},
		{
			name:    "TestDeleteById_Positive-002",
			inputId: "id002",
			urls: []entity.Url{
				{
					Id:      "id001",
					LongUrl: "https://google.com",
					Title:   "Google",
				},
				{
					Id:      "id002",
					LongUrl: "https://facebook.com",
					Title:   "Facebook",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup data in db (using gorm)
			for _, url := range test.urls {
				database.Create(&url)
			}

			// delete data after usage in db (using gorm)
			defer func() {
				for _, url := range test.urls {
					database.Delete(&url)
				}
			}()

			// test (@test)
			err := urlRepository.DeleteById(test.inputId)
			assert.NoErrorf(t, err, "error on %s", test.name)

			// verify the data is deleted
			var result entity.Url
			database.First(&result, "id = ?", result.Id)
			assert.Emptyf(t, result, "not empty on %s", test.name)
		})
	}
}

func TestDeleteById_Negative(t *testing.T) {
	urlRepository := repository.NewUrlRepository(database)

	// tests table (@datapoint)
	tests := []struct {
		name    string
		inputId string
		urls    []entity.Url
	}{
		{
			name:    "TestDeleteById_Positive-001",
			inputId: "id010",
			urls: []entity.Url{
				{
					Id:      "id001",
					LongUrl: "https://google.com",
					Title:   "Google",
				},
				{
					Id:      "id002",
					LongUrl: "https://facebook.com",
					Title:   "Facebook",
				},
				{
					Id:      "id003",
					LongUrl: "https://manjaro.org",
					Title:   "Manjaro",
				},
			},
		},
		{
			name:    "TestDeleteById_Positive-002",
			inputId: "id010",
			urls: []entity.Url{
				{
					Id:      "id001",
					LongUrl: "https://google.com",
					Title:   "Google",
				},
				{
					Id:      "id002",
					LongUrl: "https://facebook.com",
					Title:   "Facebook",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// setup data in db (using gorm)
			for _, url := range test.urls {
				database.Create(&url)
			}

			// delete data after usage in db (using gorm)
			defer func() {
				for _, url := range test.urls {
					database.Delete(&url)
				}
			}()

			// test (@test)
			err := urlRepository.DeleteById(test.inputId)
			assert.Errorf(t, err, "not error on %s", test.name)
		})
	}
}
