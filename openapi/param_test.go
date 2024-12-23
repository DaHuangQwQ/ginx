package openapi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_parsePathParams(t *testing.T) {
	testCases := []struct {
		name string

		path string

		res []string
	}{
		{
			name: "normal",
			path: "/user/:id",
			res:  []string{"id"},
		},
		{
			name: "empty",
			path: "/user",
			res:  []string(nil),
		},
		{
			name: "some",
			path: "/user/:id/:name",
			res:  []string{"id", "name"},
		},
	}
	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			res := parsePathParams(testCase.path)
			require.Equal(t, res, testCase.res)
		})
	}
}
