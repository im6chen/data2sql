package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
)

func sql2022() {
	//建立数据库连接
	// conn, err := sql.Open("mysql", "用户名:密码@tcp(机器IP:端口)/mysql?charset=utf8")
	conn, err := sql.Open("mysql", "liu:0imchen@@tcp(10.0.0.50:3306)/just4live?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	//数据库操作
	res, err := conn.Exec(
		`create table if not exists car2022(
			车型ID		INT,
			车型		VARCHAR(30),
			燃料类别	VARCHAR(10),
			品牌		VARCHAR(10),
			级别		VARCHAR(10),
			合资自主	VARCHAR(10),
			国别	    VARCHAR(10),
			整理品牌	VARCHAR(16),
			大区		VARCHAR(10),
			省份		VARCHAR(10),
			城市级别	VARCHAR(3),
			城市		VARCHAR(10),
			2022y1		INT,
			2022y2		INT,
			2022y3		INT
		);`,
	)
	if err != nil {
		fmt.Println("创建table失败:", err)
	}

	log.Println("RES:::", res, "\n", "err:::", err)
	sql, err := conn.Prepare("INSERT INTO car2022 ( `车型ID`,`车型`, `燃料类别`, `品牌`, `级别`, `合资自主`, `国别`, `整理品牌`, `大区`, `省份`, `城市级别`, `城市`, `2022y1`, `2022y2`, `2022y3`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	log.Println("SQL:::", sql, "\n", "err:::", err)
	defer sql.Close()

	// 读取excel
	f, err := excelize.OpenFile("./excel/2022y.xlsx")
	if err != nil {
		fmt.Println("读取excel错误:", err)
		return
	}

	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	count := 0
	fmt.Println(count)
	for _, row := range rows[1:] {
		res, err := sql.Exec(row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11], row[12], row[13], row[14]) //string格式
		if err != nil {
			log.Println(res, err)
		} else {
			count++
			log.Println("success")
		}
	}

	fmt.Println(count)
}
