package go_file_db

import "fmt"

/**
CREATE TABLE `tbl_file` (
    `id` INT(11) NOT NULL AUTO_INCREMENT,
    `file_sha1` CHAR(40) NOT NULL DEFAULT '' COMMENT '文件hash',
    `file_name` VARCHAR(256) NOT NULL DEFAULT '' COMMENT '文件名',
    `file_size` BIGINT(20) DEFAULT '0' COMMENT '文件大小',
    `file_addr` VARCHAR(1024) NOT NULL DEFAULT '' COMMENT '文件存储位置',
    `create_at` DATETIME DEFAULT NOW() COMMENT '创建日期',
    `update_at` DATETIME DEFAULT NOW() ON UPDATE CURRENT_TIMESTAMP() COMMENT '修改日期',
    `status` INT(11) NOT NULL DEFAULT '0' COMMENT '状态(可用/禁用/已删除等)',
    `ext1` INT(11) DEFAULT '0' COMMENT '备用字段',
    `ext2` INT(11) DEFAULT '0' COMMENT '备用字段',
    PRIMARY KEY (`id`),
    UNIQUE KEY `idx_file_hash` (`file_sha1`),
    KEY `idx_status` (`status`)
) ENGINE=INNODB DEFAULT CHARSET=utf8;
*/

// 新增文件元数据(数据库)
func InsertFileMetaDB(fmeta FileMeta) bool {
	return OnFileUploadFinished(fmeta.FileSha1, fmeta.FileName, fmeta.FileSize, fmeta.Location)
}

// 更新文件元数据(数据库)
func UpdateFileMetaDB(fileName string, fileHash string) (err error) {
    fmt.Println(fileHash)
    fmt.Println(fileName)
	err = UpdateFile("tbl_file", []string{`file_name`}, fileName, fileHash)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	// err = Update("tbl_user_file", []string{`file_name`})
	return
}

// 获取文件元数据(数据库)
func GetFileMetaDB(fileSha1 string) (FileMeta, error) {
	sqlStr := "select file_sha1, file_name, file_size, file_addr from tbl_file " +
		" where file_sha1=? and status=1 limit 1"
	var fmeta FileMeta
	if err := db.Get(&fmeta, sqlStr, fileSha1); err != nil {
		fmt.Printf("查询文件信息失败(数据库), err:%v\n", err)
		return fmeta, err
	}
	return fmeta, nil
}
