# 📟 ASCII Art Generator
> A Go application that transforms terminal input into stylized ASCII art using pre-configured banner templates.

---

### 📋 Description
Developed exclusively in **Go** using standard packages, the project features a modular architecture. It maps terminal input to graphic representations using ASCII characters based on specific banner templates.

### 🏗️ Architecture
| Component | Role |
| :--- | :--- |
| `main.go` | Application entry point and Core rendering & logic  Banner template parsing
| `main_test.go` | Test files |
| `standard.txt` | Default font template |
| `shadow.txt` | Shadow font template |
| `thinkertoy.txt` | Thinkertoy font template |

---

### 🚀 Usage
```bash
$ go run . [flag]