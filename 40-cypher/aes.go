package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
)

// Genera una clave AES de 32 bytes (AES-256)
func generateAESKey() []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatalf("Error generando clave AES: %v", err)
	}
	return key
}

// Guardar clave AES en archivo
func saveAESKeyToFile(key []byte, filename string) {
	encodedKey := base64.StdEncoding.EncodeToString(key)
	err := ioutil.WriteFile(filename, []byte(encodedKey), 0600) // ⚠️ Permisos seguros
	if err != nil {
		log.Fatalf("Error guardando la clave AES: %v", err)
	}
}

// Cargar clave AES desde archivo
func loadAESKeyFromFile(filename string) []byte {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Error leyendo la clave AES: %v", err)
	}

	key, err := base64.StdEncoding.DecodeString(string(data))
	if err != nil {
		log.Fatalf("Error decodificando la clave AES: %v", err)
	}

	return key
}

// Cifrar con AES-GCM
func encryptAES(key, plaintext []byte) string {
	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Error creando el cipher AES: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("Error creando AES-GCM: %v", err)
	}

	nonce := make([]byte, aesGCM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		log.Fatalf("Error generando nonce: %v", err)
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil)
	return base64.StdEncoding.EncodeToString(ciphertext)
}

// Descifrar con AES-GCM
func decryptAES(key []byte, encryptedBase64 string) string {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedBase64)
	if err != nil {
		log.Fatalf("Error decodificando Base64: %v", err)
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		log.Fatalf("Error creando el cipher AES: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		log.Fatalf("Error creando AES-GCM: %v", err)
	}

	nonceSize := aesGCM.NonceSize()
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		log.Fatalf("Error descifrando: %v", err)
	}
	return string(plaintext)
}
