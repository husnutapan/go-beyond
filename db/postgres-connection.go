package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

const (
	dbhost = "localhost"
	dbport = "5432"
	dbuser = "postgres"
	dbpass = "root"
	dbname = "test-go"
)

type employee struct {
	ID          int
	Name        string
	CreatedDate time.Time
}

type employees struct {
	Employees []employee
}

func main() {
	beforeDbConfig()
	http.HandleFunc("/api/employee/", repoHandler)
}

func repoHandler(writer http.ResponseWriter, request *http.Request) {
	employees := employees{}
	data := queryEmployees(&employees)

	out, _ := json.Marshal(data)
	fmt.Println(string(out))

}

func queryEmployees(employees *employees) employees {
	rows, _ := db.Query(`
		SELECT
			id,
			name,
			created_date,
		FROM employee`)

	defer rows.Close()

	for rows.Next() {
		employee := employee{}
		rows.Scan(
			&employee.ID,
			&employee.Name,
			&employee.CreatedDate,
		)
		employees.Employees = append(employees.Employees, employee)
	}
	return *employees
}

func beforeDbConfig() map[string]string {
	configuration := make(map[string]string)
	configuration[dbhost] = dbhost
	configuration[dbport] = dbport
	configuration[dbuser] = dbuser
	configuration[dbpass] = dbpass
	configuration[dbname] = dbname
	return configuration
}
