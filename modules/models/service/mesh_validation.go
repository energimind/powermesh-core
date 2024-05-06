package service

import (
	"github.com/energimind/powermesh-core/errorz"
	"github.com/energimind/powermesh-core/modules/models"
)

func validateModelID(id string) error {
	return requireString(id, "model id")
}

func validateNodeID(id string) error {
	return requireString(id, "node id")
}

func validateRelationID(id string) error {
	return requireString(id, "relation id")
}

func validateNodeData(data models.NodeData) error {
	if err := validateKind(data.Kind); err != nil {
		return err
	}

	if err := validatePropBag(data.Props); err != nil {
		return err
	}

	return nil
}

func validateRelationData(data models.RelationData) error {
	if err := validateKind(data.Kind); err != nil {
		return err
	}

	if err := validateRelationSource(data.From); err != nil {
		return err
	}

	if err := validateRelationTarget(data.To); err != nil {
		return err
	}

	if err := validatePropBag(data.Props); err != nil {
		return err
	}

	return nil
}

func validateKind(kind string) error {
	return requireString(kind, "kind")
}

func validateRelationSource(from string) error {
	return requireString(from, "relation source node")
}

func validateRelationTarget(to string) error {
	return requireString(to, "relation target node")
}

func validatePropBag(bag models.PropBag) error {
	for k, v := range bag {
		if k == "" {
			return errorz.NewValidationError("property bag key is required")
		}

		if err := validatePropSection(v); err != nil {
			return err
		}
	}

	return nil
}

func validatePropSection(section models.PropSection) error {
	for k := range section {
		if k == "" {
			return errorz.NewValidationError("property section key is required")
		}
	}

	return nil
}
