package dto

import (
	"fake.com/medium/models"
)

func ToTipoDocumento(input TipoDocumentoInput) models.TipoDocumento {
	return models.TipoDocumento{
		Id:          input.Id,
		Nombre:      input.Nombre,
		Abreviatura: input.Abreviatura,
	}
}

func ToTipoDocumentoDTO(tipoDocumento models.TipoDocumento) TipoDocumentoInput {

	return TipoDocumentoInput{
		Id:          tipoDocumento.Id,
		Nombre:      tipoDocumento.Nombre,
		Abreviatura: tipoDocumento.Abreviatura,
	}
}

func ToTiposDocumentoDTOs(tiposDocumento []models.TipoDocumento) []TipoDocumentoInput {
	tipoDocumentosDTO := make([]TipoDocumentoInput, len(tiposDocumento))
	for i, itm := range tiposDocumento {
		tipoDocumentosDTO[i] = ToTipoDocumentoDTO(itm)
	}
	return tipoDocumentosDTO
}
