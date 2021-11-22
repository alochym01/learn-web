package sqlite3

import (
	"database/sql"

	"github.com/alochym01/web-w-gin/domain"
	"github.com/alochym01/web-w-gin/errs"
	"go.uber.org/zap"
)

type UserSQL struct {
	db     *sql.DB
	logger *zap.Logger
}

func NewUserSQL(DB *sql.DB, l *zap.Logger) UserSQL {
	return UserSQL{
		db:     DB,
		logger: l,
	}
}

// FindAll ...
func (u UserSQL) FindAll() ([]domain.User, *errs.AppErr) {
	sqlstmt := "select * from users"

	rows, err := u.db.Query(sqlstmt)

	// check err from server DB and Query DB
	if err != nil {
		u.logger.Error(err.Error())
		return nil, errs.ServerError()
	}

	users := []domain.User{}

	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.ID, &user.Email, &user.FullName, &user.Password)
		// check err from server DB and Scan function
		if err != nil {
			u.logger.Error(err.Error())
			return nil, errs.ServerError()
		}
		users = append(users, user)
	}
	return users, nil
}

// FindByID ...
func (u UserSQL) FindByID(id int) (*domain.User, *errs.AppErr) {
	sqlstmt := "select * from users where id=?"

	row := u.db.QueryRow(sqlstmt, id)
	user := domain.User{}
	err := row.Scan(&user.ID, &user.Email, &user.FullName, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			u.logger.Warn(err.Error())
			return nil, errs.NotFound()
		}
		u.logger.Error(err.Error())
		return nil, errs.ServerError()
	}

	return &user, nil
}

// FindByEmail ...
func (u UserSQL) FindByEmail(email string) (*domain.User, *errs.AppErr) {
	sqlstmt := "select * from users where email=?"

	row := u.db.QueryRow(sqlstmt, email)
	user := domain.User{}
	err := row.Scan(&user.ID, &user.Email, &user.FullName)
	if err != nil {
		if err == sql.ErrNoRows {
			u.logger.Warn(err.Error())
			return nil, errs.NotFound()
		}
		u.logger.Error(err.Error())
		return nil, errs.ServerError()
	}

	return &user, nil
}

// Create ...
func (u UserSQL) Create(user domain.User) (*int64, *errs.AppErr) {
	sqlstmt := "INSERT INTO users(email, fullname, password) VALUES(?,?,?)"

	// Execute SQL Statements
	result, err := u.db.Exec(sqlstmt, user.Email, user.FullName, user.Password)

	// check err from server DB and Query DB
	if err != nil {
		u.logger.Error(err.Error())
		return nil, errs.ServerError()
	}

	rowID, err := result.LastInsertId()

	// error check for LastInsertId function
	if err != nil {
		u.logger.Error(err.Error())
		return nil, errs.ServerError()
	}

	return &rowID, nil
}

// Update ...
func (u UserSQL) Update(user domain.User) *errs.AppErr {
	sqlstmt := "UPDATE users SET email=?, fullname=?, password=? where id=?"

	// Execute SQL Statements
	result, err := u.db.Exec(sqlstmt, user.Email, user.FullName, user.Password, user.ID)

	// check err from server DB and Query DB
	if err != nil {
		u.logger.Error(err.Error())
		return errs.ServerError()
	}

	rowCount, err := result.RowsAffected()

	// error check for RowsAffected function
	if err != nil {
		u.logger.Error(err.Error())
		return errs.ServerError()
	}

	// there is no row found
	if rowCount == 0 {
		u.logger.Warn("Record Not Found")
		return errs.NotFound()
	}

	return nil
}

// Delete ...
func (u UserSQL) Delete(id string) *errs.AppErr {
	sqlstmt := "DELETE FROM users where id=?"

	// Execute SQL Statements
	result, err := u.db.Exec(sqlstmt, id)

	// check err from server DB and Query DB
	if err != nil {
		u.logger.Error(err.Error())
		return errs.ServerError()
	}

	rowCount, err := result.RowsAffected()

	// error check for RowsAffected function
	if err != nil {
		u.logger.Error(err.Error())
		return errs.ServerError()
	}

	// there is no row found
	if rowCount == 0 {
		u.logger.Warn("Record Not Found")
		return errs.NotFound()
	}

	return nil
}
