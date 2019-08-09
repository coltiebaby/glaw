// Our "bucket" that is sometimes leaky.
package jar

type Jar interface {
	// Give something to the jar
	Give()
	// Take something from the jar
	Take()
	// Check to see if the jar is full
	Full() bool
}
