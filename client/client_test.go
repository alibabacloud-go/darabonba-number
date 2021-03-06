package client

import (
	"testing"

	"github.com/alibabacloud-go/tea/tea"
	"github.com/alibabacloud-go/tea/utils"
)

func Test_ParseInt(t *testing.T) {
	num := ParseInt(tea.String("2"))
	utils.AssertEqual(t, 2, tea.IntValue(num))
}

func Test_ParseLong(t *testing.T) {
	num := ParseLong(tea.String("2"))
	utils.AssertEqual(t, int64(2), tea.Int64Value(num))
}

func Test_ParseFloat(t *testing.T) {
	num := ParseFloat(tea.String("2.0"))
	utils.AssertEqual(t, float32(2), tea.Float32Value(num))
}

func Test_ParseDouble(t *testing.T) {
	num := ParseDouble(tea.String("2.00"))
	utils.AssertEqual(t, 2.00, tea.Float64Value(num))
}
