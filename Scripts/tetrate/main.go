// package main
package tetrate

// map[string]string

type SetInterface interface {
	Add()
	Remove()
	Contains()
	Len()
}

type CustomString struct {
	items []string
}

func CreateCustomerString() *CustomString {
	return &CustomString{}
}

func (c *CustomString) Add(value string) {
	c.items = append(c.items, value)
}

func (c *CustomString) Remove(key string) string {
	found := false
	var newStringList []string
	for i := 0; i < len(c.items); i++ {
		if key == c.items[i] {
			found = true
			continue
		} else {
			newStringList = append(newStringList, c.items[i])
		}
	}
	if found == false {
		return "Key is Not Present"
	}
	c.items = newStringList
	return "Key is Removed"
}

func (c *CustomString) Contains(key string) bool {
	for i := 0; i < len(c.items); i++ {
		if key == c.items[i] {
			return true
		}
	}
	return false
}

func (c *CustomString) c() int {
	return len(c.items)
}
