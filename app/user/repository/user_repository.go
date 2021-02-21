package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/iamaul/evonix-backend-api/app/models"
	"github.com/iamaul/evonix-backend-api/app/user"
	"github.com/sirupsen/logrus"
)

type userRepository struct {
	DbConnection *sql.DB
}

func NewUserRepository(connection *sql.DB) user.Repository {
	return &userRepository{connection}
}

func (u *userRepository) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.User, error) {
	rows, err := u.DbConnection.QueryContext(ctx, query, args...)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	defer rows.Close()

	payload := make([]*models.User, 0)
	for rows.Next() {
		data := new(models.User)
		err = rows.Scan(
			&data.ID,
			&data.Name,
			&data.Email,
			&data.EmailVerified,
			&data.RegisteredDate,
			&data.RegisterIP,
			&data.UCPLoginIP,
			&data.LoginIP,
			&data.Admin,
			&data.AdminDivision,
			&data.Helper,
			&data.LastLogin,
			&data.Status,
			&data.DelayCharacterDeletion,
			&data.Blocked,
			&data.LastBlockIssuer,
			&data.LastBlockReason,
		)

		if err != nil {
			logrus.Error(err)
			return nil, err
		}

		payload = append(payload, data)
	}

	return payload, nil
}

func (u *userRepository) GetByID(ctx context.Context, id int64) (res *models.User, err error) {
	query := `SELECT id, name, email, email_verified, registered_date, register_ip, ucp_login_ip, 
		login_ip, admin, admin_division, helper, lastlogin, status, delay_character_deletion,
		blocked, lastblock_issuer, lastblock_reason
	FROM users WHERE id=?`

	result, err := u.fetch(ctx, query, id)
	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		res = result[0]
	} else {
		return nil, errors.New("User not found")
	}

	return
}

func (u *userRepository) GetByName(ctx context.Context, name string) (res *models.User, err error) {
	query := `SELECT id, name, email, email_verified, registered_date, register_ip, ucp_login_ip, 
		login_ip, admin, admin_division, helper, lastlogin, status, delay_character_deletion,
		blocked, lastblock_issuer, lastblock_reason 
	FROM users WHERE name=?`

	result, err := u.fetch(ctx, query, name)
	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		res = result[0]
	} else {
		return nil, errors.New("User not found")
	}

	return
}

func (u *userRepository) GetByEmail(ctx context.Context, email string) (res *models.User, err error) {
	query := `SELECT id, name, email, email_verified, registered_date, register_ip, ucp_login_ip, 
		login_ip, admin, admin_division, helper, lastlogin, status, delay_character_deletion,
		blocked, lastblock_issuer, lastblock_reason 
	FROM users WHERE email=?`

	result, err := u.fetch(ctx, query, email)
	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		res = result[0]
	} else {
		return nil, errors.New("User not found")
	}

	return
}

func (u *userRepository) Store(ctx context.Context, um *models.User) error {
	query := `INSERT users SET name=?, email=?, password=?, registered_date=?, register_ip=?`

	stmt, err := u.DbConnection.PrepareContext(ctx, query)
	if err != nil {
		return err
	}

	res, err := stmt.ExecContext(ctx, um.Name, um.Email, um.Password, um.RegisteredDate, um.RegisterIP)
	if err != nil {
		return err
	}

	lastID, err := res.LastInsertId()
	if err != nil {
		return err
	}

	um.ID = lastID

	return nil
}
