package main

import (
	"fmt"
	"log"

	"mercury_algo_backend/infra/mysql"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// Load database configuration
	cfg, err := rds.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// Build DSN from config
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
	)

	// Connect to database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Connected to database successfully")

	// Configure the generator
	g := gen.NewGenerator(gen.Config{
		// Output directory for generated code
		OutPath: "./infra/mysql/models/query",

		// Output directory for generated model structs
		ModelPkgPath: "./infra/mysql/models",

		// Generate mode
		Mode: gen.WithDefaultQuery | gen.WithQueryInterface | gen.WithoutContext,

		// Generate model global configuration
		FieldNullable:     true,  // generate pointer for nullable fields
		FieldCoverable:    false, // generate cover methods for struct fields
		FieldSignable:     false, // generate signatures for struct fields
		FieldWithIndexTag: true,  // generate index tags for struct fields
		FieldWithTypeTag:  true,  // generate type tags for struct fields
	})

	// Use the connected database
	g.UseDB(db)

	// Generate basic CRUD methods for all tables
	// Option 1: Generate from ALL tables in the database
	allTables := g.GenerateAllTable()

	// Apply basic CRUD methods
	g.ApplyBasic(allTables...)

	// Execute the generation
	g.Execute()

	log.Println("Code generation completed successfully!")
	log.Println("Generated files:")
	log.Println("  - Models in: ./models/")
	log.Println("  - Query methods in: ./models/query/")
}
