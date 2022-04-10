package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/xuri/excelize/v2"
)

func main() {
	//建立数据库连接
	// conn, err := sql.Open("mysql", "用户名:密码@tcp(机器IP:端口)/mysql?charset=utf8")
	conn, err := sql.Open("mysql", "liu:0imchen@@tcp(10.0.0.50:3306)/just4live?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	//数据库操作
	res, err := conn.Exec(
		`create table if not exists car2021(
			车型ID		INT,
			车型		VARCHAR(30),
			燃料类别	VARCHAR(10),
			品牌		VARCHAR(10),
			级别		VARCHAR(10),
			合资自主	VARCHAR(10),
			国别		VARCHAR(10),
			整理品牌	VARCHAR(16),
			大区		VARCHAR(10),
			省区		VARCHAR(10),
			省份		VARCHAR(10),
			城市群		VARCHAR(10),
			城市级别	VARCHAR(3),
			城市		VARCHAR(10),
			城市简称	VARCHAR(10),
			2021y1		INT,
			2021y2		INT,
			2021y3		INT,
			2021y4		INT,
			2021y5		INT,
			2021y6		INT,
			2021y7		INT,
			2021y8		INT,
			2021y9		INT,
			2021y10	INT,
			2021y11	INT,
			2021y12	INT,
			2021y		INT
		);`,
	)
	if err != nil {
		fmt.Println("创建table失败:", err)
	}

	log.Println(res, err)
	sql, err := conn.Prepare("INSERT INTO car2021 (`车型ID`,`车型`,`燃料类别`,`品牌`,`级别`,`合资自主`,`国别`,`整理品牌`,`大区`,`省区`,`省份`,`城市群`,`城市级别`,`城市`,`城市简称`,`2021y1`,`2021y2`,`2021y3`,`2021y4`,`2021y5`,`2021y6`,`2021y7`,`2021y8`,`2021y9`,`2021y10`,`2021y11`,`2021y12`,`2021y`) VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
	log.Println(sql, err)
	defer sql.Close()

	// 读取excel
	f, err := excelize.OpenFile("./excel/2021y.xlsx")
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
		res, err := sql.Exec(row[0], row[1], row[2], row[3], row[4], row[5], row[6], row[7], row[8], row[9], row[10], row[11], row[12], row[13], row[14], row[15], row[16], row[17], row[18], row[19], row[20], row[21], row[22], row[23], row[24], row[25], row[26], row[27]) //string格式
		if err != nil {
			log.Println(res, err)
		} else {
			count += 1
			log.Println("success")
		}
	}

	fmt.Println(count)
	// fmt.Printf("%T", rows[4]) //string[]格式
	//这是main分支哦,搞清楚!
}
