package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Todo struct {
	Id        int    `form:"id"`
	InCharge  string `form:"inCharge"`
	Content   string `form:"content"`
	Status    int    `form:"status"`
	CreatedAt time.Time
}

func (t *Todo) CreatedAtS() string {
	return t.CreatedAt.Format("2006-01-02")
}

var todos []Todo
var maxId int = 0

func increment() int {
	maxId += 1
	return maxId
}

func GetTodoData(c *gin.Context) {
	var b Todo
	if err := c.Bind(&b); err != nil {
		fmt.Errorf("%#v", err)
	}
	b.Id = increment()
	b.Status = 0
	b.CreatedAt = time.Now()

	todos = append(todos, b)
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todos": todos,
	})
}

func GetDoneTodoData(c *gin.Context) {
	var b Todo
	if err := c.Bind(&b); err != nil {
		fmt.Errorf("%#v", err)
	}
	var s int

	if b.Status == 0 {
		s = 1
	} else {
		s = 0
	}

	for idx, t := range todos {
		if t.Id == b.Id {
			todos[idx].Status = s
		}
	}
	c.HTML(http.StatusOK, "index.html", map[string]interface{}{
		"todos": todos,
	})
}

func main() {
	fmt.Println("sss")
	r := gin.Default()
	r.LoadHTMLFiles("./template/index.html")

	todos = []Todo{
		Todo{
			Id:        increment(),
			InCharge:  "Lea",
			Content:   "cook",
			Status:    0,
			CreatedAt: time.Now(),
		},
		Todo{
			Id:        increment(),
			InCharge:  "Koro",
			Content:   "play",
			Status:    1,
			CreatedAt: time.Now(),
		},
		Todo{
			Id:        increment(),
			InCharge:  "Lea",
			Content:   "sleep",
			Status:    0,
			CreatedAt: time.Now(),
		},
	}

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello World",
		})
	})

	r.GET("/todo", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", map[string]interface{}{
			"todos": todos,
		})
	})

	r.GET("/post-todo", GetTodoData)
	r.GET("/done", GetDoneTodoData)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
