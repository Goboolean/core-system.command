package adapter

import (
	grpcapi "github.com/Goboolean/command-server/api/grpc"
	"sync"
)

type MetadataCommandAdapter struct {
	grpcapi.UnimplementedMetadataServiceServer
}

var (
	instance *MetadataCommandAdapter
	once     sync.Once
)

func New() *MetadataCommandAdapter {
	once.Do(func() {
		instance = &MetadataCommandAdapter{}
	})

	return instance
}
