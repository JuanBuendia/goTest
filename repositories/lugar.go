package repositories

import (
	"fake.com/medium/models"
	"gorm.io/gorm"
)

type LugarRepository interface {
	GetDepartamentos() []models.Lugar
	GetMunicipiosDepartamentoById(id int) []models.Lugar
	GetMunicipiosDepartamentoByIgac(codigoIgac string) []models.Lugar
	FindMunicipioByNombre(nombre string) (lugares []models.Lugar)
	Save(lugar models.Lugar) (models.Lugar, error)
	FindByID(id int) (models.Lugar, error)
	Update(lugar models.Lugar) (models.Lugar, error)
	Delete(id int) error
}

type lugarRepository struct {
	Conn *gorm.DB
}

func NewLugarRepository(db *gorm.DB) lugarRepository {
	return lugarRepository{Conn: db}
}

func (r lugarRepository) GetDepartamentos() []models.Lugar {
	var lugares []models.Lugar
	r.Conn.Table("lugares").Where("tipo='D'").Find(&lugares)
	return lugares
}

func (r lugarRepository) GetMunicipiosDepartamentoById(id int) []models.Lugar {
	var lugares []models.Lugar
	r.Conn.Table("lugares").Where("lug_id=?", id).Find(&lugares)
	return lugares
}

func (r lugarRepository) GetMunicipiosDepartamentoByIgac(codigoIgac string) []models.Lugar {
	var lugares []models.Lugar
	sql := `
	select l.id, l.nombre, l.lug_id, l.tipo, l.codigo_igac 
	from lugares l
	where l.lug_id = (select id from lugares l2 where codigo_igac = ?)`
	r.Conn.Raw(sql, codigoIgac).Scan(&lugares)
	return lugares
}

func (r lugarRepository) FindMunicipioByNombre(nombre string) (lugares []models.Lugar) {
	nombre = "%" + nombre + "%"
	sql := `
	select l.id, l.nombre, l.lug_id, l.tipo, l.codigo_igac 
	from lugares l 
	where tipo = 'M'
	and lower(unaccent(l.nombre)) like  lower(unaccent(?))`
	r.Conn.Raw(sql, nombre).Scan(&lugares)
	return lugares
}

func (r lugarRepository) Save(lugar models.Lugar) (models.Lugar, error) {
	err := r.Conn.Table("lugares").Create(&lugar).Error
	return lugar, err
}

func (r lugarRepository) FindByID(id int) (models.Lugar, error) {
	var lugar models.Lugar
	err := r.Conn.Table("lugares").Where("id=?", id).Find(&lugar).Error
	return lugar, err
}

func (r lugarRepository) Update(lugar models.Lugar) (models.Lugar, error) {
	err := r.Conn.Table("lugares").Where("id = ?", lugar.Id).Updates(&lugar).Error
	return lugar, err
}

func (r lugarRepository) Delete(id int) error {
	return r.Conn.Table("lugares").Where("id=?", id).Delete(models.Lugar{}).Error
}
