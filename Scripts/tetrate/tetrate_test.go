package tetrate

import (
	"testing"
)

func TestAdd(t *testing.T) {
	customSet := CreateCustomerString()
	customSet.Add("vikas")
	customSet.Add("Yaro")
	customSet.Add("Yaro")
	//assert.Equal(t, 2, customSet.Len())
}

// shape --> area()
// rectangle circle area() , area()
