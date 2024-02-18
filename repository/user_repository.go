package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/devanfer02/nosudes-be/domain"
	helper "github.com/devanfer02/nosudes-be/utils"
	"github.com/devanfer02/nosudes-be/utils/layers"
	"github.com/devanfer02/nosudes-be/utils/logger"
	"github.com/jmoiron/sqlx"
)

type mysqlUserRepository struct {
	Conn *sqlx.DB
}

func NewMysqlUserRepository(conn *sqlx.DB) domain.UserRepository {
	return &mysqlUserRepository{
		conn,
	}
}

func (m *mysqlUserRepository) FetchAll(ctx context.Context) ([]domain.User, error) {
	query := "SELECT * FROM users"

	users := make([]domain.User, 0)

	err := m.Conn.SelectContext(ctx, &users, query)

	if err != nil {
		logger.ErrLog(layers.Repository, "failed to exec query", err)
		return nil, domain.ErrInternalServer
	}

	return users, nil
}

func (m *mysqlUserRepository) FetchOneByArg(ctx context.Context, param, arg string) (domain.User, error) {
	query := fmt.Sprintf("SELECT * FROM users WHERE %s = ? LIMIT 1", param)

	user := domain.User{}

	err := m.Conn.GetContext(ctx, &user, query, arg)

	if err != nil {

		if err == sql.ErrNoRows {
			return domain.User{}, domain.ErrNotFound
		}

		logger.ErrLog(layers.Repository, "failed to exec query", err)
		return domain.User{}, domain.ErrInternalServer
	}

	return user, nil
}

func (m *mysqlUserRepository) InsertUser(ctx context.Context, user *domain.UserPayload) error {
	query := "INSERT INTO users (user_id, fullname, email, password, created_at, updated_at) VALUES (?, ?, ?, ?, ?, ?)"

	currTime := helper.CurrentTime()
	
	return execStatement(m.Conn, ctx, query, user.ID, user.Fullname, user.Email, user.Password, currTime, currTime)
}

func (m *mysqlUserRepository) UpdateUser(ctx context.Context, user *domain.UserPayload) error {
	query := "UPDATE users SET fullname = ?, email = ?, password = ?, updated_at = ? WHERE user_id = ?"

	currTime := helper.CurrentTime()

	return execStatement(m.Conn, ctx, query, user.Fullname, user.Email, user.Password, currTime, user.ID)
}

func (m *mysqlUserRepository) DeleteUser(ctx context.Context, id string) error {
	query := "DELETE FROM users WHERE user_id = ?"

	return execStatement(m.Conn, ctx, query, id)
}
