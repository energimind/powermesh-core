package mongo

import "github.com/energimind/powermesh-core/services/models"

var (
	validModelModel = models.Model{
		ID:   "model-id",
		Code: "model-code",
		Name: "model-name",
	}
	validStoreModel = storeModel{
		ID:   validModelModel.ID,
		Code: validModelModel.Code,
		Name: validModelModel.Name,
	}
)
