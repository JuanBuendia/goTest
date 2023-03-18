package repositories

import (
	"fake.com/medium/models"
	"gorm.io/gorm"
)

type RolRepository interface {
	All() (categorias []models.Rol)
	FindByID(id int) (models.Rol, error)
}

type rolRepository struct {
	Conn *gorm.DB
}

func NewRolRepository(db *gorm.DB) rolRepository {
	return rolRepository{Conn: db}
}

func (r rolRepository) All() (roles []models.Rol) {
	r.Conn.Table("roles").Find(&roles)
	return roles
}

func (r rolRepository) FindByID(id int) (models.Rol, error) {
	var rol models.Rol
	err := r.Conn.Table("roles").Where("id=?", id).Find(&rol).Error
	return rol, err
}
