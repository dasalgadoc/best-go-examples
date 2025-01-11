package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"io/ioutil"
	"log"
)

// Generar claves RSA y guardarlas en archivos
func generateRSAKeys() {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		log.Fatalf("Error generando clave RSA: %v", err)
	}

	privatePEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})
	ioutil.WriteFile("private.pem", privatePEM, 0600)

	publicASN1, err := x509.MarshalPKIXPublicKey(&privateKey.PublicKey)
	if err != nil {
		log.Fatalf("Error exportando clave pública: %v", err)
	}

	publicPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: publicASN1,
	})
	ioutil.WriteFile("public.pem", publicPEM, 0644)
}

// Cargar clave pública
func loadPublicKey(path string) *rsa.PublicKey {
	data, _ := ioutil.ReadFile(path)
	block, _ := pem.Decode(data)
	key, _ := x509.ParsePKIXPublicKey(block.Bytes)
	return key.(*rsa.PublicKey)
}

// Cargar clave privada
func loadPrivateKey(path string) *rsa.PrivateKey {
	data, _ := ioutil.ReadFile(path)
	block, _ := pem.Decode(data)
	key, _ := x509.ParsePKCS1PrivateKey(block.Bytes)
	return key
}

// Cifrar con RSA
func encryptRSA(publicKey *rsa.PublicKey, message string) string {
	ciphertext, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, []byte(message))
	if err != nil {
		log.Fatalf("Error cifrando con RSA: %v", err)
	}
	return base64.StdEncoding.EncodeToString(ciphertext)
}

// Descifrar con RSA
func decryptRSA(privateKey *rsa.PrivateKey, encryptedBase64 string) string {
	ciphertext, _ := base64.StdEncoding.DecodeString(encryptedBase64)
	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		log.Fatalf("Error descifrando con RSA: %v", err)
	}
	return string(plaintext)
}
