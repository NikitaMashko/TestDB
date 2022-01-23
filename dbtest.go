package main
import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "nikita"
	password = "0000"
	dbname   = "productdb"
)

type Product struct {
	Id int
	Model string
	Company string
	Price int
}
var database *sql.DB

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	rows, err := database.Query("select * from Product")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	products := []Product{}

	for rows.Next(){
		p := Product{}
		err := rows.Scan(&p.Id, &p.Model, &p.Company, &p.Price)
		if err != nil{
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}

	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, products)
}

func main() {

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
	}
	database = db
	defer db.Close()
	http.HandleFunc("/", IndexHandler)

	fmt.Println("Server is listening...")
	http.ListenAndServe(":8080", nil)
}