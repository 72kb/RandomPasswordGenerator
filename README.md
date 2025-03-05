# Random Password Generator

A simple Go-based password generator that can be used as both a CLI tool and a reusable package.

## Features

- Generates secure random passwords.
- Ensures at least one lowercase, uppercase, digit, and special character.
- Supports password lengths between 12 and 20 characters.
- Can be used as a command-line tool or imported as a Go package.

---

## 1️⃣ Using as a CLI Tool

### **Installation**

Clone the repository:

```sh
git clone https://github.com/72kb/RandomPasswordGenerator.git
cd RandomPasswordGenerator
```

Build the executable:

```sh
go build -o passwordgen
```

### **Usage**

Run the CLI with a specified password length:

```sh
./passwordgen -length=16
```

To save the password to a file:

```sh
./passwordgen -length=18 mypassword.txt
```

---

## 2️⃣ Using as a Go Package

### **Installation**

Fetch the package using `go get`:

```sh
go get github.com/72kb/RandomPasswordGenerator/randomizer
```

### **Import & Usage**

Use it in your Go project:

```go
package main

import (
    "fmt"
    "github.com/72kb/RandomPasswordGenerator/randomizer"
)

func main() {
    password := randomizer.generatePassword(16)
    fmt.Println("Generated Password:", password)
}
```

---

## Project Structure

```
RandomPasswordGenerator/
│── randomizer/         # Go package folder
│   ├── randomizer.go   # Contains password generation logic
│── main.go             # CLI tool
│── go.mod              # Go module file
│── README.md           # Documentation
```

---

## License

This project is open-source under the [MIT License](LICENSE).

