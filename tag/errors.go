package tag

import "github.com/wutipong/mangaweb3-backend/errors"

var ErrTagNotFound = errors.New(1_000_000, "tag '%s' not found.")
