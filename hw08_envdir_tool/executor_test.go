package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunCmd(t *testing.T) {
	env, err := ReadDir("./testdata/env")
	require.NoError(t, err)

	require.Equal(t, -1, RunCmd([]string{"wdsf"}, env))
}
