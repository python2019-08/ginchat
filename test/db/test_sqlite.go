package dbtest

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func DB_sqlite3() {
	// 打开数据库连接
	db, err := sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表和清空表数据
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS foo (
			id INTEGER PRIMARY KEY,
			name TEXT
		);
		DELETE FROM foo;
	`)
	if err != nil {
		log.Fatal(err)
	}

	// 开启事务
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback() // 注意：如果出现错误，及时回滚事务

	// 准备插入语句
	stmt, err := tx.Prepare("INSERT INTO foo(id, name) VALUES(?, ?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	// 执行插入操作
	for i := 0; i < 100; i++ {
		_, err = stmt.Exec(i, fmt.Sprintf("こんにちわ世界%03d", i))
		if err != nil {
			log.Fatal(err)
		}
	}

	// 提交事务
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}

	// 查询数据并打印
	rows, err := db.Query("SELECT id, name FROM foo")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
