package model

import "gorm.io/gorm"

type ModelEndpoint struct {
	gorm.Model
	Name string  `gorm:"indexUnique;not null"`
	Desc *string `gorm:"default:null"`

	ModelRouting        *ModelRouting
	ModelRoutingTargets []*ModelRoutingTarget `gorm:"polymorphic:Target;"`
}

func (ModelEndpoint) TableName() string {
	return "model_endpoints"
}
