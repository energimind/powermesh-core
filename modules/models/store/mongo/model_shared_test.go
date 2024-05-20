package mongo

import "github.com/energimind/powermesh-core/modules/models"

var (
	validModelModel = models.Model{
		ID:          "model-id",
		Code:        "model-code",
		Name:        "model-name",
		Description: "model-description",
	}
	validStoreModel = storeModel{
		ID:          validModelModel.ID,
		Code:        validModelModel.Code,
		Name:        validModelModel.Name,
		Description: validModelModel.Description,
	}
)
