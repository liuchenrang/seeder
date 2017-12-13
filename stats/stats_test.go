package stats

import (
	"testing"
	"seeder/bootstrap"
	"seeder/logger"
	"seeder/config"
	"github.com/liuchenrang/log4go"
	"fmt"
)

func TestStats_Dig(t *testing.T) {
	// Different allocations should not be equal.
	Application := bootstrap.NewApplication()
	seederConfig := config.NewSeederConfig("../seeder.yaml")
	Application.Set("globalSeederConfig", seederConfig)
	Application.Set("globalLogger", SeederLogger.NewLogger4g(log4go.DEBUG, seederConfig))


	st := &Stats{}
	st.Dig()
	st.Dig()
	st.Dig()
	st.Dig()
	st.Dig()
	fmt.Printf("total=%d",st.GetTotal())
}
