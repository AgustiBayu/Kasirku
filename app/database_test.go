package app_test

import (
	"kasirku/app"
	"testing"
)

func TestDBConnetion(t *testing.T) {
	db := app.DB()
	if db == nil {
		t.Fatal("expected db connection, got nil")
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal("Failed get *sql.DB")
	}

	if err := sqlDB.Ping(); err != nil {
		t.Fatalf("failed to ping database: %v", err)
	}
}
