package model

import (
	"database/sql"
	"fmt"
	"time"
)

type Customer struct {
	Id            string
	Name          string
	Phone         string
	Email         string
	Address       string
	Join_date     time.Time
	Active_member bool
	Gender        string
}

func AddCustomer(customer Customer) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `INSERT INTO customer (id, name, phone, email, address, join_date, active_member, gender)
	VALUES ($1, $2, $3, $4, $5, $6, $7, $8);`
	result, err := db.Exec(sqlStatement, customer.Id, customer.Name, customer.Phone, customer.Email, customer.Address, customer.Join_date, customer.Active_member, customer.Gender)

	if err != nil {
		return err
	} else {
		fmt.Println("Succesfully Added Customer")
	}

	fmt.Println(result)
	return err
}

func UpdateCustomer(customer Customer) error {
	db := ConnectDB()
	defer db.Close()

	_, err := GetCustomerById(customer.Id)
	if err != nil{
		return err 
	}

	sqlStatement := `UPDATE customer SET name = $2, phone = $3, email = $4, address = $5, join_date = $6, active_member = $7, gender = $8 WHERE id = $1`
	result, err := db.Exec(sqlStatement, customer.Id, customer.Name, customer.Phone, customer.Email, customer.Address, customer.Join_date, customer.Active_member, customer.Gender)

	if err != nil {
		return err
	}else{
		fmt.Println("Succesfully Update Customer!")
	}

	fmt.Println(result)
	return nil 
}

func DeleteCustomer(id string) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `DELETE customer WHERE id = $1`
	result, err := db.Exec(sqlStatement, id)

	if err != nil {
		return err
	}else{
		fmt.Println("Succesfully Delete Customer!")
	}

	fmt.Println(result)
	return nil 
}

func GetAllCustomer() []Customer {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM customer`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	customers := ScanCustomer(rows)
	return customers
}

func GetCustomerById(id string) (Customer, error) {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM customer WHERE id = $1`

	customer := Customer{}
	err := db.QueryRow(sqlStatement, id).Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.Email, &customer.Address, &customer.Join_date, &customer.Active_member, &customer.Gender)

	return customer, err 
}

func ScanCustomer(rows *sql.Rows) []Customer {
	customers := []Customer{}

	for rows.Next() {
		customer := Customer{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Phone, &customer.Email, &customer.Address, &customer.Join_date, &customer.Active_member, &customer.Gender)
		if err != nil{
			panic(err)
		}

		customers = append(customers, customer)
	}
	err := rows.Err()
	if err != nil{
		panic(err)
	}
	return customers
}

func SearchBy(name string) []Customer {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM customer WHERE name LIKE $1`

	rows, err := db.Query(sqlStatement, "%"+name+"%")

	if err != nil{
		panic(err)
	}
	defer rows.Close()
	customers := ScanCustomer(rows)
	return customers
}