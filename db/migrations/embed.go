// Package migrations embeds the versioned SQL migrations so a deployed binary
// can migrate itself without shipping the db/migrations directory alongside it.
package migrations

import "embed"

// FS holds every migration file in this directory.
//
//go:embed *.sql
var FS embed.FS
