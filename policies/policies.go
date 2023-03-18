package policies

import (
	"fmt"

	"fake.com/medium/models"
	"gorm.io/gorm"
)

type Authorized interface {
	Me(username string) models.Usuario
}

type authorized struct {
	Conn *gorm.DB
}

func NewAuthorized(db *gorm.DB) authorized {
	return authorized{Conn: db}

}

func (r *authorized) Me(username string) models.Usuario {
	var usuario models.Usuario
	err := r.Conn.Raw(`
				select u.id, u.tipo_doc, u.lug_id, u.rol_id, u.numero_documento, u.numero_contacto, u.direccion, u.nombres, u.apellidos, u.email, u.username, u.hashed_password, u.password_changed_at, u.created_at, u.activo, u.updated_at , r.nombre   as rol
				from usuarios u
				inner join roles r    on    r.id =  u.rol_id
				where u.username = ?
	`, username).Scan(&usuario).Error
	if err != nil {
		fmt.Println(err.Error)
	}

	return usuario
}
