// Every Go program is made up of packages.
// Programs start running in package main.
package main

// grouped/parenthesized/factored imports
import (
	"fmt"
	"math/rand"
)

// multiple imports
import "os"
import "math"

// identifiers starts with Capital letters are exported while the one starts with small letters are unexported
var ExportedIdentifier string   // available outside package
var unExportedIdentifier string // unavailable outside package

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println("PI value is", math.Pi)
	fmt.Println("Current Dir is ", os.Chdir("C:\\MahaSecure"))
}
