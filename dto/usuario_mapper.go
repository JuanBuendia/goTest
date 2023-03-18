package dto

import (
	"fake.com/medium/models"
)

func ToUsuario(input UsuarioInput) models.Usuario {
	return models.Usuario{
		Id:                input.Id,
		TipoDoc:           input.TipoDoc,
		LugId:             input.LugId,
		RolId:             input.RolId,
		NumeroDocumento:   input.NumeroDocumento,
		NumeroContacto:    input.NumeroContacto,
		Direccion:         input.Direccion,
		Nombres:           input.Nombres,
		Apellidos:         input.Apellidos,
		Email:             input.Email,
		Username:          input.Username,
		Activo:            input.Activo,
		HashedPassword:    input.HashedPassword,
		PasswordChangedAt: input.PasswordChangedAt,
		CreatedAt:         input.CreatedAt,
		UpdatedAt:         input.UpdatedAt,
	}
}

func UsuarioUpdateToUsuario(input UsuarioUpdateInput) models.Usuario {
	return models.Usuario{
		Id:              input.Id,
		TipoDoc:         input.TipoDoc,
		LugId:           input.LugId,
		RolId:           input.RolId,
		NumeroDocumento: input.NumeroDocumento,
		NumeroContacto:  input.NumeroContacto,
		Direccion:       input.Direccion,
		Nombres:         input.Nombres,
		Apellidos:       input.Apellidos,
		Email:           input.Email,
		Activo:          input.Activo,
		UpdatedAt:       input.UpdatedAt,
	}
}

func ToUsuarioDTO(usuario models.Usuario) UsuarioInput {
	return UsuarioInput{
		Id:                usuario.Id,
		TipoDoc:           usuario.TipoDoc,
		LugId:             usuario.LugId,
		RolId:             usuario.RolId,
		NumeroDocumento:   usuario.NumeroDocumento,
		NumeroContacto:    usuario.NumeroContacto,
		Direccion:         usuario.Direccion,
		Nombres:           usuario.Nombres,
		Apellidos:         usuario.Apellidos,
		Email:             usuario.Email,
		Username:          usuario.Username,
		Activo:            usuario.Activo,
		PasswordChangedAt: usuario.PasswordChangedAt,
		CreatedAt:         usuario.CreatedAt,
		UpdatedAt:         usuario.UpdatedAt,
	}
}

func ToUsuarioDTOs(usuarios []models.Usuario) []UsuarioInput {
	usuariosdto := make([]UsuarioInput, len(usuarios))
	for i, itm := range usuarios {
		usuariosdto[i] = ToUsuarioDTO(itm)
	}
	return usuariosdto
}
