# Email-Api

[![Github Actions Status](https://github.com/gbrlsnchs/jwt/workflows/Linux,%20macOS%20and%20Windows/badge.svg)](https://github.com/gbrlsnchs/jwt/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/gbrlsnchs/jwt)](https://goreportcard.com/report/github.com/gbrlsnchs/jwt)
[![GoDoc](https://godoc.org/github.com/gbrlsnchs/jwt?status.svg)](https://pkg.go.dev/github.com/gbrlsnchs/jwt/v3)
[![Version compatibility with Go 1.11 onward using modules](https://img.shields.io/badge/compatible%20with-go1.11+-5272b4.svg)](https://github.com/gbrlsnchs/jwt#installing)

## About The Project

This mail api is the backend for sending emails using gmail smtp server.

## Getting Started

### Setup
Need to create .env file to set Gmail API KEY.

```bash
EMAIL_APIKEY=<API-KEY>
```

<!-- USAGE EXAMPLES -->
## Usage

There are swagger api for testing. It can use after running app. The URL is `http://localhost:8080/swagger/index.html`. Default username is `admin` and Default password is `password`.

* Running App
```bash
make run
```