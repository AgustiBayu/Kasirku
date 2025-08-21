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

	err := db.Ping()
	if err != nil {
		t.Fatalf("failed to ping database: %v", err)
	}
}
