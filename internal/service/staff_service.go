package service

import (
	"errors"
	"fmt"

	"agnos/internal/model"
	"agnos/pkg/auth"

	"gorm.io/gorm"
)

type StaffService struct {
	db *gorm.DB
}

func NewStaffService(db *gorm.DB) *StaffService {
	return &StaffService{db: db}
}

// CreateStaff สมัคร Staff ใหม่
func (s *StaffService) CreateStaff(req model.CreateStaffRequest) (*model.Staff, error) {
	// 1. ค้นหา Hospital จาก Code ที่ส่งมา (เช่น "HOSP_A")
	var hospital model.Hospital
	if err := s.db.Where("code = ?", req.Hospital).First(&hospital).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("hospital not found")
		}
		return nil, err
	}

	// 2. Hash Password
	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}

	// 3. บันทึก Staff ลง DB
	staff := model.Staff{
		Username: req.Username,
		PasswordHash: hashedPassword,
		HospitalID: hospital.ID,
	}

	if err := s.db.Create(&staff).Error; err != nil {
		return nil, fmt.Errorf("username already exists or DB error: %w", err)
	}

	return &staff, nil
}

// LoginStaff ตรวจสอบสิทธิ์และ คืนค่า JWT Token
func (s *StaffService) LoginStaff(req model.LoginRequest) (string, error) {
	// 1. ค้นหา Hospital
	var hospital model.Hospital
	if err := s.db.Where("code = ?", req.Hospital).First(&hospital).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	// 2. ค้นหา Staff จาก Username + HospitalID
	var staff model.Staff
	if err := s.db.Where("username = ? AND hospital_id = ?", req.Username, hospital.ID).First(&staff).Error; err != nil {
		return "", errors.New("invalid credentials")
	}

	// 3. ตรวจสอบ Password
	if !auth.CheckPasswordHash(req.Password, staff.PasswordHash) {
		return "", errors.New("invalid credentials")
	}

	// 4. ออก JWT Token
	token, err := auth.GenerateToken(staff.ID, staff.Username, staff.HospitalID)
	if err != nil {
		return "", fmt.Errorf("failed to generate token: %w", err)
	}

	return token, nil
}