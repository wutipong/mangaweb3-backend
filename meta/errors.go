package meta

import "github.com/wutipong/mangaweb3-backend/errors"

var ErrMetaDataNotFound = errors.New(2_000_000, "metadata for '%s' not found.")
