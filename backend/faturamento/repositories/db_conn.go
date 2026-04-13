package repositories

import (
	"database/sql"
	"fmt"
	"os"
	_ "github.com/lib/pq"
)

var (
	dbUser,dbName,dbPassword,dbHost string
	dbPort int
)


func init(){
	dbUser = os.Getenv("POSTGRES_USER")
	dbName = os.Getenv("POSTGRES_DB")
	dbPassword = os.Getenv("POSTGRES_PASSWORD")
	// dbPort=5432
	dbHost = "db"
}

func ConnectDB()(*sql.DB,error){
	psqlInfo := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
	dbUser,dbPassword,dbHost,dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Conected to "+dbName)
	return db,nil
}