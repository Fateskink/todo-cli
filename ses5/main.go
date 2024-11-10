package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/pterm/pterm"
)

type Geo struct {
	Lat string `json:"lat"`
	Lng string `json:"lng"`
}

type Address struct {
	Street  string `json:"street"`
	Suite   string `json:"suite"`
	City    string `json:"city"`
	Zipcode string `json:"zipcode"`
	Geo     Geo    `json:"geo"`
}

type Company struct {
	Name        string `json:"name"`
	CatchPhrase string `json:"catchPhrase"`
	Bs          string `json:"bs"`
}

type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Username string  `json:"username"`
	Email    string  `json:"email"`
	Address  Address `json:"address"`
	Phone    string  `json:"phone"`
	Website  string  `json:"website"`
	Company  Company `json:"company"`
}

func main() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/users")
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	bs, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var results []User
	err = json.Unmarshal(bs, &results)
	if err != nil {
		log.Fatal(err)
	}

	tableData := pterm.TableData{
		{"Name", "Username", "Email", "Address"},
	}

	for _, record := range results {
		tableData = append(tableData, []string{
			record.Name, record.Username, record.Email, record.Address.Street,
		})
	}

	err = pterm.DefaultTable.WithHasHeader().WithBoxed().WithData(tableData).Render()
	if err != nil {
		log.Fatal(err)
	}
}
