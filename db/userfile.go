package go_file_db

import (
	"fmt"
	"time"
)

// 更新/新增用户文件表
func OnUserFileUploadFinished(username, filehash, filename string, filesize int64) bool {
	err := Insert("tbl_user_file",
		[]string{`user_name`, `file_sha1`, `file_name`, `file_size`, `upload_at`, `last_update`, `status`},
		username, filehash, filename, filesize, time.Now(), time.Now(), 1)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}

// 删除用户的文件
func DeleteUserFile(username, filehash string) (err error) {
	// 查询是否存在
	sqlStr := `select * from tbl_user_file where user_name="` + username + `" and file_sha1="` + filehash + `"`
	var user UserFile
	if err := db.Get(&user, sqlStr, username, filehash); err != nil {
		fmt.Printf("查询文件信息失败(数据库), err:%v\n", err)
		return err
	}
	// 删除
	err = Delete("tbl_user_file", []string{`file_sha1`, `user_name`}, filehash, username)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// 移入回收站(status=0)
func RemoveUserFile(username, filehash string) (err error) {
	// 查询是否存在
	sqlStr := `select * from tbl_user_file where user_name="` + username + `" and file_sha1="` + filehash + `"`
	var user UserFile
	if err := db.Get(&user, sqlStr, username, filehash); err != nil {
		fmt.Printf("查询文件信息失败(数据库), err:%v\n", err)
		return err
	}
	// 删除
	sqlStr = "Update tbl_user_file set status=0 where file_sha1 = ? and user_name = ?"
	fmt.Println(sqlStr)
	_, err = db.Exec(sqlStr, filehash, username) // 要用...展开
	if err != nil {
		fmt.Println(err)
		fmt.Println("移入回收站失败")
		return
	}
	return nil
}

// 根据username查询用户文件信息
func GetUserFileMetaDB(username string) (files []*UserFile, err error) {
	sqlStr := "select * from tbl_user_file where user_name='" + username + "' and status=1 limit 20"
	if err := db.Select(&files, sqlStr); err != nil {
		fmt.Printf("查询用户文件信息失败(数据库), err:%v\n", err)
		return files, err
	}
	return files, err
}
