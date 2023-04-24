package go_file_web

import (
	"fmt"
	db "go_file_db"
	utils "go_file_util"
	"io"
	"io/ioutil"

	// "io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	// "strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// 上传文件接口
func AddFileHandler(c *gin.Context) {
	// 接收文件
	file, head, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "文件上传失败"})
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
		fmt.Printf("Failed to create temporary file,err:%s\n", err.Error())
	}
	defer newFile.Close()
	// 复制文件流内容到本地
	fileMeta.FileSize, err = io.Copy(newFile, file)
	if err != nil {
		fmt.Printf("Failed to save data into file,err%s\n", err.Error())
		return
	}
	// 打开
	newFile.Seek(0, 0)
	// 求hash
	fileMeta.FileSha1 = utils.FileSha1(newFile)
	// 新增文件表记录
	_ = db.InsertFileMetaDB(fileMeta)
	// 更新用户文件表记录
	username := c.PostForm("user")
	// fmt.Println(username)
	suc := db.OnUserFileUploadFinished(
		username, fileMeta.FileSha1, fileMeta.FileName, fileMeta.FileSize)
	fmt.Println(suc)
	if !suc {
		c.JSON(http.StatusOK, gin.H{
			"status": 500,
			"msg":    "文件上传失败",
		})
	}
	// 通用响应
	utils.R(c, err, "文件上传失败", fileMeta)
}

// 根据文件hash查询文件元数据
func GetMetaByFilehashHandler(c *gin.Context) {
	filehash := c.Query("filehash")
	fmeta, err := db.GetFileMetaDB(filehash)
	// 通用响应
	utils.R(c, err, "查询文件元数据失败", fmeta)
}

// 根据文件hash下载文件
func DownloadByFilehashHandler(c *gin.Context) {
	filehash := c.Query("filehash")
	fm, err := db.GetFileMetaDB(filehash)
	// 打开文件
	f, err := os.Open(fm.Location)
	if err != nil {
		fmt.Print(err.Error())
		utils.R(c, err, "下载失败: 找不到该文件", fm)
		return
	}
	defer f.Close()
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Print(err.Error())
		utils.R(c, err, "下载失败: 无法读取文件", f)
		return
	}
	// 设置头,让浏览器知道是下载
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Length", "-1")
	c.Header("Transfer-Encoding", "true")
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+fm.FileName) // 用来指定下载下来的文件名
	c.Header("Content-Transfer-Encoding", "binary")
	// c.Header("Content-Disposition", "attachment;filename=\""+fm.FileName+"\"")
	// 通用响应
	// utils.R(c, err, "查询文件元数据失败", data)
	// c.Data(http.StatusOK, "application/octet-stream", data)
	c.Data(200, "application/octet-stream", data) // r是文件的文件reader流指针
}



//根据filehash修改文件名
func UpdateFilenameByFilehashHandler(c *gin.Context) {
	var params struct {
		Filehash string `json:"filehash"`
		Filename string `json:"filename"`
	}
	//     // 绑定json和结构体(接收json,数据放入结构体)
	if err := c.BindJSON(&params); err != nil {
		return
	}
	filehash := params.Filehash
	filename := params.Filename
	fmt.Println(filehash)
	fmt.Println(filename)
	err := db.UpdateFileMetaDB(filename, filehash)
	utils.R(c, err, "重命名文件失败", "重命名文件成功")
}
// func DelUserHandler(c *gin.Context) {
// 	// 从url获取参数
// 	idStr := c.Query("uid")
// 	// fmt.Println(idStr)
// 	uid, err := strconv.ParseInt(idStr, 10, 64)
// 	err = db.Delete("mal_user", uid)
// 	// 通用响应
// 	utils.R(c, err, "删除角色失败", "删除角色成功")
// }
// func GetOneUserHandler(c *gin.Context) {
// 	// 从url获取参数
// 	idStr := c.Query("uid")
// 	fmt.Println(idStr)
// 	uid, _ := strconv.ParseInt(idStr, 10, 64)
// 	one, err2 := db.GetUserById(uid)
// 	// 通用响应
// 	utils.R(c, err2, "查询角色失败", one)
// }
// func UptUserHandler(c *gin.Context) {
// 	// 从url获取参数
// 	// uid := c.PostForm("uid")
// 	// uname := c.PostForm("uname")
// 	// upass := c.PostForm("upass")
// 	// ridStr := c.PostForm("rid")
// 	user := db.MalUser{}
// 	//绑定json和结构体
// 	if err := c.BindJSON(&user); err != nil {
// 		return
// 	}
// 	uname := user.Uname
// 	upass := user.Upass
// 	rid := user.Rid
// 	uid := user.Id
// 	// fmt.Println(idStr, UserName)
// 	// rid, _ := strconv.ParseInt(ridStr, 10, 64)
// 	err := db.UptUserById(strconv.FormatInt(uid, 10), []string{"uname", "upass", "rid"}, uname, upass, rid)
// 	// 通用响应
// 	utils.R(c, err, "修改角色失败", "修改角色成功")
// }
func registerFile(middles ...gin.HandlerFunc) {
	// 创建路由组v1/user
	file := DefineRouteGroup(v1, "file", r)
	// 添加中间件
	if middles != nil {
		file.Use(middles...)
	}
	// 添加
	file.POST("upload", AddFileHandler)
	// 根据文件hash获取元数据
	file.GET("meta", GetMetaByFilehashHandler)
	// 根据filehash下载文件
	file.GET("download", DownloadByFilehashHandler)
	// 根据filehash修改文件名
	file.PUT("update", UpdateFilenameByFilehashHandler)   
}
 