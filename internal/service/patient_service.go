package service

import (
	"context"
	"strings"

	"hospital-middleware/internal/model"

	"gorm.io/gorm"
)

type PatientService struct {
	db *gorm.DB
}

func NewPatientService(db *gorm.DB) *PatientService {
	return &PatientService{db: db}
}

func (s *PatientService) SearchPatients(ctx context.Context, hospitalID uint, req model.PatientSearchRequest) ([]model.Patient, error) {
	var patients []model.Patient

	// 1. บังคับ Filter เงื่อนไข hospital_id เสมอเพื่อความปลอดภัย
	query := s.db.WithContext(ctx).Where("hospital_id = ?", hospitalID)

	// 2. Dynamic Query: ฟิลด์ไหนมีค่าส่งมา ค่อยเติมเงื่อนไข
	if req.NationalID != "" {
		query = query.Where("national_id = ?", req.NationalID)
	}
	if req.PassportID != "" {
		query = query.Where("passport_id = ?", req.PassportID)
	}
	if req.PhoneNumber != "" {
		query = query.Where("phone_number = ?", req.PhoneNumber)
	}
	if req.Email != "" {
		query = query.Where("email = ?", req.Email)
	}
	if req.DateOfBirth != "" {
		query = query.Where("date_of_birth = ?", req.DateOfBirth)
	}

	// ค้นหาตามชื่อ (ค้นหาได้ทั้ง TH และ EN)
	if req.FirstName != "" {
		likePattern := "%" + strings.ToLower(req.FirstName) + "%"
		query = query.Where("LOWER(first_name_th) LIKE ? OR LOWER(first_name_en) LIKE ?", likePattern, likePattern)
	}
	if req.MiddleName != "" {
		likePattern := "%" + strings.ToLower(req.MiddleName) + "%"
		query = query.Where("LOWER(middle_name_th) LIKE ? OR LOWER(middle_name_en) LIKE ?", likePattern, likePattern)
	}
	if req.LastName != "" {
		likePattern := "%" + strings.ToLower(req.LastName) + "%"
		query = query.Where("LOWER(last_name_th) LIKE ? OR LOWER(last_name_en) LIKE ?", likePattern, likePattern)
	}

	if err := query.Find(&patients).Error; err != nil {
		return nil, err
	}

	return patients, nil
}