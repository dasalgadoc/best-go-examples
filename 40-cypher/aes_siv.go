package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
)

// encryptAESSIV vanilla implementation of AES SIV encryption
// AES-CTR (cypher) + AES-CMAC (authenticator)
func encryptAESSIV(key, plaintext []byte) (string, error) {
	// Paso 1: Generar el CMAC (IV Syhthetic)
	iv := generateCMAC(key, plaintext)

	// Paso 2: Cifrar con AES-CTR usando el IV generado
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCTR(block, iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	// Paso 3: Concatenar IV + Ciphertext
	finalBytes := append(iv, ciphertext...)

	// Paso 4: Convertir a Base64 para almacenarlo en la BD
	return base64.StdEncoding.EncodeToString(finalBytes), nil
}

func decryptAESSIV(key []byte, encryptedBase64 string) (string, error) {
	// Paso 1: Decodificar el mensaje cifrado
	encryptedData, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		return "", err
	}

	// Paso 2: Separar el IV del mensaje cifrado
	iv := []byte(encryptedData[:aes.BlockSize])
	ciphertext := []byte(encryptedData[aes.BlockSize:])

	// Paso 3: Descifrar con AES-CTR usando el IV
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	stream := cipher.NewCTR(block, iv)
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return string(plaintext), nil
}

func generateCMAC(key, plaintext []byte) []byte {
	mac := hmac.New(sha256.New, key)
	mac.Write(plaintext)

	return mac.Sum(nil)[:aes.BlockSize]
}
