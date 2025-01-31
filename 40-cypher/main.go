package main

import "fmt"

func main() {
	// ===== AES =====
	fmt.Println("===== AES-GCM (Simétrico) =====")
	keyAES := generateAESKey()

	message := "Mensaje secreto con AES 🔐"
	fmt.Println("Is Encrypted:", isEncrypted(keyAES, message))

	encryptedAES := encryptAES(keyAES, []byte(message))
	fmt.Println("🔒 Cifrado AES:", encryptedAES)
	fmt.Println("Is Encrypted:", isEncrypted(keyAES, encryptedAES))

	decrypt := decryptAES(keyAES, encryptedAES)
	fmt.Println("🔓 Descifrado AES:")
	fmt.Println("Is Encrypted:", isEncrypted(keyAES, decrypt))

	// ===== RSA =====
	fmt.Println("\n===== RSA (Asimétrico) =====")
	generateRSAKeys()
	publicKey := loadPublicKey("public.pem")
	privateKey := loadPrivateKey("private.pem")

	encryptedRSA := encryptRSA(publicKey, message)
	fmt.Println("🔒 Cifrado RSA:", encryptedRSA)
	fmt.Println("🔓 Descifrado RSA:", decryptRSA(privateKey, encryptedRSA))
}
