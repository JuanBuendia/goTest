package dto

import (
	"time"
)

type LugarInput struct {
	Id         int    `json:"id" `
	Nombre     string `json:"nombre" `
	LugId      int    `json:"lugar_id" `
	Tipo       string `json:"tipo" `
	CodigoIgac string `json:"codigo_igac" `
}

type RolInput struct {
	Id     int    `json:"id" `
	Nombre string `json:"nombre" `
}

type TipoDocumentoInput struct {
	Id          int    `json:"id" `
	Nombre      string `json:"nombre"`
	Abreviatura string `json:"abreviatura" `
}

type UsuarioInput struct {
	Id                int       `json:"id" gorm:"primaryKey;unique;not null"`
	TipoDoc           int       `json:"tipo_doc"  binding:"required"`
	LugId             int       `json:"lugar_id" `
	RolId             int       `json:"rol_id"`
	NumeroDocumento   string    `json:"numero_documento" binding:"required"`
	NumeroContacto    string    `json:"numero_contacto"  binding:"required"`
	Direccion         string    `json:"direccion"        binding:"required"`
	Nombres           string    `json:"nombres" `
	Apellidos         string    `json:"apellidos" `
	Email             string    `json:"email"    binding:"required,email"`
	Username          string    `json:"username" binding:"required"`
	Activo            bool      `json:"activo"`
	HashedPassword    string    `json:"hashed_password"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
	Password          string    `json:"password" binding:"required,min=6"`
}

type UsuarioResponse struct {
	UsuarioInput
	Municipio           string `json:"municipio"`
	Departamento        string `json:"departamento"`
	TipoDocumentoNombre string `json:"tipo_documento_nombre"`
}

type UsuarioUpdateInput struct {
	Id              int       `json:"id" gorm:"primaryKey;unique;not null"`
	TipoDoc         int       `json:"tipo_doc"  binding:"required"`
	LugId           int       `json:"lugar_id" `
	RolId           int       `json:"rol_id"`
	NumeroDocumento string    `json:"numero_documento" binding:"required"`
	NumeroContacto  string    `json:"numero_contacto"  binding:"required"`
	Direccion       string    `json:"direccion"        binding:"required"`
	Nombres         string    `json:"nombres" `
	Apellidos       string    `json:"apellidos" `
	Email           string    `json:"email"    binding:"required,email"`
	Activo          bool      `json:"activo"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type HablemosInput struct {
	Nombres   string `json:"nombres" binding:"required"`
	Seudonimo string `json:"seudonimo" `
	Email     string `json:"email" binding:"required,email"`
	Ciudad    string `json:"ciudad" binding:"required"`
	Tema      string `json:"tema" binding:"required"`
	Mensaje   string `json:"mensaje" binding:"required"`
}

/*login*/
type LoginUserRequest struct {
	Username string `json:"username" binding:"required,alphanum"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginUserResponse struct {
	AccessToken string       `json:"access_token"`
	User        UserResponse `json:"user"`
}

type UserResponse struct {
	Id                int       `json:"id"`
	Username          string    `json:"username"`
	Nombres           string    `json:"nombres"`
	Apellidos         string    `json:"apellidos"`
	Email             string    `json:"email"`
	Rol               string    `json:"rol"`
	PasswordChangedAt time.Time `json:"password_changed_at"`
	CreatedAt         time.Time `json:"created_at"`
}

type CategoriaInput struct {
	Id          int    `json:"id"`
	Nombre      string `json:"nombre"`
	Descripcion string `json:"descripcion"`
}
