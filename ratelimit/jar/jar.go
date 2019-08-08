package jar

type Jar interface {
	Give()
	Take()
	Full() bool
}
