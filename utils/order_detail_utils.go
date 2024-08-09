package utils

import (
	"bufio"
	"go-ndry/model"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ValidateOrderDetailId(id string) error {
	hasPrefix := strings.HasPrefix(id, "D")
	if !hasPrefix {
		return fmt.Errorf("the first character must be '%s'", "D")
	}
	return nil
}

func ValidateQuantity(quantity string) (int, error) { 
	newQuantity, err := strconv.Atoi(quantity) 
	if err != nil {                            
		return 0, err 
	}

	return newQuantity, nil 
}

func ValidateSubtotal(subtotal string) (float64, error) {
	return strconv.ParseFloat(subtotal, 64)
}

func CheckOrderDetailId(id string) error {
	_, err := model.GetOrderDetailById(id)
	if err != nil {
		return errors.New("the order_detail doesn't exist")
	}
	return nil
}

func ViewOrderDetail() {
	orderDetails := model.GetAllOrderDetail()

	fmt.Println()
	fmt.Println("OrderDetails : ")
	if len(orderDetails) < 1 {
		fmt.Println("No OrderDetail data (empty table)")
	} else {
		for _, orderDetail := range orderDetails {
			fmt.Printf("%s %s %s %d %v \n", orderDetail.Order_id,
				orderDetail.Service_id, orderDetail.Customer_id, orderDetail.Quantity, orderDetail.Subtotal)
		}
	}
	fmt.Println()
}

func ViewAllOrderDetailId() {
	orderDetails := model.GetAllOrderDetail()

	fmt.Println("AvailableOrderDetail: ")
	for _, orderDetail := range orderDetails {
		fmt.Printf("%s \n", orderDetail.Order_id)
	}
}

func CreateOrderDetail() model.OrderDetail {
	var newOrderDetail model.OrderDetail

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(strings.Repeat("=", 14), "Add New Order Detail", strings.Repeat("=", 14))
	fmt.Println("Enter order Details")

	fmt.Println("Example format for Id 'D001'")
	for {
		fmt.Print("Order Detail Id : ")
		scanner.Scan()
		newOrderDetail.Customer_id = scanner.Text()
		err := ValidateOrderId(newOrderDetail.Order_id)
		err2 := CheckOrderId(newOrderDetail.Customer_id)
		if err != nil {
			fmt.Println(err)
		} else if err2 != nil {
			fmt.Println(err2)
		} else {
			break
		}

	}

	fmt.Println()
	for {
		fmt.Print("OrderDetail id : ")
		scanner.Scan()
		newOrderDetail.Customer_id = scanner.Text()
		err := ValidateOrderDetailId(newOrderDetail.Customer_id)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}
	}
	fmt.Println()
	fmt.Print("Example format for Service_id 'S001'\n")
	for {
		fmt.Print("OrderDetail Service_id  : ")
		scanner.Scan()
		newOrderDetail.Service_id = scanner.Text()
		err := ValidateServiceId(newOrderDetail.Service_id)
		err2 := CheckServiceId(newOrderDetail.Service_id)
		if err != nil {
			fmt.Println(err)
		} else if err2 != nil {
			fmt.Println(err2)
		} else {
			break
		}
	}

	for {
		fmt.Print("Quantity : ")
		scanner.Scan()
		quantity, err := ValidateQuantity(scanner.Text())
		if err != nil {
			fmt.Println("Quantity must be consist of numbers!")
		} else {
			newOrderDetail.Quantity = quantity
			break
		}
	}

	for {
		fmt.Print("Subtotal : ")
		scanner.Scan()
		subtotal := scanner.Text()
		newSubtotal, err := ValidateSubtotal(subtotal)
		
		if err != nil {
			fmt.Println("price must be consist")
		}else{
			newOrderDetail.Subtotal = newSubtotal
			break
		}
	}

	return newOrderDetail 
}
