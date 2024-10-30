package handler

import (
	"echo-clean/domain/entity"
	serviceErrors "echo-clean/domain/error"
	"echo-clean/domain/service/mocks"
	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchArticle(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockArticleService := new(mocks.ArticleService)

		// Fake new article
		var fakeArticle entity.Article
		_ = faker.FakeData(&fakeArticle)

		// Sure no error
		fakeListArticle := make([]*entity.Article, 0)
		fakeListArticle = append(fakeListArticle, &fakeArticle)

		// Given data for Article Service Mock
		mockArticleService.On("Fetch").Return(fakeListArticle, nil)

		// Construct the handler
		articleHandler := NewArticleHandler(mockArticleService)

		// Create the request
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/articles", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		err := articleHandler.FetchArticle(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), fakeArticle.Content)

		mockArticleService.AssertExpectations(t)
	})

	t.Run("error-failed", func(t *testing.T) {
		mockArticleService := new(mocks.ArticleService)

		fakeError := []serviceErrors.Error{serviceErrors.E10001}

		// Given data for Article Service Mock
		mockArticleService.On("Fetch").Return(nil, fakeError)

		// Construct the handler
		articleHandler := NewArticleHandler(mockArticleService)

		// Create the request
		e := echo.New()
		req := httptest.NewRequest(http.MethodGet, "/articles", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		// Call function
		_ = articleHandler.FetchArticle(c)

		// Assert the result
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), fakeError[0].ToJsonError())
	})
}
