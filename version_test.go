package libclimate_test

import (
	"github.com/stretchr/testify/require"
	libclimate "github.com/synesissoftware/libCLImate.Go"

	"testing"
)

const (
	Expected_VersionMajor uint16 = 0
	Expected_VersionMinor uint16 = 7
	Expected_VersionPatch uint16 = 0
	Expected_VersionAB    uint16 = 0x4003
)

func Test_Version_Elements(t *testing.T) {
	require.Equal(t, Expected_VersionMajor, libclimate.VersionMajor)
	require.Equal(t, Expected_VersionMinor, libclimate.VersionMinor)
	require.Equal(t, Expected_VersionPatch, libclimate.VersionPatch)
	require.Equal(t, Expected_VersionAB, libclimate.VersionAB)
}

func Test_Version(t *testing.T) {
	require.Equal(t, uint64(0x0000_0007_0000_4003), libclimate.Version)
}

func Test_Version_String(t *testing.T) {
	require.Equal(t, "0.7.0-alpha3", libclimate.VersionString())
}
