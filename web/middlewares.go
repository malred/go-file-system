package go_file_web

import (
	"github.com/gin-gonic/gin"
	utils "go_file_util"
	"net/http"
	"strconv"
	"time"
)

//解决跨域问题
func Core() gin.HandlerFunc {
	return func(c *gin.Context) {
		method := c.Request.Method
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Headers", "*")
		c.Header("Access-Control-Allow-Methods", "*")
		c.Header("Access-Control-Expose-Headers", "Content-Length,Access-Control-Allow-Origin,Access-Control-Allow-Headers,Content-Type")
		c.Header("Access-Control-Max-Age", "3600")
		c.Header("Access-Control-Allow-Credentials", "true")
		//放行索引options
		if method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
		}
		//处理请求
		c.Next()
	}
}

// 权限认证(验证token)
func Token() gin.HandlerFunc {
	return func(c *gin.Context) {
		// for k, v := range c.Request.Header {
		// 	fmt.Println(k, v)
		// }
		secret := c.Request.Header["Secret"] // 获取前端传来的secret
		token := c.Request.Header["Token"]
		if len(token) == 0 {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "访问未授权",
			})
			return
		}
		// 转为分钟
		timeInt64 := strconv.FormatInt(time.Now().UnixNano()/1e6/1000/60, 10)
		// 当前时间戳(单位:分钟)+Ip对应的Token -> MD5
		md5Str := utils.MD5_malred(timeInt64 + TokenMap[c.ClientIP()])
		// fmt.Println(TokenMap[c.ClientIP()], timeInt64)
		// fmt.Println(timeInt64 + TokenMap[c.ClientIP()])
		// fmt.Println(md5Str, secret[0])
		// 如果和前端用同样的算法加密的md5字符串不同
		if md5Str != secret[0] {
			// 验证不通过，不再调用后续的函数处理
			c.Abort()
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    401,
				"message": "访问未授权",
			})
			return
		}
		// 验证jwt
		// utils.ParseJWT(secret[0][8:11]+secret[0][19:22], token[0])
		//处理请求
		c.Next()
	}
}

// 阻止缓存响应
func NoCache() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
		ctx.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
		ctx.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
		ctx.Next()
	}
}

// 响应 options 请求, 并退出
// func Options() gin.HandlerFunc {
//     return func(ctx *gin.Context) {
//         if ctx.Request.Method != "OPTIONS" {
//             ctx.Next()
//         } else {
//             ctx.Header("Access-Control-Allow-Origin", "*")
//             ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,PATCH,DELETE,OPTIONS")
//             ctx.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
//             ctx.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
//             ctx.Header("Content-Type", "application/json")
//             ctx.AbortWithStatus(200)
//         }
//     }
// }

// 安全设置
func Secure() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Header("Access-Control-Allow-Origin", "*")
		ctx.Header("X-Frame-Options", "DENY")
		ctx.Header("X-Content-Type-Options", "nosniff")
		ctx.Header("X-XSS-Protection", "1; mode=block")
		if ctx.Request.TLS != nil {
			ctx.Header("Strict-Transport-Security", "max-age=31536000")
		}

		// Also consider adding Content-Security-Policy headers
		// ctx.Header("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")
	}
}

// 权限控制(token携带当前用户的权限信息,过滤低于指定权限的请求)
