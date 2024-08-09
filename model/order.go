package model

import (
	"database/sql"
	f "fmt"
	t "time"
)

type Order struct {
	Order_id     string
	Customer_id  string
	Order_date   t.Time
	Total_amount int
	Status       string
}

func AddOrder(order Order) error {
	db := ConnectDB()
	defer db.Close()
	
	sqlStatement := `INSERT INTO "order" (order_id, customer_id, order_date, total_amount, status)
	VALUES ($1, $2, $3, $4, $5);`
	result, err := db.Exec(sqlStatement, order.Order_id, order.Customer_id, order.Order_date, order.Total_amount, order.Status)

	if err != nil{
		return err
	}else{
		f.Println("Successfuly Added Order!")
	}

	f.Println(result)
	return err
}

func UpdateOrder(order Order) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `UPDATE "order" SET order_id = $1, order_date = $3, total_amount = $4, status = $5 WHERE customer_id = $2`
	result, err := db.Exec(sqlStatement, order.Order_id, order.Customer_id, order.Order_date, order.Total_amount, order.Status)

	if err != nil {
		return err
	}else{
		f.Println("Successfuly Update Order!")
	}

	f.Println(result)
	return err
}

func GetAllOrder() []Order {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM "order"`

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	Orders := scanOrder(rows)
	return Orders
}

func GetOrderById(id string) (Order, error) {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM "order" WHERE id = $2`

	order := Order{}
	err := db.QueryRow(sqlStatement, id).Scan(&order.Order_id, &order.Order_date, &order.Total_amount, &order.Status)

	return order, err
}

func scanOrder(rows *sql.Rows) []Order{
	orders := []Order{}

	for rows.Next(){
		order := Order{}
		err := rows.Scan(&order.Order_id, &order.Order_date, &order.Total_amount, &order.Status)
		if err != nil {
			panic(err)
		}

		orders = append(orders, order)
	}

	err := rows.Err()
	if err != nil{
		panic(err)
	}
	return orders
}
