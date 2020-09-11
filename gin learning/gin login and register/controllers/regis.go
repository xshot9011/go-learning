package controllers

import (
	"net/http"
	"time"
	"unicode"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/go-playground/validator.v9"
	"gorm.io/gorm"
)

var dbConnect *gorm.DB

const saltSize = 32

// InitiateDB > We are creating an instance of our DB to avoid too many connections.
func InitiateDB(db *gorm.DB) {
	dbConnect = db
}

// User model for db
type User struct {
	UUID            string    `gorm:"primaryKey" form:"-"`
	Fullname        string    `gorm:"not null;size:256" form:"fullname" validate:"required,min=1,max=256"`
	Email           string    `gorm:"not null;unique;size:256" form:"email" validate:"required,min=4,max=256,email"`
	Password        string    `gorm:"-" form:"password" validate:"required,min=8,eqfield=PasswordConfirm"`
	PasswordConfirm string    `gorm:"-" form:"password_confirm" validate:"required,min=8"`
	PasswordHash    string    `gorm:"not null;size:256" form:"-"`
	Created         time.Time `gorm:"autoCreateTime:milli" form:"-"`
	Updated         time.Time `gorm:"autoUpdateTime:milli" form:"-"`
}

// CreateuserTable create user table from struct
func CreateuserTable(db *gorm.DB) {
	isExist := db.Migrator().HasTable(&User{})
	if !isExist {
		db.Migrator().CreateTable(&User{})
	}
}

// InsertUserToDatabase > insert validated data and tranform to appropriate form then into database
func (user *User) InsertUserToDatabase() error {
	result := dbConnect.Model(User{}).Create(user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Registration > make registration for user
func Registration(c *gin.Context) {
	v := validator.New()
	user := User{
		UUID:            uuid.New().String(),
		Fullname:        c.PostForm("fullname"),
		Email:           c.PostForm("email"),
		Password:        c.PostForm("password"),
		PasswordConfirm: c.PostForm("password_confirm"),
	}
	// check if password is secure or not
	isSecure := isSecurePassword(user.Password)
	if !isSecure {
		c.JSON(http.StatusBadRequest, gin.H{
			"password": "password is not secure",
		})
		return
	}
	// validate data and get error
	err := v.Struct(user)
	if err != nil {
		errorMsg := make(map[string]string)
		for _, e := range err.(validator.ValidationErrors) {
			errorMsg[e.Field()] = e.ActualTag()
		}
		c.JSON(http.StatusBadRequest, errorMsg)
		return
	}
	// check if the email is exist
	var temp string
	row := dbConnect.Table("users").Where("email = ?", user.Email).Select("UUID").Row()
	row.Scan(&temp)
	if len(temp) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Email is already used",
		})
		return
	}
	// hash the password and sert salt to user
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}
	user.PasswordHash = string(passwordHash)
	now := time.Now().UTC()
	user.Created = now
	user.Updated = now
	// Insert data to database
	err = user.InsertUserToDatabase()
	if err != nil {
		c.JSON(http.StatusCreated, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"msg":  "Successfully create account",
		"data": user,
	})
}

func isSecurePassword(pass string) bool {
	var (
		upp, low, num, sym bool
		tot                uint8
	)

	for _, char := range pass {
		switch {
		case unicode.IsUpper(char):
			upp = true
			tot++
		case unicode.IsLower(char):
			low = true
			tot++
		case unicode.IsNumber(char):
			num = true
			tot++
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			sym = true
			tot++
		default:
			return false
		}
	}

	if !upp || !low || !num || !sym || tot < 8 {
		return false
	}

	return true
}

/*
example code
// GetAllTodos return all row of todo in db
func GetAllTodos(c *gin.Context) {
	var todos []ToDo
	err := dbConnect.Model(&todos).Select()
	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "All Todos",
		"data":    todos,
	})
	return
}

// CreateTodo create a obj todo row
func CreateTodo(c *gin.Context) {
	var todo ToDo
	c.BindJSON(&todo)

	title := todo.Title
	body := todo.Body
	completed := todo.Completed
	id := guuid.New().String()

	insertError := dbConnect.Insert(&ToDo{
		ID:        id,
		Title:     title,
		Body:      body,
		Completed: completed,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if insertError != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"message": "Todo created Successfully",
	})
	return
}

// GetSingleTodo return the infomation about specific id
func GetSingleTodo(c *gin.Context) {
	todoID := c.Param("todoId")
	todo := &ToDo{ID: todoID}
	err := dbConnect.Select(todo)
	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Todo not found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Single Todo",
		"data":    todo,
	})
	return
}

// EditTodo like patch method
func EditTodo(c *gin.Context) {
	todoID := c.Param("todoId")
	var todo ToDo
	c.BindJSON(&todo)
	completed := todo.Completed
	_, err := dbConnect.Model(&ToDo{}).Set("completed = ?", completed).Where("id = ?", todoID).Update()
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo Edited Successfully",
	})
	return
}

// DeleteTodo delete specific item
func DeleteTodo(c *gin.Context) {
	todoID := c.Param("todoId")
	todo := &ToDo{ID: todoID}
	err := dbConnect.Delete(todo)
	if err != nil {
		log.Printf("Error while deleting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Something went wrong",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Todo deleted successfully",
	})
	return
}

*/
