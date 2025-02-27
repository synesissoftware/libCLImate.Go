package libclimate_test

import (
	"github.com/stretchr/testify/require"
	libclimate "github.com/synesissoftware/libCLImate.Go"

	"testing"
)

func Test_INIT_Flags_1(t *testing.T) {
	require.Equal(t, int64(0), int64(libclimate.InitFlag_None))

	require.NotEqual(t, libclimate.InitFlag_None, libclimate.InitFlag_PanicOnFailure)
	require.NotEqual(t, libclimate.InitFlag_None, libclimate.InitFlag_NoHelpFlag)
	require.NotEqual(t, libclimate.InitFlag_None, libclimate.InitFlag_NoVersionFlag)

	require.NotEqual(t, libclimate.InitFlag_PanicOnFailure, libclimate.InitFlag_NoHelpFlag)
	require.NotEqual(t, libclimate.InitFlag_PanicOnFailure, libclimate.InitFlag_NoVersionFlag)

	require.NotEqual(t, libclimate.InitFlag_NoHelpFlag, libclimate.InitFlag_NoVersionFlag)

	require.Equal(t, int64(0), int64(libclimate.InitFlag_PanicOnFailure&libclimate.InitFlag_NoHelpFlag&libclimate.InitFlag_NoVersionFlag))
}

func Test_PARSE_Flags_1(t *testing.T) {
	require.Equal(t, int64(0), int64(libclimate.ParseFlag_None))

	require.NotEqual(t, libclimate.ParseFlag_None, libclimate.ParseFlag_PanicOnFailure)
	require.NotEqual(t, libclimate.ParseFlag_None, libclimate.ParseFlag_DontCheckUnused)

	require.NotEqual(t, libclimate.ParseFlag_PanicOnFailure, libclimate.ParseFlag_DontCheckUnused)

	require.Equal(t, int64(0), int64(libclimate.ParseFlag_PanicOnFailure&libclimate.ParseFlag_DontCheckUnused))
}
