package utils

import (
	"github.com/bwmarrin/snowflake"
)

type UidGenerator interface {
	GenerateUID() int64
}

type uidGenerator struct {
	node *snowflake.Node
}

func NewUidGenerator(machineID int64) (UidGenerator, error) {
	node, err := snowflake.NewNode(machineID)
	if err != nil {
		return nil, err
	}
	return &uidGenerator{node: node}, nil
}

func (s *uidGenerator) GenerateUID() int64 {
	// Generate a new snowflake ID
	id := s.node.Generate()
	// Convert the ID to an int64
	return id.Int64()
}
