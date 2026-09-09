package data

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
)

const sqlLogPrefix = "\x1b[38;5;39m[SQL]\x1b[0m"

type sqlQueryTraceKey struct{}

type sqlQueryTrace struct {
	startedAt time.Time
	sql       string
}

// sqlTracer logs SQL text and elapsed time in debug mode. Query arguments are
// deliberately omitted because they can contain API keys and other secrets.
type sqlTracer struct {
	log *zap.SugaredLogger
}

func (t sqlTracer) TraceQueryStart(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryStartData) context.Context {
	return context.WithValue(ctx, sqlQueryTraceKey{}, sqlQueryTrace{
		startedAt: time.Now(),
		sql:       data.SQL,
	})
}

func (t sqlTracer) TraceQueryEnd(ctx context.Context, _ *pgx.Conn, data pgx.TraceQueryEndData) {
	trace, ok := ctx.Value(sqlQueryTraceKey{}).(sqlQueryTrace)
	if !ok {
		return
	}

	elapsed := time.Since(trace.startedAt)
	if data.Err != nil {
		t.log.Errorf("%s %s (%s): %v", sqlLogPrefix, trace.sql, elapsed, data.Err)
		return
	}
	t.log.Infof("%s %s (%s)", sqlLogPrefix, trace.sql, elapsed)
}
