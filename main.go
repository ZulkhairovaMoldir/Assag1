package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

//Task1

type Book struct {
	ID         string
	Title      string
	Author     string
	IsBorrowed bool
}

type Library struct {
	Books map[string]Book
}

func (l *Library) AddBook(book Book) {
	if _, exists := l.Books[book.ID]; exists {
		fmt.Println("A book  already exists.")
		return
	}
	l.Books[book.ID] = book
	fmt.Println("Book added successfully.")
}

func (l *Library) BorrowBook(id string) {
	book, exists := l.Books[id]
	if !exists {
		fmt.Println("Book not found.")
		return
	}
	if book.IsBorrowed {
		fmt.Println("Book is already borrowed.")
		return
	}
	book.IsBorrowed = true
	l.Books[id] = book
	fmt.Println("Book borrowed successfully.")
}

func (l *Library) ReturnBook(id string) {
	book, exists := l.Books[id]
	if !exists {
		fmt.Println("Book not found")
		return
	}
	if !book.IsBorrowed {
		fmt.Println("Book is not borrowed")
		return
	}
	book.IsBorrowed = false
	l.Books[id] = book
	fmt.Println("Book returned successfully.")
}

func (l *Library) ListBooks() {
	fmt.Println("Available books:")
	for _, book := range l.Books {
		if !book.IsBorrowed {
			fmt.Printf("ID: %s, Title: %s, Author: %s\n", book.ID, book.Title, book.Author)
		}
	}
}

// Task2
type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Length float64
	Width  float64
}

type Circle struct {
	Radius float64
}

type Square struct {
	Length float64
}

