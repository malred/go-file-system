package go_file_db

import (
	"fmt"
)

// 新增user
func UserSignup(username string, password string) bool {
	err := Insert("tbl_user", []string{`user_name`, `user_pwd`}, username, password)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return false
}

// 登录
// func UserSignin(username string, encpwd string) bool {

// 	stmt, err := f_db.DBConn().Prepare(
// 		"select * from tbl_user where user_name=? limit 1",
// 	)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return false
// 	}
// 	defer stmt.Close()
// 	rows, err := stmt.Query(username)
// 	// fmt.Println(username)
// 	// fmt.Println(encpwd)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return false
// 	} else if rows == nil {
// 		fmt.Println("Username not found!")
// 		return false
// 	}
// 	pRows := f_db.ParseRows(rows)
// 	if len(pRows) > 0 && string(pRows[0]["user_pwd"].([]byte)) == encpwd {
// 		return true
// 	}
// 	return false
// }

// // 刷新token
// func UpdateToken(username string, token string) bool {
// 	stmt, err := f_db.DBConn().Prepare(
// 		"replace into tbl_user_token (`user_name`, `user_token` ) values(?,?)",
// 	)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return false
// 	}
// 	defer stmt.Close()
// 	_, err = stmt.Exec(username, token)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return false
// 	}
// 	return true
// }

// // 查询用户信息
// func GetUserInfo(username string) (User, error) {
// 	user := User{}
// 	stmt, err := f_db.DBConn().Prepare(
// 		"select user_name,signup_at from tbl_user where user_name=? limit 1",
// 	)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return user, err
// 	}
// 	defer stmt.Close()
// 	err = stmt.QueryRow(username).Scan(&user.Username, &user.SignupAt)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return user, err
// 	}
// 	return user, nil
// }
