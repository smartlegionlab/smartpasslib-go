/**
 * SmartPassLib v4.0.0 - Go smart password generator
 * Cross-platform deterministic password generation
 * Same secret + same length = same password across all platforms
 * Decentralized by design — no central servers, no cloud dependency, no third-party trust required
 *
 * Compatible with smartpasslib Python/JS/Kotlin/C# implementations
 *
 * Key derivation:
 * - Private key: 15-30 iterations (dynamic, deterministic per secret)
 * - Public key: 45-60 iterations (dynamic, deterministic per secret)
 *
 * Secret phrase:
 *   - is not transferred anywhere
 *   - is not stored anywhere
 *   - is required to generate the private key when creating a smart password
 *   - minimum 12 characters (enforced)
 *
 * Password length:
 *   - minimum 12 characters (enforced)
 *   - maximum 100 characters (enforced)
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
    "strconv"
)

const Version = "4.0.0"

const Chars = "!@#$%^&*()_+-=[]{};:,.<>?/ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz"

func sha256Hash(text string) string {
    hash := sha256.Sum256([]byte(text))
    return hex.EncodeToString(hash[:])
}

func validateSecret(secret string) error {
    if len(secret) < 12 {
        return fmt.Errorf("secret phrase must be at least 12 characters. Current: %d", len(secret))
    }
    return nil
}

func validatePasswordLength(length int) error {
    if length < 12 {
        return fmt.Errorf("password length must be at least 12 characters. Current: %d", length)
    }
    if length > 100 {
        return fmt.Errorf("password length cannot exceed 100 characters. Current: %d", length)
    }
    return nil
}

func validateCodeLength(length int) error {
    if length < 4 {
        return fmt.Errorf("code length must be at least 4 characters. Current: %d", length)
    }
    if length > 100 {
        return fmt.Errorf("code length cannot exceed 100 characters. Current: %d", length)
    }
    return nil
}

func getStepsFromSecret(secret string, minSteps, maxSteps int, salt string) (int, error) {
    hashValue := sha256Hash(fmt.Sprintf("%s:%s", secret, salt))
    if len(hashValue) < 8 {
        return 0, fmt.Errorf("hash too short")
    }
    hashInt, err := strconv.ParseInt(hashValue[:8], 16, 64)
    if err != nil {
        return 0, err
    }
    return minSteps + int(hashInt)%(maxSteps-minSteps+1), nil
}

func generateKey(secret string, steps int, salt string) (string, error) {
    if err := validateSecret(secret); err != nil {
        return "", err
    }
    allHash := sha256Hash(fmt.Sprintf("%s:%s", secret, salt))
    for i := 0; i < steps; i++ {
        allHash = sha256Hash(fmt.Sprintf("%s:%d", allHash, i))
    }
    return allHash, nil
}

func GeneratePrivateKey(secret string) (string, error) {
    if err := validateSecret(secret); err != nil {
        return "", err
    }
    steps, err := getStepsFromSecret(secret, 15, 30, "private")
    if err != nil {
        return "", err
    }
    return generateKey(secret, steps, "private")
}

func GeneratePublicKey(secret string) (string, error) {
    if err := validateSecret(secret); err != nil {
        return "", err
    }
    steps, err := getStepsFromSecret(secret, 45, 60, "public")
    if err != nil {
        return "", err
    }
    return generateKey(secret, steps, "public")
}

func VerifySecret(secret, publicKey string) (bool, error) {
    computedKey, err := GeneratePublicKey(secret)
    if err != nil {
        return false, err
    }
    return computedKey == publicKey, nil
}

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

func generatePasswordFromPrivateKey(privateKey string, length int) (string, error) {
    if err := validatePasswordLength(length); err != nil {
        return "", err
    }
    result := make([]byte, 0, length)
    counter := 0
    for len(result) < length {
        hashHex := sha256Hash(fmt.Sprintf("%s:%d", privateKey, counter))
        hashBytes, err := hexToBytes(hashHex)
        if err != nil {
            return "", err
        }
        for _, b := range hashBytes {
            if len(result) < length {
                idx := int(b) & 0xFF
                result = append(result, Chars[idx%len(Chars)])
            }
        }
        counter++
    }
    return string(result), nil
}

func GenerateSmartPassword(secret string, length int) (string, error) {
    if err := validateSecret(secret); err != nil {
        return "", err
    }
    if err := validatePasswordLength(length); err != nil {
        return "", err
    }
    privateKey, err := GeneratePrivateKey(secret)
    if err != nil {
        return "", err
    }
    return generatePasswordFromPrivateKey(privateKey, length)
}

func GenerateStrongPassword(length int) (string, error) {
    if err := validatePasswordLength(length); err != nil {
        return "", err
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

func GenerateBasePassword(length int) (string, error) {
    return GenerateStrongPassword(length)
}

func GenerateCode(length int) (string, error) {
    if err := validateCodeLength(length); err != nil {
        return "", err
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