package main

import "fmt"

func main() {
	// ===== AES =====
	fmt.Println("===== AES-GCM (Simétrico) =====")
	keyAES := generateAESKey()
	message := "Mensaje secreto con AES 🔐"
	encryptedAES := encryptAES(keyAES, []byte(message))
	fmt.Println("🔒 Cifrado AES:", encryptedAES)
	fmt.Println("🔓 Descifrado AES:", decryptAES(keyAES, encryptedAES))

	// ===== RSA =====
	fmt.Println("\n===== RSA (Asimétrico) =====")
	generateRSAKeys()
	publicKey := loadPublicKey("public.pem")
	privateKey := loadPrivateKey("private.pem")

	encryptedRSA := encryptRSA(publicKey, message)
	fmt.Println("🔒 Cifrado RSA:", encryptedRSA)
	fmt.Println("🔓 Descifrado RSA:", decryptRSA(privateKey, encryptedRSA))
}
