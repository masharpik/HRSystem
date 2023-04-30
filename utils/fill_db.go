package utils

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	"gorm.io/gorm"
)

const familyCompositionsCount = 4

func fillFamilyCompositions(db *gorm.DB) (err error) {
	familyCompositions := []FamilyComposition{
		{Name: "Простая элементарная"},
		{Name: "Простая составная"},
		{Name: "Сложная"},
		{Name: "Однополая"},
	}
	err = db.Create(&familyCompositions).Error
	return err
}

func fillEnterprisesDepartmentsPositions(db *gorm.DB, enterprisesCount, departmentsCount, positionCount int) (err error) {
	enterprises := make([]Enterprise, 0, enterprisesCount)

	for i := 1; i <= enterprisesCount; i++ {
		departments := make([]Department, 0, departmentsCount)

		for j := 1; j <= departmentsCount; j++ {
			positions := make([]Position, 0, positionCount)

			for k := 1; k <= positionCount; k++ {
				positions = append(positions, Position{
					Name:        fmt.Sprintf("Позиция %d отдела %d компании %d", k, j, i),
					Description: "Описание",
					Wage:        float64(rand.Intn(100000)),
				})
			}

			departments = append(departments, Department{
				Name:      fmt.Sprintf("Отдел %d компании %d", j, i),
				Positions: positions,
			})
		}

		enterprises = append(enterprises, Enterprise{
			Name:           fmt.Sprintf("Компания %d", i),
			StaffUnits:     rand.Intn(100),
			SystemPassword: "password",
			Departments:    departments,
		})
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err = tx.CreateInBatches(&enterprises, 1000).Error; err != nil {
			return err
		}

		return nil
	})

	return err
}

func fillEmployeesContractsPassportsIdentities(db *gorm.DB, employeesCount, positionCount int) (err error) {
	var employees = make([]Employee, 0, employeesCount)

	for i := 1; i <= employeesCount; i++ {
		Identity := Identity{
			Surname:             "Surname",
			Name:                "Name",
			Middlename:          "Middlename",
			Birth:               time.Now(),
			FamilyCompositionID: uint(rand.Intn(familyCompositionsCount) + 1),
		}

		var contract Contract
		contract = Contract{
			TIN:                  i,
			DateHiring:           time.Now(),
			PlaceHiring:          "Someplace",
			Workplace:            "Workplace",
			ContractExpiry:       time.Now().Add(time.Hour * 24 * 365),
			WorkMode:             "Full-time",
			AdditionalConditions: "None",
			RepresentativeID:     nil,
		}

		passport := Passport{
			Series:       rand.Intn(9999) + 1,
			Number:       rand.Intn(999999) + 1,
			UssiedBy:     "UssiedBy",
			IssueDate:    time.Now(),
			Registration: "Some registration",
		}

		employees = append(employees, Employee{
			PositionID: uint(rand.Intn(positionCount) + 1),
			Identity:   Identity,
			Contracts:  contract,
			Passports:  passport,
		})
	}

	err = db.Transaction(func(tx *gorm.DB) error {
		if err := tx.CreateInBatches(&employees, 1000).Error; err != nil {
			return err
		}

		return nil
	})

	return
}

func fillSalaryChanges(db *gorm.DB, salaryChangesCount int) (err error) {
	var salaryChanges []SalaryChanges
	for i := 1; i <= salaryChangesCount; i++ {
		salaryChanges = append(salaryChanges, SalaryChanges{
			Amount:     float64(rand.Intn(100000)),
			ChangeDate: time.Now(),
			EmployeeID: uint(i),
		})
	}

	err = db.Create(&salaryChanges).Error

	return
}

func fillCareer(db *gorm.DB, careerCount int) (err error) {
	var career []Career
	for i := 1; i <= careerCount; i++ {
		var jsonData []byte

		jsonData, err = json.Marshal([]uint{uint(i), uint(i + 1)})
		career = append(career, Career{
			PositionIDs: string(jsonData),
			EmployeeID:  uint(i + 1),
		})
	}
	err = db.Create(&career).Error

	return
}

func fillEditors(db *gorm.DB, editorsCount int) (err error) {
	var editors []Editor
	for i := 1; i <= editorsCount; i++ {
		editors = append(editors, Editor{
			Login:        fmt.Sprintf("login%d", i),
			Password:     "password",
			EnterpriseID: uint(i),
		})
	}
	err = db.Create(&editors).Error

	return
}

func fillHREmployees(db *gorm.DB, hrEmployeesCount int) (err error) {
	var hrEmployees []HREmployee
	for i := 1; i <= hrEmployeesCount; i++ {
		hrEmployees = append(hrEmployees, HREmployee{
			Login:        fmt.Sprintf("login%d", i),
			Password:     "password",
			EnterpriseID: uint(i),
			Approval:     i%2 == 0,
		})
	}
	err = db.Create(&hrEmployees).Error

	return
}

func fillOrders(db *gorm.DB, ordersCount int) (err error) {
	var orders []Order
	for i := 1; i <= ordersCount; i++ {
		orders = append(orders, Order{
			Title: fmt.Sprintf("Название отчета №%d", i),
			Text: fmt.Sprintf("Текст отчета №%d", i),
			Date: time.Now(),
			HREmployeeID:  uint(i),
		})
	}
	err = db.Create(&orders).Error

	return
}

func FillDB(db *gorm.DB) (err error) {
	if err = fillFamilyCompositions(db); err != nil {
		return err
	}
	log.Println("Таблица family_compositions заполнилась")

	enterprisesCount, departmentsCount, positionCount := 10, 5, 10
	if err = fillEnterprisesDepartmentsPositions(db, enterprisesCount, departmentsCount, positionCount); err != nil {
		return err
	}
	log.Println("Таблицы enterprises, departments, positions заполнились")

	employeesCount := 1000000
	if err = fillEmployeesContractsPassportsIdentities(db, employeesCount, enterprisesCount*departmentsCount*positionCount); err != nil {
		return err
	}
	log.Println("Таблицы employees, contracts, passposts, identities заполнились")

	if err = fillSalaryChanges(db, 10); err != nil {
		return err
	}
	log.Println("Таблица salary_changes заполнилась")

	if err = fillCareer(db, 10); err != nil {
		return err
	}
	log.Println("Таблица career заполнилась")

	if err = fillEditors(db, 10); err != nil {
		return err
	}
	log.Println("Таблица editors заполнилась")

	if err = fillHREmployees(db, 10); err != nil {
		return err
	}
	log.Println("Таблица hr_employees заполнилась")

	if err = fillOrders(db, 10); err != nil {
		return err
	}
	log.Println("Таблица orders заполнилась")

	return
}
