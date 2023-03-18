package dto

import (
	"fake.com/medium/models"
)

func ToLugar(input LugarInput) models.Lugar {
	return models.Lugar{
		Id:         input.Id,
		Nombre:     input.Nombre,
		LugId:      input.LugId,
		Tipo:       input.Tipo,
		CodigoIgac: input.CodigoIgac,
	}
}

func ToLugarDTO(lugar models.Lugar) LugarInput {
	return LugarInput{
		Id:         lugar.Id,
		Nombre:     lugar.Nombre,
		LugId:      lugar.LugId,
		Tipo:       lugar.Tipo,
		CodigoIgac: lugar.CodigoIgac,
	}
}

func ToLugarDTOs(lugares []models.Lugar) []LugarInput {
	lugaresDTO := make([]LugarInput, len(lugares))
	for i, itm := range lugares {
		lugaresDTO[i] = ToLugarDTO(itm)
	}
	return lugaresDTO
}
