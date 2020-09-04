package models

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofrs/uuid"
	"github.com/jackc/pgx"
	"golang.org/x/crypto/bcrypt"
)

var (
	tokenSecret = []byte(os.Getenv("TOKEN_SECRET"))
)

// User form
type User struct {
	ID              uuid.UUID `json:"id"`
	Created         time.Time `json:"_"`
	Updated         time.Time `json:"-"`
	Email           string    `json:"email"`
	PasswordHash    string    `json:"-"` // use when send info to db
	Password        string    `json:"password"`
	PasswordConfirm string    `json:"password_confirm"`
}

// Register tide with User
func (u *User) Register(conn *pgx.Conn) error {
	if len(u.Password) < 4 || len(u.PasswordConfirm) > 11 {
		return fmt.Errorf("Password must be at least 4 characters long")
	}
	if u.Password != u.PasswordConfirm {
		return fmt.Errorf("Password not match")
	}
	if len(u.Email) < 4 {
		return fmt.Errorf("Email must be at least 4 characters long")
	}

	u.Email = strings.ToLower(u.Email)
	row := conn.QueryRow(context.Background(), "SELECT id from Account WHERE email = $1", u.Email)
	userLookup := User{}
	err := row.Scan(&userLookup)
	if err != pgx.ErrNoRows {
		fmt.Println("found user")
		fmt.Println("userLookup.Email")
		return fmt.Errorf("A user with this email already exists.")
	}
	// hash the password
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("There was an error creating your account.")
	}
	u.PasswordHash = string(passwordHash)
	now := time.Now()
	_, err = conn.Exec(context.Background(), "INSERT INTO Account (created, updated, email, password_hash) VALUES($1, $2, $3, $4)", now, now, u.Email, passwordHash)
	// save to the database >> create account

	return err
}

// GetAuthToken return JWT token for user
func (u *User) GetAuthToken(string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = u.ID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodES256, &claims)
	authToken, err := token.SignedString(tokenSecret)
	return authToken, err
}
