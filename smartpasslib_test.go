package smartpasslib

import (
    "strings"
    "testing"
)

func TestVersion(t *testing.T) {
    if Version != "1.0.0" {
        t.Errorf("Version mismatch: got %s, want 1.0.0", Version)
    }
}

func TestGenerateSmartPassword(t *testing.T) {
    secret := "MyCatHippo2026"
    password, err := GenerateSmartPassword(secret, 16)
    if err != nil {
        t.Fatalf("Error: %v", err)
    }
    if len(password) != 16 {
        t.Errorf("Password length mismatch: got %d, want 16", len(password))
    }
    t.Logf("✅ Smart password: %s", password)
}

func TestDeterminism(t *testing.T) {
    secret := "TestSecret123!@#"
    pwd1, _ := GenerateSmartPassword(secret, 20)
    pwd2, _ := GenerateSmartPassword(secret, 20)
    if pwd1 != pwd2 {
        t.Errorf("Determinism failed: %s != %s", pwd1, pwd2)
    }
    t.Logf("✅ Determinism OK: %s", pwd1)
}

func TestDifferentSecrets(t *testing.T) {
    pwd1, _ := GenerateSmartPassword("SecretOne123456", 16)
    pwd2, _ := GenerateSmartPassword("SecretTwo123456", 16)
    if pwd1 == pwd2 {
        t.Errorf("Different secrets produced same password")
    }
    t.Logf("✅ Secret A: %s", pwd1)
    t.Logf("✅ Secret B: %s", pwd2)
}

func TestDifferentLengths(t *testing.T) {
    secret := "LengthTestSecret!@#"
    pwd12, _ := GenerateSmartPassword(secret, 12)
    pwd24, _ := GenerateSmartPassword(secret, 24)
    if len(pwd12) != 12 || len(pwd24) != 24 {
        t.Errorf("Length mismatch: got %d and %d", len(pwd12), len(pwd24))
    }
    t.Logf("✅ Length 12: %s", pwd12)
    t.Logf("✅ Length 24: %s", pwd24)
}

func TestPublicPrivateKeys(t *testing.T) {
    secret := "KeyTestSecret!@#"
    pubKey, _ := GeneratePublicKey(secret)
    privKey, _ := GeneratePrivateKey(secret)

    if pubKey == privKey {
        t.Errorf("Public and private keys are the same")
    }
    t.Logf("✅ Public key: %s...", pubKey[:32])
    t.Logf("✅ Private key: %s...", privKey[:32])

    isValid, _ := VerifySecret(secret, pubKey)
    if !isValid {
        t.Errorf("Verification failed for correct secret")
    }
    t.Logf("✅ Verification: %v", isValid)
}

func TestStrongRandomPassword(t *testing.T) {
    pwd, err := GenerateStrongPassword(16)
    if err != nil {
        t.Fatalf("Error: %v", err)
    }
    if len(pwd) != 16 {
        t.Errorf("Length mismatch: got %d, want 16", len(pwd))
    }
    t.Logf("✅ Strong random: %s", pwd)
}

func TestBaseRandomPassword(t *testing.T) {
    pwd, err := GenerateBasePassword(16)
    if err != nil {
        t.Fatalf("Error: %v", err)
    }
    if len(pwd) != 16 {
        t.Errorf("Length mismatch: got %d, want 16", len(pwd))
    }
    t.Logf("✅ Base random: %s", pwd)
}

func TestCodeGeneration(t *testing.T) {
    code, err := GenerateCode(8)
    if err != nil {
        t.Fatalf("Error: %v", err)
    }
    if len(code) != 8 {
        t.Errorf("Code length mismatch: got %d, want 8", len(code))
    }
    t.Logf("✅ Auth code: %s", code)
}

func TestSecretMinLength(t *testing.T) {
    _, err := GenerateSmartPassword("short", 16)
    if err == nil {
        t.Errorf("Should reject secret shorter than 12 chars")
    }
    t.Logf("✅ Secret too short rejected: %v", err)
}

func TestCodeMinLength(t *testing.T) {
    _, err := GenerateCode(2)
    if err == nil {
        t.Errorf("Should reject code shorter than 4 chars")
    }
    t.Logf("✅ Code too short rejected: %v", err)
}

func TestPasswordMinLength(t *testing.T) {
    _, err := GenerateStrongPassword(5)
    if err == nil {
        t.Errorf("Should reject password shorter than 12 chars")
    }
    t.Logf("✅ Password too short rejected: %v", err)
}

func TestAllowedChars(t *testing.T) {
    secret := "MyCatHippo2026"
    password, _ := GenerateSmartPassword(secret, 16)
    for _, c := range password {
        if !strings.ContainsRune(Chars, c) {
            t.Errorf("Invalid character in password: %c", c)
        }
    }
    t.Logf("✅ All characters allowed")
}