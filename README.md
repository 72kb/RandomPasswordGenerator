# Secure Password Generator (Go)

A lightweight and secure password generator written in **Go**. This tool generates random passwords of configurable length and complexity, ensuring strong security practices using **cryptographic randomness (`crypto/rand`)**.

## 🚀 Features
- ✅ Generates passwords between **12-20 characters** (configurable).
- ✅ Includes at least **one uppercase, lowercase, digit, and special character**.
- ✅ Uses **cryptographically secure randomization**.
- ✅ Supports **shuffling for unpredictability**.
- ✅ Optimized for performance using **`strings.Builder`**.
- ✅ Can be used as a **CLI tool**.

## 🛠 Installation
Ensure you have **Go installed**. If not, [download it here](https://go.dev/dl/).

```sh
# Clone the repository
git clone https://github.com/yourusername/password-generator-go.git
cd password-generator-go

# Build the project
go build -o password-generator

# Run the program
./password-generator
```

## 📌 Usage

### **1️⃣ Generate a Password in Code**
```go
package main

import (
    "fmt"
)

func main() {
    password := generatePassword(16) // Generate a 16-character password
    fmt.Println("Generated Password:", password)
}
```

### **2️⃣ Run as a CLI Tool**
```sh
./password-generator -length=14
```

### **3️⃣ Example Output**
```
Generated Password: G7!xPq19B4$
```

## 🔧 Configuration
Modify `generatePasswordWithConfig` to customize password rules:
```go
config := PasswordConfig{
    IncludeLowercase: true,
    IncludeUppercase: true,
    IncludeDigits:    true,
    IncludeSpecial:   false, // No special characters
    Length:           14,
}
password := generatePasswordWithConfig(config)
```

## 🔒 Security Considerations
- Uses **`crypto/rand`** for **secure randomness**.
- **Avoids predictable patterns** by shuffling characters.
- **Never store passwords** in plaintext; use a **password manager**.

## 📜 License
This project is **open-source** under the MIT License.

## 💡 Contributing
Pull requests are welcome! Feel free to open issues for improvements or feature requests.

## 📬 Contact
Have questions or suggestions? Reach out on **GitHub Issues**!

---
💻 **Happy coding!** 🚀


