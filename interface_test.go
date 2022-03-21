package interface_for_test

import "testing"

type Programmer interface {
	writeHelloworld() string
}
type goprogrammer struct {
	first int
}

func (tmp_obj *goprogrammer) writeHelloworld() string {
	return "fmt.Println(\"hello world\")"
}
func TestClient(t *testing.T) {
	var p Programmer
	p = new(goprogrammer)
	t.Log(p.writeHelloworld())
}
