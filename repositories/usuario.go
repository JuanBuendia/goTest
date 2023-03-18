package repositories

import (
	"time"

	"fake.com/medium/models"
	"gorm.io/gorm"
)

/*
 1	ADMIN
 2	SOLICITANTE
 3	JURIDICA
 4	TECNICA
 5	INFORMATIVO
*/

type UsuarioRepository interface {
	Save(user models.Usuario) (models.Usuario, error)
	FindAll() ([]models.Usuario, error)
	FindByRol(rolId int) (models.Usuario, error)
	FindByUserName(username string) (models.Usuario, error)
	FindUser(word string) ([]models.Usuario, error)
	FindByDocumento(documento string) (models.Usuario, error)
	FindByID(id int) (models.Usuario, error)
	Update(user models.Usuario) (models.Usuario, error)
	UserExists(user models.Usuario) (models.Usuario, error)
	GetTecnicosAndJuridicos() ([]models.Usuario, error)
}

type usuarioRepository struct {
	Conn *gorm.DB
}

func NewUsuarioRepository(db *gorm.DB) usuarioRepository {
	return usuarioRepository{Conn: db}

}

func (r usuarioRepository) Save(user models.Usuario) (models.Usuario, error) {
	user.CreatedAt = time.Now()
	err := r.Conn.Save(&user).Error
	return user, err
}

func (r usuarioRepository) FindAll() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := r.Conn.Find(&usuarios).Error
	return usuarios, err
}

func (r usuarioRepository) UserExists(user models.Usuario) (models.Usuario, error) {
	err := r.Conn.Raw(`
				select u.id, u.email , u.numero_documento , u.numero_contacto , r.nombre as rol
				from usuarios u 
				inner join roles r ON   r.id  = u.rol_id 
				where u.email = ?
				or  u.username = ?
				or u.numero_documento = ?
				limit 1
	`, user.Email, user.Username, user.NumeroDocumento).Scan(&user).Error

	return user, err
}

func (r usuarioRepository) FindByRol(rolId int) (models.Usuario, error) {
	var user = models.Usuario{}
	err := r.Conn.Where("rol_id= ?", rolId).First(&user).Error
	return user, err
}

func (r usuarioRepository) FindByUserName(username string) (models.Usuario, error) {
	var user = models.Usuario{}
	err := r.Conn.Raw(`
			select u.id, u.tipo_doc, u.lug_id, u.rol_id, u.numero_documento, u.numero_contacto, u.direccion, u.nombres, u.apellidos, u.email, u.username, u.hashed_password, u.password_changed_at, u.created_at, u.activo, u.updated_at, r.nombre as rol
			from usuarios u 
			inner join roles r ON   r.id  = u.rol_id 
			where u.username = ?
			limit 1
	`, username).Scan(&user).Error
	return user, err
}

func (r usuarioRepository) FindByID(id int) (models.Usuario, error) {
	var usuario models.Usuario
	err := r.Conn.Table("usuarios").Where("id=?", id).Find(&usuario).Error
	return usuario, err
}

func (r usuarioRepository) Update(user models.Usuario) (models.Usuario, error) {
	user.UpdatedAt = time.Now()
	err := r.Conn.Table("usuarios").Where("id = ?", user.Id).Updates(&user).Error
	return user, err
}

func (r usuarioRepository) FindUser(word string) ([]models.Usuario, error) {
	var usuarios []models.Usuario
	word = "%" + word + "%"
	sql := `
		select id, tipo_doc, lug_id, rol_id, numero_documento, numero_contacto, direccion, nombres, apellidos, email, username, hashed_password, password_changed_at, created_at, activo, updated_at 
		from usuarios
		where lower(unaccent(nombres)) like lower(unaccent(@word))
		or lower(unaccent(apellidos)) like lower(unaccent(@word))
		or numero_documento  like @word
		limit 5
	`
	err := r.Conn.Raw(sql, map[string]interface{}{"word": word}).Scan(&usuarios).Error
	return usuarios, err
}

func (r usuarioRepository) FindByDocumento(documento string) (models.Usuario, error) {
	var usuario models.Usuario
	err := r.Conn.Table("usuarios").Where("numero_documento=?", documento).Find(&usuario).Error
	return usuario, err
}

func (r usuarioRepository) GetTecnicosAndJuridicos() ([]models.Usuario, error) {
	var usuarios []models.Usuario
	err := r.Conn.Raw(`
			select u.id, u.tipo_doc, u.lug_id, u.rol_id, u.numero_documento, u.numero_contacto, u.direccion, u.nombres, u.apellidos, u.email, u.username,  u.password_changed_at, u.created_at, u.activo, u.updated_at, r.nombre as rol
			from usuarios u 
			inner join roles r ON r.id  = u.rol_id 
			where u.rol_id = 3
			or    u.rol_id  = 4
	`).Scan(&usuarios).Error
	return usuarios, err
}
