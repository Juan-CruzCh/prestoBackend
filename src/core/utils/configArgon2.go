package utils

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"fmt"
	"strings"

	"golang.org/x/crypto/argon2"
)

type configArgon2 struct {
	memory      uint32
	iterations  uint32
	parallelism uint8
	saltLength  uint32
	keyLength   uint32
}

var defaultConfig = configArgon2{
	memory:      64 * 1024,
	iterations:  10,
	parallelism: 2,
	saltLength:  16,
	keyLength:   32,
}

func EncriptarPassword(password string) (string, error) {
	salt, err := generarRandomBytes(defaultConfig.saltLength)
	if err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, defaultConfig.iterations, defaultConfig.memory, defaultConfig.parallelism, defaultConfig.keyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	encodedHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, defaultConfig.memory, defaultConfig.iterations, defaultConfig.parallelism, b64Salt, b64Hash)

	return encodedHash, nil
}

func generarRandomBytes(length uint32) ([]byte, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	return b, err
}

func ComparePasswordAndHash(password, encodedHash string) (bool, error) {
	cfg, salt, hash, err := decodeHash(encodedHash)
	if err != nil {
		return false, err
	}

	newHash := argon2.IDKey([]byte(password), salt, cfg.iterations, cfg.memory, cfg.parallelism, cfg.keyLength)

	if subtle.ConstantTimeCompare(hash, newHash) == 1 {
		return true, nil
	}

	return false, nil
}

func decodeHash(encodedHash string) (*configArgon2, []byte, []byte, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return nil, nil, nil, fmt.Errorf("invalid hash format")
	}

	var version int
	_, err := fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil || version != argon2.Version {
		return nil, nil, nil, fmt.Errorf("incompatible Argon2 version")
	}

	var cfg configArgon2
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &cfg.memory, &cfg.iterations, &cfg.parallelism)
	if err != nil {
		return nil, nil, nil, fmt.Errorf("invalid Argon2 parameters")
	}

	salt, err := base64.RawStdEncoding.Strict().DecodeString(parts[4])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("invalid base64 salt")
	}
	cfg.saltLength = uint32(len(salt))

	hash, err := base64.RawStdEncoding.Strict().DecodeString(parts[5])
	if err != nil {
		return nil, nil, nil, fmt.Errorf("invalid base64 hash")
	}
	cfg.keyLength = uint32(len(hash))

	return &cfg, salt, hash, nil
}
