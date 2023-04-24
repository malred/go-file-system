package go_file_util

// 拼接sql语句的value
// len是语句有几个参数
func ValueStr(len int) (values string) {
	// 拼接 values(?,?,...)
	values = "("
	for i := 0; i < len-1; i++ {
		values += "?"
		values += ","
	}
	values += "?"
	values += ")"
	return
}

// 拼接sql语句update的param
// params是参数名数组
func UptParamsStr(params []string) (paramStr string) {
	// 拼接参数列表 xxx=?,xxx=?
	paramStr = ""
	for i := 0; i < len(params)-1; i++ {
		paramStr += params[i]
		paramStr += "=?,"
	}
	paramStr += params[len(params)-1]
	paramStr += "=?"
	return
}

// 拼接sql语句的param
// params是参数名数组
func ParamsStr(params []string) (paramStr string) {
	// 拼接 表名(参数1,参数2,...)
	paramStr = "("
	for i := 0; i < len(params)-1; i++ {
		paramStr += params[i]
		paramStr += ","
	}
	paramStr += params[len(params)-1]
	paramStr += ")"
	return
}
