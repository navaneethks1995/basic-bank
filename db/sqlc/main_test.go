package db

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
)

const POSTGRESQL_URI = "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable"

var testQueries *Queries
var connPool *pgxpool.Pool

// TestMain sets up a test database connection and runs the tests.
func TestMain(m *testing.M) {
	var err error
	connPool, err = pgxpool.New(context.Background(), POSTGRESQL_URI)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	if connPool == nil {
		fmt.Fprintln(os.Stderr, "connPool should not be nil")
		os.Exit(1)
	}

	testQueries = New(connPool)

	if testQueries == nil {
		fmt.Fprintln(os.Stderr, "testQueries should not be nil")
		os.Exit(1)
	}

	// close the db connection pool after all tests have finished
	defer connPool.Close()

	os.Exit(m.Run())
}
