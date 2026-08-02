package service_test

import (
	"context"
	"testing"

	"hospital-middleware/internal/model"
	"hospital-middleware/internal/service"

	"github.com/stretchr/testify/assert"
)

// 🛡️ Positive & Negative Test: Multi-tenancy Isolation Check
func TestSearchPatients_HospitalIsolation(t *testing.T) {
	db := setupTestDB(t)
	patientService := service.NewPatientService(db)

	// ดึง Hospital ID
	var hospA, hospB model.Hospital
	db.Where("code = ?", "HOSP_A").First(&hospA)
	db.Where("code = ?", "HOSP_B").First(&hospB)

	fnTH := "สมชาย"
	fnEN := "Somchai"

	// สร้างผู้ป่วยประจำ Hospital A และ Hospital B
	patientA := model.Patient{
		HospitalID:  hospA.ID,
		PatientHN:   "HN-A-001",
		FirstNameTH: &fnTH,
		FirstNameEN: &fnEN,
	}
	patientB := model.Patient{
		HospitalID:  hospB.ID,
		PatientHN:   "HN-B-001",
		FirstNameTH: &fnTH,
		FirstNameEN: &fnEN,
	}

	db.Create(&patientA)
	db.Create(&patientB)

	// 🟢 Case 1: Staff ของ Hospital A ค้นหาคำว่า "Somchai"
	// ต้องเจอเฉพาะผู้ป่วยของ Hospital A เท่านั้น
	resultsA, err := patientService.SearchPatients(context.Background(), hospA.ID, model.PatientSearchRequest{
		FirstName: "Somchai",
	})

	assert.NoError(t, err)
	assert.Len(t, resultsA, 1) // ต้องเจอแค่ 1 คน
	assert.Equal(t, "HN-A-001", resultsA[0].PatientHN)

	// 🔴 Case 2: พิสูจน์ Isolation - ผลลัพธ์ของ Hospital A ต้องไม่มีผู้ป่วยของ Hospital B ติดมาด้วย
	for _, p := range resultsA {
		assert.NotEqual(t, "HN-B-001", p.PatientHN, "Staff Hosp A should not see Hosp B patients")
	}
}