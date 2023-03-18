package routes

import (
	"fake.com/medium/api"
	"github.com/gin-gonic/gin"
)

func AuthHandlers(g gin.IRoutes, api api.UserAPI) {
	g.POST("/users/login", api.LoginUser)
}

func TipoTipoDocumentoHandlers(g gin.IRoutes, api api.TipoDocumentoApi) {
	g.GET("/tipos_documento", api.FindAll)
	g.GET("/tipos_documento/:id", api.FindByID)
	g.POST("/tipos_documento", api.Save)
	g.PUT("/tipos_documento", api.Update)
	g.DELETE("/tipos_documento/:id", api.Delete)
}

func UsersHandlers(g gin.IRoutes, rg *gin.RouterGroup, api api.UserAPI) {
	g.POST("/registro_solicitante", api.RegistrarSolcitante)
	g.PUT("/usuarios", api.Update)
	g.GET("/usuarios/buscar/:word", api.FindUser)
	g.GET("/usuarios/buscar_documento/:documento", api.FindByDocumento)
	rg.GET("/usuarios/tecnicos_juridicos", api.GetTecnicosAndJuridicos)
	rg.GET("/usuarios", api.FindAll)
	g.GET("/usuarios/:id", api.FindById)

	//	g.GET("/usuarios", api.FindAll)
	//g.GET("/usuarios/user_name/:user_name", api.FindByUserName)

}

func LugaresHandlers(g gin.IRoutes, api api.LugarApi) {
	g.DELETE("/lugares", api.Delete)
	g.POST("/lugares", api.Save)
	g.GET("/lugares/:id", api.FindByID)
	g.GET("/departamentos", api.GetDepartamentos)
	g.GET("/departamentos/:id/municipios", api.GetMunicipiosDepartamentoById)
	g.GET("/departamentos_igac/:codigo_igac/municipios", api.GetMunicipiosDepartamentoByIgac)
	g.GET("/municipios/:nombre", api.FindMunicipioByNombre)
}

func RolHandlers(g gin.IRoutes, api api.RolApi) {
	g.GET("/roles", api.FindAll)
	g.GET("/roles/:id", api.FindByID)

}

func CategoriaHandlers(g gin.IRoutes, api api.CategoriaApi) {
	g.GET("/categorias", api.FindAll)

}
