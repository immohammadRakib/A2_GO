package main

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"

)


var DB *pgxpool.Pool

func InitDB() {
	ctx := context.Background()

	connStr := "postgres://postgres:secret_password@localhost:5432/tech_tracker?sslmode=disable"

	var err error

	DB, err = pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("Facing Problem to Connect Database: %v\n", err)
	}

	_, err = DB.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			name VARCHAR(150) NOT NULL,
			email VARCHAR(255) NOT NULL UNIQUE,
			password VARCHAR(255) NOT NULL,
			role VARCHAR(20) NOT NULL DEFAULT 'contributor' CHECK (role IN ('contributor', 'maintainer')),
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);

		CREATE TABLE IF NOT EXISTS issues (
			id SERIAL PRIMARY KEY,
			title VARCHAR(150) NOT NULL,
			description TEXT NOT NULL,
			type VARCHAR(20) NOT NULL CHECK (type IN ('bug', 'feature_request')),
			status VARCHAR(20) NOT NULL DEFAULT 'open' CHECK (status IN ('open', 'in_progress', 'resolved')),
			reporter_id INT NOT NULL,
			created_at TIMESTAMP NOT NULL DEFAULT NOW(),
			updated_at TIMESTAMP NOT NULL DEFAULT NOW()
		);
	`)
	
	if err != nil {
		log.Fatalf("Facing Issue To Connect Database: %v\n", err)
	}

	log.Println("Database connected & tables verified successfully!")
}
