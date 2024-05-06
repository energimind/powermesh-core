package mongo

import "github.com/energimind/powermesh-core/modules/models"

func toStoreModel(m models.Model) storeModel {
	return storeModel{
		ID:   m.ID,
		Code: m.Code,
		Name: m.Name,
	}
}

func fromStoreModel(m storeModel) models.Model {
	return models.Model{
		ID:   m.ID,
		Code: m.Code,
		Name: m.Name,
	}
}
