package go_file_web

import (
	db "go_file_db"
	utils "go_file_util"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/gin-gonic/gin"
)

// 获取批量文件元数据
func GetUserFileByUserName(c *gin.Context) {
	username := c.Query("username")
	files, err := db.GetUserFileMetaDB(username)
	utils.R(c, err, "查询用户文件信息失败", files)
}

// 根据文件hash删除文件
func DeleteByFilehashHandler(c *gin.Context) {
	filehash := c.Query("filehash")
	username := c.Query("username")
	var err error
	if filehash != "" && username != "" {
		err = db.DeleteUserFile(username, filehash)
		if err != nil {
			utils.R(c, err, "删除文件失败", nil)
			return
		}
	} else {
		utils.R(c, err, "用户名或文件hash为空!", nil)
		return
	}
	utils.R(c, err, "删除文件失败", "删除文件成功")
}

// 移动文件到回收站
func RemoveUserFileByFilehash(c *gin.Context) {
	filehash := c.Query("filehash")
	username := c.Query("username")
	var err error
	if filehash != "" && username != "" {
		err = db.RemoveUserFile(username, filehash)
		if err != nil {
			utils.R(c, err, "移入回收站失败", nil)
			return
		}
	} else {
		utils.R(c, err, "用户名或文件hash为空!", nil)
		return
	}
	utils.R(c, err, "移入回收站失败", "移入回收站成功")
}

// 用户上传或秒传
func UserUploadFile(c *gin.Context) {
	username := c.PostForm("username")
	// todo校验用户是否存在

	// 接收文件
	file, head, err := c.Request.FormFile("file")
	if err != nil {
		utils.R(c, err, "文件上传失败", nil)
		return
	}
	defer file.Close()
	// 创建文件元数据
	fileMeta := db.FileMeta{
		FileName: head.Filename,
		Location: filepath.Join(utils.GetCurrentAbPath(), head.Filename),
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}
	// 创建临时文件(tmp目录必须事先存在!) -> 报错:
	// The system cannot find the path specified.Failed to save data into file,errinvalid argument
	newFile, err := os.Create(fileMeta.Location)
	if err != nil {
		utils.R(c, err, "上传失败", nil)
	}
	defer newFile.Close()
	// 求hash
	fileMeta.FileSha1 = utils.FileSha1(newFile)
	// 查询是否已经有该文件(根据hash比对)
	fmeta, err := db.GetFileMetaDB(fileMeta.FileSha1)
	if fmeta != (db.FileMeta{}) {
		// 插入到用户表
		suc := db.OnUserFileUploadFinished(
			username, fmeta.FileSha1, fmeta.FileName, fmeta.FileSize)
		if suc {
			utils.R(c, err, "上传失败", "上传成功")
			return
		} else {
			c.JSON(500, gin.H{
				"status": 500,
				"msg":    "文件上传失败",
			})
			return
		}
	}
	// 复制文件流内容到本地
	fileMeta.FileSize, err = io.Copy(newFile, file)
	if err != nil {
		utils.R(c, err, "上传失败", nil)
	}
	// 打开
	newFile.Seek(0, 0)
	// 新增文件表记录
	_ = db.InsertFileMetaDB(fileMeta)
	// 插入到用户表
	suc := db.OnUserFileUploadFinished(
		username, fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize)
	// fmt.Println(username)
	if !suc {
		c.JSON(500, gin.H{
			"status": 500,
			"msg":    "文件上传失败",
		})
		return
	}
	utils.R(c, err, "上传失败", "上传成功")
}
func registerUser(middles ...gin.HandlerFunc) {
	// 创建路由组v1/user
	user := DefineRouteGroup(v1, "user", r)
	// 添加中间件
	if middles != nil {
		user.Use(middles...)
	}
	// username唯一标识
	// 根据username查询其所有的文件
	user.GET("query", GetUserFileByUserName)
	// 文件移入回收站
	user.DELETE("remove", RemoveUserFileByFilehash)
	// 根据filehash删除文件
	user.DELETE("delete", DeleteByFilehashHandler)
	// 用户秒传或上传
	user.POST("upload", UserUploadFile)
}
