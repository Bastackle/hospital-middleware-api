package database

import (
	"log"
	"time"

	"agnos/internal/model"

	"gorm.io/gorm"
)

func SeedHospitals(db *gorm.DB) {
	hospitals := []model.Hospital{
		{Code: "HOSP_A", Name: "Hospital A"},
		{Code: "HOSP_B", Name: "Hospital B"},
	}

	for _, h := range hospitals {
		var count int64
		db.Model(&model.Hospital{}).Where("code = ?", h.Code).Count(&count)
		if count == 0 {
			if err := db.Create(&h).Error; err != nil {
				log.Printf("Failed to seed hospital %s: %v\n", h.Code, err)
			} else {
				log.Printf("Seeded hospital: %s\n", h.Code)
			}
		}
	}

	SeedPatients(db)
}

func SeedPatients(db *gorm.DB) {
	var hospA, hospB model.Hospital
	if err := db.Where("code = ?", "HOSP_A").First(&hospA).Error; err != nil {
		return
	}
	if err := db.Where("code = ?", "HOSP_B").First(&hospB).Error; err != nil {
		return
	}

	dob1 := time.Date(1990, 1, 15, 0, 0, 0, 0, time.UTC)
	dob2 := time.Date(1985, 5, 20, 0, 0, 0, 0, time.UTC)

	// แปลง String เป็น Pointer สั้นๆ
	strPtr := func(s string) *string { return &s }

	patients := []model.Patient{
		// 🏥 ผู้ป่วยฝั่ง Hospital A
		{
			HospitalID:  hospA.ID,
			PatientHN:   "HN-A-001",
			NationalID:  strPtr("1100100100123"),
			PassportID:  nil,
			FirstNameTH: strPtr("สมชาย"),
			LastNameTH:  strPtr("ใจดี"),
			FirstNameEN: strPtr("Somchai"),
			LastNameEN:  strPtr("Jaidee"),
			DateOfBirth: &dob1,
			PhoneNumber: strPtr("0812345678"),
			Email:       strPtr("somchai@example.com"),
			Gender:      strPtr("M"),
		},
		{
			HospitalID:  hospA.ID,
			PatientHN:   "HN-A-002",
			NationalID:  nil,
			PassportID:  strPtr("A12345678"),
			FirstNameTH: strPtr("จอห์น"),
			LastNameTH:  strPtr("โดว์"),
			FirstNameEN: strPtr("John"),
			LastNameEN:  strPtr("Doe"),
			DateOfBirth: &dob2,
			PhoneNumber: strPtr("0898765432"),
			Email:       strPtr("john@example.com"),
			Gender:      strPtr("M"),
		},
		// 🏥 ผู้ป่วยฝั่ง Hospital B
		{
			HospitalID:  hospB.ID,
			PatientHN:   "HN-B-001",
			NationalID:  strPtr("2200200200456"),
			PassportID:  nil,
			FirstNameTH: strPtr("สมหญิง"),
			LastNameTH:  strPtr("รักดี"),
			FirstNameEN: strPtr("Somying"),
			LastNameEN:  strPtr("Rakdee"),
			DateOfBirth: &dob1,
			PhoneNumber: strPtr("0865554321"),
			Email:       strPtr("somying@example.com"),
			Gender:      strPtr("F"),
		},
	}

	for _, p := range patients {
		var count int64
		db.Model(&model.Patient{}).Where("hospital_id = ? AND patient_hn = ?", p.HospitalID, p.PatientHN).Count(&count)
		if count == 0 {
			if err := db.Create(&p).Error; err != nil {
				log.Printf("Failed to seed patient %s: %v\n", p.PatientHN, err)
			} else {
				log.Printf("Seeded patient: %s (%s)\n", p.PatientHN, *p.FirstNameEN)
			}
		}
	}
}
