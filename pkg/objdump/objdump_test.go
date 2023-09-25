package objdump

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const (
	exampleObjdump = `00000000007d11b0 g     F .text	0000000000000570 exmaple.com/whatever.(*Example).Range
00000000007d1720 g     F .text	0000000000000110 example.com/whatever.(*Example).Range.func3
00000000007d1830 g     F .text	0000000000000070 example.com/whatever.(*Example).Range.func2
00000000007d18a0 g     F .text	0000000000000070 example.com/whatever.(*Example).Range.func4
00000000007d1910 g     F .text	0000000000000370 example.com/whatever.(*Example).Range.func1
00000000007d1dc0 g     F .text	00000000000000e0 example.com/whatever.(*Example).DeleteRange`
)

func TestFindAddrs(t *testing.T) {
	addrs, err := findAddrs([]byte(exampleObjdump), []string{"(*Example).Range", "(*Example).DeleteRange"})
	require.NoError(t, err)
	require.Equal(t, addrs["(*Example).Range"], "00000000007d11b0")
	require.Equal(t, addrs["(*Example).DeleteRange"], "00000000007d1dc0")
}
