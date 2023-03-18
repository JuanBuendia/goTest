package dto

import (
	"fake.com/medium/models"
)

func ToRol(input RolInput) models.Rol {
	return models.Rol{
		Id:     input.Id,
		Nombre: input.Nombre,
	}
}

func ToRolDTO(Rol models.Rol) RolInput {

	return RolInput{
		Id:     Rol.Id,
		Nombre: Rol.Nombre,
	}
}

func ToRolDTOs(Rols []models.Rol) []RolInput {
	RolsDTO := make([]RolInput, len(Rols))
	for i, itm := range Rols {
		RolsDTO[i] = ToRolDTO(itm)
	}
	return RolsDTO
}
