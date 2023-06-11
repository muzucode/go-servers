package main

import (
	"log"
	db "muzucode/goroutines/database"
	"muzucode/goroutines/environments"
	"muzucode/goroutines/groups"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/spf13/cobra"
)

func main() {

	// Load .env file vars
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal(err)
	}

	// Connect to DB
	db.ConnectToDatabase()

	// Create root command
	rootCmd := &cobra.Command{
		Use:   "fawn",
		Short: "Manage groups",
	}

	// Load commands
	groups.LoadGroupCommands(rootCmd)
	environments.LoadEnvironmentCommands(rootCmd)

	//

	if err := rootCmd.Execute(); err != nil {
		log.Fatalf("Failed to execute command: %v", err)
	}
}
