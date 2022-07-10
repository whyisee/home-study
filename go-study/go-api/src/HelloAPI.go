package main
import (
	"strings"
	"strconv"
	"github.com/gin-gonic/gin"
	)

func main() {
    r := gin.Default()
    r.GET("/users", listUser)
	r.GET("/users/:id", getUser)

    r.Run(":8080")
}
func listUser(c *gin.Context) {
    c.JSON(200, users)
}

func getUser(c *gin.Context) {
	id := c.Param("id")
	var user User
	found := false
	//类似于数据库的SQL查询
	for _, u := range users {
	   if strings.EqualFold(id, strconv.Itoa(u.ID)) {
		  user = u
		  found = true
		  break
	   }
	}
	if found {
	   c.JSON(200, user)
	} else {
	   c.JSON(404, gin.H{
		  "message": "用户不存在",
	   })
	}
 }
//数据源，模拟MySQL中的数据
var users = []User{
    {ID: 1, Name: "张三"},
    {ID: 2, Name: "李四"},
    {ID: 3, Name: "王五"},
}

//用户
type User struct {
    ID   int
    Name string
}