package go_file_db

import (
	"fmt"
	// "time"
)

// 文件上传完成执行
func OnFileUploadFinished(filehash string, filename string, filesize int64, fileaddr string) bool {
	err := Insert("tbl_file", []string{"file_sha1", "file_name", `file_size`, `file_addr`, "status"},
		filehash, filename, filesize, fileaddr, 1)
	if err != nil {
		fmt.Println(err.Error())
		return false
	}
	return true
}
