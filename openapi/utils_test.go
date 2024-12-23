package openapi

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func demo() {

}

func TestFuncName(t *testing.T) {
	testCases := []struct {
		name string

		arg any

		res string
	}{
		{
			name: "normal",
			arg:  demo,
			res:  "github.com/DaHuangQwQ/ginx/openapi.demo",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			name := FuncName(tc.arg)
			require.Equal(t, name, tc.res)
		})
	}
}

func Test_camelToHuman(t *testing.T) {
	testCases := []struct {
		name string

		str string

		res string
	}{
		{
			name: "normal",
			str:  "DaHuang",
			res:  "da huang",
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			res := camelToHuman(tc.str)
			require.Equal(t, res, tc.res)
		})
	}
}
