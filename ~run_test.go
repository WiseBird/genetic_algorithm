package genetic_algorithm

import ( 
	"testing"
	. "gopkg.in/check.v1"
	log "github.com/cihub/seelog"
)

func Test(t *testing.T) {
	log.ReplaceLogger(log.Disabled)
	TestingT(t)
}

/*
mockgen -source="selector_base.go" -destination="~mock_selector_base_vm.go" -package="genetic_algorithm"
*/