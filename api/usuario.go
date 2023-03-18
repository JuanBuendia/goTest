package api

import (
	"fmt"
	"net/http"
	"strconv"

	"fake.com/medium/dto"
	"fake.com/medium/response"
	service "fake.com/medium/services"
	"fake.com/medium/validators"
	"github.com/gin-gonic/gin"
)

type UserAPI struct {
	usuarioService service.UsuarioService
}

func NewUserAPI(service service.UsuarioService) UserAPI {
	return UserAPI{
		usuarioService: service,
	}
}

func (api *UserAPI) FindByDocumento(c *gin.Context) {
	documento := c.Param("documento")
	usuario, err := api.usuarioService.FindByDocumento(documento)
	if err != nil {
		fmt.Println(err.Error)
		return
	}
	if usuario.NumeroDocumento != documento {
		c.JSON(http.StatusNotFound, gin.H{"error": "No data Found"})
		return
	}
	c.JSON(http.StatusOK, dto.ToUsuarioDTO(usuario))
}

func (api *UserAPI) FindUser(c *gin.Context) {
	word := c.Param("word")
	usuarios, err := api.usuarioService.FindUser(word)
	if err != nil {
		fmt.Println(err.Error)
		return
	}
	if len(usuarios) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No data Found"})
		return
	}
	c.JSON(http.StatusOK, dto.ToUsuarioDTOs(usuarios))
}

func (api *UserAPI) GetTecnicosAndJuridicos(c *gin.Context) {
	usuarios, err := api.usuarioService.GetTecnicosAndJuridicos()
	if err != nil {
		fmt.Println(err.Error)
		return
	}
	if len(usuarios) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "No data Found"})
		return
	}
	c.JSON(http.StatusOK, dto.ToUsuarioDTOs(usuarios))
}

func (api *UserAPI) RegistrarSolcitante(c *gin.Context) {
	var input dto.UsuarioInput
	err := validators.ValidateInput(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usuario := dto.ToUsuario(input)
	password := input.Password
	usuario, err = api.usuarioService.RegistrarSolcitante(usuario, password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"estado": "Se ha registrado el usuario con éxito"})
}

func (api *UserAPI) Update(c *gin.Context) {
	var input dto.UsuarioUpdateInput
	err := validators.ValidateInput(c, &input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	usuario := dto.UsuarioUpdateToUsuario(input)
	usuario, err = api.usuarioService.Update(usuario)
	c.JSON(http.StatusOK, dto.ToUsuarioDTO(usuario))
}

func (api *UserAPI) LoginUser(c *gin.Context) {
	var req dto.LoginUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse(err))
		return
	}
	rsp, err := api.usuarioService.LoginUser(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, rsp)
}

func (api *UserAPI) FindAll(c *gin.Context) {
	lista, err := api.usuarioService.FindAll()
	if err != nil {
		fmt.Println(err.Error)
	}
	c.JSON(http.StatusOK, dto.ToUsuarioDTOs(lista))
}

func (api *UserAPI) FindById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	usuario, err := api.usuarioService.FindById(id)
	if err != nil {
		fmt.Println(err.Error)
		return
	}
	c.JSON(http.StatusOK, dto.ToUsuarioDTO(usuario))
}
