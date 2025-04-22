package maintenance

import (
	"os"

	"github.com/wutipong/mangaweb3-backend/configuration"
)

func PurgeCache() error {
	c := configuration.Get()

	return os.RemoveAll(c.CachePath)
}
