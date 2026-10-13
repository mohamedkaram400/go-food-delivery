package tests

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/joho/godotenv"
	"github.com/mohamed-karam/go-food-delivery/identity-service/conn"
	"github.com/mohamed-karam/go-food-delivery/identity-service/entity"
	"github.com/mohamed-karam/go-food-delivery/identity-service/repo"
	"gorm.io/gorm"
)

func TestRegisterTransactionRollback(t *testing.T) {

	// dsn := "root:qazwsx123@tcp(localhost:3306)/identity_service?charset=utf8mb4&parseTime=True&loc=Local"
	if err := godotenv.Load("../.env"); err != nil {
        t.Fatalf("failed to load .env: %v", err)
    }
	dsn := os.Getenv("DSN")
	if dsn == "" {
		t.Fatal("DSN must point to a dedicated test database")
	}

	db, err := conn.ConnectMySQL(dsn)
	if err != nil {
		t.Fatalf("failed to connect to test database: %v", err)
	}

	identityRepo := repo.NewIdentityRepo(db)
	ctx := context.Background()

	// Use a unique email so repeated test runs do not collide.
	email := fmt.Sprintf("rollback-%d@example.com", time.Now().UnixNano())

	user := &entity.User{
		Name:         "Rollback Test",
		Email:        email,
		Phone:        fmt.Sprintf("9%d", time.Now().UnixNano()%1000000000),
		RoleID:       1,
		PasswordHash: "test-password-hash",
	}

	err = identityRepo.Transaction(ctx, func(tx *gorm.DB) error {
		if err := identityRepo.Register(tx, user); err != nil {
			return err
		}

		// Simulate a failure after the user has been inserted.
		return errors.New("simulated refresh token storage failure")
	})

	if err == nil {
		t.Fatal("expected transaction to fail")
	}

	// The transaction should have rolled back the inserted user.
	var count int64
	err = db.WithContext(ctx).
		Model(&entity.User{}).
		Where("email = ?", email).
		Count(&count).
		Error

	if err != nil {
		t.Fatalf("failed to check whether user exists: %v", err)
	}

	if count != 0 {
		t.Fatalf("expected user insert to be rolled back, but found %d user(s)", count)
	}

	t.Log("Transaction rollback verified successfully")
}