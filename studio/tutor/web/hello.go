package main

import (
	"github.com/kataras/iris"
	"time"
	"fastweb/studio/tutor/web/model"
)

func main() {

	/*mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"localhost"},
		Timeout:  30 * time.Second,
		Database: "fastweb",
		Username: "fastweb",
		Password: "fastweb",
	}

	mongoSession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		fmt.Println("mongodb", err)
		os.Exit(1)
	}
	mongoSession.SetMode(mgo.Monotonic, true);
	defer mongoSession.Close()*/


	//iris.Use(logger.New())
	iris.Get("/hi_json", func(c *iris.Context) {
		c.JSON(200, iris.Map{
			"Name" : "iris",
			"Released":"13 March 2016",
		})
	})
	iris.Get("/user/save", func(c *iris.Context) {
		user := model.User{}
		user.Name = "wangcj"
		user.Memo = "test"
		user.Password = []byte{'1', '2', '3', '@', '4', '5', '6'}
		user.LoginName = "wangcj"
		user.Sex = model.Male
		user.Birthday, _ = time.Parse("2006-01-02", "1980-03-18")
		if success := model.AddUser(user); success == "false" {
			c.Text(200, "save failure")
		} else {
			c.Text(200, "save success:" + success)
		}
	})

	iris.Get("/user/all", func(c *iris.Context) {
		c.JSON(200, model.GetAllUser())
	})

	iris.Get("/user/get/:id", func(c *iris.Context) {
		userId := c.Param("id")
		c.JSON(200, model.GetUserById(userId))
	})

	iris.Get("/user/update/:id", func(c * iris.Context) {
		userId := c.Param("id");
		model.UpdateUser(model.GetUserById(userId))
		c.JSON(200, model.GetUserById(userId))
	})

	logme := func(ctx *iris.Context) {
		println("request to /hi")
		ctx.Next()
	}
	iris.Get("/hi", logme, hi)
	iris.Static("/public", "./static", 1)
	iris.Listen(":8080")
}

func hi(c *iris.Context) {
	c.MustRender("hi.html", struct{ Name string }{Name: "iris"})
}
