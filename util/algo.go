package util

import (
	"bytes"
	"crypto/ecdsa"
	"github.com/dgrijalva/jwt-go"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
	"goweb/mem"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

func SignJwt(id string, userName string, exp int64, secret string) (tokenString string, err error) {
	if secret == "" {
		secret = viper.GetString("jwtSecret")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":       id,
		"username": userName,
		"nbf":      time.Now().Unix(),
		"iat":      time.Now().Unix(),
		"exp":      time.Now().Unix() + exp,
	})
	tokenString, err = token.SignedString([]byte(secret))
	return
}

func ParseJwt(tokenString string, secret string) (id string, userName string, err error) {
	id = ""
	userName = ""
	err = nil

	if secret == "" {
		secret = viper.GetString("jwtSecret")
	}

	token, err := jwt.Parse(tokenString, secretFunc(secret))

	if err != nil {
		return
	} else {
		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			err = mem.TypeAssertError
			return
		}

		if !token.Valid {
			err = mem.TokenNotValid
			return
		}
		return claims["id"].(string), claims["username"].(string), err
	}
}

func Bcrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

func CompareBcrypt(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func RandomHex(num int) string {
	var hexString string = "abcdef"
	hexLen := 16
	var rs string = ""

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < num; i++ {
		rdnNum := rand.Intn(hexLen)
		if rdnNum <= 9 {
			rs += strconv.Itoa(rdnNum)
		} else {
			rs += string(hexString[rdnNum-10])
		}
	}
	return rs
}

func GetPrivateKeyECDSA() *ecdsa.PrivateKey {
	privateKey, err := crypto.HexToECDSA(viper.GetString("privateKey"))

	if err != nil {
		panic(err)
	}
	return privateKey
}

func GetPublicKeyECDSA(privateKeyECDSA *ecdsa.PrivateKey) *ecdsa.PublicKey {
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)

	if !ok {
		panic(mem.EcdsaError)
	}
	return publicKeyECDSA
}

func SignECDSA(content string) ([]byte, *ecdsa.PublicKey, error) {
	privateKeyECDSA := GetPrivateKeyECDSA()
	publicKeyECDSA := GetPublicKeyECDSA(privateKeyECDSA)

	hashContent := crypto.Keccak256(Str2Bytes(content))
	signature, err := crypto.Sign(hashContent, privateKeyECDSA)
	if err != nil {
		return nil, nil, err
	}
	return signature, publicKeyECDSA, nil
}

func VerifyECDSA(content string, signature []byte, publicKeyECDSA *ecdsa.PublicKey) (bool, error) {
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	hashContent := crypto.Keccak256(Str2Bytes(content))
	sigPublicKey, err := crypto.Ecrecover(hashContent, signature)

	if err != nil {
		return false, err
	}
	matched := bytes.Equal(sigPublicKey, publicKeyBytes)
	return matched, nil
}

func DecryptEcies(content []byte) (string, error) {
	privateKeyECDSA := GetPrivateKeyECDSA()
	privateKeyEcies := ecies.ImportECDSA(privateKeyECDSA)
	dect, err := privateKeyEcies.Decrypt(content, nil, nil)
	return Bytes2Str(dect), err
}

func EncryptEcies(content string, publicKeyECDSA *ecdsa.PublicKey) ([]byte, error) {
	publicKeyEcies := ecies.ImportECDSAPublic(publicKeyECDSA)
	ct, err := ecies.Encrypt(strings.NewReader(RandomHex(64)), publicKeyEcies, Str2Bytes(content), nil, nil)
	return ct, err
}
