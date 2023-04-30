package utils

import (
	"time"
)

type Enterprise struct {
	ID             uint `gorm:"primary hey"`
	Name           string
	StaffUnits     int
	SystemPassword string
	Departments    []Department `gorm:"foreignKey:EnterpriseID"`
}

type Department struct {
	ID           uint `gorm:"primary hey"`
	Name         string
	EnterpriseID uint
	Positions    []Position `gorm:"foreignKey:DepartmentID"`
}

type Position struct {
	ID           uint `gorm:"primary hey"`
	Name         string
	Description  string
	Wage         float64
	DepartmentID uint
}

type Employee struct {
	ID         uint `gorm:"primary hey"`
	PositionID uint
	Contracts  Contract `gorm:"foreignKey:EmployeeID"`
	Passports  Passport `gorm:"foreignKey:EmployeeID"`
	Identity   Identity `gorm:"foreignKey:EmployeeID"`
}

type Contract struct {
	ID                   uint `gorm:"primary hey"`
	TIN                  int
	RepresentativeID     *uint
	DateHiring           time.Time
	PlaceHiring          string
	Workplace            string
	ContractExpiry       time.Time
	WorkMode             string
	AdditionalConditions string
	EmployeeID           uint
}

type Passport struct {
	ID           uint `gorm:"primary hey"`
	Series       int
	Number       int
	UssiedBy     string
	IssueDate    time.Time
	Registration string
	EmployeeID   uint
}

type Identity struct {
	ID                  uint `gorm:"primary hey"`
	Surname             string
	Name                string
	Middlename          string
	Birth               time.Time
	FamilyCompositionID uint
	EmployeeID          uint
}

type FamilyComposition struct {
	ID   uint `gorm:"primary hey"`
	Name string
}

type SalaryChanges struct {
	ID         uint `gorm:"primary hey"`
	Amount     float64
	ChangeDate time.Time
	EmployeeID uint
}

type Career struct {
	ID          uint   `gorm:"primary hey"`
	PositionIDs string `gorm:"type:json"`
	EmployeeID  uint
}

func (Career) TableName() string {
	return "career"
}

type Editor struct {
	ID           uint `gorm:"primary hey"`
	Login        string
	Password     string
	Avatar       string
	EnterpriseID uint
}

type HREmployee struct {
	ID           uint `gorm:"primary hey"`
	Login        string
	Password     string
	Avatar       string
	EnterpriseID uint
	Approval     bool
}

type Order struct {
	ID           uint `gorm:"primary hey"`
	Title        string
	Text         string
	Date         time.Time
	HREmployeeID uint
}
