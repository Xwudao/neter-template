package ent

//go:generate go run -mod=mod entgo.io/ent/cmd/ent generate --idtype int64 --feature sql/versioned-migration ./schema
//go:generate go run entgo.io/ent/cmd/ent generate --idtype int64 --feature sql/versioned-migration ./schema
