package internal

import (
	_ "embed"
	"github.com/busy-cloud/boat/api"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"time"
)

//go:embed icon.png
var defaultIcon []byte

const gmtFormat = "Mon, 02 Jan 2006 15:04:05 GMT"

func init() {

	//启动时间作为默认图标的修改时间
	bootTime := time.Now()

	api.Register("GET", "plugin/list", func(ctx *gin.Context) {
		var as []*Plugin
		plugins.Range(func(name string, item *Plugin) bool {
			as = append(as, item)
			return true
		})
		api.OK(ctx, as)
	})

	api.Register("GET", "plugin/:app", func(ctx *gin.Context) {
		app := plugins.Load(ctx.Param("app"))
		if app != nil {
			api.Fail(ctx, "找不到插件")
			return
		}
		api.OK(ctx, app)
	})

	api.Register("GET", "plugin/:app/icon", func(ctx *gin.Context) {
		app := plugins.Load(ctx.Param("app"))
		if app != nil {
			api.Fail(ctx, "找不到插件")
			return
		}

		icon := app.Icon
		if icon == "" {
			icon = IconName
		}

		file, err := os.Open(IconName)
		if err != nil {
			//return nil, err
			//return icon, nil //使用默认图片
			ctx.Header("Last-Modified", bootTime.UTC().Format(gmtFormat))
			ctx.Header("Content-Type", "image/png")
			_, _ = ctx.Writer.Write(defaultIcon)
			return
		}
		defer file.Close()

		st, _ := file.Stat()
		buf, err := io.ReadAll(file)
		if err != nil {
			_ = ctx.Error(err)
			return
		}

		ctx.Header("Last-Modified", st.ModTime().UTC().Format(gmtFormat))
		ctx.Header("Content-Type", "image/png")
		_, _ = ctx.Writer.Write(buf)
	})
}
