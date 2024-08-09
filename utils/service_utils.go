package utils

import (
	"fmt"
	"strings"
	"bufio"
	"strconv"
	"errors"
	"go-ndry/model"
	"os"
)

func ValidateServiceId(service_id string) error {
	hasPrefix := strings.HasPrefix(service_id, "S")
	if !hasPrefix {
		return fmt.Errorf("the first character must be '%s'", "S")
	}
	return nil 
}

func ValidateServicePrice(price string) (int, error) {
	newPrice, err := strconv.Atoi(price)
	if err != nil {
		return 0, err
	}

	return newPrice, nil
}

func CheckServiceId(service_id string) error {
	_, err := model.GetServiceById(service_id)
	if err != nil {
		return errors.New("service doesn't exist")
	}
	return nil
}

func ViewService() {
	services := model.GetAllService()

	fmt.Println()
	fmt.Println("Services : ")
	for _, service := range services {
		fmt.Printf("%s, %s, %s, %v, %s\n", service.Service_id, service.Name, service.Description, service.Price, service.Duration)
	}
	fmt.Println()
}

func ViewAllServiceId() {
	services := model.GetAllService()

	fmt.Println("AvailableServices: ")
	for _, service := range services {
		fmt.Printf("%s \n", service.Service_id)
	}
}

func DeleteServiceUtil() {
	var service_id string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Please enter service id to be deleted : ")
	scanner.Scan()
	service_id = scanner.Text()
	err := CheckServiceId(service_id)
	if err != nil {
		fmt.Println(err)
	} else {
		err := model.DeleteService(service_id)
		if err != nil {
			fmt.Println("Error", err, "\n the service already has relation")
		}
	}
}

func CreateService() model.Service {
	var newService model.Service

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(strings.Repeat("=", 14), "Add New service", strings.Repeat("=", 14))
	fmt.Println("Enter Service Details")

	fmt.Println("Example format for Service ID 'S001'")
	for {
		fmt.Print("Service Id : ")
		scanner.Scan()
		newService.Service_id = scanner.Text()
		err := ValidateServiceId(newService.Service_id)
		if err != nil {
			fmt.Println(err)
		} else {
			break
		}

	}

	fmt.Print("Service Name : ")
	scanner.Scan()
	newService.Name = scanner.Text()

	fmt.Print("Service Description : ")
	scanner.Scan()
	newService.Description = scanner.Text()

	// fmt.Print("Service Satuan : ")
	// scanner.Scan()
	// newService.Satuan = scanner.Text()
 
	for {

		fmt.Print("Service Price : ")
		scanner.Scan()
		price := scanner.Text()
		newPrice, err := ValidateServicePrice(price)
		if err != nil {
			fmt.Println("Price must be consist of numbers!")
		} else {
			newService.Price = newPrice
			break
		}
	}

	fmt.Print("Service Duration : ")
	scanner.Scan()
	newService.Duration = scanner.Text()

	return newService
}
