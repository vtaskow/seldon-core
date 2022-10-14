package filters

import (
	"github.com/seldonio/seldon-core/scheduler/pkg/store"
)

type ReplicaFilter interface {
	Name() string
	Filter(model *store.ModelVersion, replica *store.ServerReplica) bool
	Description(model *store.ModelVersion, replica *store.ServerReplica) string
}

type ServerFilter interface {
	Name() string
	Filter(model *store.ModelVersion, server *store.ServerSnapshot) bool
	Description(model *store.ModelVersion, server *store.ServerSnapshot) string
}
