# Punycoder
[![GitHub release](http://img.shields.io/github/release/jakewarren/punycoder.svg?style=flat-square)](https://github.com/jakewarren/punycoder/releases])
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](https://github.com/jakewarren/punycoder/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/jakewarren/punycoder)](https://goreportcard.com/report/github.com/jakewarren/punycoder)

Converts punycode domains to unicode and vice-versa.

![Screenshot](screenshot.png)

## Installation

### Option 1: Binary

Download the latest release from [https://github.com/jakewarren/punycoder/releases](https://github.com/jakewarren/punycoder/releases)

### Option 2: From source

```
go get github.com/jakewarren/punycoder
```

## Usage

Simply provide either a punycode (IDN) or Unicode domain and it will be converted automagically.

```
$ punycoder bücher.example.com
xn--bcher-kva.example.com
$ punycoder xn--bcher-kva.example.com
bücher.example.com
```
