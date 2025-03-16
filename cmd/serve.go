package cmd

import (
	"fmt"
	"log"
	"wmjtyd-iot/internal/server"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var serveParams struct {
	cfgFile        string
	host           string
	port           int
	preforkEnabled bool
	httpEnabled    bool
	grpcEnabled    bool
	wsEnabled      bool
}

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Run a server",
	Long:  `Run web api server & mqtt client`,
	Run: func(cmd *cobra.Command, args []string) {
		// Initialize logger
		logger, err := zap.NewProduction()
		if err != nil {
			log.Fatalf("Failed to initialize logger: %v", err)
		}
		defer logger.Sync()

		// Initialize database
		db, err := gorm.Open(mysql.Open(viper.GetString("database.dsn")), &gorm.Config{})
		if err != nil {
			log.Fatalf("Failed to connect database: %v", err)
		}

		// Initialize viper
		// Always read .env file first
		viper.SetConfigFile(".env")
		if err := viper.ReadInConfig(); err == nil {
			// Set config file based on DEBUG value
			if viper.GetBool("DEBUG") {
				log.Println("Using develop config")
				viper.SetConfigName("config.develop")
			} else {
				viper.SetConfigName("config")
			}
			viper.SetConfigType("yaml")
			fmt.Println(viper.GetBool("DEBUG"))
			// Set database config from .env if exists
			if viper.IsSet("DATABASE.TYPE") {
				viper.Set("database.type", viper.GetString("DATABASE.TYPE"))
				viper.Set("database.host", viper.GetString("DATABASE.HOST"))
				viper.Set("database.port", viper.GetString("DATABASE.PORT"))
				viper.Set("database.user", viper.GetString("DATABASE.USER"))
				viper.Set("database.password", viper.GetString("DATABASE.PASSWORD"))
				viper.Set("database.name", viper.GetString("DATABASE.NAME"))
			}
		} else {
			// If .env not found, use default config
			viper.SetConfigName("config")
			viper.SetConfigType("yaml")
		}

		// Add config search paths
		viper.AddConfigPath(".")
		viper.AddConfigPath("/etc/wmjtyd-iot")
		viper.AddConfigPath("config")

		// Override with specified config file if provided
		if serveParams.cfgFile != "" {
			viper.SetConfigFile(serveParams.cfgFile)
		}

		if err := viper.ReadInConfig(); err != nil {
			log.Fatalf("Error reading config file: %v", err)
		}

		// Get prefork config from yaml, override by command line
		prefork := viper.GetBool("app.prefork")

		if serveParams.preforkEnabled {
			prefork = serveParams.preforkEnabled
		}

		host := serveParams.host
		port := serveParams.port
		if host == "" {
			host = viper.GetString("app.host")
		}
		if port == 0 {
			port = viper.GetInt("app.port")
		}

		// Start HTTP server if enabled
		if serveParams.httpEnabled {
			httpServer := server.NewHTTPServer(fmt.Sprintf("%s:%d", host, port), logger, db)
			httpServer.App = fiber.New(fiber.Config{
				Prefork: prefork,
			})
			httpServer.App.Use(cors.New())
			httpServer.App.Use(recover.New())
			go func() {
				log.Fatal(httpServer.Start())
			}()
		}

		// Block main goroutine
		select {}
	},
}

func init() {
	RootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().StringVarP(&serveParams.cfgFile, "config", "c", "", "config file (default: config.yaml)")
	serveCmd.PersistentFlags().StringVarP(&serveParams.host, "host", "", "127.0.0.1", "server host")
	serveCmd.PersistentFlags().IntVarP(&serveParams.port, "port", "", 3000, "server port")
	serveCmd.PersistentFlags().BoolVarP(&serveParams.preforkEnabled, "prefork", "", false, "enable prefork (default: false)")
	serveCmd.PersistentFlags().BoolVarP(&serveParams.httpEnabled, "http", "", true, "enable HTTP server (default: true)")
	serveCmd.PersistentFlags().BoolVarP(&serveParams.grpcEnabled, "grpc", "", false, "enable gRPC server (default: false)")
	serveCmd.PersistentFlags().BoolVarP(&serveParams.wsEnabled, "websocket", "", false, "enable WebSocket server (default: false)")
}
