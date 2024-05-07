# jwtx

## 跨域资源共享 CORS

使用官方 cors 全局中间件，并添加以下必要配置

```golang
r.Use(cors.New(cors.Config{

    // 允许跨域请求携带 Authorization 头信息
    AllowMethods:  []string{"OPTIONS"},
    AllowHeaders:  []string{"Authorization"},

    // 允许跨域请求获取 token 响应头
    ExposeHeaders: []string{"token"},
}))
```

