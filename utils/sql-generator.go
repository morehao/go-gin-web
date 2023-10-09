package utils

import "fmt"

var sql = "ALTER TABLE tblBalanceDetail%d add `cuid` varchar(50) NOT NULL COMMENT '用户设备id'"

func main2() {

	for i := 0; i < 20; i++ {
		fmt.Printf(sql, i)
	}
}
