package hash

import (
	"encoding/hex"
	conf "github.com/Skijetler/alphinium/auth/internal/config"
	"golang.org/x/crypto/argon2"
)

type PasswordHasher interface {
	Hash(password string) string
	Compare(hash, password string) bool
}

type Hasher struct {
	salt       []byte
	iterations uint32
	memory     uint32
	threads    uint8
	keyLen     uint32
}

func NewPasswordHasher(conf *conf.Config) PasswordHasher {
	if conf.Hasher.Iterations == 0 {
		conf.Hasher.Iterations = 1
	}
	if conf.Hasher.Memory == 0 {
		conf.Hasher.Memory = 64 * 1024
	}
	if conf.Hasher.Threads == 0 {
		conf.Hasher.Threads = 1
	}
	if conf.Hasher.KeyLen == 0 {
		conf.Hasher.KeyLen = 32
	}

	return &Hasher{
		salt:       []byte(conf.Hasher.Salt),
		iterations: conf.Hasher.Iterations,
		memory:     conf.Hasher.Memory,
		threads:    conf.Hasher.Threads,
		keyLen:     conf.Hasher.KeyLen,
	}
}

func (h *Hasher) Hash(password string) string {
	hash := argon2.IDKey([]byte(password), h.salt, h.iterations, h.memory, h.threads, h.keyLen)
	return hex.EncodeToString(hash)
}

func (h *Hasher) Compare(hash, password string) bool {
	passHash := h.Hash(password)
	return hash == passHash
}
