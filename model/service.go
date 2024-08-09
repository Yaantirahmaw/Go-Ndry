package model

import (
	"database/sql"
	"fmt"
)

type Service struct {
	Service_id  string
	Name        string
	Description string
	Price       int
	Duration    string
}

func AddService(service Service) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `INSERT INTO service (service_id, name, description, price, duration)
	VALUES ($1, $2, $3, $4, $5)`

	result, err := db.Exec(sqlStatement, service.Service_id, service.Name, service.Description, service.Price, service.Duration)

	if err != nil {
		return err
	} else {
		fmt.Println("Successfuly Added Customer!")
	}
	fmt.Println(result)
	return err
}

func UpdateService(service Service) error {
	db := ConnectDB()
	defer db.Close()

	_, err := GetServiceById(service.Service_id)
	if err != nil {
		return err 
	}

	sqlStatement := `UPDATE service SET service_name = $2, description = $3, price = $4, duration = $5 WHERE service_id = $1`
	result, err := db.Exec(sqlStatement, service.Service_id, service.Name, service.Description, service.Price, service.Duration)

	if err != nil{
		fmt.Println(err)
	}else{
		fmt.Println("Successfuly Update Service")
	}
	fmt.Println(result)
	return nil 
}

func DeleteService(id string) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `DELETE FROM service WHERE service_id = $1`
	result, err := db.Exec(sqlStatement, id)

	if err != nil {
		return err
	} else {
		fmt.Println("Succesfully Delete Service")
	}

	fmt.Println(result)
	return err
}

func GetAllService() []Service {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := "SELECT service_id, name, description, price, duration FROM service"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	services := scanService(rows)
	return services
}

func GetServiceById(service_id string) (Service, error) {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM service WHERE service_id = $1`

	service := Service{}
	err := db.QueryRow(sqlStatement, service_id).Scan(&service.Service_id, &service.Name, &service.Description, &service.Price, &service.Duration)

	return service, err 
}

func scanService(rows *sql.Rows) []Service {
	services := []Service{}

	for rows.Next() {
		service := Service{}
		err := rows.Scan(&service.Service_id, &service.Name, &service.Description, &service.Price, &service.Duration)
		if err != nil {
			panic(err)
		}

		services = append(services, service)
	}

	err := rows.Err()
	if err != nil {
		panic(err)
	}

	return services
}
