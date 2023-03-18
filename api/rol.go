package api

import (
	"net/http"
	"strconv"

	"fake.com/medium/dto"
	repo "fake.com/medium/repositories"
	"github.com/gin-gonic/gin"
)

type RolApi struct {
	RolRepository repo.RolRepository
}

func NewRolAPI(repo repo.RolRepository) RolApi {
	return RolApi{RolRepository: repo}
}

func (api *RolApi) FindAll(c *gin.Context) {
	lista := api.RolRepository.All()
	c.JSON(http.StatusOK, dto.ToRolDTOs(lista))
}

func (api *RolApi) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	rol, err := api.RolRepository.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if rol.Id == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, dto.ToRolDTO(rol))
}
