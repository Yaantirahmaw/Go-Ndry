package main

import (
	"bufio"
	"fmt"
	"go-ndry/model"
	"go-ndry/utils"
	"os"
	"strconv"
	"strings"
)

func main() {
	initMenu()
}

func initMenu() {
	for {
		fmt.Println(strings.Repeat("=", 50)) 
		fmt.Println(strings.Repeat("=", 14), "Go Laundry", strings.Repeat("=", 14)) 
		fmt.Println("1. View All Data")
		fmt.Println("2. Add New Data")
		fmt.Println("3. Update Data")
		fmt.Println("4. Delete Data")
		fmt.Println("5. Exit")
		fmt.Println(strings.Repeat("=", 50)) 

		scanner := bufio.NewScanner(os.Stdin) 
		fmt.Print("Enter your choice : ") 
		scanner.Scan() 
		choice, _ := strconv.Atoi(scanner.Text()) 

		switch choice { 
			case 1: 
				choice = utils.ViewPrompt() 
	
				switch choice {
				case 1:
					utils.ViewCustomer()
				case 2:
					utils.ViewService()
				case 3:
					utils.ViewOrder()
				case 4:
					utils.ViewOrderDetail()
				default:
					fmt.Println("Invalid choice. Please try again.")
					os.Exit(0)
				}
	
			case 2:
				choice = utils.AddPrompt()
	
				switch choice {
				case 1:
					customer := utils.CreateCustomer()
					err := model.AddCustomer(customer)
					if err != nil {
						fmt.Println("Error", err, "\n", "Please Try Again!")
					}
				case 2:
					service := utils.CreateService()
					err := model.AddService(service)
					if err != nil {
						fmt.Println("Error", err, "\n", "Please Try Again!")
					}
				case 3:
					order := utils.CreateOrder()
					err := model.AddOrder(order)
					if err != nil {
						fmt.Println("Error", err, "\n", "Please Try Again!")
					}
				case 4:
					orderDetail := utils.CreateOrderDetail()
					err := model.AddOrderDetail(orderDetail)
					if err != nil {
						fmt.Println("Error", err, "\n", "Please Try Again!")
					}
				default:
					fmt.Println("Invalid choice. Please try again.")
				}
	
			case 3:
				choice = utils.UpdatePrompt()
				switch choice {
				case 1:
					fmt.Println()
					utils.GetAllCustomerId()
					fmt.Println()
					fmt.Println("Enter Detail Below to Updates : ")
					customer := utils.CreateCustomer()
					err := utils.CheckCustomerId(customer.Id)
					if err != nil {
						fmt.Println("Error", err)
					} else {
						err = model.UpdateCustomer(customer)
						if err != nil {
							fmt.Println("Error", err, "\n", "Please Try Again!")
						}
					}
				case 2:
					fmt.Println()
					utils.ViewAllServiceId()
					fmt.Println()
					fmt.Println("Enter Detail Below to Updates : ")
					service := utils.CreateService()
					err := utils.CheckServiceId(service.Service_id)
					if err != nil {
						fmt.Println("Error", err, "\n", "Please Try Again!")
					} else {
						err = model.UpdateService(service)
						if err != nil {
							fmt.Println("Error", err, "\n", "Please Try Again!")
						}
					}
				case 3:
					fmt.Println()
					utils.ViewAllOrderId()
					fmt.Println("Enter Detail Below to Updates : ")
					order := utils.CreateOrder()
					err := utils.CheckOrderId(order.Customer_id)
					if err != nil {
						fmt.Println("Error", err, "\n", "Please Try Again!")
					} else {
						err := model.UpdateOrder(order)
						if err != nil {
							fmt.Println("Error", err, "\n", "Please Try Again!")
						}
					}
				case 4:
					fmt.Println()
					utils.ViewAllOrderDetailId()
					fmt.Println("Enter Detail Below to Updates : ")
					orderDetail := utils.CreateOrderDetail()
					err := utils.CheckOrderDetailId(orderDetail.Customer_id)
					if err != nil {
						fmt.Println("Error", err, "\n", "Please Try Again!")
					} else {
						err = model.UpdateOrderDetail(orderDetail)
						if err != nil {
							fmt.Println("Error", err, "\n", "Please Try Again!")
						}
					}
				default:
					fmt.Println("Invalid choice. Please try again.")
				}
			case 4:
				choice = utils.DeletePrompt()
	
				switch choice {
				case 1:
					utils.GetAllCustomerId()
					utils.DeleteCustomerUtil()
				case 2:
					utils.ViewAllServiceId()
					utils.DeleteServiceUtil()
				default:
					fmt.Println("Invalid choice. Please try again.")
				}
			case 5:
				fmt.Println("See You Later ^-^")
				os.Exit(0)
			default:
				fmt.Println("Invalid choice. Please try again.")
			}
		}
	}