package service

import (
	"echo-clean/domain/entity"
	serviceErrors "echo-clean/domain/error"
	"echo-clean/domain/repository/mocks"
	"github.com/go-faker/faker/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestFetchArticle(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockArticleRepository, _, articleServiceImpl := setup()

		// Fake new article
		var mockArticle entity.Article
		_ = faker.FakeData(&mockArticle)

		// Sure no error
		mockListArticle := make([]*entity.Article, 0)
		mockListArticle = append(mockListArticle, &mockArticle)

		// Given data for Article Repository Mock
		mockArticleRepository.On("Fetch").Return(mockListArticle, nil)

		// When: Call Fetch function
		result, e := articleServiceImpl.Fetch()

		// Then: Assert the result
		assert.Empty(t, e, nil)
		assert.Equal(t, result, mockListArticle)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockArticleRepository, _, articleServiceImpl := setup()
		// Given data for Article Repository Mock
		mockArticleRepository.On("Fetch").Return(nil, assert.AnError)

		// When: Call Fetch function
		result, e2 := articleServiceImpl.Fetch()

		// Then: Assert the result
		assert.Empty(t, result)
		assert.Equal(t, e2, []serviceErrors.Error{serviceErrors.E10001})
	})
}

func TestStore(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockArticleRepository, mockAuthorRepository, articleServiceImpl := setup()

		// Fake new article
		var mockArticle entity.Article
		var mockAuthor entity.Author
		_ = faker.FakeData(&mockArticle)
		_ = faker.FakeData(&mockAuthor)

		mockListAuthor := make([]*entity.Author, 0)
		mockListAuthor = append(mockListAuthor, &mockAuthor)

		// Given data for Article Repository Mock
		mockAuthorRepository.On("Fetch").Return(mockListAuthor, nil)
		mockArticleRepository.On("Store", mock.Anything, mock.Anything).Return(&mockArticle, nil)

		// When: Call Store function
		result, e := articleServiceImpl.Store(&mockArticle)

		// Then: Assert the result
		assert.Empty(t, e, nil)
		assert.Equal(t, result, &mockArticle)
	})

	t.Run("error-validation", func(t *testing.T) {
		_, _, articleServiceImpl := setup()

		// Fake new article
		mockArticle := entity.Article{
			Title:   "",
			Content: "",
		}

		// When: Call Store function
		result, e := articleServiceImpl.Store(&mockArticle)

		// Then: Assert the result
		assert.Empty(t, result)
		assert.Equal(t, e, []serviceErrors.Error{serviceErrors.StatusBadRequestError})
	})

	t.Run("error-failed", func(t *testing.T) {
		mockArticleRepository, mockAuthorRepository, articleServiceImpl := setup()

		// Fake new article
		var mockArticle entity.Article
		var mockAuthor entity.Author
		_ = faker.FakeData(&mockArticle)
		_ = faker.FakeData(&mockAuthor)

		mockListAuthor := make([]*entity.Author, 0)
		mockListAuthor = append(mockListAuthor, &mockAuthor)

		// Given data for Article Repository Mock
		mockAuthorRepository.On("Fetch").Return(mockListAuthor, nil)
		mockArticleRepository.On("Store", mock.Anything, mock.Anything).Return(nil, assert.AnError)

		// When: Call Store function
		result, e := articleServiceImpl.Store(&mockArticle)

		// Then: Assert the result
		assert.Empty(t, result)
		assert.Equal(t, e, []serviceErrors.Error{serviceErrors.E10001})
	})
}

func setup() (*mocks.ArticleRepository, *mocks.AuthorRepository, ArticleService) {
	mockArticleRepository := new(mocks.ArticleRepository)
	mockAuthorRepository := new(mocks.AuthorRepository)

	articleServiceImpl := NewArticleService(mockArticleRepository, mockAuthorRepository)

	return mockArticleRepository, mockAuthorRepository, articleServiceImpl
}
