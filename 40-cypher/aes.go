package main

import (
	"crypto/aes"
	"crypto/cipher"   // Para usar modos de operación como GCM (Galois/Counter Mode)
	"crypto/rand"     // Para generar números aleatorios seguros
	"encoding/base64" // Para codificar y decodificar en Base64
	"io"
	"io/ioutil"
	"log"
)

// Genera una clave AES de 32 bytes (AES-256)
func generateAESKey() []byte {
	key := make([]byte, 32)
	_, err := rand.Read(key) // Llena el slice con números aleatorios seguros
	if err != nil {
		log.Fatalf("Error generando clave AES: %v", err)
	}
	return key
}

// Guardar clave AES en archivo
func saveAESKeyToFile(key []byte, filename string) {
	encodedKey := base64.StdEncoding.EncodeToString(key)        // Codifica la clave en Base64
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
	block, err := aes.NewCipher(key) // Crea un nuevo cifrador AES
	if err != nil {
		log.Fatalf("Error creando el cipher AES: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block) // Crea un nuevo cifrador GCM para cifrado, usando el block cipher AES creado para cifrar bloques
	// Usar NewGCMWithNonceSize si se necesita un tamaño de nonce diferente a 12 bytes
	if err != nil {
		log.Fatalf("Error creando AES-GCM: %v", err)
	}

	nonce := make([]byte, aesGCM.NonceSize()) // crea un slice para el IV nonce con el tamaño requerido por el cifrador GCM, por lo general 12 bytes
	_, err = io.ReadFull(rand.Reader, nonce)  // Llena el slice NONCE con números aleatorios seguros
	if err != nil {
		log.Fatalf("Error generando nonce: %v", err)
	}

	ciphertext := aesGCM.Seal(nonce, nonce, plaintext, nil) // Cifra el mensaje con el cifrador GCM y con el nonce como IV
	return base64.StdEncoding.EncodeToString(ciphertext)    // Codifica el mensaje cifrado en Base64 nonce + ciphertext
}

// Descifrar con AES-GCM
func decryptAES(key []byte, encryptedBase64 string) string {
	ciphertext, err := base64.StdEncoding.DecodeString(encryptedBase64) // Decodifica el mensaje cifrado en Base64
	if err != nil {
		log.Fatalf("Error decodificando Base64: %v", err)
	}

	block, err := aes.NewCipher(key) // Crea un nuevo cifrador AES con la clave
	if err != nil {
		log.Fatalf("Error creando el cipher AES: %v", err)
	}

	aesGCM, err := cipher.NewGCM(block) // Crea un nuevo cifrador GCM para cifrado, usando el block cipher AES creado para cifrar bloques
	if err != nil {
		log.Fatalf("Error creando AES-GCM: %v", err)
	}

	nonceSize := aesGCM.NonceSize()                                     // Obtiene el tamaño del nonce del cifrador GCM (por lo general 12 bytes)
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:] // Separa el nonce del mensaje cifrado
	plaintext, err := aesGCM.Open(nil, nonce, ciphertext, nil)          // Descifra el mensaje con el cifrador GCM y con el nonce como IV
	if err != nil {
		log.Fatalf("Error descifrando: %v", err)
	}
	return string(plaintext)
}
