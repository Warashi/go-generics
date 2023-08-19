package genio_test

import (
	"testing"

	"github.com/Warashi/go-generics/genio"
)

func TestJSON(t *testing.T) {
	r :=
		genio.NewJSONDecoder[struct{}](
			genio.NewJSONEncoder(
				genio.NewSliceReader([]struct{}{{}, {}, {}}),
			),
		)

	result, err := genio.ReadAll(r)
	if err != nil {
		t.Fatal(err)
	}
	if len(result) != 3 {
		t.Fatal("invalid length")
	}
}
