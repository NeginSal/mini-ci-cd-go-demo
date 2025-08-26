package hello

func Hello(name string) string {
	if name == "" {
		name = "DevOps"
	}
	return "Hello, " + name + " World!"
}
