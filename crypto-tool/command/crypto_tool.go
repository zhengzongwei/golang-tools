/*
 * Copyright (c)2025. zhengzongwei
 * golang-tools is licensed under Mulan PSL v2.
 * You can use this software according to the terms and conditions of the Mulan PSL v2.
 * You may obtain a copy of Mulan PSL v2 at:
 *         http://license.coscl.org.cn/MulanPSL2
 * THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND,
 * EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT,
 * MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
 * See the Mulan PSL v2 for more details.
 *
 */

package command

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/pbkdf2"
)

const (
	saltSize   = 16
	iterations = 100000
	keySize    = 32
	nonceSize  = 12
)

type CryptoTool struct {
	salt []byte
}

func NewCryptoTool(salt string) *CryptoTool {
	if salt == "" {
		salt = generateMachineCode()
	}

	return &CryptoTool{
		salt: []byte(salt),
	}
}

func generateMachineCode() string {
	hardwareInfo := NewHardwareInfo()
	return hardwareInfo.GenerateMachineCode()

}

func (c *CryptoTool) deriveKey(salt []byte) []byte {
	return pbkdf2.Key(c.salt, salt, iterations, keySize, sha256.New)
}

func (c *CryptoTool) encrypt(data string) (string, error) {
	salt := make([]byte, saltSize)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	key := c.deriveKey(salt)
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, nonceSize)
	if _, err := rand.Read(nonce); err != nil {
		return "", err
	}

	encryptedData := aead.Seal(nil, nonce, []byte(data), nil)
	finalData := append(salt, nonce...)
	finalData = append(finalData, encryptedData...)

	return base64.URLEncoding.EncodeToString(finalData), nil
}

func (c *CryptoTool) decrypt(encodedToken string) (string, error) {
	token, err := base64.URLEncoding.DecodeString(encodedToken)
	if err != nil {
		return "", err
	}

	salt := token[:saltSize]
	nonce := token[saltSize : saltSize+nonceSize]
	encryptedData := token[saltSize+nonceSize:]

	key := c.deriveKey(salt)
	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return "", err
	}

	decryptedData, err := aead.Open(nil, nonce, encryptedData, nil)
	if err != nil {
		return "", err
	}

	return string(decryptedData), nil
}
