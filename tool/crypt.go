// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package neosqa is neoone's white box testing library.
package tool

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// DefaultEncryptKey neoone external use
const DefaultEncryptKey = "@@@@123456aa45888170bb654321@@@@"

var encryptKey = ""

func pkcs5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

// SetEncKey set the key used by NeoEnc NeoDec
func SetEncKey(key string) {
	encryptKey = key
}

func aesEnc(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blkSize := block.BlockSize()
	data = pkcs5Padding(data, blkSize)
	blkMode := cipher.NewCBCEncrypter(block, key[:blkSize])
	crypted := make([]byte, len(data))
	blkMode.CryptBlocks(crypted, data)
	return crypted, nil
}

func aesDec(data, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blkSize := block.BlockSize()
	blkMode := cipher.NewCBCDecrypter(block, key[:blkSize])
	decrypted := make([]byte, len(data))
	blkMode.CryptBlocks(decrypted, data)
	decrypted = pkcs5UnPadding(decrypted)
	return decrypted, nil
}

// NeoEnc encrypt order, data -> aes -> base64 -> result
func NeoEnc(data string) (string, error) {
	if encryptKey == "" {
		encryptKey = DefaultEncryptKey
	}
	key := []byte(encryptKey)
	encAes, err := aesEnc([]byte(data), key)
	if err != nil {
		return "", err
	}
	encBase64 := base64.StdEncoding.EncodeToString(encAes)
	return encBase64, nil
}

// NeoDec decrypt order, data -> base64 -> aes -> result
func NeoDec(data string) (string, error) {
	if encryptKey == "" {
		encryptKey = DefaultEncryptKey
	}
	key := []byte(encryptKey)
	decBase64, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", err
	}
	decAes, err := aesDec(decBase64, key)
	if err != nil {
		return "", err
	}
	return string(decAes), nil
}
