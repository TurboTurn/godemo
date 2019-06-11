package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/Go-SQL-Driver/Mysql"
	"net/http"
	"os"
	"runtime"
	"webDemo/db"
)

func find(uid string) (id int, username, departname, created string) {
	db, err := sql.Open("mysql", "root:admin@/test")
	checkErr(err)
	//查询数据
	rows, err := db.Query("select * from userinfo where uid=" + uid)
	checkErr(err)
	for rows.Next() {
		err = rows.Scan(&id, &username, &departname, &created)
		checkErr(err)
		fmt.Println(id, username, departname, created)
	}
	return id, username, departname, created
}

func checkErr(e error) {
	if e != nil {
		panic(e)
	}
}

func getuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	id, username, departname, created := find(uid)
	_, _ = fmt.Fprint(w, id, username, departname, created)
}
func modifyuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	_, _ = fmt.Fprintf(w, "you are modify user %s", uid)
}
func deleteuser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	_, _ = fmt.Fprintf(w, "you are delete user %s", uid)
}
func adduser(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	uid := params.Get(":uid")
	_, _ = fmt.Fprintf(w, "you are add user %s", uid)
}

type Book struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float64
}
type User struct {
	Title       string
	Authors     []string
	Publisher   string
	IsPublished bool
	Price       float64
}

func main() {
	db.Save()
	bo := new(Book)
	bo.Title = "rt"
	fmt.Println("bo-", bo)
	boo := Book{}
	fmt.Println("boo-", boo)
	boo1 := &Book{}
	boo1.Title = "boo1"
	fmt.Println("boo1-", boo1)
	gobook := Book{"Go语言编程",
		[]string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"},
		"ituring.com.cn",
		true,
		9.99,
	}
	user := User{"Go语言编程",
		[]string{"XuShiwei", "HughLv", "Pandaman", "GuaguaSong", "HanTuo", "BertYuan", "XuDaoli"},
		"ituring.com.cn",
		true,
		9.99}
	b, err := json.Marshal(gobook)
	if err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Fprintln(os.Stderr, err)
	}
	b, err = json.Marshal(user)
	if err == nil {
		fmt.Println(string(b))
	} else {
		fmt.Fprintln(os.Stderr, err)
	}

	m := map[int]int{3: 4, 4: 5}
	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println(runtime.GOOS)

	/*mux := routes.New()
	mux.Get("/user/:uid", getuser)
	mux.Post("/user/:uid", modifyuser)
	mux.Del("/user/:uid", deleteuser)
	mux.Put("/user/:uid", adduser)
	http.Handle("/", mux)
	addr := ":8080"
	fmt.Printf("服务器开启在%s端口\n",addr)
	http.ListenAndServe(addr, nil)*/

}
func init() {
	fmt.Println("init...", f)
	f = 78
}
func init() {
	fmt.Println("init...2", f, b)
}

var f = 4

const b = 56

//先初始化全局常量、变量，然后执行init，多个init按上下顺序执行，再执行main
