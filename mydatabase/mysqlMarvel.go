package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"crypto/md5"
	"encoding/hex"
)
/**
	连接 Mysql 数据库
 */

func main() {

	Add()
	Select()

}

/**
	添加一条数据
 */
func Add()  {

	// 获取数据库连接
	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	defer db.Close()

	checkErr(err)

	// 插入数据(两种方式都可以，都是正确的 sql 语句)
	//stmt, err := db.Prepare("INSERT INTO user_test(userName, pass) VALUES(?, ?)")
	stmt, err := db.Prepare("INSERT user_test SET userName=?, pass=?")
	checkErr(err)

	// 密码要使用 md5 加密算法
	h := md5.New()
	h.Write([]byte("123456"))
	pass := hex.EncodeToString(h.Sum(nil))

	res, err := stmt.Exec("astaxie", pass)
	checkErr(err)
	// 获取插入之后自动增长的id 值
	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
}

/**
	更新数据
 */
func Update()  {

	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	defer db.Close()

	checkErr(err)

	stmt, err := db.Prepare("UPDATE user_test SET userName=? WHERE id=?")
	res, err := stmt.Exec("marvelous", 1)
	checkErr(err)

	row, err := res.RowsAffected()
	checkErr(err)
	fmt.Println("影响的行数", row)

}

/**
	查找所有数据
 */
func Select()  {

	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	defer db.Close()

	checkErr(err)

	stmt, err := db.Prepare("SELECT * FROM user_test")
	checkErr(err)

	res, err := stmt.Query()
	checkErr(err)

	var id int
	var userName string
	var pass string

	for res.Next() {
		res.Scan(&id, &userName, &pass)
		fmt.Println(id, userName, pass)
	}
}

/**
	根据id 删除一条数据
 */
func DeleteById(id int)  {
	//删除数据
	db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	defer db.Close()

	checkErr(err)
	stmt, err := db.Prepare("DELETE FROM user_test WHERE id=?")
	checkErr(err)

	res, err := stmt.Exec(2)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)
}

/**
	检查异常
 */
func checkErrs(err error) {
	if err != nil {
		panic(err)
	}
}
