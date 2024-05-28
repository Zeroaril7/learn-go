package unitmock

import "github.com/stretchr/testify/mock"

type Repository interface {
	FindById(id int) *Category
}

type RepositoryMock struct {
	Mock mock.Mock
}

func (r *RepositoryMock) FindById(id int) *Category {
	arguments := r.Mock.Called(id)
	if arguments.Get(0) == nil {
		return nil
	}

	result := arguments.Get(0).(Category)
	return &result
}
