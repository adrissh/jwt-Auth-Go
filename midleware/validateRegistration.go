package midleware

import (
	"jwtAuthGo/response"
	"net/http"
	"regexp"
	"time"

	"github.com/gin-gonic/gin"
)

type Request struct {
	Username      string `json:"username"`
	Email         string `json:"email"`
	Password_hash string `json:"password"`
	Role          string `json:"role"`
}

func ValidateRequest(req Request) map[string]string {

	// map for store validation errors
	errors := make(map[string]string)

	// Validate username
	if req.Username == "" {
		errors["username"] = "username is required"
	} else {
		usernamePattern := `^[a-zA-Z 0-9\s]+$`
		nameRegex := regexp.MustCompile(usernamePattern)
		if !nameRegex.MatchString(req.Username) {
			errors["username"] = "invalid username format"
		}
	}
	// Validasi Email
	if req.Email == "" {
		errors["email"] = "email is required"
	} else {
		emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
		emailRegex := regexp.MustCompile(emailPattern)
		if !emailRegex.MatchString(req.Email) {
			errors["email"] = "invalid email format"
		}
	}

	// Validate Password
	if req.Password_hash == "" { // password required
		errors["password"] = "password is required"
	} else {
		// Make sure letters and numbers exist
		if matched, _ := regexp.MatchString(`[a-zA-Z]`, req.Password_hash); !matched {
			errors["password"] = "password must contain at least one letter"
		}
		// make sure numbes exist
		if matched, _ := regexp.MatchString(`\d`, req.Password_hash); !matched {
			errors["password"] = "password must contain at least one number"
		}
		// Make sure the special character exist
		if matched, _ := regexp.MatchString(`[!@#\$%\^&\*\(\)_\+\-\=\[\]\{\};:\'",<>\./?\\|]`, req.Password_hash); !matched {
			errors["password"] = "password must contain at least one special character"
		}
	}

	if len(errors) > 0 {
		return errors

	}
	return nil
}

func ValidateRegister(c *gin.Context) {
	var reqPayload Request
	if err := c.ShouldBind(&reqPayload); err != nil {
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:     "error",
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request payload",
			Timestamp:  time.Now().Format(time.RFC3339),
			Errors:     err.Error(),
		})
		return
	}

	// validate request body
	err := ValidateRequest(reqPayload)
	if err != nil {
		var errorDetails []map[string]string
		for field, message := range err {
			errorDetails = append(errorDetails, map[string]string{
				"field": field,
				"error": message,
			})
		}
		c.JSON(http.StatusBadRequest, response.ErrorResponse{
			Status:     "error",
			StatusCode: http.StatusBadRequest,
			Message:    "invalid request body",
			Timestamp:  time.Now().Format(time.RFC3339),
			Errors:     errorDetails,
		})
		c.Abort()
		return
	}
	c.Set("payload", reqPayload)
	c.Next()
}
