package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "dennis"
	password = ""
	dbname   = "googlephotos_dev"
)

// User defines the table for users
type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		" dbname=%s sslmode=disable",
		host, port, user, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&User{})

	var u User
	maxId := 3

	db.Where("id <= ?", maxId).First(&u)
	if db.Error != nil {
		panic(db.Error)
	}

	fmt.Println(u)

	// name, email := getInfo()
	// u := &User{
	// 	Name:  name,
	// 	Email: email,
	// }

	// if err = db.Create(u).Error; err != nil {
	// 	panic(err)
	// }

	// fmt.Printf("%+v\n", u)
}

func getInfo() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("What is your name?")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)
	fmt.Println("What is your email?")
	email, _ := reader.ReadString('\n')
	email = strings.TrimSpace(email)
	return name, email
}

// func main() {

// 	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+"dbname=%s sslmode=disable",
// 		host, port, user, dbname)

// 	fmt.Println(psqlInfo)

// 	db, err := sql.Open("postgres", psqlInfo)
// 	if err != nil {
// 		panic(err)
// 	}

// 	defer db.Close()

// 	err = db.Ping()
// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Successfully connected!")

// 	db.AutoMigrate(&User{})

// How to insert a row into table `users`
// var id int
// row := db.QueryRow(`
// 	INSERT INTO users(name, email)
// 	VALUES($1, $2) RETURNING id`,
// 	"Meg I", "meg@mail.com")
// err = row.Scan(&id)
// if err != nil {
// 	panic(err)
// }

// How to query a row from table `users`
// var id int
// var name, email string
// row := db.QueryRow(`
// 	SELECT id, name, email
// 	FROM users
// 	WHERE id=$1`, 2)
// err = row.Scan(&id, &name, &email)
// if err != nil {
// 	panic(err)
// }
//
// fmt.Println("ID: ", id, "Name:", name, "Email:", email)

// How to query multiple rows from table `users`
// var id int
// var name, email string
// rows, err := db.Query(`
// 	SELECT id, name, email
// 	FROM users
// 	WHERE email = $1
// 	OR ID > $2`, "dennis@mail.com", 0)
// if err != nil {
// 	panic(err)
// }

// for rows.Next() {
// 	rows.Scan(&id, &name, &email)
// 	fmt.Println("ID:", id, "Name:", name, "Email:", email)
// }

// Insert orders to into table `orders`
// var id int
// for i := 1; i < 6; i++ {
// 	// Create some fake data
// 	userID := 1
// 	if i > 3 {
// 		userID = 2
// 	}
// 	amount := 1000 * i
// 	description := fmt.Sprintf("USB-C Adapter x%d", i)

// 	err := db.QueryRow(`
// 		INSERT INTO orders (user_id, amount, description)
// 		VALUES ($1, $2, $3)
// 		RETURNING id`,
// 		userID, amount, description).Scan(&id)

// 	if err != nil {
// 		panic(err)
// 	}

// 	fmt.Println("Created an order with the ID:", id)
// }

// }
