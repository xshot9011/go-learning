package api

import (
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

// RegistrationInfomation > declare field for registration
type RegistrationInfomation struct {
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	PasswordConfirm string `json:"password_confirm"`
	PasswordHash    string `json:"-"`
}

// Register > register the user
// try to connect to database later
func (regInfo *RegistrationInfomation) Register(c *gin.Context) error {
	if len(regInfo.Password) < 4 || len(regInfo.PasswordConfirm) > 11 {
		return fmt.Errorf("Password must be at least 4 characters long")
	}
	if regInfo.Password != regInfo.PasswordConfirm {
		return fmt.Errorf("Password not match")
	}
	if len(regInfo.Email) < 4 {
		return fmt.Errorf("Email must be at least 4 characters long")
	}

	regInfo.Email = strings.ToLower(regInfo.Email)
	// hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(regInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("THERE WAS ERROR CREATING YOUR ACCOUNT")
	}
	regInfo.PasswordHash = string(passwordHash)

	return err
}
