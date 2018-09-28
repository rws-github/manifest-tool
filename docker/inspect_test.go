package docker

import "testing"

func TestValidateName(t *testing.T) {
	var crctnames = []struct {
		a string
	}{
		{"localhost:5000/hello-world"},
		{"myregistrydomain:5000/java"},
		{"docker.io/debian"},
	}
	var wrngnames = []struct {
		b string
	}{
		{"localhost:5000,hello-world"},
		{"myregistrydomain:5000&java"},
		{"docker.io@busybox"},
	}

	for _, i := range crctnames {
		res := validateName(i.a)
		if res != nil {
			t.Errorf("%s is an invalid name: %s", i.a, res)
		}
	}
	for _, j := range wrngnames {
		res := validateName(j.b)
		if res == nil {
			t.Errorf("%s is an invalid name", j.b)
		}
	}
}
