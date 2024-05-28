package unitmock

import "errors"

type Service struct {
	repository Repository
}

func (s *Service) GetById(id int) (result *Category, err error) {
	category := s.repository.FindById(id)
	if category == nil {
		return nil, errors.New("Category Not Found")
	}

	return category, nil
}
