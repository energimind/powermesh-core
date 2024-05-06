package service

import "github.com/energimind/powermesh-core/modules/models"

func modelFromData(id string, data models.ModelData) models.Model {
	return models.Model{
		ID:   id,
		Code: data.Code,
		Name: data.Name,
	}
}
