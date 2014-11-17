package genetic_algorithm

import ( 
	"testing"
	. "gopkg.in/check.v1"
	log "github.com/cihub/seelog"
)

func Test(t *testing.T) {
	defer log.Flush()

	log.ReplaceLogger(log.Disabled)
	//setupLogger()

	TestingT(t)
}
func setupLogger() {
	logger, err := log.LoggerFromConfigAsString("<seelog type=\"sync\"></seelog>")

	if err != nil {
	    panic(err)
	}

	log.ReplaceLogger(logger)
}
/*
mockgen -source="selector_base.go" -destination="~mock_selector_base_vm.go" -package="genetic_algorithm"
*/