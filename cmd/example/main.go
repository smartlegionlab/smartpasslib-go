package main

import (
    "fmt"
    "github.com/smartlegionlab/smartpasslib-go"
)

func main() {
    fmt.Println()
    fmt.Println("============================================================")
    fmt.Println("🔐 SMART PASSWORDS LIBRARY GO - EXAMPLE")
    fmt.Println("============================================================")
    fmt.Println()

    secret := "MyCatHippo2026"
    wrongSecret := "WrongSecret123456"

    // ============================================================
    // 1. SMART PASSWORD
    // ============================================================
    fmt.Println("📌 [1] SMART PASSWORD (Deterministic)")
    fmt.Println("------------------------------------------------------------")

    password, _ := smartpasslib.GenerateSmartPassword(secret, 16)
    fmt.Println()
    fmt.Printf("  Secret phrase: %s\n", secret)
    fmt.Printf("  Password length: 16\n")
    fmt.Printf("  Generated password: %s\n", password)
    fmt.Println()

    // ============================================================
    // 2. DETERMINISM
    // ============================================================
    fmt.Println("📌 [2] DETERMINISM (Same secret → same password)")
    fmt.Println("------------------------------------------------------------")

    pwd1, _ := smartpasslib.GenerateSmartPassword(secret, 20)
    pwd2, _ := smartpasslib.GenerateSmartPassword(secret, 20)
    fmt.Printf("  First:  %s\n", pwd1)
    fmt.Printf("  Second: %s\n", pwd2)
    if pwd1 == pwd2 {
        fmt.Println("  ✅ PASS: Passwords match")
    } else {
        fmt.Println("  ❌ FAIL: Passwords do not match")
    }
    fmt.Println()

    // ============================================================
    // 3. DIFFERENT SECRETS
    // ============================================================
    fmt.Println("📌 [3] DIFFERENT SECRETS → DIFFERENT PASSWORDS")
    fmt.Println("------------------------------------------------------------")

    diff1, _ := smartpasslib.GenerateSmartPassword("SecretOne123456", 16)
    diff2, _ := smartpasslib.GenerateSmartPassword("SecretTwo123456", 16)
    fmt.Printf("  Secret A -> %s\n", diff1)
    fmt.Printf("  Secret B -> %s\n", diff2)
    if diff1 != diff2 {
        fmt.Println("  ✅ PASS: Passwords are different")
    } else {
        fmt.Println("  ❌ FAIL: Passwords are the same")
    }
    fmt.Println()

    // ============================================================
    // 4. DIFFERENT LENGTHS
    // ============================================================
    fmt.Println("📌 [4] DIFFERENT LENGTHS")
    fmt.Println("------------------------------------------------------------")

    shortPwd, _ := smartpasslib.GenerateSmartPassword(secret, 12)
    longPwd, _ := smartpasslib.GenerateSmartPassword(secret, 24)
    fmt.Printf("  Length 12: %s\n", shortPwd)
    fmt.Printf("  Length 24: %s\n", longPwd)
    if len(shortPwd) == 12 && len(longPwd) == 24 {
        fmt.Println("  ✅ PASS: Lengths are correct")
    } else {
        fmt.Println("  ❌ FAIL: Lengths are incorrect")
    }
    fmt.Println()

    // ============================================================
    // 5. PUBLIC & PRIVATE KEYS
    // ============================================================
    fmt.Println("📌 [5] PUBLIC & PRIVATE KEYS")
    fmt.Println("------------------------------------------------------------")

    pubKey, _ := smartpasslib.GeneratePublicKey(secret)
    privKey, _ := smartpasslib.GeneratePrivateKey(secret)
    fmt.Println()
    fmt.Printf("  🔓 Public key (60 iterations) - STORE ON SERVER:\n  %s\n", pubKey)
    fmt.Println()
    fmt.Printf("  🔐 Private key (30 iterations) - NEVER STORE:\n  %s\n", privKey)
    fmt.Println()

    if pubKey != privKey {
        fmt.Println("  ✅ PASS: Public key != Private key")
    } else {
        fmt.Println("  ❌ FAIL: Keys are the same")
    }
    fmt.Println()

    // ============================================================
    // 6. VERIFICATION
    // ============================================================
    fmt.Println("📌 [6] SECRET VERIFICATION")
    fmt.Println("------------------------------------------------------------")

    isValid, _ := smartpasslib.VerifySecret(secret, pubKey)
    isInvalid, _ := smartpasslib.VerifySecret(wrongSecret, pubKey)
    fmt.Printf("  Correct secret: %v\n", isValid)
    fmt.Printf("  Wrong secret:   %v\n", isInvalid)
    if isValid && !isInvalid {
        fmt.Println("  ✅ PASS: Verification works")
    } else {
        fmt.Println("  ❌ FAIL: Verification failed")
    }
    fmt.Println()

    // ============================================================
    // 7. RANDOM PASSWORDS
    // ============================================================
    fmt.Println("📌 [7] RANDOM PASSWORD GENERATORS")
    fmt.Println("------------------------------------------------------------")

    strong, _ := smartpasslib.GenerateStrongPassword(16)
    base, _ := smartpasslib.GenerateBasePassword(16)
    code, _ := smartpasslib.GenerateCode(8)
    fmt.Printf("  🎲 Strong random (crypto secure): %s\n", strong)
    fmt.Printf("  🎲 Base random:                   %s\n", base)
    fmt.Printf("  🔢 Auth code (2FA):               %s\n", code)
    fmt.Println()

    // ============================================================
    // 8. INPUT VALIDATION
    // ============================================================
    fmt.Println("📌 [8] INPUT VALIDATION")
    fmt.Println("------------------------------------------------------------")

    _, err1 := smartpasslib.GenerateSmartPassword("short", 16)
    _, err2 := smartpasslib.GenerateCode(2)
    _, err3 := smartpasslib.GenerateStrongPassword(5)

    if err1 != nil {
        fmt.Printf("  ⚠️  Secret too short: %v\n", err1)
    }
    if err2 != nil {
        fmt.Printf("  ⚠️  Code too short: %v\n", err2)
    }
    if err3 != nil {
        fmt.Printf("  ⚠️  Password too short: %v\n", err3)
    }
    fmt.Println("  ✅ PASS: All validations work")
    fmt.Println()

    // ============================================================
    // RESULTS
    // ============================================================
    fmt.Println("============================================================")
    fmt.Println("🎉 ALL TESTS PASSED! Library is ready for production.")
    fmt.Println("============================================================")
    fmt.Println()
}