package service

import (
	"testing"

	"github.com/energimind/powermesh-core/modules/models"
	"github.com/stretchr/testify/require"
)

func Test_validateModelID(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateModelID("1"))
	require.Error(t, validateModelID(""))
}

func Test_validateNodeID(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateNodeID("1"))
	require.Error(t, validateNodeID(""))
}

func Test_validateRelationID(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateRelationID("1"))
	require.Error(t, validateRelationID(""))
}

func Test_validateNodeData(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		data    models.NodeData
		wantErr bool
	}{
		"valid": {
			data: validNodeData,
		},
		"invalid-kind": {
			data: models.NodeData{
				Kind:  "",
				Props: models.PropBag{},
			},
			wantErr: true,
		},
		"invalid-props": {
			data: models.NodeData{
				Kind:  "kind",
				Props: models.PropBag{"": models.PropSection{}},
			},
			wantErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateNodeData(test.data)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_validateRelationData(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		data    models.RelationData
		wantErr bool
	}{
		"valid": {
			data: validRelationData,
		},
		"invalid-kind": {
			data: models.RelationData{
				Kind:  "",
				From:  "n1",
				To:    "n2",
				Props: models.PropBag{},
			},
			wantErr: true,
		},
		"invalid-from": {
			data: models.RelationData{
				Kind:  "kind",
				From:  "",
				To:    "n2",
				Props: models.PropBag{},
			},
			wantErr: true,
		},
		"invalid-to": {
			data: models.RelationData{
				Kind:  "kind",
				From:  "n1",
				To:    "",
				Props: models.PropBag{},
			},
			wantErr: true,
		},
		"invalid-props": {
			data: models.RelationData{
				Kind:  "kind",
				From:  "n1",
				To:    "n2",
				Props: models.PropBag{"": models.PropSection{}},
			},
			wantErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validateRelationData(test.data)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_validateKind(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateKind("kind"))
	require.Error(t, validateKind(""))
}

func Test_validateRelationSource(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateRelationSource("n1"))
	require.Error(t, validateRelationSource(""))
}

func Test_validateRelationTarget(t *testing.T) {
	t.Parallel()

	require.NoError(t, validateRelationTarget("n2"))
	require.Error(t, validateRelationTarget(""))
}

func Test_validatePropBag(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		bag     models.PropBag
		wantErr bool
	}{
		"valid": {
			bag: models.PropBag{
				"key": models.PropSection{},
			},
		},
		"empty": {
			bag: models.PropBag{},
		},
		"invalid-key": {
			bag: models.PropBag{
				"": models.PropSection{},
			},
			wantErr: true,
		},
		"invalid-section": {
			bag: models.PropBag{
				"key": models.PropSection{
					"": "value",
				},
			},
			wantErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validatePropBag(test.bag)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func Test_validatePropSection(t *testing.T) {
	t.Parallel()

	tests := map[string]struct {
		section models.PropSection
		wantErr bool
	}{
		"valid": {
			section: models.PropSection{
				"key": "value",
			},
		},
		"empty": {
			section: models.PropSection{},
		},
		"invalid-key": {
			section: models.PropSection{
				"": "value",
			},
			wantErr: true,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			err := validatePropSection(test.section)

			if test.wantErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
