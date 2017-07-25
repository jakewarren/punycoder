# Punycoder

Converts punycode domains to unicode and vice-versa.

## Installation

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
