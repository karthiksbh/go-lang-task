package models

import (
	"gorm.io/gorm"
)

type ProceedingAdd struct {
	gorm.Model
	Identifier int    `json:"identifier" gorm:"uniqueIndex"`
	TypeCode   string `json:"type-code"`
	Name       string `json:"name"`
	OrgName    string `json:"orgname"`
	Address1   string `json:"address-1"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	Postcode   int    `json:"postcode"`
}

type AddressInformation struct {
	gorm.Model
	ProceedingAddID uint          `json:"proceeding-add-id"`
	ProceedingAdd   ProceedingAdd `json:"proceeding-address" gorm:"foreignKey:ProceedingAddID"`
}

type Property_vals struct {
	gorm.Model
	Identifier    int    `json:"identifier" gorm:"uniqueIndex"`
	Serial_Number int    `json:"serial-number"`
	Mark_Text     string `json:"mark-text"`
}

// Contains property information for party values
type Property_Information struct {
	gorm.Model
	Property_valsID uint          `json:"property-vals-id"`
	Property_vals   Property_vals `json:"property" gorm:"foreignKey:Property_valsID"`
}

// Contains party values
type Party_vals struct {
	gorm.Model
	Identifier             int                  `json:"identifier" gorm:"uniqueIndex"`
	Role_Code              string               `json:"role-code"`
	Name                   string               `json:"name"`
	Property_InformationID uint                 `json:"property-information-id"`
	Property_Information   Property_Information `json:"property-information" gorm:"foreignKey:Property_InformationID"`
	AddressInformationID   uint                 `json:"address-information-id"`
	AddressInformation     AddressInformation   `json:"address-information" gorm:"foreignKey:AddressInformationID"`
}

type Prosecution_Entry struct {
	gorm.Model
	Identifier   int    `json:"identifier" gorm:"uniqueIndex"`
	Code         int    `json:"code"`
	Type_Code    string `json:"type-code"`
	Date         int    `json:"date"`
	History_Text string `json:"history-text"`
}

// Main model which contains its fields and the other fields including party information and prosecution history
type Proceeding_Entry struct {
	gorm.Model
	Number                 int                  `json:"number" gorm:"uniqueIndex"`
	TypeCode               string               `json:"type-code"`
	Filing_Date            int                  `json:"filing-date"`
	Location_Code          int                  `json:"location-code"`
	Day_InLocation         int                  `json:"day-in-location"`
	Status_UpdateDate      int                  `json:"status-update-date"`
	Status_Code            int                  `json:"status-code"`
	Property_InformationID uint                 `json:"property-information-id"`
	Property_Information   Property_Information `json:"property-information" gorm:"foreignKey:Property_InformationID"`
	Prosecution_EntryID    uint                 `json:"prosecution-entry-array-id"`
	Prosecution_Entry      Prosecution_Entry    `json:"prosecution-entry" gorm:"foreignKey:Prosecution_EntryID"`
}
