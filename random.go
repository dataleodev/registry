package registry

//Rand Generator generates a random string of specified length n
//It is essential in creating reference ids and tokens
type Randomizer interface {
	Get(length int) (value string)
}
