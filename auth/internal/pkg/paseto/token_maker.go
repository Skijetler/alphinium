package paseto

import (
	"encoding/hex"
	conf "github.com/Skijetler/alphinium/auth/internal/config"
	"github.com/vk-rv/pvx"
	"time"
)

type TokenMaker interface {
	NewAccessToken(userId string) (string, error)
	NewRefreshToken(userId string) (string, error)
	ParseAccessToken(token string) (string, error)
	ParseRefreshToken(token string) (string, error)
}

type Maker struct {
	paseto        *pvx.ProtoV4Local
	accessSymKey  *pvx.SymKey
	accessTTL     time.Duration
	refreshSymKey *pvx.SymKey
	refreshTTL    time.Duration
	assertion     string
}

func NewPasetoMaker(conf *conf.Config) (TokenMaker, error) {
	ak, err := hex.DecodeString(conf.TokenMaker.AccessKey)
	if err != nil {
		return nil, err
	}

	rk, err := hex.DecodeString(conf.TokenMaker.RefreshKey)
	if err != nil {
		return nil, err
	}

	return &Maker{
		paseto:        pvx.NewPV4Local(),
		accessSymKey:  pvx.NewSymmetricKey(ak, pvx.Version4),
		accessTTL:     conf.TokenMaker.AccessTtl,
		refreshSymKey: pvx.NewSymmetricKey(rk, pvx.Version4),
		refreshTTL:    conf.TokenMaker.RefreshTtl,
		assertion:     conf.TokenMaker.Assert,
	}, nil
}

func (m *Maker) NewAccessToken(sessionId string) (string, error) {
	exp := time.Now().Add(m.accessTTL)

	return m.paseto.Encrypt(
		m.accessSymKey,
		&pvx.RegisteredClaims{Subject: sessionId, Expiration: &exp},
		pvx.WithAssert([]byte(m.assertion)),
	)
}

func (m *Maker) NewRefreshToken(sessionId string) (string, error) {
	exp := time.Now().Add(m.refreshTTL)

	return m.paseto.Encrypt(
		m.refreshSymKey,
		&pvx.RegisteredClaims{Subject: sessionId, Expiration: &exp},
		pvx.WithAssert([]byte(m.assertion)),
	)
}

func (m *Maker) ParseAccessToken(token string) (string, error) {
	var c pvx.RegisteredClaims

	err := m.paseto.Decrypt(token, m.accessSymKey, pvx.WithAssert([]byte(m.assertion))).ScanClaims(&c)
	if err != nil {
		return "", err
	}

	return c.Subject, nil
}

func (m *Maker) ParseRefreshToken(token string) (string, error) {
	var c pvx.RegisteredClaims

	err := m.paseto.Decrypt(token, m.refreshSymKey, pvx.WithAssert([]byte(m.assertion))).ScanClaims(&c)
	if err != nil {
		return "", err
	}

	return c.Subject, nil
}
