package internal

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func Proxy(ctx *gin.Context) {
	//插件 接口
	if str, has := strings.CutPrefix(ctx.Request.RequestURI, "/api/"); has {
		if app, _, has := strings.Cut(str, "/"); has {
			if p := plugins.Load(app); p != nil {
				if p.proxy == nil {
					//执行反向代理
					ctx.Abort()
					p.proxy.ServeHTTP(ctx.Writer, ctx.Request)
					return
				}
			}
		}
	}

	if ctx.Request.Method != "GET" {
		ctx.Next()
		return
	}

	//插件 前端页面
	if str, has := strings.CutPrefix(ctx.Request.RequestURI, "/plugin/"); has {
		if len(str) == 0 {
			return
		}

		var app string
		var path string

		if strings.Index(str, "/") > 0 {
			app, path, _ = strings.Cut(str, "/")
		} else {
			app = str
			path = "index.html"
		}

		if p := plugins.Load(app); p != nil {
			if p.static != nil {
				ctx.FileFromFS(path, p.static)
				ctx.Abort()
				return
			}
		}
	}

	//TODO 这个代理不太合适，先这样吧
	//插件 模板页面接口
	if str, has := strings.CutPrefix(ctx.Request.RequestURI, "/api/page/"); has {
		if app, _, has := strings.Cut(str, "/"); has {
			if p := plugins.Load(app); p != nil {
				if p.proxy == nil {
					//执行反向代理
					ctx.Abort()
					p.proxy.ServeHTTP(ctx.Writer, ctx.Request)
					return
				}
			}
		}
	}

	ctx.Next()
}
