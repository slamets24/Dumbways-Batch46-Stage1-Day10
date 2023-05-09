package main

import (
	"context"
	"fmt"
	"html/template"
	"math"
	"net/http"
	"personal-web/connection"
	"personal-web/middleware"
	"strconv"
	"time"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	connection.DatabaseConnect()

	e := echo.New()

	e.Static("/public", "public")

	// initialitation to use session
	// e.Use(session.Middleware(sessions.NewCookieStore([]byte("session"))))

	e.GET("/", home)
	e.GET("/contact", contactMe)
	e.GET("/my-project", myProject)
	// e.GET("/detail-project", detailProject)
	e.GET("/testimonial", testimonial)
	e.GET("/signUp", signup)
	e.GET("/signIn", signin)
	e.POST("/add-project", middleware.UploadFile(addProject))
	e.POST("/update", updateProject)
	e.GET("/project-detail/:id", projectDetail)
	e.GET("/delete-project/:id", deleteProject)
	e.GET("/update-project/:id", tmplUpdate)
	e.POST("/add-user", adduser)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

type Project struct {
	Id          int
	Pname       string
	Startdate   time.Time
	Enddate     time.Time
	Description string
	Tech        []string
	Image       string
}

type User struct {
	Id       int
	Name     string
	Email    string
	Password string
}

type SessionData struct {
	IsLogin bool
	Name    string
}

var userData = SessionData{}

// =================================================

func home(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func contactMe(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/contact-me.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"messege": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func myProject(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/myProject.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"messege": err.Error()})
	}

	// map(tipe data) => key and value
	data, _ := connection.Conn.Query(context.Background(), "SELECT id, pname, startdate, enddate, description, tech, image FROM tb_project")
	fmt.Println(data)

	var result []Project

	for data.Next() {
		var each = Project{}

		err := data.Scan(&each.Id, &each.Pname, &each.Startdate, &each.Enddate, &each.Description, &each.Tech, &each.Image)

		if err != nil {
			fmt.Println(err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{"messege": err.Error()})
		}

		result = append(result, each)
	}

	projects := map[string]interface{}{
		"Project": result,
	}

	return tmpl.Execute(c.Response(), projects)
}

