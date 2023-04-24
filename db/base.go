package go_file_db

import (
	"fmt"
	utils "go_file_util"

	_ "github.com/mattn/go-sqlite3"
)

// 插入数据
func Insert(tableName string, params []string, datas ...interface{}) (err error) {
	// 拼接 表名(参数1,参数2,...)
	paramStr := utils.ParamsStr(params)
	// 拼接 values(?,?,...)
	values := utils.ValueStr(len(params))
	var sqlStr = "insert into " + tableName + paramStr + " values" + values
	fmt.Println(sqlStr)
	_, err = db.Exec(sqlStr, datas...) // 要用...展开
	if err != nil {
		fmt.Println(err)
		fmt.Println("插入数据失败")
		return
	}
	return
}

// 更新数据(最后一位要传id[int64])
func Update(tableName string, params []string, datas ...interface{}) (err error) {
	// 拼接 param1=?,param2=?,
	paramStr := utils.UptParamsStr(params)
	sqlStr := "Update " + tableName + " set " + paramStr + " where id = ?"
	fmt.Println(sqlStr)
	_, err = db.Exec(sqlStr, datas...) // 要用...展开
	if err != nil {
		fmt.Println(err)
		fmt.Println("更新数据失败")
		return
	}
	return
}

// 删除数据
func Delete(tableName string, deleteBy []string, datas ...interface{}) (err error) {
	sqlStr := `delete from ` + tableName + ` where `
	for i := 0; i < len(deleteBy)-1; i++ {
		sqlStr += (deleteBy[i] + ` = ? and `)
	}
	sqlStr += deleteBy[len(deleteBy)-1] + ` = ?`
	fmt.Println(sqlStr)
	_, err = db.Exec(sqlStr, datas...) // 要用...展开
	if err != nil {
		fmt.Println(err)
		fmt.Println("删除数据失败")
		return
	}
	return
}

// 更新数据(最后一位要传file_sha1[string])
func UpdateFile(tableName string, params []string, datas ...interface{}) (err error) {
	// 拼接 param1=?,param2=?,
	paramStr := utils.UptParamsStr(params)
	sqlStr := "Update " + tableName + " set " + paramStr + " where file_sha1 = ?"
	fmt.Println(sqlStr)
	_, err = db.Exec(sqlStr, datas...) // 要用...展开
	if err != nil {
		fmt.Println(err)
		fmt.Println("更新数据失败")
		return
	}
	return
}