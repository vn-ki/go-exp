package struc

// Person is no one
type Person struct {
	Name string
}

// NamE returns the name of the person
func (p Person) NamE() string {
	return p.Name
}
