package pack22

import (
	_ "../src/github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

/**
[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
 */

func Pack22() {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mex_pscal?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	stmt, err := db.Prepare(" " +
		"INSERT INTO base_users " +
		"(id, account, mobile, email, password, status, created_at, updated_at) " +
		"VALUES " +
		"('233', 'demo100', '17700000000', 'demo@demo.com', 'demo100200300', 1, '2019-05-26', '2019-05-26') ")
	if err != nil {
		panic(err)
	}
	_, err = stmt.Exec()
	if err != nil {
		panic(err)
	}
}

func Pack22Select()  {
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/mex_pscal?charset=utf8mb4")
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(" select * from  base_users; ")
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		var id int16
		var username string
		_ = rows.Scan(&id, &username)
		fmt.Println(id, username)
	}
}