type Triangle struct {
	SideA float64
	SideB float64
	SideC float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (s Square) Area() float64 {
	return s.Length * s.Length
}

func (s Square) Perimeter() float64 {
	return 4 * s.Length
}

func (t Triangle) Area() float64 {
	s := (t.SideA + t.SideB + t.SideC) / 2
	return math.Sqrt(s * (s - t.SideA) * (s - t.SideB) * (s - t.SideC))
}

func (t Triangle) Perimeter() float64 {
	return t.SideA + t.SideB + t.SideC
}

func PrintShapeDetails(s Shape) {
	fmt.Printf("Area: %.2f\n", s.Area())
	fmt.Printf("Perimeter: %.2f\n", s.Perimeter())
}

// Task3
type Employee interface {
	GetDetails() string
}

type FullTimeEmployee struct {
	ID     uint64
	Name   string
	Salary uint32
}

type PartTimeEmployee struct {
	ID          uint64
	Name        string
	HourlyRate  uint64
	HoursWorked float32
}

func (fte FullTimeEmployee) GetDetails() string {
	return fmt.Sprintf("FullTimeEmployee - ID: %d, Name: %s, Salary: %d Tenge", fte.ID, fte.Name, fte.Salary)
}

func (pte PartTimeEmployee) GetDetails() string {
	salary := float64(pte.HourlyRate) * float64(pte.HoursWorked)
	return fmt.Sprintf("PartTimeEmployee - ID: %d, Name: %s, HourlyRate: %d Tenge, HoursWorked: %.1f, Total Salary: %.2f Tenge", pte.ID, pte.Name, pte.HourlyRate, pte.HoursWorked, salary)
}

type Company struct {
	Employees map[string]Employee
}

func (c *Company) AddEmployee(emp Employee) {
	switch e := emp.(type) {
	case FullTimeEmployee:
		c.Employees[strconv.FormatUint(e.ID, 10)] = e
	case PartTimeEmployee:
		c.Employees[strconv.FormatUint(e.ID, 10)] = e
	default:
		fmt.Println("Unknown employee type.")
	}
}

func (c *Company) ListEmployees() {
	for id, emp := range c.Employees {
		fmt.Printf("ID: %s\n%s\n", id, emp.GetDetails())
	}
}

//Task4

type BankAccount struct {
	AccountNumber string
	HolderName    string
	Balance       float64
}

func (ba *BankAccount) Deposit(amount float64) {
	if amount > 0 {
		ba.Balance += amount
		fmt.Printf("Deposited %.2f. New Balance: %.2f\n", amount, ba.Balance)
	} else {
		fmt.Println("Deposit amount must be positive.")
	}
}

func (ba *BankAccount) Withdraw(amount float64) {
	if amount > 0 {
		if ba.Balance >= amount {
			ba.Balance -= amount
			fmt.Printf("Withdrew %.2f. New Balance: %.2f\n", amount, ba.Balance)
		} else {
			fmt.Println("Insufficient balance for withdrawal.")
		}
	} else {
		fmt.Println("Withdrawal amount must be positive.")
	}
}

func (ba *BankAccount) GetBalance() {
	fmt.Printf("Current Balance: %.2f\n", ba.Balance)
}

func Transaction(account *BankAccount, transactions []float64) {
	for _, amount := range transactions {
		if amount > 0 {
			account.Deposit(amount)
		} else {
			account.Withdraw(-amount)
		}
	}
}

func main() {
	//Task1
	library := Library{Books: make(map[string]Book)}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Add")
		fmt.Println("2. Borrow")
		fmt.Println("3. Return")
		fmt.Println("4. List")
		fmt.Println("5. Exit")
		fmt.Print("Enter your choice: ")

		if !scanner.Scan() {
			break
		}
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			fmt.Print("Enter book ID: ")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter book title: ")
			scanner.Scan()
			title := strings.TrimSpace(scanner.Text())

			fmt.Print("Enter book author: ")
			scanner.Scan()
			author := strings.TrimSpace(scanner.Text())

			library.AddBook(Book{ID: id, Title: title, Author: author, IsBorrowed: false})

		case "2":
			fmt.Print("Enter book ID to borrow: ")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())
			library.BorrowBook(id)

		case "3":
			fmt.Print("Enter book ID to return: ")
			scanner.Scan()
			id := strings.TrimSpace(scanner.Text())
			library.ReturnBook(id)

		case "4":
			library.ListBooks()

		case "5":
			fmt.Println("Exiting program.")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
	//Task2
	shapes := []Shape{
		Rectangle{Length: 5, Width: 3},
		Circle{Radius: 4},
		Square{Length: 2.5},
		Triangle{SideA: 3, SideB: 4, SideC: 5},
	}

	for i, shape := range shapes {
		fmt.Printf("Shape %d:\n", i+1)
		PrintShapeDetails(shape)
		fmt.Println("---")
	}

	//Task3
	company := Company{Employees: make(map[string]Employee)}

	e1 := FullTimeEmployee{ID: 1, Name: "Moldir", Salary: 200000}
	e2 := PartTimeEmployee{ID: 2, Name: "Asya", HourlyRate: 5000, HoursWorked: 20.5}
	e3 := FullTimeEmployee{ID: 3, Name: "Dylnaz", Salary: 350000}
	e4 := PartTimeEmployee{ID: 4, Name: "Era", HourlyRate: 4500, HoursWorked: 15}

	company.AddEmployee(e1)
	company.AddEmployee(e2)
	company.AddEmployee(e3)
	company.AddEmployee(e4)

	fmt.Println("Employees:")
	company.ListEmployees()

	//Task4
	account := &BankAccount{
		AccountNumber: "12345678",
		HolderName:    "Zulkhairova Moldir",
		Balance:       0.0,
	}

	fmt.Println(" Bank Account System")
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Deposit")
		fmt.Println("2. Withdraw")
		fmt.Println("3. Get balance")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)

		var amount float64
		switch choice {
		case 1:
			fmt.Print("Enter amount to deposit: ")
			fmt.Scan(&amount)
			account.Deposit(amount)
		case 2:
			fmt.Print("Enter amount to withdraw: ")
			fmt.Scan(&amount)
			account.Withdraw(amount)
		case 3:
			account.GetBalance()
		case 4:
			fmt.Println("Exiting the system")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
