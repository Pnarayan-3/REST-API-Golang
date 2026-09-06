# Student REST API

A simple RESTful API built with **Go** and **SQLite** for managing student records.

The project demonstrates how to build a backend REST API in Go with HTTP routing, request validation, SQLite database integration, and a clean separation between handlers and storage.


## 🛠️ Tech Stack

* **Language:** Go
* **API:** REST
* **Database:** SQLite
* **Routing:** Go HTTP Router
* **Validation:** go-playground/validator
* **Configuration:** cleanenv
* **Testing:** Go testing package
* **CI/CD:** GitHub Actions

## ⚙️ Prerequisites

Make sure you have the following installed:

* Go 1.27+
* Git

Check your Go installation:

```bash
go version
```

## 📥 Clone the Repository

```bash
git clone <your-repository-url>
cd STUDENT-API
```

## 📦 Install Dependencies

```bash
go mod download
```

You can also run:

```bash
go mod tidy
```

## ▶️ Run the Application

From the project root:

```bash
go run ./cmd/student-api
```

The API will start on the configured port.

For example:

```text
http://localhost:8082
```

## 🔍 Code Quality

Format the Go code:

```bash
gofmt -w .
```

Run Go's static analysis:

```bash
go vet ./...
```

Build the application:

```bash
go build ./...
```

## 👨‍💻 Author

**Pushkar Narayan**

B.Tech – Computer Science and Information Technology

---

⭐ If you found this project useful, consider giving the repository a star.
