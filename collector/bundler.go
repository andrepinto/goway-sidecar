package collector

import (
	"time"
	"google.golang.org/api/support/bundler"
	log "github.com/sirupsen/logrus"
)

func initHttpLoggerBundler(execute SendHttpLogger, delayThreshold int, bundleCountThreshold int) *bundler.Bundler{

	htppLoggerBundler := bundler.NewBundler((*HttpLogger)(nil), func(bundle interface{}) {
		items := bundle.([]*HttpLogger)
		err := execute(items)
		if err != nil {
			log.Printf("failed to send %d traces to  server: %v", len(items), err)
		}
	})
	htppLoggerBundler.DelayThreshold = time.Duration(delayThreshold)  * time.Second
	htppLoggerBundler.BundleCountThreshold = bundleCountThreshold
	return htppLoggerBundler
}

