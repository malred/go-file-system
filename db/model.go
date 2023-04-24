package go_file_db

/*
	创建数据库
		// 进入sqlite命令行
		sqlite3 xxx.db
		// 创建数据库
		.open xxx.db
*/

// token
type User_Token struct {
	user_name string `db: user_name`
	user_pwd  string `db: user_pwd`
}

// user
type User struct {
	Username       string `db:"user_name"`
	Userpwd        string `db:"user_pwd"`
	Email          string `db:"email"`
	Phone          string `db:"phone"`
	EmailValidated int    `db:"email_validated` // 邮箱是否已验证 1-yes 0-no
	PhoneValidated int    `db:"phone_validated` // 电话是否已验证
	SignupAt       string `db:"signup_at"`
	LastActiveAt   string `db:"last_active"`
	Status         int    `db:"status"`
	Profile        string `db:"profile"`
}

// file
type TableFile struct {
	FileHash string `db:"file_sha1"`
	FileName string `db:"file_name"`
	FileSize int64  `db:"file_size"`
	FileAddr string `db:"file_addr"`
}

// 用户文件表
type UserFile struct {
	// Uid         int64  `db:"id"`
	UserName    string `db:"user_name"`
	FileHash    string `db:"file_sha1"`
	FileName    string `db:"file_name"`
	FileSize    int64  `db:"file_size"`
	UploadAt    string `db:"upload_at"`
	LastUpdated string `db:"last_update"`
	Status      string `db:"status"`
}

// FileMeta: 文件元数据结构体
type FileMeta struct {
	// 文件id(hash)
	FileSha1 string `db:"file_sha1"`
	FileName string `db:"file_name"`
	FileSize int64  `db:"file_size"`
	// 根路径
	Location string `db:"file_addr"`
	// 上传时间
	UploadAt string `db:"update_at"`
}
