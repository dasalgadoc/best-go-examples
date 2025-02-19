package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"testing"
	"time"
)

type scenario struct {
	cipherText    string
	decryptedText string
	timeElapsed   int64
}

func TestAESGCMCipher(t *testing.T) {
	text := generateRandomPlainText()
	actualScenarios := [100]scenario{}
	key := generateAESKey()

	for i := 0; i < 100; i++ {
		start := time.Now()
		encrypted := encryptAES(key, []byte(text))

		if len(encrypted) <= len(text) {
			t.Fatalf("Ciphertext length is too short: expected > %d, got %d", len(text), len(encrypted))
		}

		actualScenarios[i] = scenario{
			cipherText:    encrypted,
			decryptedText: text,
			timeElapsed:   time.Since(start).Nanoseconds(),
		}
	}

	fmt.Println("AES-GCM")
	t.Logf("Text: %s\n", text)
	t.Logf("Key: %v\n", key)
	timeElapsedSum := int64(0)
	for i := 0; i < 100; i++ {
		t.Logf("Scenario %d: %v\n", i, actualScenarios[i])
		timeElapsedSum += actualScenarios[i].timeElapsed
	}
	t.Logf("Time elapsed: %v nanoseconds\n", timeElapsedSum)
	t.Logf("Average time elapsed: %v nanoseconds\n", timeElapsedSum/100)
}

func TestAESSIVCipher(t *testing.T) {
	text := generateRandomPlainText()
	actualScenarios := [100]scenario{}
	key := generateAESKey()

	for i := 0; i < 100; i++ {
		start := time.Now()
		encrypted, err := encryptAESSIV(key, []byte(text))
		if err != nil {
			t.Fatalf("Error encrypting text: %v", err)
		}
		if len(encrypted) == 0 {
			t.Fatalf("Encrypted text is empty for iteration %d", i)
		}
		if len(encrypted) <= len(text) {
			t.Fatalf("Ciphertext length is too short: expected > %d, got %d", len(text), len(encrypted))
		}

		actualScenarios[i] = scenario{
			cipherText:    encrypted,
			decryptedText: text,
			timeElapsed:   time.Since(start).Nanoseconds(),
		}
	}

	cipherTextPivot := actualScenarios[0].cipherText
	for i := 0; i < 100; i++ {
		if actualScenarios[i].cipherText != cipherTextPivot {
			t.Fatalf("Encryption failed: expected %s, got %s", cipherTextPivot, actualScenarios[i].decryptedText)
		}
	}

	fmt.Println("AES-SIV")
	t.Logf("Text: %s\n", text)
	t.Logf("Key: %v\n", key)
	timeElapsedSum := int64(0)
	for i := 0; i < 100; i++ {
		t.Logf("Scenario %d: %v\n", i, actualScenarios[i])
		timeElapsedSum += actualScenarios[i].timeElapsed
	}
	t.Logf("Time elapsed: %v nanoseconds\n", timeElapsedSum)
	t.Logf("Average time elapsed: %v nanoseconds\n", timeElapsedSum/100)
}

