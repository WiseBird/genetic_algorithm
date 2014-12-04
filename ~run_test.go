package genetic_algorithm

import (
	log "github.com/cihub/seelog"
	. "gopkg.in/check.v1"
	"testing"
)

func Test(t *testing.T) {
	defer log.Flush()

	log.ReplaceLogger(log.Disabled)
	//setupLogger()

	TestingT(t)
}
func setupLogger() {
	logger, err := log.LoggerFromConfigAsString("<seelog minlevel=\"info\" type=\"sync\"></seelog>")

	if err != nil {
		panic(err)
	}

	log.ReplaceLogger(logger)
}

/*
mockgen -source="selector_base.go" -destination="~mock_selector_base_vm.go" -package="genetic_algorithm"
*/
