package main

import (
	"fmt"
	"database/sql"
	"log"
	_ "github.com/lib/pq"
)

var DB *sql.DB

type DashboardData struct {
	TotalPatients      int
	TotalDiseases      int
	TotalCountries     int
	LatestDiseaseName  string
	LatestCountry      string
	LatestRecordDate   string
}

func openDB() error {
	connect := "host=localhost port=5432 user=nurymalibekov dbname=db3_assignment sslmode=disable"

	var err error
	DB, err = sql.Open("postgres", connect)
	if err != nil {
		return fmt.Errorf("failed to open database: %v", err)
	}

	if err = DB.Ping(); err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	fmt.Println("Connected to the database successfully!")
	return nil
}

func closeDB() error {		
	err := DB.Close()
	if err != nil {
        log.Fatal(err)
    }
	fmt.Println("Closing database!")
	return nil
}

func getDashboardData() (DashboardData, error) {
	var data DashboardData

	// Query for total statistics
	err := DB.QueryRow("SELECT COUNT(*) FROM Patients").Scan(&data.TotalPatients)
	if err != nil {
		return data, err
	}

	err = DB.QueryRow("SELECT COUNT(*) FROM DiseaseType").Scan(&data.TotalDiseases)
	if err != nil {
		return data, err
	}

	err = DB.QueryRow("SELECT COUNT(*) FROM Country").Scan(&data.TotalCountries)
	if err != nil {
		return data, err
	}

	// Query for the latest disease outbreak record
	err = DB.QueryRow("SELECT disease_code, cname, first_enc_date FROM Discover ORDER BY first_enc_date DESC LIMIT 1").Scan(&data.LatestDiseaseName, &data.LatestCountry, &data.LatestRecordDate)
	if err != nil {
		return data, err
	}

	return data, nil
}