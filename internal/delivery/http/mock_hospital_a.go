package http

import (
	"net/http"

	"hospital-middleware/internal/model"

	"github.com/gin-gonic/gin"
)

// MockHospitalASearchHandler godoc
// @Summary      Search patient from Hospital A HIS (Mock API)
// @Description  Mock External HIS API for Hospital A using national_id or passport_id
// @Tags         external-his
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "National ID or Passport ID"
// @Success      200  {object}  model.HospitalAPatientResponse
// @Failure      404  {object}  map[string]interface{}
// @Router       /patient/search/{id} [get]
func MockHospitalASearchHandler(c *gin.Context) {
	id := c.Param("id")

	// สมมติว่าหาเจอ ถ้า id ตรงกับที่กำหนด
	if id == "1100100100123" || id == "A12345678" {
		fnTH := "สมชาย"
		lnTH := "ใจดี"
		fnEN := "Somchai"
		lnEN := "Jaidee"
		dob := "1990-01-15"
		natID := "1100100100123"
		phone := "0812345678"
		email := "somchai@example.com"
		gender := "M"

		c.JSON(http.StatusOK, model.HospitalAPatientResponse{
			FirstNameTH: &fnTH,
			LastNameTH:  &lnTH,
			FirstNameEN: &fnEN,
			LastNameEN:  &lnEN,
			DateOfBirth: &dob,
			PatientHN:   "HN-A-MOCK-999",
			NationalID:  &natID,
			PhoneNumber: &phone,
			Email:       &email,
			Gender:      &gender,
		})
		return
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "patient not found in Hospital A"})
}