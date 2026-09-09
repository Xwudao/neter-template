package data

import (
	"context"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5"
	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"
)

func TestSQLTracerLogsColoredSQLWithoutArguments(t *testing.T) {
	core, logs := observer.New(zap.InfoLevel)
	tracer := sqlTracer{log: zap.New(core).Sugar()}

	ctx := tracer.TraceQueryStart(context.Background(), nil, pgx.TraceQueryStartData{
		SQL:  "SELECT * FROM users WHERE api_key = $1",
		Args: []any{"secret-api-key"},
	})
	tracer.TraceQueryEnd(ctx, nil, pgx.TraceQueryEndData{})

	entries := logs.All()
	if len(entries) != 1 {
		t.Fatalf("logged entries = %d, want 1", len(entries))
	}
	message := entries[0].Message
	if !strings.HasPrefix(message, sqlLogPrefix+" SELECT * FROM users") {
		t.Errorf("message = %q, want colored SQL prefix", message)
	}
	if strings.Contains(message, "secret-api-key") {
		t.Errorf("message must not contain query arguments: %q", message)
	}
}
