package tests

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/core/vm"
)

func TestEWASMVM(t *testing.T) {
	t.Parallel()
	st := new(testMatcher)
	st.walk(t, stateTestDir+"/stEWASMTests", func(t *testing.T, name string, test *StateTest) {
		for _, subtest := range test.Subtests() {
			subtest := subtest
			key := fmt.Sprintf("%s/%d", subtest.Fork, subtest.Index)
			name := name + "/" + key
			t.Run(key, func(t *testing.T) {
				if subtest.Fork == "Constantinople" {
					t.Skip("constantinople not supported yet")
				}
				withTrace(t, test.gasLimit(subtest), func(vmconfig vm.Config) error {
					_, err := test.Run(subtest, vmconfig)
					return st.checkFailure(t, name, err)
				})
			})
		}
	})
}