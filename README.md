# chacha32go

A lightweight Go implementation of the ChaCha32 stream cipher.  
Compatible with C and Arduino versions – designed for secure communication between microcontrollers and Go backends.

## 📦 Installation

```bash
go get github.com/regimantas/chacha32go/chacha32
```

## 🔐 Usage example

```go
package main

import (
    "fmt"
    "github.com/regimantas/chacha32go/chacha32"
)

func main() {
    // Message with '\0' at the end (as in C/Arduino)
    message := append([]byte("Hello ChaCha32!"), 0)

    // 32-byte key (same as Arduino/C)
    key := []byte{
        0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07,
        0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F,
        0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17,
        0x18, 0x19, 0x1A, 0x1B, 0x1C, 0x1D, 0x1E, 0x1F,
    }

    // 12-byte nonce (same as Arduino/C)
    nonce := []byte{
        0x00, 0x00, 0x00, 0x09,
        0x00, 0x00, 0x00, 0x4A,
        0x00, 0x00, 0x00, 0x00,
    }

    // Encrypt the message
    ciphertext := chacha32.Encrypt(key, nonce, message)

    // Print encrypted bytes
    fmt.Printf("Encrypted: ")
    for _, b := range ciphertext {
        fmt.Printf("%02X ", b)
    }
    fmt.Println()

    // Decrypt the message
    decrypted := chacha32.Decrypt(key, nonce, ciphertext)

    // Print decrypted text
    fmt.Printf("Decrypted: %s\n", decrypted)
}
```

---

## 🟢 Minimal usage example

```go
package main

import (
    "fmt"
    "github.com/regimantas/chacha32go/chacha32"
)

func main() {
    key := make([]byte, 32)   // 32 zero bytes
    nonce := make([]byte, 12) // 12 zero bytes
    message := []byte("Hello, world!")

    ciphertext := chacha32.Encrypt(key, nonce, message)
    decrypted := chacha32.Decrypt(key, nonce, ciphertext)

    fmt.Printf("Encrypted: %x\n", ciphertext)
    fmt.Printf("Decrypted: %s\n", decrypted)
}
```

---

## 🔗 Compatible Arduino Library

**[ChaCha32Arduino – Arduino compatible library](https://github.com/regimantas/ChaCha32Arduino/)**

---


## 📄 License

MIT