func testimonial(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func projectDetail(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tmpl, err := template.ParseFiles("views/detailProject.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	var Projectdetail = Project{}

	err = connection.Conn.QueryRow(context.Background(), " SELECT * FROM tb_project WHERE id = $1", id).Scan(&Projectdetail.Id, &Projectdetail.Pname, &Projectdetail.Startdate, &Projectdetail.Enddate, &Projectdetail.Description, &Projectdetail.Tech, &Projectdetail.Image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	data := map[string]interface{}{
		"Project": Projectdetail,
	}

	return tmpl.Execute(c.Response(), data)
}

func tmplUpdate(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	tmpl, err := template.ParseFiles("views/updateProject.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	var tplupdate = Project{}

	err = connection.Conn.QueryRow(context.Background(), " SELECT * FROM tb_project WHERE id = $1", id).Scan(&tplupdate.Id, &tplupdate.Pname, &tplupdate.Startdate, &tplupdate.Enddate, &tplupdate.Description, &tplupdate.Tech, &tplupdate.Image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	data := map[string]interface{}{
		"Project": tplupdate,
	}

	return tmpl.Execute(c.Response(), data)
}

func signup(c echo.Context) error {
	var tmpl, err = template.ParseFiles("views/signup.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"messege": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

func signin(c echo.Context) error {
	// sess, _ := session.Get("session", c)
	// flash := map[string]interface{}{
	// 	"FlashStatus":  sess.Values["alertStatus"], // true / false
	// 	"FlashMessage": sess.Values["message"],     // "Register success"
	// }

	// delete(sess.Values, "message")
	// delete(sess.Values, "alertStatus")

	var tmpl, err = template.ParseFiles("views/signin.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"messege": err.Error()})
	}

	return tmpl.Execute(c.Response(), nil)
}

// ==============================================================================

func addProject(c echo.Context) error {
	pName := c.FormValue("projectname")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	description := c.FormValue("description")
	nodeBox := c.FormValue("nodeBox")
	nextBox := c.FormValue("nextBox")
	reactBox := c.FormValue("reactBox")
	typeScriptBox := c.FormValue("typeScriptBox")
	image := c.Get("dataFile").(string)

	var tech []string
	if nodeBox == "NodeJs" {
		tech = append(tech, "NodeJs")
	} else {
		tech = append(tech, "")
	}
	if nextBox == "NextJs" {
		tech = append(tech, "NextJs")
	} else {
		tech = append(tech, "")
	}
	if reactBox == "ReactJs" {
		tech = append(tech, "ReactJs")
	} else {
		tech = append(tech, "")
	}
	if typeScriptBox == "TypeScript" {
		tech = append(tech, "TypeScript")
	} else {
		tech = append(tech, "")
	}

	_, err := connection.Conn.Exec(context.Background(), " INSERT INTO tb_project (pname, startdate, enddate, description, tech, image) VALUES ($1, $2, $3, $4, $5, $6)", pName, startDate, endDate, description, tech, image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/my-project")
}

func deleteProject(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))

	_, err := connection.Conn.Exec(context.Background(), " DELETE FROM tb_project WHERE id = $1", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/my-project")
}

func updateProject(c echo.Context) error {
	id := c.FormValue("id")
	pName := c.FormValue("projectname")
	startDate := c.FormValue("startDate")
	endDate := c.FormValue("endDate")
	description := c.FormValue("description")
	nodeBox := c.FormValue("nodeBox")
	nextBox := c.FormValue("nextBox")
	reactBox := c.FormValue("reactBox")
	typeScriptBox := c.FormValue("typeScriptBox")
	image := c.FormValue("image")

	var tech []string
	if nodeBox == "NodeJs" {
		tech = append(tech, "NodeJs")
	} else {
		tech = append(tech, "")
	}
	if nextBox == "NextJs" {
		tech = append(tech, "NextJs")
	} else {
		tech = append(tech, "")
	}
	if reactBox == "ReactJs" {
		tech = append(tech, "ReactJs")
	} else {
		tech = append(tech, "")
	}
	if typeScriptBox == "TypeScript" {
		tech = append(tech, "TypeScript")
	} else {
		tech = append(tech, "")
	}

	_, err := connection.Conn.Exec(context.Background(), "UPDATE tb_project SET id=$1 ,pname=$2, startdate=$3, enddate=$4, description=$5, tech=$6, image=$7 WHERE id = $1", id, pName, startDate, endDate, description, tech, image)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message ": err.Error()})
	}

	return c.Redirect(http.StatusMovedPermanently, "/my-project")
}

// =======================================

func adduser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	password := c.FormValue("password")

	passwordHash, _ := bcrypt.GenerateFromPassword([]byte(password), 10)

	_, err := connection.Conn.Exec(context.Background(), " INSERT INTO tb_user (name, email, password ) VALUES ($1, $2, $3)", name, email, passwordHash)

	if err != nil {
		redirectWithMessage(c, "Register failed, please try again :)", false, "/signUp")
	}

	return c.Redirect(http.StatusMovedPermanently, "/signIn")
}

//================================================================

func redirectWithMessage(c echo.Context, message string, status bool, path string) error {
	sess, _ := session.Get("session", c)
	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, path)
}

func createduration(Startdate string, Enddate string) string {
	startTime, _ := time.Parse("2000-09-20", Startdate)
	endTime, _ := time.Parse("2000-09-20", Enddate)

	duration := endTime.Sub(startTime)
	days := int(duration.Hours() / 24)
	mounths := int(math.Floor(float64(days) / 30))
	years := int(math.Floor(float64(mounths) / 12))

	if days > 0 && days <= 29 {
		return fmt.Sprint("%d Hari", days)
	} else if days >= 30 && mounths <= 12 {
		return fmt.Sprint("%d Bulan", mounths)
	} else if mounths >= 12 {
		return fmt.Sprint("%d Hari", years)
	} else if days >= 0 && days <= 24 {
		return "1 Hari"
	}
	return ""
}
