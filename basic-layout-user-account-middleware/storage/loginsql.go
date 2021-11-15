package storage

import (
	"database/sql"

	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/domain"
	"github.com/alochym01/learn-web/basic-layout-user-account-middleware/errs"
	"go.uber.org/zap"
)

type LoginSQL struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewLoginSQL(DB *sql.DB, l *zap.Logger) LoginSQL {
	return LoginSQL{
		db:     DB,
		logger: l,
	}
}

// FindByEmail ...
func (l LoginSQL) FindByEmail(temp domain.LoginRequest) (*domain.Login, *errs.Error) {
	sqlstmt := "select email, password from users where email=?"

	row := l.db.QueryRow(sqlstmt, temp.Email)
	login := domain.Login{}
	err := row.Scan(&login.Email, &login.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			l.logger.Error(err.Error())
			return nil, errs.NotFound("Record Not Found")
		}
		l.logger.Error(err.Error())
		return nil, errs.ServerError("Server Internal Error")
	}

	if login.Password != temp.Password {
		l.logger.Warn("Password not match")
		return nil, errs.PasswordError("User/Password not match")
	}

	return &login, nil
}
