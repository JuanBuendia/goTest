package api

import (
	"net/http"

	"fake.com/medium/dto"
	repo "fake.com/medium/repositories"
	"github.com/gin-gonic/gin"
)

type CategoriaApi struct {
	CategoriaRepository repo.CategoriaRepository
}

func NewCategoriaAPI(repo repo.CategoriaRepository) CategoriaApi {
	return CategoriaApi{CategoriaRepository: repo}
}

func (api *CategoriaApi) FindAll(c *gin.Context) {
	lista := api.CategoriaRepository.All()
	c.JSON(http.StatusOK, dto.ToCategoriaDTOs(lista))
}
