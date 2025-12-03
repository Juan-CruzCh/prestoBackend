package service

import "prestoBackend/src/module/lectura/repository"

type LecturaService struct {
	Repository repository.LecturaRepository
}

func NewLecturaService(repo repository.LecturaRepository) *LecturaService {
	return &LecturaService{
		Repository: repo,
	}
}
func (s *LecturaService) ListarLectura() {

}
