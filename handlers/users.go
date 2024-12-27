package handlers

import (
	"context"
	"go-server/config"
	"go-server/models"
	"html/template"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Template struct {
    templates *template.Template
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}



func GetUsers (c echo.Context) error{
	renderer := &Template{
		templates: template.Must(template.ParseGlob("./public/views/*.html")),
	}
	e := echo.New()
	e.Renderer = renderer

	var users []models.User
	
	db, err := config.Conn()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close(context.Background())

	rows, err := db.Query(context.Background(), "SELECT * FROM users")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Salary); err != nil{
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		users = append(users, user)
	}
	rows.Close()

	return c.Render(http.StatusOK, "all-users", map[string]interface{}{
		"users": users,
	})
}

func GetUser (c echo.Context) error{
	var id string;
	id = c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, "")
	}
	
	db, err := config.Conn()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	defer db.Close(context.Background())	

	row := db.QueryRow(context.Background(), "SELECT * FROM users where id=$1", (id))
	var user models.User;
	if err := row.Scan(&user.ID, &user.Name, &user.Age, &user.Salary); err != nil{
		return c.JSON(http.StatusNoContent, "No rows found")
	}

	return c.Render(http.StatusOK, "user", map[string]interface{}{
		"user": user,
	})
}

func SearchUser (c echo.Context) error{
	var users []models.User

	var name string = c.FormValue("name")
	var age string = c.FormValue("age")
	var salary string = c.FormValue("salary")

	var sql = "SELECT * FROM users where 1=1"

	if name != ""{
		sql = sql + " AND name ILIKE '"+name+"%'"
	}

	if age != ""{
		sql = sql + " AND age="+age
	}

	if salary != ""{
		sql = sql + " AND salary="+salary
	}
	
	db, err := config.Conn()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, "")
	}
	defer db.Close(context.Background())	
	
	rows, err := db.Query(context.Background(), sql)
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age, &user.Salary); err != nil{
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		users = append(users, user)
	}
	rows.Close()

	return c.Render(http.StatusOK, "all-users", map[string]interface{}{
		"users": users,
	})
	
}

func AddUser (c echo.Context) error{
	var name string = c.FormValue("name")
	var age string = c.FormValue("age")
	var salary string = c.FormValue("salary")

	if name == "" || age == "" || salary == "" {
		return c.String(http.StatusOK, "Invalid credentials") 
	}

	db, err := config.Conn()
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	defer db.Close(context.Background())
	
	res, err := db.Exec(context.Background(), "INSERT INTO users (name, age, salary) VALUES ($1, $2, $3)", (name), (age), (salary))
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	if res.RowsAffected() > 0 {
		return c.String(http.StatusOK, "user created") 
	}
	
	return c.String(http.StatusInternalServerError, err.Error())
}

func UpdateUser (c echo.Context) error{
	var id string = c.Param("id");
	var name string = c.FormValue("name")
	var age string = c.FormValue("age")
	var salary string = c.FormValue("salary")


	if id == "" || (name == "" && age == "" && salary == "") {
		return c.JSON(http.StatusBadRequest, "")
	}
	
	var sql = "UPDATE users SET "

	if name != "" {
		sql = sql + "name='"+name+"'"
	}

	if age != "" {
		sql = sql + "age="+age
	}

	if salary != "" {
		sql = sql + "salary="+salary
	}

	db, err := config.Conn()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close(context.Background())

	res, err := db.Exec(context.Background(), sql)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if res.RowsAffected() > 0 {
		return c.JSON(http.StatusOK, "User updated")
	}
	
	return c.JSON(http.StatusInternalServerError, err.Error())
}

func DeleteUser (c echo.Context) error{
	var id string = c.Param("id")

	if id == "" {
		return c.JSON(http.StatusBadRequest, "Missing attributes")
	}

	db, err := config.Conn()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer db.Close(context.Background())

	res, err := db.Exec(context.Background(), "DELETE FROM users WHERE id=$1", (id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	if res.RowsAffected() > 0 {
		return c.Render(http.StatusOK, "user", map[string]interface{}{
			"user": nil,
		})
	}

	return c.JSON(http.StatusInternalServerError, err.Error())
}

