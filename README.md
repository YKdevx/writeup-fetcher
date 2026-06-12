# writeup-fetcher

A Go-based automation tool that collects cybersecurity writeups from RSS feeds, processes article metadata, and stores them for research workflows.

---

## 🚀 Overview

**Writeup Fetcher** is a backend-focused project designed to demonstrate how to build a clean and modular data ingestion pipeline in Go. It retrieves RSS feeds from security-related sources, parses XML content, and outputs structured article information.

The project emphasizes simplicity, separation of concerns, and extensibility in a real-world backend workflow.

---

## 🏗️ Architecture

The system follows a simple pipeline design:

```
RSS Source → HTTP Fetcher → XML Parser → Data Model → Output
```

### Layers:

* **Fetcher Layer**: Handles HTTP requests and retrieves raw RSS data.
* **Parser Layer**: Converts XML data into structured Go structs.
* **Model Layer**: Defines internal data representations.
* **Application Layer**: Orchestrates the workflow in `main.go`.

---

## 🧰 Tech Stack

* Go (Golang)
* net/http (HTTP client)
* encoding/xml (RSS parsing)
* Standard library only (no external dependencies)

---

## 📦 Features

* Fetch RSS feeds from cybersecurity sources
* Parse XML-based RSS data
* Extract article metadata (title, link, publication date)
* Clean modular architecture
* Lightweight and dependency-free implementation

---

## ▶️ How to Run
**1. Clone the repository**
```bash
git clone https://github.com/YKdevx/writeup-fetcher.git
cd writeup-fetcher
```
**2. Initialize Go modules (if needed!)**
```bash
go mod tidy
```
**3. Run the application**
```bash
go run ./cmd
```
**4. Expected output**
You should see a list of article titles printed in the terminal:
```bash
How I Found a XSS Vulnerability
Understanding SQL Injection in 2026
Bug Bounty Tips for Beginners
...
```
## ⚠️ Requirements
- Go 1.18+
- Internet connection (RSS source must be reachable)
- 
---

## 🎯 Purpose

This project was built as part of my learning journey in backend development and cybersecurity automation. It demonstrates my ability to:

* Work with HTTP-based data ingestion
* Parse structured XML data
* Design modular and maintainable Go applications
* Build real-world backend pipelines

---

## 🔮 Future Improvements

* Add persistent storage (JSON / SQLite)
* Improve error handling and logging
* Introduce filtering by tags (e.g., XSS, SQLi, RCE)
* Add scheduling (cron-based automation)
* Extend parser to support additional RSS formats

---

## ⚠️ Notes

* The system currently relies on external RSS availability.
* Designed for simplicity and educational purposes.
* Architecture is intentionally kept minimal but extensible.

---

## 👨‍💻 Author

Built as a backend + cybersecurity learning project focused on Go and automation workflows.
