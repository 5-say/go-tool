# responsex

## 自定义

```go
import (
	"github.com/5-say/go-tool/gin/responsex"
	"github.com/gin-gonic/gin"
)

func CustomResponse(c *gin.Context) responsex.ResponseT {
	return responsex.ResponseT{
		C: c,

		DefaultSuccessHTTPCode: 200,
		DefaultErrorHTTPCode:   422,

		SuccessDataWarp: func(data interface{}) interface{} {
			return gin.H{
				"message":  "success",
				"resource": data,
			}
		},

		ErrorDataWarp: func(data interface{}) interface{} {
			switch data := data.(type) {
			case string:
				return gin.H{
					"message":  data,
					"resource": gin.H{},
				}
			case error:
				return gin.H{
					"message":  data.Error(),
					"resource": gin.H{},
				}
			default:
				return data
			}
		},
	}
}
```

## 调用

```go
func Login(c *gin.Context) {
    // 初始化响应
    var (
        r = CustomResponse(c)
    )

    // 构造 错误响应、写入日志 并 返回
    if err != nil {
        r.Error("error message").Logf("error id %v", 1)
        return
    }

    // 构造 正确响应
    r.Success(gin.H{})
}
```
