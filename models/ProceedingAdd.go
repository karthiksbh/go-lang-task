package models

import "gorm.io/gorm"

type ProceedingAdd struct {
	gorm.Model
	Identifier int    `json:"identifier" gorm:"primaryKey"`
	TypeCode   string `json:"type-code"`
	Name       string `json:"name"`
	OrgName    string `json:"orgname"`
	Address1   string `json:"address-1"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	Postcode   int    `json:"postcode"`
}
