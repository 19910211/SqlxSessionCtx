package sessionctx

import (
	"context"
	"go-zero/core/stores/sqlx"
)

const _sqlxSessionKeyCtx = "__sqlxSessionKeyCtx"

// New sqlx session 上下文
func New(session sqlx.Session) context.Context {
	return NewCtx(context.Background(), session)
}

// NewCtx sqlx session 上下文
func NewCtx(ctx context.Context, session sqlx.Session) context.Context {
	return context.WithValue(ctx, _sqlxSessionKeyCtx, session)
}

// GetSession 获取sqlx session 上下文
func GetSession(sqlSessionCtx context.Context, defaultSession sqlx.Session) sqlx.Session {
	if sqlSession, ok := sqlSessionCtx.Value(_sqlxSessionKeyCtx).(sqlx.Session); ok {
		if sqlSession != nil {
			return sqlSession
		}
	}

	return defaultSession
}
