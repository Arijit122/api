package main

import (
	"fmt"

	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/mvc"
	"github.com/kataras/iris/v12/sessions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	cookieNameForSessionID = "mycookiesessionnameid"
	sess                   = sessions.New(sessions.Config{Cookie: cookieNameForSessionID})
)

type register struct {
	Username  string `gorm:"column:username"`
	Password  string `gorm:"column:password"`
	Firstname string `gorm:"column:firstname"`
	Lastname  string `gorm:"column:lastname"`
	Age       int    `gorm:"column:age"`
}

type User struct {
	Id     int    `gorm:"primary key,autoIncrement:true"`
	Uname  string `gorm:"unique"`
	Passwd string
	Fname  string
	Lname  string
	Ages   int
}

func DB() gorm.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "postgres"
		password = "Arijit123@"
		dbname   = "Authentication"
	)
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d",
		host, user, password, dbname, port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	} else {
		fmt.Println("Database Connected...")
	}

	db.AutoMigrate(&User{})
	db.AutoMigrate(&register{})
	return *db
}

func main() {
	db := DB()
	fmt.Println(db)
	app := iris.New()
	mvc.Configure(app.Party("/root"), myMVC)
	app.Listen(":8080")

}

func myMVC(app *mvc.Application) {

	app.Handle(new(MyController))
}

type MyController struct{}

func (m *MyController) BeforeActivation(b mvc.BeforeActivation) {

	b.Handle("POST", "/register/{username:string}/{password:string}/{firstname:string}/{lastname:string}/{age:int}", "Register")
	// b.Handle("POST", "/login/{username:string}/{password:string}", "Login")
	// b.Handle("GET", "/logout", "Logout")
	// b.Handle("GET", "/data", "Data")

}

func (m *MyController) Register(ctx iris.Context, username string, password string, firstname string, lastname string, age int) any {
	fmt.Println("")
	fmt.Println("Register users")
	fmt.Println("")

	un := username
	pw := password
	fn := firstname
	ln := lastname
	ag := age
	usr := User{
		Uname:  un,
		Passwd: pw,
		Fname:  fn,
		Lname:  ln,
		Ages:   ag,
	}

	db := DB()
	db.Create(&usr)
	fmt.Println("")
	fmt.Println("Creating a new user")

	return "Hi,Welcome new user"
}
