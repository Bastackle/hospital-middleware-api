package service_test

import (
	"testing"

	"hospital-middleware/internal/model"
	"hospital-middleware/internal/service"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Helper function สำหรับสร้าง In-memory DB ไว้ใช้ทดสอบ
func setupTestDB(t *testing.T) *gorm.DB {
	dsn := "file:" + t.Name() + "?mode=memory&cache=shared"
	db, err := gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	assert.NoError(t, err)

	// Auto Migrate
	err = db.AutoMigrate(&model.Hospital{}, &model.Staff{}, &model.Patient{})
	assert.NoError(t, err)

	// Seed Hospital ตั้งต้น
	db.Create(&model.Hospital{Code: "HOSP_A", Name: "Hospital A"})
	db.Create(&model.Hospital{Code: "HOSP_B", Name: "Hospital B"})

	return db
}

// 🟢 Positive Test: Create Staff Success
func TestCreateStaff_Success(t *testing.T) {
	db := setupTestDB(t)
	staffService := service.NewStaffService(db)

	req := model.CreateStaffRequest{
		Username: "doctor_a",
		Password: "password123",
		Hospital: "HOSP_A",
	}

	staff, err := staffService.CreateStaff(req)

	assert.NoError(t, err)
	assert.NotNil(t, staff)
	assert.Equal(t, "doctor_a", staff.Username)
}

// 🔴 Negative Test: Login Invalid Password
func TestLoginStaff_InvalidPassword(t *testing.T) {
	db := setupTestDB(t)
	staffService := service.NewStaffService(db)

	// 1. สมัคร Staff ก่อน
	_, err := staffService.CreateStaff(model.CreateStaffRequest{
		Username: "doctor_b",
		Password: "password123",
		Hospital: "HOSP_A",
	})
	assert.NoError(t, err)

	// 2. ลอง Login ด้วย Password ที่ผิด
	loginReq := model.LoginRequest{
		Username: "doctor_b",
		Password: "wrong_password",
		Hospital: "HOSP_A",
	}

	token, err := staffService.LoginStaff(loginReq)

	assert.Error(t, err)
	assert.Empty(t, token)
	assert.Equal(t, "invalid credentials", err.Error())
}