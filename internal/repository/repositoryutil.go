package repository

import (
	"fmt"
	"strings"
)

func InsertQuery(tableName string, cols string) string {
	numberOfMarks := strings.Count(cols, ",")
	values := "( "
	for i := 0; i <= numberOfMarks; i++ {
		if i != 0 {
			values += ","
		}
		values += " ?"
	}
	values += ", NOW(), NOW())"
	cols = "(" + cols + ", created_at, updated_at)"
	return fmt.Sprintf(
		"INSERT INTO %s %s VALUES %s",
		tableName, cols, values,
	)
}
