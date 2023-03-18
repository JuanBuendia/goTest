package repositories

import (
	"fake.com/medium/models"
	"gorm.io/gorm"
)

type CategoriaRepository interface {
	All() (categorias []models.Categoria)
}

type cateriaRepository struct {
	Conn *gorm.DB
}

func NewCategoriaRepository(db *gorm.DB) cateriaRepository {
	return cateriaRepository{Conn: db}
}

func (r cateriaRepository) All() (categorias []models.Categoria) {
	r.Conn.Table("categorias").Find(&categorias)
	return categorias
}
