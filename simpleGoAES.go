package simpleGoAES

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5" // I use md5 on the password as an stupid way to ensure that AES gets the 32 char password it needs
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"io"
	"log"
)

//IsEncryptionOn : A simple system for turning on and off encryption in case you need to see the plain text results in a database for testing (default is On)
var IsEncryptionOn = true

//Encrypt : AES256 encryption function to work with strings
//returns a base64 encoded string safe to write/save most anywhere
//(depends on EncryptByteArray)
func Encrypt(key, stringToEncrypt string) (base64String string, err error) {
	if IsEncryptionOn == false {
		return stringToEncrypt, nil
	}
	key = md5TheKey(key)
	keyBytes := []byte(key)
	stringToEncryptBytes := []byte(stringToEncrypt)
	encryptedByteArray, err := EncryptByteArray(keyBytes, stringToEncryptBytes)
	encryptedString := base64.StdEncoding.EncodeToString(encryptedByteArray)
	return encryptedString, err
}

//Decrypt : AES256 decryption function to work with strings
//(depends on DecryptByteArray)
func Decrypt(key, base64StringToDecrypt string) (decodedString string, err error) {
	if IsEncryptionOn == false {
		return base64StringToDecrypt, nil
	}
	key = md5TheKey(key)
	keyBytes := []byte(key)
	stringToDecryptBytes, err := base64.StdEncoding.DecodeString(base64StringToDecrypt)
	if err != nil {
		log.Fatal(err)
		return "", err
	}
	decryptedByteArray, err := DecryptByteArray(keyBytes, stringToDecryptBytes)
	decryptedString := string(decryptedByteArray[:])
	return decryptedString, err
}

//EncryptByteArray : AES256 encryption function to work with byte arrays
//This expects the key as an array with 32 bytes in it
func EncryptByteArray(key, byteArrayToEncrypt []byte) ([]byte, error) {
	if IsEncryptionOn {
		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}
		ciphertext := make([]byte, aes.BlockSize+len(byteArrayToEncrypt))
		iv := ciphertext[:aes.BlockSize]
		if _, err := io.ReadFull(rand.Reader, iv); err != nil {
			return nil, err
		}
		cfb := cipher.NewCFBEncrypter(block, iv)
		cfb.XORKeyStream(ciphertext[aes.BlockSize:], byteArrayToEncrypt)
		return ciphertext, nil
	}
	// Encryption was off, just return the string
	return byteArrayToEncrypt, nil
}

//DecryptByteArray : AES256 encryption function to work with byte arrays
//This expects the key as an array with 32 bytes in it
func DecryptByteArray(key, byteArrayToDecrypt []byte) ([]byte, error) {
	if IsEncryptionOn {
		block, err := aes.NewCipher(key)
		if err != nil {
			return nil, err
		}
		if len(byteArrayToDecrypt) < aes.BlockSize {
			return nil, errors.New("ciphertext too short")
		}
		iv := byteArrayToDecrypt[:aes.BlockSize]
		byteArrayToDecrypt = byteArrayToDecrypt[aes.BlockSize:]
		cfb := cipher.NewCFBDecrypter(block, iv)
		cfb.XORKeyStream(byteArrayToDecrypt, byteArrayToDecrypt)
		return byteArrayToDecrypt, nil
	}
	return byteArrayToDecrypt, nil
}

// md5 on the passkey
func md5TheKey(key string) string {
	m := md5.New()
	m.Write([]byte(key))
	return hex.EncodeToString(m.Sum(nil))
}
