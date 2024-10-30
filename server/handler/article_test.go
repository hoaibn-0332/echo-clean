package handler

import (
	"echo-clean/domain/entity"
	serviceErrors "echo-clean/domain/error"
	"echo-clean/domain/service/mocks"
	"github.com/go-faker/faker/v4"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
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

func TestStoreArticle(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		mockArticleService := new(mocks.ArticleService)

		// Fake new article
		var fakeArticle entity.Article
		_ = faker.FakeData(&fakeArticle)

		// Given data for Article Service Mock
		mockArticleService.On("Store", mock.Anything).Return(&fakeArticle, nil)

		// Construct the handler
		articleHandler := NewArticleHandler(mockArticleService)

		// Create the request
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/articles",
			strings.NewReader(`{"content":"`+fakeArticle.Content+`", "title":"`+fakeArticle.Title+`"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles")

		err := articleHandler.Store(c)
		assert.NoError(t, err)
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), fakeArticle.Content)

		mockArticleService.AssertExpectations(t)
	})

	t.Run("error-bind", func(t *testing.T) {
		mockArticleService := new(mocks.ArticleService)

		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/articles", strings.NewReader("invalid data")) // Dữ liệu không hợp lệ
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		articleHandler := &ArticleHandler{
			Service: mockArticleService,
		}

		err := articleHandler.Store(c)

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusUnsupportedMediaType, rec.Code)
			assert.Contains(t, rec.Body.String(), serviceErrors.UnsupportedMediaTypeError.Message)
		}
	})

	t.Run("error-failed", func(t *testing.T) {
		mockArticleService := new(mocks.ArticleService)

		// Fake new article
		var fakeArticle entity.Article
		_ = faker.FakeData(&fakeArticle)

		fakeError := []serviceErrors.Error{serviceErrors.E10001}

		// Given data for Article Service Mock
		mockArticleService.On("Store", mock.Anything).Return(nil, fakeError)

		// Construct the handler
		articleHandler := NewArticleHandler(mockArticleService)

		// Create the request
		e := echo.New()
		req := httptest.NewRequest(http.MethodPost, "/articles",
			strings.NewReader(`{"content":"`+fakeArticle.Content+`", "title":"`+fakeArticle.Title+`"}`))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)
		c.SetPath("/articles")

		// Call function
		_ = articleHandler.Store(c)

		// Assert the result
		assert.Equal(t, http.StatusBadRequest, rec.Code)
		assert.Contains(t, rec.Body.String(), fakeError[0].ToJsonError())
	})
}
