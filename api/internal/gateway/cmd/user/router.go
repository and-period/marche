package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"strconv"

	"github.com/and-period/furumaru/api/pkg/cors"
	"github.com/and-period/furumaru/api/pkg/jst"
	"github.com/gin-contrib/gzip"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/line/line-bot-sdk-go/v7/linebot"
	"github.com/newrelic/go-agent/v3/integrations/nrgin"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func newRouter(reg *registry, logger *zap.Logger) *gin.Engine {
	opts := make([]gin.HandlerFunc, 0)
	opts = append(opts, nrgin.Middleware(reg.newRelic))
	opts = append(opts, accessLogger(logger, reg))
	opts = append(opts, cors.NewGinMiddleware())
	opts = append(opts, gzip.Gzip(gzip.DefaultCompression))
	opts = append(opts, ginzap.RecoveryWithZap(logger, true))

	rt := gin.New()
	rt.Use(opts...)

	reg.v1.Routes(rt.Group(""))

	// other routes
	rt.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "ok")
	})
	rt.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, "not found")
	})

	return rt
}

type wrapResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w *wrapResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func accessLogger(logger *zap.Logger, reg *registry) gin.HandlerFunc {
	skipPaths := map[string]bool{
		"/health": true,
	}

	return func(ctx *gin.Context) {
		w := &wrapResponseWriter{
			ResponseWriter: ctx.Writer,
			body:           bytes.NewBufferString(""),
		}
		ctx.Writer = w

		start := jst.Now()
		method := ctx.Request.Method
		path := ctx.Request.URL.Path
		ctx.Next()

		if _, ok := skipPaths[path]; ok {
			return
		}

		end := jst.Now()
		status := ctx.Writer.Status()

		fields := []zapcore.Field{
			zap.Int("status", status),
			zap.String("method", method),
			zap.String("path", path),
			zap.String("query", ctx.Request.URL.RawQuery),
			zap.String("ip", ctx.ClientIP()),
			zap.String("user-agent", ctx.Request.UserAgent()),
			zap.Duration("latency", end.Sub(start)),
			zap.String("time", end.Format("2006-01-02 15:04:05")),
			zap.String("userId", ctx.GetHeader("adminId")),
		}

		// ~ 499
		if status < 500 {
			logger.Info(path, fields...)
			return
		}

		// 500 ~
		res := w.body.String()
		fields = append(fields, zap.String("response", res))
		fields = append(fields, zap.Strings("errors", ctx.Errors.Errors()))
		logger.Error(path, fields...)

		if reg.line == nil {
			return
		}

		altText := fmt.Sprintf("[ふるマルアラート] %s", reg.appName)
		components := []linebot.FlexComponent{
			newAlertContent("service", reg.appName),
			newAlertContent("env", reg.env),
			newAlertContent("status", strconv.FormatInt(int64(status), 10)),
			newAlertContent("method", method),
			newAlertContent("path", path),
			newAlertContent("detail", res),
		}
		_ = reg.line.PushMessage(ctx, newAlertMessage(altText, components))
	}
}

func newAlertMessage(altText string, components []linebot.FlexComponent) *linebot.FlexMessage {
	container := &linebot.BubbleContainer{
		Type:      linebot.FlexContainerTypeBubble,
		Direction: linebot.FlexBubbleDirectionTypeLTR,
		Header: &linebot.BoxComponent{
			Type:   linebot.FlexComponentTypeBox,
			Layout: linebot.FlexBoxLayoutTypeVertical,
			Contents: []linebot.FlexComponent{
				&linebot.TextComponent{
					Type:   linebot.FlexComponentTypeText,
					Text:   "ふるマル APIアラート",
					Size:   linebot.FlexTextSizeTypeXl,
					Weight: linebot.FlexTextWeightTypeBold,
					Color:  "#FFFFFF", // white
				},
			},
		},
		Body: &linebot.BoxComponent{
			Type:     linebot.FlexComponentTypeBox,
			Layout:   linebot.FlexBoxLayoutTypeVertical,
			Contents: components,
		},
		Size: linebot.FlexBubbleSizeTypeGiga,
		Styles: &linebot.BubbleStyle{
			Header: &linebot.BlockStyle{
				BackgroundColor: "#F44336", // red
			},
		},
	}
	return linebot.NewFlexMessage(altText, container)
}

func newAlertContent(field, value string) *linebot.BoxComponent {
	const (
		fcolor = "#aaaaaa" // gray
		vcolor = "#666666" // black
	)
	var (
		fflex = 1
		vflex = 4
	)
	return &linebot.BoxComponent{
		Type:   linebot.FlexComponentTypeBox,
		Layout: linebot.FlexBoxLayoutTypeBaseline,
		Contents: []linebot.FlexComponent{
			&linebot.TextComponent{
				Type:  linebot.FlexComponentTypeText,
				Text:  field,
				Size:  linebot.FlexTextSizeTypeSm,
				Color: fcolor,
				Flex:  &fflex,
			},
			&linebot.TextComponent{
				Type:  linebot.FlexComponentTypeText,
				Text:  value,
				Wrap:  true,
				Size:  linebot.FlexTextSizeTypeSm,
				Color: vcolor,
				Flex:  &vflex,
			},
		},
	}
}
