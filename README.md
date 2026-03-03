
# GoLang
Go (Golang) is a high-performance programming language created by Google, widely used for building web services, microservices, and distributed systems. Its lightweight concurrency model, efficient memory management, and fast compilation make it a strong choice for cloud platforms and large-scale systems

# Information

Go is expressive, concise, clean, and efficient

los números que acompañan al tipo de dato son bits y el valor es equivalente a 2 elevato a la #numero de bits

# Go on Windows --- Professional Guide (Run, Debug, Test, Build Like a Senior Engineer)

**Author mindset:** 10+ years Go engineering experience\
**Platform:** Windows (PowerShell / CMD)\
**Audience:** Developers who want production-level workflows

------------------------------------------------------------------------

# Table of Contents

1.  Installing Go on Windows
2.  Understanding Your Go Environment
3.  Creating a Professional Project Structure
4.  Running Go Projects (Development Mode)
5.  Building Executables (Production Mode)
6.  Debugging with Delve (CLI + VS Code)
7.  Testing (Unit, Integration, Race Detection)
8.  Benchmarks & Performance Analysis
9.  Coverage & Reports
10. Advanced Pro Tips
11. Command Cheat Sheet

------------------------------------------------------------------------

# 1. Installing Go on Windows

## Step 1 --- Download & Install

1.  Download the Windows MSI installer from the official Go website.
2.  Run the installer (default: `C:\Program Files\Go`).
3.  Ensure the installer adds Go to your system PATH.

## Step 2 --- Verify Installation

Open PowerShell or CMD:

``` bash
go version
```

Expected output:

``` bash
go version go1.xx.x windows/amd64
```

------------------------------------------------------------------------

# 2. Understanding Your Go Environment

Inspect your Go environment:

``` bash
go env
```

Important variables:

  Variable     Meaning
  ------------ -----------------------------------
  GOPATH       Workspace for caches and binaries
  GOMOD        Active module path
  GOCACHE      Build cache location
  GOMODCACHE   Module dependency cache

------------------------------------------------------------------------

# 3. Creating a Professional Project Structure

Initialize a new module:

``` bash
go mod init example.com/your-app
go mod tidy
```

Recommended structure:

    your-app/
      go.mod
      cmd/
        your-app/
          main.go
      internal/
      pkg/

-   `cmd/` → Application entry points\
-   `internal/` → Private business logic\
-   `pkg/` → Public reusable packages

------------------------------------------------------------------------

# 4. Running Go Projects (Development Mode)

## Run a Single File

``` bash
go run main.go
```

## Run Entire Module

``` bash
go run .
```

## Pass CLI Arguments

``` bash
go run . --port=8080 --mode=dev
```

------------------------------------------------------------------------

# 5. Building Executables (Production Mode)

## Basic Build

``` bash
go build
```

## Custom Output Name

``` bash
go build -o your-app.exe .
```

Run executable:

``` bash
.\your-app.exe
```

## Optimized Production Build

``` bash
go build -trimpath -ldflags="-s -w" -o your-app.exe .
```

------------------------------------------------------------------------

# 6. Debugging with Delve

## Install Delve

``` bash
go install github.com/go-delve/delve/cmd/dlv@latest
```

Verify:

``` bash
dlv version
```

## Debug in Terminal

``` bash
dlv debug .
```

Common commands inside Delve:

-   break main.main
-   continue
-   next
-   step
-   locals
-   print variableName
-   goroutines

------------------------------------------------------------------------

# VS Code Debug Configuration

Create `.vscode/launch.json`:

``` json
{
  "version": "0.2.0",
  "configurations": [
    {
      "name": "Go: Debug",
      "type": "go",
      "request": "launch",
      "mode": "debug",
      "program": "${workspaceFolder}"
    }
  ]
}
```

------------------------------------------------------------------------

# 7. Testing in Go

Run all tests:

``` bash
go test ./...
```

Verbose mode:

``` bash
go test -v ./...
```

Race detection:

``` bash
go test -race ./...
```

Run specific test:

``` bash
go test -run TestName ./...
```

Short mode:

``` bash
go test -short ./...
```

------------------------------------------------------------------------

# 8. Benchmarks

Run benchmarks:

``` bash
go test -bench=. -benchmem ./...
```

Run specific benchmark:

``` bash
go test -bench=BenchmarkName -benchmem ./...
```

------------------------------------------------------------------------

# 9. Coverage

Generate coverage profile:

``` bash
go test -coverprofile=coverage.out ./...
```

Generate HTML report:

``` bash
go tool cover -html=coverage.out
```

------------------------------------------------------------------------

# 10. Advanced Pro Tips

• Always use `context.Context` for cancellation in production systems.\
• Run `go test -race` regularly in concurrent projects.\
• Keep modules clean with `go mod tidy`.\
• Use `internal/` to protect business logic boundaries.\
• Benchmark before optimizing.\
• Prefer small, composable packages.\
• Avoid unnecessary global state.

------------------------------------------------------------------------

# 11. Command Cheat Sheet

### Run

``` bash
go run .
```

### Build

``` bash
go build -o app.exe .
```

### Test

``` bash
go test ./...
```

### Race Detection

``` bash
go test -race ./...
```

### Coverage

``` bash
go test -cover ./...
```

### Benchmarks

``` bash
go test -bench=. -benchmem ./...
```

### Debug

``` bash
dlv debug .
```

------------------------------------------------------------------------

# Final Words

This workflow is what real Go engineers use in production environments
on Windows.\
Master these commands, structure, debugging tools, and testing
strategies --- and you are operating at a senior level.

Happy coding 🚀



