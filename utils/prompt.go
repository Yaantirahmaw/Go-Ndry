package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func AddPrompt() int {
	fmt.Println(strings.Repeat("=", 14), "Add Data", strings.Repeat("=", 14))
	fmt.Println("1. Add Customer")
	fmt.Println("2. Add Service")
	fmt.Println("3. Add Order")
	fmt.Println("4. Add Order Detail")
	fmt.Print("Enter Your Choice : ")

	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())

	return choice
}

func ViewPrompt() int {
	fmt.Println(strings.Repeat("=", 16), "View Data", strings.Repeat("=", 16))
	fmt.Println("1. View Customer")
	fmt.Println("2. View Service")
	fmt.Println("3. View Order")
	fmt.Println("4. View Order Detail")
	fmt.Print("Enter Your Choice : ")

	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())

	return choice
}

func UpdatePrompt() int {
	fmt.Println(strings.Repeat("=", 14), "Update Data", strings.Repeat("=", 14))
	fmt.Println("1. Update Customer")
	fmt.Println("2. Update Service")
	fmt.Println("3. Update Order")
	fmt.Println("4. Update Order Detail")
	fmt.Print("Enter your choice : ")

	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())

	return choice
}

func DeletePrompt() int {
	fmt.Println(strings.Repeat("=", 16), "Delete Data", strings.Repeat("=", 16))
	fmt.Println("1. Delete Customer")
	fmt.Println("2. Delete Service")
	fmt.Print("Enter Your Choice : ")

	scanner.Scan()
	choice, _ := strconv.Atoi(scanner.Text())

	return choice
}