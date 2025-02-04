package go_say_hello

func SayHello() string {
	return "Hello"
}

func SayHelloTo(fromName string, toName string) string {
	return "Hello," + toName + " my name is " + fromName
}
