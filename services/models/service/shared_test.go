package service

import (
	"strconv"
	"sync/atomic"

	"github.com/energimind/powermesh-core/access"
)

var (
	adminActor   = access.Actor{Role: access.RoleAdmin}
	validModelID = "1"
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
