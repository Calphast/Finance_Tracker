/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"finance-tracker/cmd"
	"finance-tracker/db"
)

func main() {
	db.InitDB()
	cmd.Execute()
}
