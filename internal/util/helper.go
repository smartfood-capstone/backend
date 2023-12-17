package util

import (
	"database/sql"
	"log"

	"github.com/labstack/echo/v4"
)

func CommitOrRollback(tx *sql.Tx) {
	if err := recover(); err != nil {
		errorRollback := tx.Rollback()
		log.Fatal(errorRollback)
		panic(err)
	} else if err != nil {
		tx.Rollback()
	} else {
		tx.Commit()
	}
}

func MakeResponse(status int, message string, err error, data any) echo.Map {
	if err != nil {
		return echo.Map{
			"status":  status,
			"message": message,
			"data": echo.Map{
				"error": err.Error(),
			},
		}
	}
	return echo.Map{
		"status":  status,
		"message": message,
		"data":    data,
	}
}
