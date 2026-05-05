package smartpasslib

import (
    "strings"
    "testing"
)

func TestAll(t *testing.T) {
    println()
    println(strings.Repeat("=", 60))
    println("SmartPassLib Go v4.0.0 - Generator Test")
    println(strings.Repeat("=", 60))
    println()

    secret1 := "MyCatHippo2026"
    lengths := []int{12, 16, 20, 24}

    println("Secret phrase 1:", secret1)
    println()

    for _, length := range lengths {
        password, err := GenerateSmartPassword(secret1, length)
        if err != nil {
            t.Fatalf("Error: %v", err)
        }
        println("Length", length, ":", password)
    }

    println()
    println("Public/Private keys for secret 1:")
    publicKey1, _ := GeneratePublicKey(secret1)
    privateKey1, _ := GeneratePrivateKey(secret1)
    println("Public key: ", publicKey1)
    println("Private key:", privateKey1)
    if publicKey1 != privateKey1 {
        println("Keys are different: YES")
    } else {
        println("Keys are different: NO")
    }

    println()
    println(strings.Repeat("-", 60))

    secret2 := "TestSecret2026!"
    println()
    println("Secret phrase 2:", secret2)
    println()

    passwordA, _ := GenerateSmartPassword(secret2, 16)
    passwordB, _ := GenerateSmartPassword(secret2, 16)

    println("Determinism test:")
    if passwordA == passwordB {
        println("Same secret + same length = SAME")
    } else {
        println("Same secret + same length = DIFFERENT")
    }
    println("Password:", passwordA)

    println()
    println("Public/Private keys for secret 2:")
    publicKey2, _ := GeneratePublicKey(secret2)
    privateKey2, _ := GeneratePrivateKey(secret2)
    println("Public key: ", publicKey2)
    println("Private key:", privateKey2)
    if publicKey2 != privateKey2 {
        println("Keys are different: YES")
    } else {
        println("Keys are different: NO")
    }

    println()
    println(strings.Repeat("=", 60))
    println("Test complete")
    println(strings.Repeat("=", 60))
}