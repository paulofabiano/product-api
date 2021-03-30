package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "nics",
		Price: 10.0,
		SKU:   "abs-asdj-eoks",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
