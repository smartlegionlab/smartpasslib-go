# SmartPassLib Go <sup>v1.0.1</sup>

**Go implementation of deterministic smart password generator. Same secret + same length = same password across all platforms (Python, JS, Kotlin, Go).**

---

[![GitHub top language](https://img.shields.io/github/languages/top/smartlegionlab/smartpasslib-go)](https://github.com/smartlegionlab/smartpasslib-go)
[![GitHub license](https://img.shields.io/github/license/smartlegionlab/smartpasslib-go)](https://github.com/smartlegionlab/smartpasslib-go/blob/master/LICENSE)
[![GitHub release](https://img.shields.io/github/v/release/smartlegionlab/smartpasslib-go)](https://github.com/smartlegionlab/smartpasslib-go/)
[![GitHub stars](https://img.shields.io/github/stars/smartlegionlab/smartpasslib-go?style=social)](https://github.com/smartlegionlab/smartpasslib-go/stargazers)
[![GitHub forks](https://img.shields.io/github/forks/smartlegionlab/smartpasslib-go?style=social)](https://github.com/smartlegionlab/smartpasslib-go/network/members)

---

## ⚠️ Disclaimer

**By using this software, you agree to the full disclaimer terms.**

**Summary:** Software provided "AS IS" without warranty. You assume all risks.

**Full legal disclaimer:** See [DISCLAIMER.md](https://github.com/smartlegionlab/smartpasslib-go/blob/master/DISCLAIMER.md)

---

## Core Principles

- **Deterministic Generation**: Same secret + same length = same password, every time
- **Zero Storage**: Passwords exist only when generated, never stored
- **Cross-Platform**: Compatible with Python, JS, Kotlin implementations
- **Crypto Secure**: Uses crypto/rand for random generation

## Key Features

- **Smart Password Generation**: Deterministic from secret phrase
- **Public/Private Key System**: 30 iterations for private key, 60 for public key
- **Secret Verification**: Verify secret without exposing it
- **Random Password Generation**: Cryptographically secure random passwords
- **Authentication Codes**: Short codes for 2FA/MFA (4-20 chars)
- **No External Dependencies**: Pure Go, uses standard crypto

## Security Model

- **Proof of Knowledge**: Public keys verify secrets without exposing them
- **Deterministic Certainty**: Mathematical certainty in password regeneration
- **Ephemeral Passwords**: Passwords exist only in memory during generation
- **Local Computation**: No data leaves your device
- **No Recovery Backdoors**: Lost secret = permanently lost passwords (by design)

---

## Research Paradigms & Publications

- **[Pointer-Based Security Paradigm](https://doi.org/10.5281/zenodo.17204738)** - Architectural Shift from Data Protection to Data Non-Existence
- **[Local Data Regeneration Paradigm](https://doi.org/10.5281/zenodo.17264327)** - Ontological Shift from Data Transmission to Synchronous State Discovery

---

## Technical Foundation

**Key derivation (same as Python/JS/Kotlin versions):**

| Key Type    | Iterations | Purpose                            |
|-------------|------------|------------------------------------|
| Private Key | 30         | Password generation (never stored) |
| Public Key  | 60         | Verification (stored on server)    |

**Character Set:** `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789!@#$&*-_`

## Installation

```bash
go get github.com/smartlegionlab/smartpasslib-go
```

## Quick Usage

### Generate Smart Password
```go
package main

import (
    "fmt"
    "github.com/smartlegionlab/smartpasslib-go"
)

func main() {
    secret := "MyCatHippo2026"
    length := 16
    
    password, _ := smartpasslib.GenerateSmartPassword(secret, length)
    fmt.Println(password) // e.g., "jrh_E5V!2#neNjnP"
}
```

### Generate Public/Private Keys
```go
secret := "MyCatHippo2026"

publicKey, _ := smartpasslib.GeneratePublicKey(secret)
privateKey, _ := smartpasslib.GeneratePrivateKey(secret)

fmt.Println("Public Key (store on server):", publicKey)
fmt.Println("Private Key (never store):", privateKey)
```

### Verify Secret Against Public Key
```go
secret := "MyCatHippo2026"
storedPublicKey := "..." // from server

isValid, _ := smartpasslib.VerifySecret(secret, storedPublicKey)
if isValid {
    password, _ := smartpasslib.GenerateSmartPassword(secret, 16)
}
```

### Generate Random Passwords
```go
// Strong random (cryptographically secure)
strong, _ := smartpasslib.GenerateStrongPassword(20)

// Base random
base, _ := smartpasslib.GenerateBasePassword(16)

// Authentication code (4-20 chars)
code, _ := smartpasslib.GenerateCode(8)
```

## API Reference

### Constants

| Constant  | Type   | Description                       |
|-----------|--------|-----------------------------------|
| `Version` | string | Library version                   |
| `Chars`   | string | Character set used for generation |

### Functions

| Function                                      | Parameters        | Returns         | Description                      |
|-----------------------------------------------|-------------------|-----------------|----------------------------------|
| `GeneratePrivateKey(secret)`                  | secret: string    | (string, error) | Private key (30 iterations)      |
| `GeneratePublicKey(secret)`                   | secret: string    | (string, error) | Public key (60 iterations)       |
| `VerifySecret(secret, publicKey)`             | secret, publicKey | (bool, error)   | Verify secret matches public key |
| `GenerateSmartPassword(secret, length)`       | secret, length    | (string, error) | Deterministic password           |
| `GenerateStrongPassword(length)`              | length            | (string, error) | Cryptographically random         |
| `GenerateBasePassword(length)`                | length            | (string, error) | Simple random password           |
| `GenerateCode(length)`                        | length            | (string, error) | Short code (4-20 chars)          |

### Input Validation

| Parameter       | Minimum  | Maximum    |
|-----------------|----------|------------|
| Secret phrase   | 12 chars | unlimited  |
| Password length | 12 chars | 1000 chars |
| Code length     | 4 chars  | 20 chars   |

## Security Requirements

### Secret Phrase
- **Minimum 12 characters** (enforced)
- Case-sensitive
- Use mix of: uppercase, lowercase, numbers, symbols, emoji, or Cyrillic
- Never store digitally
- **NEVER use your password description as secret phrase**

### Strong Secret Examples
```
✅ "MyCatHippo2026"          — mixed case + numbers
✅ "P@ssw0rd!LongSecret"     — special chars + numbers + length
✅ "КотБегемот2026НаДиете"   — Cyrillic + numbers
✅ "GitHubPersonal2026!"     — description + extra chars (but not the description alone)
```

### Weak Secret Examples (avoid)
```
❌ "GitHub Account"          — using description as secret (weak!)
❌ "password"                — dictionary word, too short
❌ "1234567890"              — only digits, too short
❌ "qwerty123"               — keyboard pattern
❌ Same as description       — never use the same value as password description
```

## Cross-Platform Compatibility

SmartPassLib Go produces **identical passwords** to:

| Platform   | Repository                                                                                                                |
|------------|:--------------------------------------------------------------------------------------------------------------------------|
| Python     | [smartpasslib](https://github.com/smartlegionlab/smartpasslib)                                                            |
| JavaScript | [smartpasslib-js](https://github.com/smartlegionlab/smartpasslib-js)                                                      |
| Kotlin     | [smartpasslib-kotlin](https://github.com/smartlegionlab/smartpasslib-kotlin)                                              |
| Go         | [smartpasslib-go](https://github.com/smartlegionlab/smartpasslib-go)                                                      |
| Web        | [Web Manager](https://github.com/smartlegionlab/smart-password-manager-web)                                               |
| Android    | [Android Manager](https://github.com/smartlegionlab/smart-password-manager-android)                                       |
| Desktop    | [Desktop Manager](https://github.com/smartlegionlab/smart-password-manager-desktop)                                       |
| CLI        | [CLI PassMan](https://github.com/smartlegionlab/clipassman) / [CLI PassGen](https://github.com/smartlegionlab/clipassgen) |

## Testing

### Install Go

```bash
# Arch Linux
sudo pacman -S go

# Ubuntu/Debian
sudo apt install golang
```

### Run example

```bash
go run ./cmd/example/main.go
````

### Run tests

```bash
go test -v
```

### Run test script

```bash
chmod +x test.sh
./test.sh
```

## Ecosystem

**Core Libraries:**
- **[smartpasslib](https://github.com/smartlegionlab/smartpasslib)** - Python implementation
- **[smartpasslib-js](https://github.com/smartlegionlab/smartpasslib-js)** - JavaScript implementation
- **[smartpasslib-kotlin](https://github.com/smartlegionlab/smartpasslib-kotlin)** - Kotlin implementation
- **[smartpasslib-go](https://github.com/smartlegionlab/smartpasslib-go)** - Go implementation

**Applications:**
- **[Desktop Manager](https://github.com/smartlegionlab/smart-password-manager-desktop)** - Cross-platform desktop app
- **[CLI PassMan](https://github.com/smartlegionlab/clipassman)** - Console password manager
- **[CLI PassGen](https://github.com/smartlegionlab/clipassgen)** - Console password generator
- **[Web Manager](https://github.com/smartlegionlab/smart-password-manager-web)** - Web interface
- **[Android Manager](https://github.com/smartlegionlab/smart-password-manager-android)** - Mobile Android app

## License

**[BSD 3-Clause License](LICENSE)**

Copyright (©) 2026, [Alexander Suvorov](https://github.com/smartlegionlab)

## Author

**Alexander Suvorov** - [GitHub](https://github.com/smartlegionlab)

---

## Support

- **Issues**: [GitHub Issues](https://github.com/smartlegionlab/smartpasslib-go/issues)
- **Documentation**: This [README](README.md)

---

