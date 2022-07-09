package main

// 写不动了 写个连接mysql 压压惊
import (
    "database/sql"
    //"encoding/json"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/pkg/errors"
    "strings"
)

//数据库配置
const (
    userName = "root"
    password = "123456"
    ip       = "10.100.175.100"
    port     = "3001"
    dbName   = "dmp_test"
)
//Db数据库连接池
var DB *sql.DB

//注意方法名大写，就是public
func InitDB() {
    //构建连接："用户名:密码@tcp(IP:端口)/数据库?charset=utf8"
    path := strings.Join([]string{userName, ":", password, "@tcp(", ip, ":", port, ")/", dbName, "?charset=utf8"}, "")
    //打开数据库,前者是驱动名，所以要导入： _ "github.com/go-sql-driver/mysql"
    DB, _ = sql.Open("mysql", path)
    //设置数据库最大连接数
    DB.SetConnMaxLifetime(100)
    //设置上数据库最大闲置连接数
    DB.SetMaxIdleConns(10)
    //验证连接
    if err := DB.Ping(); err != nil {
        fmt.Println("open database fail")
        return
    }
    fmt.Println("connnect success")
}
//查询操作
func Query() {
    rows, err := DB.Query("SELECT 1 a,2 b")
    if err == nil {
        errors.New("query incur error")
    }
	defer rows.Close()

	cols, err := rows.Columns()
	if err != nil{
		errors.New("err")
    }
    fmt.Println(cols)
    vals := make([][]byte, len(cols))
    scans := make([]interface{}, len(cols))

    for i := range vals{
        scans[i] = &vals[i]
    }

    var results []map[string]string

    for rows.Next(){
        err = rows.Scan(scans...)
        if err != nil{
			errors.New("err")
        }

        row := make(map[string]string)
        for k, v := range vals{
            key := cols[k]
            row[key] = string(v)
        }
        results = append(results, row)
    }

    for k, v :=range results{
        fmt.Println(k, v)
    }

    rows.Close()


}

func main(){
    InitDB()
    Query()
    defer DB.Close()
}