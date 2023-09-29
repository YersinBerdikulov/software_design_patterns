package main

import (
	"fmt"
)

type DBConnector interface {
	Connect()
}

type DBConnection struct {
	connector DBConnector
}

func (con DBConnection) DBConnect() {
	con.connector.Connect()
}

type MySQLConnection struct {
	ConnectionInfo string
}

func (con MySQLConnection) Connect() {
	fmt.Println("MySQL: " + con.ConnectionInfo)
}

type PostgreSQLConnection struct {
	ConnectionInfo string
}

func (con PostgreSQLConnection) Connect() {
	fmt.Println("PostgreSQL: " + con.ConnectionInfo)
}

type SQLServerConnection struct {
	ConnectionInfo string
}

func (con SQLServerConnection) Connect() {
	fmt.Println("SQL Server: " + con.ConnectionInfo)
}

type MongoDBConnection struct {
	ConnectionInfo string
}

func (con MongoDBConnection) Connect() {
	fmt.Println("MongoDB: " + con.ConnectionInfo)
}

func main() {
	MySQLConnection := MySQLConnection{ConnectionInfo: "MySQL is connected"}
	con := DBConnection{connector: MySQLConnection}
	con.DBConnect()


}

package main

func main() {
	product := "vehicle"
	payWay := 3

	var payment Payment
	switch payWay {
	case 1:
		payment = NewCardPayment("12345", "12345")
	case 2:
		payment = NewPayPalPayment()
	case 3:
		payment = NewQIWIPayment()
	}

	processOrder(product, payment)
}

func processOrder(product string, payment Payment) {
	err := payment.Pay()
	if err != nil {
		return
	}
}


type Payment interface {
	Pay() error
}


type cardPayment struct {
	cardNumber, cvv string
}

func NewCardPayment(cardNumber, cvv string) Payment {
	return &cardPayment {
		cardNumber: cardNumber,
		cvv: cvv,
	}
}

func (p *cardPayment) Pay() error {
	return nil
}

type payPalPayment struct {
}

func NewPayPalPayment() Payment {
	return &payPalPayment{}
}

func (p *payPalPayment) Pay() error {
	return nil
}

type qiwiPayment struct {
}

func NewQIWIPayment() Payment {
	return &qiwiPayment{}
}

func (p *qiwiPayment) Pay() error {
	return nil
}