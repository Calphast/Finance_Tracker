package db

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3" // Import the SQLite driver
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "finance.db")
	if err != nil {
		log.Fatal(err)
	}

	createTables()
}

func createTables() {
	query := `
	CREATE TABLE IF NOT EXISTS transactions (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		type TEXT CHECK(type IN ('income', 'expense')) NOT NULL,
		amount REAL NOT NULL,
		date TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);`
	_, err := DB.Exec(query)
	if err != nil {
		log.Fatal("Error creating table:", err)
	}
}

func AddTransactions(txType string, amount float64) error {
	query := `INSERT INTO transactions (type, amount) VALUES (?, ?)`
	_, err := DB.Exec(query, txType, amount)
	return err
}

func GetSummary() (float64, float64, float64, error) {
	var income, expenses float64

	queryIncome := `SELECT SUM(amount) FROM transactions WHERE type = 'income'`
	queryExpense := `SELECT SUM(amount) FROM transactions WHERE type = 'expense'`

	err := DB.QueryRow(queryIncome).Scan(&income)
	if err != nil && err != sql.ErrNoRows {
		return 0, 0, 0, err
	}

	err = DB.QueryRow(queryExpense).Scan(&expenses)
	if err != nil && err != sql.ErrNoRows {
		return 0, 0, 0, err
	}

	balance := income - expenses
	return income, expenses, balance, nil
}
