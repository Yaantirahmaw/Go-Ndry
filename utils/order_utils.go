package utils

import (
	"bufio"
	"errors"
	"fmt"
	"go-ndry/model"
	"os"
	"strings"
)

func ValidateOrderId(order_id string) error {
	hasPrefix := strings.HasPrefix(order_id, "D")
	if !hasPrefix {
		return fmt.Errorf("the first character must be '%s'", "D")
	}
	return nil
}

func CheckOrderId(order_id string) error {
	_, err := model.GetOrderById(order_id)
	if err != nil {
		return errors.New("order doesn't exist")
	}
	return nil
}

func ViewOrder() {
	orders := model.GetAllOrder()

	fmt.Println()
	fmt.Println("Orders: ")
	if len(orders) < 1 {
		fmt.Println("No orders data (empty table)")
	} else {
		for _, order := range orders {
			fmt.Printf("%s, %s, %s, %v, %s", order.Order_id, order.Customer_id, order.Order_date, order.Total_amount, order.Status)
		}
	}
	fmt.Println()
}

func ViewAllOrderId() {
	orders := model.GetAllOrder()

	fmt.Println("Available Orders: ")
	for _, order := range orders {
		fmt.Printf("%s \n", order.Order_id)
	}
}

func CreateOrder() model.Order {
	var newOrder model.Order

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(strings.Repeat("=", 5), "Add New Order", strings.Repeat("=", 5))
	fmt.Println("Enter Order Details")

	fmt.Println("Example format for Id: 'D001'")
	for {
		fmt.Print("Order ID: ")
		scanner.Scan()
		newOrder.Order_id = scanner.Text()

		err := ValidateOrderId(newOrder.Order_id)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}

	fmt.Println()
	fmt.Print("Format yyyy-mm-dd, example: 2023-12-02\n")
	for {
		fmt.Print("Order Date: ")
		scanner.Scan()
		fmt.Println(scanner.Text())
		date, err := DateValidation(scanner.Text())

		if err != nil {
			fmt.Println("Please enter a valid date!")
		} else {
			newOrder.Order_date = date
			break
		}
	}

	fmt.Println("Example for Customer ID: 'C001'")
	for {
		fmt.Print("Customer ID: ")
		scanner.Scan()
		newOrder.Customer_id = scanner.Text()
		err := ValidationCustomerId(newOrder.Customer_id)
		err2 := CheckCustomerId(newOrder.Customer_id)

		if err != nil {
			fmt.Println(err)
		} else if err2 != nil {
			fmt.Println(err2)
		}else{
			break
		}
	}
	return newOrder
}
