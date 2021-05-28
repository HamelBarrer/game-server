package security

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"strings"

	"golang.org/x/crypto/argon2"
)

type Hash struct {
	salt    uint32
	time    uint32
	memory  uint32
	threads uint8
	keyLen  uint32
}

func GenerateHash(password string) (string, error) {
	h := &Hash{
		salt:    16,
		time:    1,
		memory:  64 * 1024,
		threads: 4,
		keyLen:  32,
	}

	salt, err := genereteBytes(h.salt)
	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, h.time, h.memory, h.threads, h.keyLen)
	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	mergeHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s", argon2.Version, h.memory, h.time, h.threads, b64Salt, b64Hash)

	return mergeHash, nil
}

func genereteBytes(salt uint32) ([]byte, error) {
	b := make([]byte, salt)
	if _, err := rand.Read(b); err != nil {
		return nil, err
	}

	return b, nil
}

func CompareHash(password, codingHash string) (bool, error) {
	p, salt, hash, err := decodeHash(codingHash)
	if err != nil {
		return false, err
	}

	new_hash := argon2.IDKey([]byte(password), salt, p.time, p.memory, p.threads, p.keyLen)

	if subtle.ConstantTimeCompare(hash, new_hash) == 1 {
		return true, nil
	}

	return false, nil
}

func decodeHash(codingHash string) (*Hash, []byte, []byte, error) {
	vals := strings.Split(codingHash, "$")
	if len(vals) != 6 {
		fmt.Println(vals)
		return nil, nil, nil, errors.New("invalid hash")
	}

	var version int
	_, err := fmt.Sscanf(vals[2], "v=%d", &version)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != argon2.Version {
		return nil, nil, nil, err
	}

	h := &Hash{}
	_, err = fmt.Sscanf(vals[3], "m=%d,t=%d,p=%d", &h.memory, &h.time, &h.threads)
	if err != nil {
		return nil, nil, nil, err
	}

	salt, err := base64.RawStdEncoding.Strict().DecodeString(vals[4])
	if err != nil {
		return nil, nil, nil, err
	}

	h.salt = uint32(len(salt))

	hash, err := base64.RawStdEncoding.Strict().DecodeString(vals[5])
	if err != nil {
		return nil, nil, nil, err
	}

	h.keyLen = uint32(len(hash))

	return h, salt, hash, nil
}
