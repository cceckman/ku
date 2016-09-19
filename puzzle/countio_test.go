// 2016-09-19 cceckman <charles@cceckman.com>
package puzzle

import(
	"bytes"
	"strings"
	"testing"
)

func TestCountingReader(t *testing.T) {
	for expIn, expOut := range map[string]string{
		"abcde": "abc",
		"qwerty": "qwert",
	}{
		r := strings.NewReader(expIn)
		cr := NewCountingReader(r)

		b := make([]byte, len(expOut))

		n, err := cr.Read(b)
		if err != nil {
			t.Errorf("non-nil error encountered for '%s': %v", expIn, err)
		}
		if n != len(expOut) {
			t.Errorf("expected length mismatch for '%s': got: %d expected: %d", expIn, n, len(expOut))
		}
		if cr.Count != n {
			t.Errorf("expected count mismatch for '%s': got: %d expected: %d", expIn, cr.Count, n)
		}

		for i := range expOut {
			if b[i] != expOut[i] {
				t.Errorf("expected character mismatch for '%s' at %d: got: %c expected: %c",
					expIn, i, b[i], expOut[i])
			}
		}

	}
}

func TestCountingWriter(t *testing.T) {
	for _, expOut := range []string{
		"abcde",
		"qwerty",
	}{
		w := bytes.NewBufferString("")
		cw := NewCountingWriter(w)

		n, err := cw.Write([]byte(expOut))
		if err != nil {
			t.Errorf("non-nil error encountered for '%s': %v", expOut, err)
		}
		if n != len(expOut) {
			t.Errorf("expected length mismatch for '%s': got: %d expected: %d", expOut, n, len(expOut))
		}
		if cw.Count != n {
			t.Errorf("expected count mismatch for '%s': got: %d expected: %d", expOut, cw.Count, n)
		}

		realOut := w.Bytes()
		for i := range expOut {
			if realOut[i] != expOut[i] {
				t.Errorf("expected character mismatch for '%s' at %d: got: %c expected: %c",
					expOut, i, realOut[i], expOut[i])
			}
		}
	}
}
