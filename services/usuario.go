package services

import (
	"errors"
	"fmt"
	"time"

	dto "fake.com/medium/dto"
	"fake.com/medium/models"
	repo "fake.com/medium/repositories"
	"fake.com/medium/tokenauth"
	"fake.com/medium/utils"
)

const (
	ADMIN       = 1
	SOLICITANTE = 2
	JURIDICA    = 3
	TECNICA     = 4
	INFORMATIVO = 5
)

type UsuarioService interface {
	RegistrarSolcitante(usuario models.Usuario, password string) (models.Usuario, error)
	FindByUserName(username string) (models.Usuario, error)
	LoginUser(req dto.LoginUserRequest) (dto.LoginUserResponse, error)
	FindUser(word string) ([]models.Usuario, error)
	FindAll() ([]models.Usuario, error)
	FindByDocumento(documento string) (models.Usuario, error)
	FindById(id int) (models.Usuario, error)
	GetTecnicosAndJuridicos() ([]models.Usuario, error)
	Update(user models.Usuario) (models.Usuario, error)
}

type usuarioService struct {
	usuarioRepository   repo.UsuarioRepository
	tokenMaker          tokenauth.Maker
	accessTokenDuration time.Duration
}

func NewUsuarioService(repoUser repo.UsuarioRepository, token tokenauth.Maker, duration time.Duration) *usuarioService {
	return &usuarioService{
		usuarioRepository:   repoUser,
		tokenMaker:          token,
		accessTokenDuration: duration,
	}
}

func NewUsuarioServiceDB(repoUser repo.UsuarioRepository) usuarioService {
	return usuarioService{
		usuarioRepository: repoUser,
	}
}

func (s *usuarioService) GetTecnicosAndJuridicos() ([]models.Usuario, error) {
	return s.usuarioRepository.GetTecnicosAndJuridicos()
}

func (s *usuarioService) FindUser(word string) ([]models.Usuario, error) {
	return s.usuarioRepository.FindUser(word)
}

func (s *usuarioService) FindAll() ([]models.Usuario, error) {
	return s.usuarioRepository.FindAll()
}

func (s *usuarioService) Update(user models.Usuario) (models.Usuario, error) {
	return s.usuarioRepository.Update(user)
}

func (s *usuarioService) RegistrarSolcitante(usuario models.Usuario, password string) (models.Usuario, error) {
	usuarioBD, _ := s.usuarioRepository.UserExists(usuario)
	if usuarioBD.Id > 0 {
		err := errors.New("El usuario se encuentra registrado")
		return models.Usuario{}, err
	}
	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		fmt.Print("Hashed")
		fmt.Print(err.Error)
	}

	usuario.HashedPassword = hashedPassword
	usuario.RolId = SOLICITANTE
	usuario.Activo = true
	usuario, err = s.usuarioRepository.Save(usuario)
	return usuario, err
}

func (s *usuarioService) RegistrarUsuarioPropietario(usuario models.Usuario) (models.Usuario, error) {
	usuarioBD, _ := s.usuarioRepository.FindByDocumento(usuario.NumeroDocumento)
	if usuarioBD.Id > 0 {
		return usuarioBD, nil
	}
	usuario.RolId = INFORMATIVO
	return s.usuarioRepository.Save(usuario)
}

func (s *usuarioService) FindByUserName(username string) (models.Usuario, error) {
	user, err := s.usuarioRepository.FindByUserName(username)
	return user, err
}

func (s *usuarioService) FindByDocumento(documento string) (models.Usuario, error) {
	user, err := s.usuarioRepository.FindByDocumento(documento)
	return user, err
}

func (s *usuarioService) FindById(id int) (models.Usuario, error) {
	user, err := s.usuarioRepository.FindByID(id)
	return user, err
}

func (s *usuarioService) LoginUser(req dto.LoginUserRequest) (dto.LoginUserResponse, error) {
	user, err := s.usuarioRepository.FindByUserName(req.Username)
	if err != nil {
		err := errors.New("Usuario o contraseña no coinciden")
		return dto.LoginUserResponse{}, err
	}
	err = errors.New("Usuario o contraseña no coinciden")
	if (user == models.Usuario{}) {
		return dto.LoginUserResponse{}, err
	}

	error := utils.CheckPassword(req.Password, user.HashedPassword)

	if error != nil {
		return dto.LoginUserResponse{}, err
	}

	accessToken, err := s.tokenMaker.CreateToken(
		user.Username,
		s.accessTokenDuration,
	)

	if err != nil {
		return dto.LoginUserResponse{}, err
	}
	rsp := dto.LoginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	return rsp, nil

}

func newUserResponse(user models.Usuario) dto.UserResponse {
	return dto.UserResponse{
		Id:                user.Id,
		Username:          user.Username,
		Nombres:           user.Nombres,
		Apellidos:         user.Apellidos,
		Email:             user.Email,
		Rol:               user.Rol,
		PasswordChangedAt: user.PasswordChangedAt,
		CreatedAt:         user.CreatedAt,
	}
}
