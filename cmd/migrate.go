package cmd

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/cobra"
)

var migrateParams struct {
	host     string
	port     int
	user     string
	filePath string
}

var migrateCmd = &cobra.Command{
	Use:   "migrate",
	Short: "Migrate database",
	Long:  `Migrate database by creating new database, importing SQL file and setting up user permissions`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("Enter database password: ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		password := scanner.Text()

		// Create database connection
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8&parseTime=True&loc=Local", migrateParams.user, password, migrateParams.host, migrateParams.port)
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			fmt.Printf("Error connecting to database: %v\n", err)
			return
		}
		defer db.Close()

		// Create database
		_, err = db.Exec("CREATE DATABASE IF NOT EXISTS wmjtyd-iot")
		if err != nil {
			fmt.Printf("Error creating database: %v\n", err)
			return
		}

		// Import SQL file
		fileContent, err := os.ReadFile(migrateParams.filePath)
		if err != nil {
			fmt.Printf("Error reading SQL file: %v\n", err)
			return
		}

		_, err = db.Exec("USE wmjtyd-iot")
		if err != nil {
			fmt.Printf("Error switching to wmjtyd-iot database: %v\n", err)
			return
		}

		_, err = db.Exec(string(fileContent))
		if err != nil {
			fmt.Printf("Error executing SQL file: %v\n", err)
			return
		}

		// Create user and grant privileges
		_, err = db.Exec("CREATE USER IF NOT EXISTS 'wmjtyd-iot'@'localhost' IDENTIFIED BY 'wmjtyd-iot'")
		if err != nil {
			fmt.Printf("Error creating user: %v\n", err)
			return
		}

		_, err = db.Exec("GRANT ALL PRIVILEGES ON wmjtyd-iot.* TO 'wmjtyd-iot'@'localhost'")
		if err != nil {
			fmt.Printf("Error granting privileges: %v\n", err)
			return
		}

		fmt.Println("Database migration completed successfully")
	},
}

func init() {
	RootCmd.AddCommand(migrateCmd)
	migrateCmd.Flags().StringVarP(&migrateParams.user, "user", "", "", "Database user with create database right")
	migrateCmd.Flags().StringVarP(&migrateParams.filePath, "file", "", "", "Database file to load")
	migrateCmd.PersistentFlags().StringVarP(&migrateParams.host, "host", "", "127.0.0.1", "database server host")
	migrateCmd.PersistentFlags().IntVarP(&migrateParams.port, "port", "", 3306, "database server port")
	migrateCmd.MarkFlagRequired("user")
	migrateCmd.MarkFlagRequired("file")
}
