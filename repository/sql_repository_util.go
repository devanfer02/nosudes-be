package repository

import (
	"fmt"
	"context"

	"github.com/devanfer02/nosudes-be/utils/logger"
	"github.com/devanfer02/nosudes-be/utils/layers"
	"github.com/devanfer02/nosudes-be/domain"

	"github.com/jmoiron/sqlx"
)

func execStatement(conn *sqlx.DB, ctx context.Context, query string, args ...interface{}) error {
	stmt, err := conn.PrepareContext(ctx, query)

	if err != nil {
		logger.ErrLog(layers.Repository, "failed to prepare statement", err)
		return domain.ErrInternalServer
	}

	rows, err := stmt.ExecContext(ctx, args...)

	if err != nil {
		if domain.IsSQLUniqueViolation(err) {
			return domain.ErrConflict
		}
		
		logger.ErrLog(layers.Repository, "failed to execute statement", err)
		return domain.ErrInternalServer
	}

	affected, _ := rows.RowsAffected()

	if affected == 0 {
		return domain.ErrNotFound
	}

	if affected > 1 {
		logger.ErrLog(layers.Repository, "internal server error", fmt.Errorf("weird behaviour, affected more than 1"))
		return domain.ErrInternalServer
	}

	return nil
}