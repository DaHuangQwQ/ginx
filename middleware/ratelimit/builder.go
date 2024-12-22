package ratelimit

import (
	_ "embed"
	"fmt"
	"net/http"

	"github.com/DaHuangQwQ/ginx"

	"github.com/gin-gonic/gin"
)

type Option func(builder *Builder)

type Builder struct {
	prefix  string
	limiter ginx.Limiter
	l       ginx.Logger
}

func NewBuilder(limiter ginx.Limiter, l ginx.Logger, opts ...Option) *Builder {
	res := &Builder{
		prefix:  "ip-limiter",
		limiter: limiter,
		l:       l,
	}
	for _, opt := range opts {
		opt(res)
	}

	return res
}

func (b *Builder) Prefix(prefix string) *Builder {
	b.prefix = prefix
	return b
}

func (b *Builder) Build() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		limited, err := b.limiter.Limit(ctx, fmt.Sprintf("%s:%s", b.prefix, ctx.ClientIP()))
		if err != nil {
			b.l.Warn("限流 redis 宕机了")
			ctx.AbortWithStatus(http.StatusInternalServerError)
			return
		}
		if limited {
			b.l.Warn("限流了" + ctx.ClientIP())
			ctx.AbortWithStatus(http.StatusTooManyRequests)
			return
		}
		ctx.Next()
	}
}

func WithPrefixName(name string) Option {
	return func(builder *Builder) {
		builder.prefix = name
	}
}
