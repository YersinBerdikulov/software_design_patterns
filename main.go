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
