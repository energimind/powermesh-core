package service

import (
	"strconv"
	"sync/atomic"

	"github.com/energimind/powermesh-core/access"
	"github.com/energimind/powermesh-core/services/models"
)

var (
	adminActor     = access.Actor{Role: access.RoleAdmin}
	validModelID   = "1"
	validModelData = models.ModelData{
		Code: "code1",
		Name: "name1",
	}
	validModel = models.Model{
		ID:   validModelID,
		Code: validModelData.Code,
		Name: validModelData.Name,
	}
	missingModelID = "missing"
)

type testIDGenerator struct {
	idCounter atomic.Int64
}

// Ensure that the testIDGenerator implements the idGenerator interface.
var _ idGenerator = (*testIDGenerator)(nil)

func newTestIDGenerator() *testIDGenerator {
	return &testIDGenerator{}
}

func (g *testIDGenerator) GenerateID() string {
	return strconv.FormatInt(g.idCounter.Add(1), 10)
}
