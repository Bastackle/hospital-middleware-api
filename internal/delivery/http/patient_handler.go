package http

import (
	"net/http"

	"agnos/internal/model"
	"agnos/internal/service"

	"github.com/gin-gonic/gin"
)

type PatientHandler struct {
	patientService *service.PatientService
}

func NewPatientHandler(patientService *service.PatientService) *PatientHandler {
	return &PatientHandler{patientService: patientService}
}

// SearchPatients godoc
// @Summary      Search patients
// @Description  Search patients belonging to the same hospital as the logged-in staff
// @Tags         patient
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        request body model.PatientSearchRequest false "Search Criteria"
// @Success      200 {object} map[string]interface{}
// @Failure      401 {object} map[string]interface{}
// @Router       /patient/search [post]
func (h *PatientHandler) SearchPatients(c *gin.Context) {
	// ดึง hospital_id ที่แกะได้จาก JWT Token ใน Middleware
	hospitalIDVal, exists := c.Get("hospital_id")
	if !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized context"})
		return
	}
	hospitalID := hospitalIDVal.(uint)

	var req model.PatientSearchRequest
	// ใช้ ShouldBindJSON (ถ้าร่าง Body มาว่างๆ ก็ยังถือว่าผ่านเพราะทุกฟิลด์ optional)
	if err := c.ShouldBindJSON(&req); err != nil && err.Error() != "EOF" {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	patients, err := h.patientService.SearchPatients(c.Request.Context(), hospitalID, req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search patients"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": patients,
	})
}