package cases

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dustin/go-humanize"

	"github.com/ipfs/testground/plans/smlbench"
	"github.com/ipfs/testground/sdk/iptb"
	"github.com/ipfs/testground/sdk/runtime"
)

// simpleAddTC is a simple test that adds a file of the specified size to an IPFS node. It measures time to add.
type simpleAddTC struct {
	SizeBytes int64
}

var _ smlbench.SmallBenchmarksTestCase = (*simpleAddTC)(nil)

func (tc *simpleAddTC) Name() string {
	h := strings.ReplaceAll(strings.ToLower(humanize.IBytes(uint64(tc.SizeBytes))), " ", "")
	return fmt.Sprintf("simple-add-%s", h)
}

func (tc *simpleAddTC) Configure(runenv *runtime.RunEnv, spec *iptb.TestEnsembleSpec) {
	spec.AddNodesDefaultConfig(iptb.NodeOpts{Initialize: true, Start: true}, "adder")
}

func (tc *simpleAddTC) Execute(runenv *runtime.RunEnv, ensemble *iptb.TestEnsemble) {
	node := ensemble.GetNode("adder")
	client := node.Client()

	file := smlbench.TempRandFile(runenv, ensemble.TempDir(), tc.SizeBytes)
	defer os.Remove(file.Name())

	tstarted := time.Now()
	_, err := client.Add(file)
	if err != nil {
		runenv.Abort(err)
		return
	}

	runenv.EmitMetric(smlbench.MetricTimeToAdd, float64(time.Now().Sub(tstarted)/time.Millisecond))
	runenv.OK()
}
