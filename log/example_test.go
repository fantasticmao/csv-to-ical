package log

func ExampleInfo() {
	Info("hello %v", "world")
	// Output:
	// hello world
}
