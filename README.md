# SmartPassLib Go <sup>v4.0.0</sup>

---

**Smart Passwords Library**: Cryptographic password generation and management without storage. 
Generate passwords from secrets, verify knowledge without exposure, manage metadata securely.

**Now with Cross-Platform Determinism**: Same secret + same parameters = identical password on 
**Go, C#, Python, Kotlin, JavaScript** and any language with SHA-256.

**Decentralized by Design**: Unlike traditional password managers that store encrypted vaults on central servers, 
smartpasslib stores nothing. Your secrets never leave your device. Passwords are regenerated on-demand — 
**no cloud, no database, no trust required**.

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

## 🔄 Breaking Change (v4.0.0)

> **⚠️ This version is NOT backward compatible with v1.x.x**

Passwords generated with older versions **cannot be regenerated** with v4.0.0.

📖 **Full migration instructions** → see [MIGRATION.md](https://github.com/smartlegionlab/smartpasslib-go/blob/master/MIGRATION.md)

---

## Core Principles

- **Zero-Storage Security**: No passwords or secret phrases are ever stored or transmitted
- **Decentralized Architecture**: No central servers, no cloud dependency, no third-party trust required
- **Cross-Platform Deterministic Generation**: Identical secret + parameters = identical password **on any language** (SHA-256 based)
- **Metadata Only**: Store only verification metadata (public keys, descriptions, lengths)
- **On-Demand Regeneration**: Passwords are recalculated when needed, never retrieved from storage
- **Cryptographically Secure**: Uses `crypto/rand` and SHA-256

## Key Features

- **Decentralized & Serverless**: No central database, no cloud lock-in, complete user sovereignty
- **Smart Password Generation**: Deterministic from secret phrase
- **Public/Private Key System**: 15-30 iterations for private key, 45-60 for public key (dynamic per secret)
- **Secret Verification**: Verify secret without exposing it
- **Random Password Generation**: Cryptographically secure random passwords
- **Authentication Codes**: Short codes for 2FA/MFA (4-100 chars)
- **No External Dependencies**: Pure Go, uses standard crypto

## Security Model

- **Proof of Knowledge**: Public keys verify secrets without exposing them
- **Decentralized Trust**: No third party needed — you control your secrets completely
- **Deterministic Security**: Same input = same output, always reproducible across platforms
- **Dynamic Iteration Counts**: Private key uses 15-30 iterations, public key uses 45-60 iterations (deterministic per secret)
- **No Vulnerable Metadata Storage**: Only public keys and descriptions can be stored (optional)
- **Zero Storage of Secrets**: Secret phrases exist only in your memory, private keys are derived on-demand and never persisted
- **No Recovery Backdoors**: Lost secret = permanently lost passwords (by design)

---

## Research Paradigms & Publications

- **[Pointer-Based Security Paradigm](https://doi.org/10.5281/zenodo.17204738)** - Architectural Shift from Data Protection to Data Non-Existence
- **[Local Data Regeneration Paradigm](https://doi.org/10.5281/zenodo.17264327)** - Ontological Shift from Data Transmission to Synchronous State Discovery

---

## Technical Foundation

**Key derivation (same as Python/JS/Kotlin/C# versions v4.0.0):**

| Key Type    | Iterations              | Purpose                                                 |
|-------------|-------------------------|---------------------------------------------------------|
| Private Key | 15-30 (dynamic)         | Password generation (never stored, never transmitted)   |
| Public Key  | 45-60 (dynamic)         | Verification (stored locally)                           |

**Character Set:** `!@#$%^&*()_+-=[]{};:,.<>?/ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789abcdefghijklmnopqrstuvwxyz`

**Validation Rules:**
- Secret phrase: minimum 12 characters
- Password length: 12-100 characters
- Code length: 4-100 characters

**Decentralized Architecture**:
- No central authority required
- Metadata can be synced via any channel (USB, cloud, even paper)
- Your security depends only on your secret phrase, not on any service provider
- Works offline — no internet connection required

## Installation

```bash
go get github.com/smartlegionlab/smartpasslib-go@v4.0.0
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
    secret := "MyStrongSecretPhrase2026!"
    length := 16
    
    password, _ := smartpasslib.GenerateSmartPassword(secret, length)
    fmt.Println(password)
}
```

### Generate Public/Private Keys
```go
secret := "MyStrongSecretPhrase2026!"

publicKey, _ := smartpasslib.GeneratePublicKey(secret)
privateKey, _ := smartpasslib.GeneratePrivateKey(secret)

fmt.Println("Public Key (store locally):", publicKey)
fmt.Println("Private Key (never store):", privateKey)
```

### Verify Secret Against Public Key
```go
secret := "MyStrongSecretPhrase2026!"
storedPublicKey := "..." // from local

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

// Authentication code (4-100 chars)
code, _ := smartpasslib.GenerateCode(8)
```

## API Reference

### Constants

| Constant  | Type   | Description                       |
|-----------|--------|-----------------------------------|
| `Version` | string | Library version (4.0.0)           |
| `Chars`   | string | Character set used for generation |

### Functions

| Function                                      | Parameters        | Returns         | Description                      |
|-----------------------------------------------|-------------------|-----------------|----------------------------------|
| `GeneratePrivateKey(secret)`                  | secret: string    | (string, error) | Private key (15-30 iterations)   |
| `GeneratePublicKey(secret)`                   | secret: string    | (string, error) | Public key (45-60 iterations)    |
| `VerifySecret(secret, publicKey)`             | secret, publicKey | (bool, error)   | Verify secret matches public key |
| `GenerateSmartPassword(secret, length)`       | secret, length    | (string, error) | Deterministic password           |
| `GenerateStrongPassword(length)`              | length            | (string, error) | Cryptographically random         |
| `GenerateBasePassword(length)`                | length            | (string, error) | Simple random password           |
| `GenerateCode(length)`                        | length            | (string, error) | Short code (4-100 chars)         |

### Input Validation

| Parameter       | Minimum  | Maximum    |
|-----------------|----------|------------|
| Secret phrase   | 12 chars | unlimited  |
| Password length | 12 chars | 100 chars  |
| Code length     | 4 chars  | 100 chars  |

## Security Requirements

### Secret Phrase
- **Minimum 12 characters** (enforced)
- Case-sensitive
- Use mix of: uppercase, lowercase, numbers, symbols
- Never store digitally
- **NEVER use your password description as secret phrase**

### Strong Secret Examples
```
✅ "MyStrongSecretPhrase2026!"   — mixed case + numbers + symbols
✅ "P@ssw0rd!LongSecret"         — special chars + numbers + length
✅ "GitHubPersonal2026!"         — description + extra chars
```

### Weak Secret Examples (avoid)
```
❌ "short"                       — too short, returns error
❌ "GitHub Account"              — using description as secret (weak!)
❌ "password"                    — dictionary word, too short
❌ "1234567890"                  — only digits, too short
```

### Decentralized Nature

**There is no "forgot password" button.** This is by design:

- No central server can reset your passwords
- No support team can recover your access
- Your secret phrase is the ONLY key

**This is the price of true decentralization** — you are completely in control.

## Cross-Platform Implementations

The same deterministic algorithm is available in multiple languages.
SmartPassLib Go produces **identical passwords** to:

| Language   | Repository                                                                   |
|------------|:-----------------------------------------------------------------------------|
| Python     | [smartpasslib](https://github.com/smartlegionlab/smartpasslib)               |
| JavaScript | [smartpasslib-js](https://github.com/smartlegionlab/smartpasslib-js)         |
| Kotlin     | [smartpasslib-kotlin](https://github.com/smartlegionlab/smartpasslib-kotlin) |
| C#         | [smartpasslib-csharp](https://github.com/smartlegionlab/smartpasslib-csharp) |

## Testing

### Install Go

```bash
# Arch Linux
sudo pacman -S go

# Ubuntu/Debian
sudo apt install golang
```

### Run tests

```bash
go test -v
```

## Ecosystem

**Core Libraries:**
- **[smartpasslib](https://github.com/smartlegionlab/smartpasslib)** - Python
- **[smartpasslib-js](https://github.com/smartlegionlab/smartpasslib-js)** - JavaScript
- **[smartpasslib-kotlin](https://github.com/smartlegionlab/smartpasslib-kotlin)** - Kotlin
- **[smartpasslib-go](https://github.com/smartlegionlab/smartpasslib-go)** - Go (this)
- **[smartpasslib-csharp](https://github.com/smartlegionlab/smartpasslib-csharp)** - C#

**CLI Applications:**
- **[CLI Smart Password Manager (Python)](https://github.com/smartlegionlab/clipassman)**
- **[CLI Smart Password Generator (Python)](https://github.com/smartlegionlab/clipassgen)**
- **[CLI Smart Password Manager (C#)](https://github.com/smartlegionlab/SmartPasswordManagerCsharpCli)**
- **[CLI Smart Password Generator (C#)](https://github.com/smartlegionlab/SmartPasswordGeneratorCsharpCli)** 

**Desktop Applications:**
- **[Desktop Smart Password Manager (Python)](https://github.com/smartlegionlab/smart-password-manager-desktop)**
- **[Desktop Smart Password Manager (C#)](https://github.com/smartlegionlab/SmartPasswordManagerCsharpDesktop)**

**Other:**
- **[Web Smart Password Manager](https://github.com/smartlegionlab/smart-password-manager-web)**
- **[Android Smart Password Manager](https://github.com/smartlegionlab/smart-password-manager-android)**

## License

**[BSD 3-Clause License](https://github.com/smartlegionlab/smartpasslib-go/blob/master/LICENSE)**

Copyright (©) 2026, [Alexander Suvorov](https://github.com/smartlegionlab)

## Author

**Alexander Suvorov** - [GitHub](https://github.com/smartlegionlab)

---

## Support

- **Issues**: [GitHub Issues](https://github.com/smartlegionlab/smartpasslib-go/issues)
- **Documentation**: This [README](https://github.com/smartlegionlab/smartpasslib-go/blob/master/README.md)

---

