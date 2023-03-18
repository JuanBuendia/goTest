package models

import (
	"time"
)

type Lugar struct {
	Id         int
	Nombre     string
	LugId      int
	Tipo       string
	CodigoIgac string
}

type TipoDocumento struct {
	Id          int
	Nombre      string
	Abreviatura string
}

type Usuario struct {
	Id                int
	TipoDoc           int
	LugId             int
	NumeroDocumento   string
	NumeroContacto    string
	Direccion         string
	Nombres           string
	Apellidos         string
	Email             string
	Username          string
	Activo            bool
	HashedPassword    string
	RolId             int
	Rol               string    `gorm:"<-:false"`
	PasswordChangedAt time.Time `gorm:"<-:false"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
}

type Rol struct {
	Id     int
	Nombre string
}

type Categoria struct {
	Id          int
	Nombre      string
	Descripcion string
}
