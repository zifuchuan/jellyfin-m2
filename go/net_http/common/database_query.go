package common

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

func initMysql() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:mysql@123@tcp(127.0.0.1:3306)/testdb")
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err
}

func DatabaseQuery01(text string) (string, error) {
	db, err := initMysql()
	if err != nil {
		return "", err
	}
	defer safeClose(db)

	rows, err := db.Query(fmt.Sprintf("select * from t_py_class where name = '%s'", text))
	if err != nil {
		return "", err
	}
	defer safeClose(rows)

	rowItems := make([]string, 0)
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return "", err
		}
		rowItems = append(rowItems, name)
	}
	err = rows.Err()
	if err != nil {
		return "", err
	}
	return strings.Join(rowItems, "\n"), nil
}

func DatabaseQuery02(text string) (string, error) {
	db, err := initMysql()
	if err != nil {
		return "", err
	}
	defer safeClose(db)

	row := db.QueryRow(fmt.Sprintf("select * from t_py_class where name = '%s'", text))
	var name string
	err = row.Scan(&name)
	if err != nil {
		return "", err
	}

	return name, nil
}

func DatabaseQuery03(text string) (string, error) {
	db, err := initMysql()
	if err != nil {
		return "", err
	}
	defer safeClose(db)

	result, err := db.Exec(fmt.Sprintf("insert into t_py_class values('%s')", text))
	if err != nil {
		return "", err
	}
	ret := fmt.Sprintf("%v", result)

	return ret, nil
}

func DatabaseQuery04(text string) (string, error) {
	db, err := initMysql()
	if err != nil {
		return "", err
	}
	defer safeClose(db)

	stmt, err := db.Prepare(fmt.Sprintf("delete from t_py_class where name = '%s'", text))
	if err != nil {
		return "", err
	}
	res, err := stmt.Exec()
	if err != nil {
		return "", err
	}
	_ = res
	// res.LastInsertId()
	// res.RowsAffected()

	return "ok", nil
}

func DatabaseQuery05(text string) (string, error) {
	db, err := initMysql()
	if err != nil {
		return "", err
	}
	defer safeClose(db)

	stmt, err := db.Prepare("delete from t_py_class where name = ?")
	if err != nil {
		return "", err
	}
	res, err := stmt.Exec(text)
	if err != nil {
		return "", err
	}
	_ = res
	// res.LastInsertId()
	// res.RowsAffected()

	return "ok", nil
}

func DatabaseQuery06(text string) (string, error) {
	db, err := initMysql()
	if err != nil {
		return "", err
	}
	defer safeClose(db)
	ctx := context.Background()
	rows, err := db.QueryContext(ctx, fmt.Sprintf("select * from t_py_class where name = '%s'", text))
	if err != nil {
		return "", err
	}
	defer safeClose(rows)

	rowItems := make([]string, 0)
	for rows.Next() {
		var name string
		err = rows.Scan(&name)
		if err != nil {
			return "", err
		}
		rowItems = append(rowItems, name)
	}
	err = rows.Err()
	if err != nil {
		return "", err
	}
	return strings.Join(rowItems, "\n"), nil
}

func DatabaseQuery07(text string) (string, error) {
	return "", nil
}

func DatabaseQuery08(text string) (string, error) {
	return "", nil
}

func DatabaseQuery09(text string) (string, error) {
	return "", nil
}

func DatabaseQuery10(text string) (string, error) {
	return "", nil
}
