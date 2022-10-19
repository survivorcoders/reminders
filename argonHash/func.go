package argonHash

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"golang.org/x/crypto/argon2"
	"strings"
)

// init error value
var (
	ErrInvalidHash         = errors.New("the encoded hash is not in the correct format")
	ErrIncompatibleVersion = errors.New("incompatible version of argon2")
)

// GenerateFromPassword

// Storing Passwords in most cases you'll want to store the salt and specific
// parameters that you used alongside the hashed password,
// so that it can be reproducibly verified at a later point.
// The standard way to do this is to create an encoded representation
// of the hashed password which looks like this: ->

// $argon2id$v=19$m=65536,t=3,p=2$c29tZXNhbHQ$RdescudvJCsgt3ub+b+dWRWJTmaaJObG

// $argon2id — the variant of Argon2 being used.
// $v=19 — the version of Argon2 being used.
// $m=65536,t=3,p=2 — the memory (m), iterations (t) and parallelism (p) parameters being used.
// $c29tZXNhbHQ — the base64-encoded salt, using standard base64-encoding and no padding.
// $RdescudvJCsgt3ub+b+dWRWJTmaaJObG — the base64-encoded hashed password (derived key), using standard base64-encoding and no padding.

// GenerateFromPassword() function so that it returns a string in this format:

func GenerateFromPassword(password string, p *Params) (encodedHash string, err error) {
	salt, err := generateRandomBytes(p.SaltLength)
	if err != nil {
		return "", err
	}
	hash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	// Base64 encode the salt and hashed password.
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	// Return a string using the standard encoded hash representation.
	encodedHash = fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, p.Memory, p.Iterations, p.Parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

// ComparePasswordAndHash
// Verifying Passwords
// Check whether this new hash is the same as the original one.

func ComparePasswordAndHash(password, encodedHash string) (match bool, err error) {
	// Extract the parameters, salt and derived key from the encoded password
	// hash.
	p, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	// Derive the key from the other password using the same parameters.
	otherHash := argon2.IDKey([]byte(password), salt, p.Iterations, p.Memory, p.Parallelism, p.KeyLength)

	// Check that the contents of the hashed passwords are identical. Note
	// that we are using the subtle.ConstantTimeCompare() function for this
	// to help prevent timing attacks.
	if subtle.ConstantTimeCompare(hash, otherHash) == 1 {
		return true, nil
	}
	return false, nil
}

// DecodeHash
// Extract the salt and parameters from the encoded password hash stored in the database.
// Derive the hash of the plaintext password using the exact same Argon2 variant, version, salt and parameters.

func decodeHash(encodedHash string) (p *Params, salt, hash []byte, err error) {
	//the hashed password :=$argon2id$v=19$m=65536,t=3,p=2$c29tZXNhbHQ$RdescudvJCsgt3ub+b+dWRWJTmaaJObG

	val := strings.Split(encodedHash, "$")
	if len(val) != 6 {
		return nil, nil, nil, ErrInvalidHash
	}

	var version int
	_, err = fmt.Sscanf(val[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, ErrIncompatibleVersion
	}

	p = &Params{}
	_, err = fmt.Sscanf(val[3], "m=%d,t=%d,p=%d", &p.Memory, &p.Iterations, &p.Parallelism)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err = base64.RawStdEncoding.Strict().DecodeString(val[4])
	if err != nil {
		return nil, nil, nil, err
	}
	p.SaltLength = uint32(len(salt))

	hash, err = base64.RawStdEncoding.Strict().DecodeString(val[5])
	if err != nil {
		return nil, nil, nil, err
	}
	p.KeyLength = uint32(len(hash))

	return p, salt, hash, nil
}

// generateRandomBytes() function.
// In this we're using Go's crypto/rand package to generate a cryptographically
// secure random salt, rather than using a fixed salt or a pseudo-random salt.

func generateRandomBytes(n uint32) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}
