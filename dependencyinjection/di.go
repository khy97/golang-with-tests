package dependencyinjection

import (
	"fmt"
	"io"
	"os"
)

// Greet - Greets the person
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func main() {
	Greet(os.Stdout, "Elodie")
}
