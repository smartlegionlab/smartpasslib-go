/**
 * SmartPassLib v1.0.2 - Go smart password generator
 * Cross-platform deterministic password generation
 * Same secret + same length = same password across all platforms
 * Decentralized by design — no central servers, no cloud dependency, no third-party trust required
 *
 * Compatible with smartpasslib Python/JS/Kotlin/C# implementations
 *
 * Key derivation:
 * - Private key: 30 iterations of SHA-256 (used for password generation)
 * - Public key: 60 iterations of SHA-256 (used for verification, stored locally)
 *
 * Secret phrase:
 *   - is not transferred anywhere
 *   - is not stored anywhere
 *   - is required to generate the private key when creating a smart password
 *
 * Ecosystem:
 *   - Core library (Python): https://github.com/smartlegionlab/smartpasslib
 *   - Core library (JS): https://github.com/smartlegionlab/smartpasslib-js
 *   - Core library (Kotlin): https://github.com/smartlegionlab/smartpasslib-kotlin
 *   - Core library (Go): https://github.com/smartlegionlab/smartpasslib-go
 *   - Core library (C#): https://github.com/smartlegionlab/smartpasslib-csharp
 *   - Desktop Python: https://github.com/smartlegionlab/smart-password-manager-desktop
 *   - Desktop C#: https://github.com/smartlegionlab/SmartPasswordManagerCsharpDesktop
 *   - CLI Manager Python: https://github.com/smartlegionlab/clipassman
 *   - CLI Manager C#: https://github.com/smartlegionlab/SmartPasswordManagerCsharpCli
 *   - CLI Generator Python: https://github.com/smartlegionlab/clipassgen
 *   - CLI Generator C#: https://github.com/smartlegionlab/SmartPasswordGeneratorCsharpCli
 *   - Web: https://github.com/smartlegionlab/smart-password-manager-web
 *   - Android: https://github.com/smartlegionlab/smart-password-manager-android
 *
 * Author: Alexander Suvorov https://github.com/smartlegionlab
 * License: BSD 3-Clause https://github.com/smartlegionlab/smartpasslib-go/blob/master/LICENSE
 * Copyright (c) 2026, Alexander Suvorov. All rights reserved.
 */

package smartpasslib

import (
    "crypto/rand"
    "crypto/sha256"
    "encoding/hex"
    "fmt"
)

// Version
const Version = "1.0.2"

// Character set for password generation (must match other implementations)
const Chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$&*-_"

// Iteration counts
const (
    privateIterations = 30 // For private key (password generation)
    publicIterations  = 60 // For public key (verification, stored on server)
)

// sha256 hash function
func sha256Hash(text string) string {
    hash := sha256.Sum256([]byte(text))
    return hex.EncodeToString(hash[:])
}

// generateKey generates a key from secret phrase with specified number of iterations
func generateKey(secret string, iterations int) (string, error) {
    if len(secret) < 12 {
        return "", fmt.Errorf("secret phrase must be at least 12 characters. Current: %d", len(secret))
    }

    allHash := sha256Hash(secret)

    for i := 0; i < iterations; i++ {
        tempString := fmt.Sprintf("%s:%s:%d", allHash, secret, i)
        allHash = sha256Hash(tempString)
    }

    return allHash, nil
}

// GeneratePrivateKey generates private key from secret phrase (30 iterations)
// Used for password generation, never stored or transmitted
func GeneratePrivateKey(secret string) (string, error) {
    return generateKey(secret, privateIterations)
}

// GeneratePublicKey generates public key from secret phrase (60 iterations)
// Used for verification, stored on server
func GeneratePublicKey(secret string) (string, error) {
    return generateKey(secret, publicIterations)
}

// VerifySecret verifies that a secret phrase matches a stored public key
func VerifySecret(secret, publicKey string) (bool, error) {
    computedKey, err := GeneratePublicKey(secret)
    if err != nil {
        return false, err
    }
    return computedKey == publicKey, nil
}

// hexToBytes converts hex string to byte slice
func hexToBytes(hexStr string) ([]byte, error) {
    bytes := make([]byte, len(hexStr)/2)
    for i := 0; i < len(hexStr); i += 2 {
        b, err := hex.DecodeString(hexStr[i : i+2])
        if err != nil {
            return nil, err
        }
        bytes[i/2] = b[0]
    }
    return bytes, nil
}

// generatePasswordFromPrivateKey generates deterministic smart password from private key
func generatePasswordFromPrivateKey(privateKey string, length int) (string, error) {
    if length < 12 || length > 1000 {
        return "", fmt.Errorf("password length must be between 12 and 1000. Current: %d", length)
    }

    result := make([]byte, 0, length)
    counter := 0

    for len(result) < length {
        data := fmt.Sprintf("%s:%d", privateKey, counter)
        hashHex := sha256Hash(data)
        hashBytes, err := hexToBytes(hashHex)
        if err != nil {
            return "", err
        }

        for _, b := range hashBytes {
            if len(result) < length {
                idx := int(b) & 0xFF
                result = append(result, Chars[idx%len(Chars)])
            } else {
                break
            }
        }
        counter++
    }

    return string(result), nil
}

// GenerateSmartPassword generates deterministic smart password directly from secret phrase
func GenerateSmartPassword(secret string, length int) (string, error) {
    if len(secret) < 12 {
        return "", fmt.Errorf("secret phrase must be at least 12 characters. Current: %d", len(secret))
    }
    if length < 12 || length > 1000 {
        return "", fmt.Errorf("password length must be between 12 and 1000. Current: %d", length)
    }

    privateKey, err := GeneratePrivateKey(secret)
    if err != nil {
        return "", err
    }
    return generatePasswordFromPrivateKey(privateKey, length)
}

// GenerateStrongPassword generates cryptographically secure random password
func GenerateStrongPassword(length int) (string, error) {
    if length < 12 || length > 1000 {
        return "", fmt.Errorf("password length must be between 12 and 1000. Current: %d", length)
    }

    bytes := make([]byte, length)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }

    result := make([]byte, length)
    for i, b := range bytes {
        result[i] = Chars[int(b)%len(Chars)]
    }
    return string(result), nil
}

// GenerateBasePassword generates base random password
func GenerateBasePassword(length int) (string, error) {
    return GenerateStrongPassword(length)
}

// GenerateCode generates authentication code (shorter, for 2FA)
func GenerateCode(length int) (string, error) {
    if length < 4 || length > 20 {
        return "", fmt.Errorf("code length must be between 4 and 20. Current: %d", length)
    }

    bytes := make([]byte, length)
    if _, err := rand.Read(bytes); err != nil {
        return "", err
    }

    result := make([]byte, length)
    for i, b := range bytes {
        result[i] = Chars[int(b)%len(Chars)]
    }
    return string(result), nil
}