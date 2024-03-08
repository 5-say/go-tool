package change

import (
	"testing"
)

func Test_To(t *testing.T) {
	var aa = int(-8)
	// result, err := To[uint32](&aa)
	// assert.Equal(t, err, nil)
	// assert.Equal(t, result, uint32(23))
	t.Log(ToInt64("-32"))
	t.Log(ToFloat64(-64.56789))
	t.Log(ToUint64(&aa))
	t.Fail()
}
