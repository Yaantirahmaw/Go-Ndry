package model

import (
	"database/sql"
	"fmt"
)

type OrderDetail struct{
	Detail_id string
	Order_id string
	Service_id string 
	Customer_id string 
	Quantity int
	Subtotal float64
}

func AddOrderDetail(orderDetail OrderDetail) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `INSERT INTO order_detail (order_id, service_id, customer_id, quantity, subtotal) 
    VALUES ($1, $2, $3, $4, $5);`
	result, err := db.Exec(sqlStatement, orderDetail.Order_id, orderDetail.Service_id, orderDetail.Customer_id, orderDetail.Quantity, orderDetail.Subtotal)

	if err != nil {
		return err
	} else {
		fmt.Println("Succesfully Added OrderDetail")
	}

	fmt.Println(result)
	return err
}

func UpdateOrderDetail(orderDetail OrderDetail) error {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `UPDATE order_detail SET order_id = $1, service_id = $2, quantity = $4, subtotal = $5 WHERE customer_id = $3`
	result, err := db.Exec(sqlStatement, orderDetail.Customer_id, orderDetail.Order_id,
		orderDetail.Service_id, orderDetail.Quantity)

	if err != nil {
		return err
	} else {
		fmt.Println("Succesfully Updated OrderDetail")
	}

	fmt.Println(result)
	return err
}

func GetAllOrderDetail() []OrderDetail {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := "SELECT order_id, service_id, customer_id, quantity, subtotal FROM order_detail"

	rows, err := db.Query(sqlStatement)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	orderDetails := scanOrderDetail(rows)
	return orderDetails
}

func GetOrderDetailById(customer_id string) (OrderDetail, error) {
	db := ConnectDB()
	defer db.Close()

	sqlStatement := `SELECT * FROM order_detail WHERE customer_id = $3`

	orderDetail := OrderDetail{}
	err := db.QueryRow(sqlStatement, customer_id).Scan(&orderDetail.Order_id, orderDetail.Service_id, &orderDetail.Customer_id, &orderDetail.Quantity, &orderDetail.Subtotal)

	return orderDetail, err
}

func scanOrderDetail(rows *sql.Rows) []OrderDetail {
	orderDetails := []OrderDetail{}

	for rows.Next() {
		orderDetail := OrderDetail{}
		err := rows.Scan(&orderDetail.Order_id, &orderDetail.Service_id, &orderDetail.Customer_id, &orderDetail.Quantity, &orderDetail.Subtotal)
		if err != nil {
			panic(err)
		}

		orderDetails = append(orderDetails, orderDetail)
	}

	err := rows.Err()
	if err != nil {
		panic(err)
	}

	return orderDetails
}
