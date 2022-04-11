package database

import (
	"database/sql"
	"fmt"

	"github.com/gh0st3e/TodoForTechArt/internal/util"
	"github.com/pkg/errors"
)

func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open(util.DriverName, fmt.Sprintf("%s%s", util.ConnectionPrefix, util.Address)) // Подключение к БД
	if err != nil {
		return nil, errors.Wrap(err, "repository.Connect.Open couldn't connect to sql")
	}

	err = db.Ping()
	if err != nil {
		return nil, errors.Wrap(err, "repository.Connect.Ping couldn't ping database")
	}

	return db, nil
}
