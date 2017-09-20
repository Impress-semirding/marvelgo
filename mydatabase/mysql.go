package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
)
/*
	db.Prepare()函数用来返回准备要执行的 sql 操作，然后返回准备完毕的执行状态。
	db.Query()函数用来直接执行 Sql 返回 Rows 结果。
	stmt.Exec()函数用来执行 stmt 准备好的 SQL 语句
 */

func main() {

	startTime := time.Now().UnixNano()
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/golang?charset=utf8")
	checkErr(err)

	sql := "INSERT userinfo SET username = ?, departName = ?, created = ?"
	insertss(db, sql)

	sql = "UPDATE userinfo SET username = ? WHERE uid > 1 AND uid < 11"
	update(db, sql)

	sql = "DELETE FROM userinfo WHERE uid=?"
	deletes(db, sql)

	sql = "SELECT * FROM userinfo WHERE uid = ?"
	findById(db, sql)

	endTime := time.Now().UnixNano()
	fmt.Println(endTime-startTime)

}

func insertss(db *sql.DB, sql string) (err error) {
	defer db.Close()
	// 插入一条数据
	stmt, err := db.Prepare(sql)
	checkErr(err)

	res, err := stmt.Exec("marvel", "研发部门", "2017-07-20") // 返回值是 Result, error
	checkErr(err)
	// 获取自增长的 id
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println("id:", id)
	return
}

/*
	更新一条数据
 */
func update(db *sql.DB, sql string) (err error) {
	defer db.Close()

	stmt, err := db.Prepare(sql)
	checkErr(err)

	res, err := stmt.Exec("marvel")
	checkErr(err)

	// 影响的行数
	rows, err := res.RowsAffected()
	checkErr(err)

	fmt.Println("影响的行数", rows)
	return
}

/**
	删除一条数据
 */
func deletes(db *sql.DB, sql string) (err error) {
	defer db.Close()

	//删除数据
	stmt, err := db.Prepare(sql)
	checkErr(err)

	res, err := stmt.Exec(10)
	checkErr(err)

	affect, err := res.RowsAffected()

	checkErr(err)
	fmt.Println(affect)

	return
}

/**
	查找指定id 的信息
 */
func findById(db *sql.DB, sql string) (err error) {
	defer db.Close()

	stmt, err := db.Prepare(sql)
	stmt.Query()

	row, err := db.Query(sql, 1)
	checkErr(err)

	for row.Next() {
		var uid int
		var username string
		var departName string
		var created string

		row.Scan(&uid, &username, &departName, &created)

		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(departName)
		fmt.Println(created)
	}

	return
}

func checkErr(err error)  {
	if err != nil {
		fmt.Println(err)
	}
}

