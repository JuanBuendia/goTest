package dto

import (
	"fake.com/medium/models"
)

func ToCategoria(input CategoriaInput) models.Categoria {
	return models.Categoria{
		Id:          input.Id,
		Nombre:      input.Nombre,
		Descripcion: input.Nombre,
	}
}

func ToCategoriaDTO(model models.Categoria) CategoriaInput {

	return CategoriaInput{
		Id:          model.Id,
		Nombre:      model.Nombre,
		Descripcion: model.Descripcion,
	}
}

func ToCategoriaDTOs(categorias []models.Categoria) []CategoriaInput {
	categoriasDTO := make([]CategoriaInput, len(categorias))
	for i, itm := range categorias {
		categoriasDTO[i] = ToCategoriaDTO(itm)
	}
	return categoriasDTO
}
