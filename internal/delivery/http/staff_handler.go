package http

import (
	"net/http"

	"hospital-middleware/internal/model"
	"hospital-middleware/internal/service"

	"github.com/gin-gonic/gin"
)

type StaffHandler struct {
	staffService *service.StaffService
}

func NewStaffHandler(staffService *service.StaffService) *StaffHandler {
	return &StaffHandler{staffService: staffService}
}

// CreateStaff godoc
// @Summary      Create a new hospital staff member
// @Description  Create staff with username, password, and hospital code
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param        request body model.CreateStaffRequest true "Staff Details"
// @Success      201 {object} map[string]interface{}
// @Failure      400 {object} map[string]interface{}
// @Router       /staff/create [post]
func (h *StaffHandler) CreateStaff(c *gin.Context) {
	var req model.CreateStaffRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	staff, err := h.staffService.CreateStaff(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": "Staff created successfully",
		"data": gin.H{
			"id": staff.ID,
			"username": staff.Username,
			"hospital_id": staff.HospitalID,
		},
	})
}

// Login godoc
// @Summary      Staff Login
// @Description  Authenticate staff and return JWT token
// @Tags         staff
// @Accept       json
// @Produce      json
// @Param        request body model.LoginRequest true "Login Credentials"
// @Success      200 {object} map[string]string
// @Failure      401 {object} map[string]interface{}
// @Router       /staff/login [post]
func (h *StaffHandler) Login(c *gin.Context) {
	var req model.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	token, err := h.staffService.LoginStaff(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}