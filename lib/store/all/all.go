// Package all is a meta-package that imports all store implementations.
//
// This is a HACK to make tests work consistently.
package all

import (
	_ "github.com/ToastyTheBot/anubis/lib/store/bbolt"
	_ "github.com/ToastyTheBot/anubis/lib/store/memory"
	_ "github.com/ToastyTheBot/anubis/lib/store/s3api"
	_ "github.com/ToastyTheBot/anubis/lib/store/valkey"
)
