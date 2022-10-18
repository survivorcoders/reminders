package argonHash

//The Argon2 algorithm accepts a number of configurable parameters:

// Params
// Memory — The amount of memory used by the algorithm (in kilobytes).
// Iterations — The number of iterations (or passes) over the memory.
// Parallelism — The number of threads (or lanes) used by the algorithm.
// Salt length — Length of the random salt. 16 bytes is recommended for password hashing.
// Key length — Length of the generated key (or password hash). 16 bytes or more is recommended.
type Params struct {
	Memory      uint32
	Iterations  uint32
	Parallelism uint8
	SaltLength  uint32
	KeyLength   uint32
}
