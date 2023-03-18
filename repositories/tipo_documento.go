package repositories

import (
	"fake.com/medium/models"
	"gorm.io/gorm"
)

type TipoDocumentoRepository interface {
	All() (tiposDocumento []models.TipoDocumento)
	Save(tipoDocumento models.TipoDocumento) (models.TipoDocumento, error)
	FindByID(id int) (models.TipoDocumento, error)
	Update(tipoDocumento models.TipoDocumento) (models.TipoDocumento, error)
	Delete(id int) error
}

type tipoDocumentoRepository struct {
	Conn *gorm.DB
}

func NewTipoDocumentoRepository(db *gorm.DB) tipoDocumentoRepository {
	return tipoDocumentoRepository{Conn: db}
}

func (r tipoDocumentoRepository) All() (tiposDocumento []models.TipoDocumento) {
	r.Conn.Table("tipo_documento").Find(&tiposDocumento)
	return tiposDocumento
}

func (r tipoDocumentoRepository) Save(tipoDocumento models.TipoDocumento) (models.TipoDocumento, error) {
	err := r.Conn.Table("tipo_documento").Create(&tipoDocumento).Error
	return tipoDocumento, err
}

func (r tipoDocumentoRepository) FindByID(id int) (models.TipoDocumento, error) {
	var tipoDocumento models.TipoDocumento
	err := r.Conn.Table("tipo_documento").Where("id=?", id).Find(&tipoDocumento).Error
	return tipoDocumento, err
}

func (r tipoDocumentoRepository) Update(tipoDocumento models.TipoDocumento) (models.TipoDocumento, error) {
	err := r.Conn.Table("tipo_documento").Where("id = ?", tipoDocumento.Id).Updates(&tipoDocumento).Error
	return tipoDocumento, err
}

func (r tipoDocumentoRepository) Delete(id int) error {
	return r.Conn.Table("tipo_documento").Where("id=?", id).Delete(models.TipoDocumento{}).Error
}
