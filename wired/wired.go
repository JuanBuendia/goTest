package wired

import (
	"time"

	"fake.com/medium/api"
	"fake.com/medium/policies"
	repo "fake.com/medium/repositories"
	"fake.com/medium/services"
	"fake.com/medium/tokenauth"
	"gorm.io/gorm"
)

func InitLugarAPI(db *gorm.DB) api.LugarApi {
	lugarRepository := repo.NewLugarRepository(db)
	lugarApi := api.NewLugarAPI(lugarRepository)
	return lugarApi
}

func InitRolAPI(db *gorm.DB) api.RolApi {
	rolRepository := repo.NewRolRepository(db)
	rolApi := api.NewRolAPI(rolRepository)
	return rolApi
}

func InitTipoDocumentoAPI(db *gorm.DB) api.TipoDocumentoApi {
	tipoDocumentoRepository := repo.NewTipoDocumentoRepository(db)
	tipoDocumentoApi := api.NewTipoDocumentoAPI(tipoDocumentoRepository)
	return tipoDocumentoApi
}

func InitUserAPI(db *gorm.DB, tokenMaker tokenauth.Maker, duration time.Duration) api.UserAPI {
	usuarioRepository := repo.NewUsuarioRepository(db)
	usuarioService := services.NewUsuarioService(usuarioRepository, tokenMaker, duration)
	userApi := api.NewUserAPI(usuarioService)
	return userApi
}

func InitAuthorized(db *gorm.DB) policies.Authorized {
	authorized := policies.NewAuthorized(db)
	return &authorized
}

func InitCategoriaAPI(db *gorm.DB) api.CategoriaApi {
	categoriaRepository := repo.NewCategoriaRepository(db)
	categoriaApi := api.NewCategoriaAPI(categoriaRepository)
	return categoriaApi
}