func TestAESGCMDecrypt(t *testing.T) {
	text := generateRandomPlainText()
	key := generateAESKey()
	actualScenarios := [100]scenario{}

	for i := 0; i < 100; i++ {
		encrypted := encryptAES(key, []byte(text))

		if len(encrypted) == 0 {
			t.Fatalf("Encrypted text is empty for iteration %d", i)
		}
		if len(encrypted) <= len(text) {
			t.Fatalf("Ciphertext length is too short: expected > %d, got %d", len(text), len(encrypted))
		}

		actualScenarios[i] = scenario{
			cipherText: encrypted,
		}
	}

	for i := 0; i < 100; i++ {
		start := time.Now()
		decrypt := decryptAES(key, actualScenarios[i].cipherText)
		actualScenarios[i].timeElapsed = time.Since(start).Nanoseconds()
		actualScenarios[i].decryptedText = decrypt
	}

	decryptPivot := actualScenarios[0].decryptedText

	for i := 0; i < 100; i++ {
		if actualScenarios[i].decryptedText != decryptPivot {
			t.Fatalf("Decryption failed: expected %s, got %s", decryptPivot, actualScenarios[i].decryptedText)
		}
	}

	fmt.Println("AES-GCM")
	t.Logf("Text: %s\n", text)
	t.Logf("Key: %v\n", key)
	timeElapsedSum := int64(0)
	for i := 0; i < 100; i++ {
		t.Logf("Scenario %d: %v\n", i, actualScenarios[i])
		timeElapsedSum += actualScenarios[i].timeElapsed
	}
	t.Logf("Time elapsed: %v nanoseconds\n", timeElapsedSum)
	t.Logf("Average time elapsed: %v nanoseconds\n", timeElapsedSum/100)
}

func TestAESSIVDecrypt(t *testing.T) {
	text := generateRandomPlainText()
	key := generateAESKey()
	actualScenarios := [100]scenario{}

	for i := 0; i < 100; i++ {
		encrypted, err := encryptAESSIV(key, []byte(text))
		if err != nil {
			t.Fatalf("Error encrypting text: %v", err)
		}
		if len(encrypted) == 0 {
			t.Fatalf("Encrypted text is empty for iteration %d", i)
		}
		if len(encrypted) <= len(text) {
			t.Fatalf("Ciphertext length is too short: expected > %d, got %d", len(text), len(encrypted))
		}

		actualScenarios[i] = scenario{
			cipherText: encrypted,
		}
	}

	for i := 0; i < 100; i++ {
		start := time.Now()
		decrypt, err := decryptAESSIV(key, actualScenarios[i].cipherText)
		if err != nil {
			t.Fatalf("Error encrypting text: %v", err)
		}

		actualScenarios[i].timeElapsed = time.Since(start).Nanoseconds()
		actualScenarios[i].decryptedText = decrypt
	}

	decryptPivot := actualScenarios[0].decryptedText

	for i := 0; i < 100; i++ {
		if actualScenarios[i].decryptedText != decryptPivot {
			t.Fatalf("Decryption failed: expected %s, got %s", decryptPivot, actualScenarios[i].decryptedText)
		}
	}

	fmt.Println("AES-SIV")
	t.Logf("Text: %s\n", text)
	t.Logf("Key: %v\n", key)
	timeElapsedSum := int64(0)
	for i := 0; i < 100; i++ {
		t.Logf("Scenario %d: %v\n", i, actualScenarios[i])
		timeElapsedSum += actualScenarios[i].timeElapsed
	}
	t.Logf("Time elapsed: %v nanoseconds\n", timeElapsedSum)
	t.Logf("Average time elapsed: %v nanoseconds\n", timeElapsedSum/100)
}

func TestAESSIVIdempotence(t *testing.T) {
	text := generateRandomPlainText()
	key := generateAESKey()

	encrypted1, err1 := encryptAESSIV(key, []byte(text))
	encrypted2, err2 := encryptAESSIV(key, []byte(text))

	if err1 != nil || err2 != nil {
		t.Fatalf("Encryption failed: err1=%v, err2=%v", err1, err2)
	}

	if len(encrypted1) == 0 || len(encrypted2) == 0 {
		t.Fatalf("Encrypted text is empty")
	}

	if encrypted1 != encrypted2 {
		t.Fatalf("AES-SIV should be deterministic, but results differ: %s vs %s", encrypted1, encrypted2)
	}
}

func generateRandomPlainText() string {
	rand.Seed(time.Now().UnixNano()) // Inicializa la semilla para números aleatorios únicos

	var sb strings.Builder
	for i := 0; i < 10; i++ {
		sb.WriteString(strconv.Itoa(rand.Intn(10))) // rand.Intn(10) genera un número entre 0 y 9
	}

	return sb.String()
}
