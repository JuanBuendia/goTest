package api

import (
	"net/http"
	"strconv"

	"fake.com/medium/dto"
	repo "fake.com/medium/repositories"
	"fake.com/medium/validators"
	"github.com/gin-gonic/gin"
)

type LugarApi struct {
	LugarRepository repo.LugarRepository
}

func NewLugarAPI(repo repo.LugarRepository) LugarApi {
	return LugarApi{LugarRepository: repo}
}

func (api *LugarApi) GetDepartamentos(c *gin.Context) {
	departamentos := api.LugarRepository.GetDepartamentos()
	c.JSON(http.StatusOK, dto.ToLugarDTOs(departamentos))
}

func (api *LugarApi) GetMunicipiosDepartamentoById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	lugares := api.LugarRepository.GetMunicipiosDepartamentoById(id)
	c.JSON(http.StatusOK, dto.ToLugarDTOs(lugares))
}

func (api *LugarApi) GetMunicipiosDepartamentoByIgac(c *gin.Context) {
	codigoIgac := c.Param("codigo_igac")
	lugares := api.LugarRepository.GetMunicipiosDepartamentoByIgac(codigoIgac)
	c.JSON(http.StatusOK, dto.ToLugarDTOs(lugares))
}

func (api *LugarApi) FindMunicipioByNombre(c *gin.Context) {
	nombre := c.Param("nombre")
	municipios := api.LugarRepository.FindMunicipioByNombre(nombre)
	c.JSON(http.StatusOK, dto.ToLugarDTOs(municipios))
}

func (api *LugarApi) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	lugar, err := api.LugarRepository.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if lugar.Id == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, dto.ToLugarDTO(lugar))
}

func (api *LugarApi) Save(c *gin.Context) {
	var input dto.LugarInput
	err := validators.ValidateInput(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	lugar := dto.ToLugar(input)
	lugar, err = api.LugarRepository.Save(lugar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.ToLugarDTO(lugar))
}

func (api *LugarApi) Update(c *gin.Context) {
	var input dto.LugarInput
	err := validators.ValidateInput(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	lugar := dto.ToLugar(input)
	lugar, err = api.LugarRepository.Update(lugar)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.ToLugarDTO(lugar))
}

func (api *LugarApi) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tipo, err := api.LugarRepository.FindByID(id)
	if tipo.Id == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	api.LugarRepository.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "register delete"})
}
