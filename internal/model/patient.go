package model

import "time"

type Patient struct {
	ID           uint       `gorm:"primaryKey" json:"id"`
	HospitalID   uint       `gorm:"not null;index" json:"hospital_id"`
	PatientHN    string     `gorm:"not null" json:"patient_hn"`
	NationalID   *string    `gorm:"index" json:"national_id"`
	PassportID   *string    `gorm:"index" json:"passport_id"`
	FirstNameTH  *string    `json:"first_name_th"`
	MiddleNameTH *string    `json:"middle_name_th"`
	LastNameTH   *string    `json:"last_name_th"`
	FirstNameEN  *string    `json:"first_name_en"`
	MiddleNameEN *string    `json:"middle_name_en"`
	LastNameEN   *string    `json:"last_name_en"`
	DateOfBirth  *time.Time `json:"date_of_birth"`
	PhoneNumber  *string    `json:"phone_number"`
	Email        *string    `json:"email"`
	Gender       *string    `gorm:"type:varchar(1)" json:"gender"`
	CreatedAt    time.Time  `json:"created_at"`
	UpdatedAt    time.Time  `json:"updated_at"`
}

type PatientSearchRequest struct {
	NationalID  string `json:"national_id"`
	PassportID  string `json:"passport_id"`
	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	LastName    string `json:"last_name"`
	DateOfBirth string `json:"date_of_birth"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
}
