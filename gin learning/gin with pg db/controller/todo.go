package controller

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
	"github.com/go-pg/pg/orm"
	guuid "github.com/google/uuid"
)

var dbConnect *pg.DB

// ToDo ...
type ToDo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title" binding:"Required"`
	Body      string    `json:"body" binding:"Required"`
	Completed string    `json:"completed" binding:"Required"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// CreateTodoTable create table on db server
func CreateTodoTable(db *pg.DB) error {
	options := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	err := db.CreateTable(&ToDo{}, options)
	if err != nil {
		log.Printf("Error while creating todo table, Reason: %v\n", err)
		return err
	}
	log.Printf("Todo table created")
	return nil
}

// InitiateDB > We are creating an instance of our DB to avoid too many connections.
func InitiateDB(db *pg.DB) {
	dbConnect = db
}

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
