package unitmock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var repo = &RepositoryMock{Mock: mock.Mock{}}
var service = Service{repository: repo}

func TestGetCategoryNotFound(t *testing.T) {
	repo.Mock.On("FindById", 1).Return(nil)

	category, err := service.GetById(1)

	assert.NotNil(t, err)
	assert.Nil(t, category)
}

func TestGetCategorySucess(t *testing.T) {

	category := Category{
		Id:   1,
		Name: "Income",
	}

	repo.Mock.On("FindById", 1).Return(category)

	result, err := service.GetById(1)

	assert.NotNil(t, result)
	assert.Nil(t, err)
	assert.Equal(t, result, &category)
}
