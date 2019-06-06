package main

import (
	"database/sql"
	"fmt"
	_ "github.com/Go-SQL-Driver/Mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@/test")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("insert userinfo set username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("yishuang", "人力资源平台部", "2019-6-1")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)
	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("yishuangupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("select * from userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var departname string
		var created string
		err = rows.Scan(&uid, &username, &departname, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(departname)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()

}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}
