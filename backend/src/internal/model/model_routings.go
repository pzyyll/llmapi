package model

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type ModelRouting struct {
	gorm.Model
	ModelEndpointID     uint `gorm:"index"`
	Strategy            uint `gorm:"default:0"`
	OverrideParams      datatypes.JSONType[any]
	ModelRoutingTargets []*ModelRoutingTarget `gorm:"OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (ModelRouting) TableName() string {
	return "model_routings"
}

type ModelRoutingTarget struct {
	gorm.Model
	Weight         uint `gorm:"default:0"`
	Priority       uint `gorm:"default:0"`
	OverrideParams datatypes.JSONType[any]
	IsFallback     bool `gorm:"default:false"`
	ModelRoutingID uint `gorm:"index"`

	TargetID   uint   `gorm:"index"`
	TargetType string `gorm:"index"`
	Target     any    `gorm:"-"`
}

func (ModelRoutingTarget) TableName() string {
	return "model_routing_targets"
}
