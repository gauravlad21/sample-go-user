package dbhelper

import (
	"context"
	"database/sql"
	"fmt"
)

func StartTransaction(ctx context.Context, db *sql.DB) (*sql.Tx, error) {
	return db.Begin()
}

func EndTransaction(ctx context.Context, tx *sql.Tx, txErr error) error {
	if txErr != nil {
		// error in transaction: rollback
		err := tx.Rollback()
		if err != nil {
			fmt.Printf("unable to rollback with error: %v", err)
		}
	} else {
		err := tx.Commit()
		if err != nil {
			fmt.Printf("unable to commit with error: %v", err)
		}
	}
	return nil
}
