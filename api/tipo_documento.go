package api

import (
	"net/http"
	"strconv"

	"fake.com/medium/dto"
	repo "fake.com/medium/repositories"
	"fake.com/medium/validators"
	"github.com/gin-gonic/gin"
)

type TipoDocumentoApi struct {
	TipoDocumentoRepository repo.TipoDocumentoRepository
}

func NewTipoDocumentoAPI(repo repo.TipoDocumentoRepository) TipoDocumentoApi {
	return TipoDocumentoApi{TipoDocumentoRepository: repo}
}

func (api *TipoDocumentoApi) FindAll(c *gin.Context) {
	lista := api.TipoDocumentoRepository.All()
	c.JSON(http.StatusOK, dto.ToTiposDocumentoDTOs(lista))
}

func (api *TipoDocumentoApi) FindByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tipoDocumento, err := api.TipoDocumentoRepository.FindByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if tipoDocumento.Id == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	c.JSON(http.StatusOK, dto.ToTipoDocumentoDTO(tipoDocumento))
}

func (api *TipoDocumentoApi) Save(c *gin.Context) {
	var input dto.TipoDocumentoInput
	err := validators.ValidateInput(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tipoDocumento := dto.ToTipoDocumento(input)
	tipoDocumento, err = api.TipoDocumentoRepository.Save(tipoDocumento)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.ToTipoDocumentoDTO(tipoDocumento))
}

func (api *TipoDocumentoApi) Update(c *gin.Context) {
	var input dto.TipoDocumentoInput
	err := validators.ValidateInput(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	tipoDocumento := dto.ToTipoDocumento(input)
	tipoDocumento, err = api.TipoDocumentoRepository.Update(tipoDocumento)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.ToTipoDocumentoDTO(tipoDocumento))
}

func (api *TipoDocumentoApi) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	tipo, err := api.TipoDocumentoRepository.FindByID(id)
	if tipo.Id == 0 {
		c.Status(http.StatusNotFound)
		return
	}
	api.TipoDocumentoRepository.Delete(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "register delete"})
}
