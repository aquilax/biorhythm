# Biorhythm

A small Go library that calculates pseudo-scientific biorhythm values from a birth date.

## Overview

This package exposes three cycle values based on the number of days since birth:

- Physical: 23-day cycle
- Emotional: 28-day cycle
- Intellectual: 33-day cycle

Each value is calculated with a sine wave and returns a number in the range roughly [-1, 1].

## Installation

```bash
go get github.com/aquilax/biorhythm
```

## Usage

```go
package main

import (
    "fmt"
    "time"

    "github.com/aquilax/biorhythm"
)

func main() {
    birthDate := time.Date(1990, 1, 1, 0, 0, 0, 0, time.UTC)
    now := time.Date(2024, 1, 15, 0, 0, 0, 0, time.UTC)

    calc := biorhythm.New(birthDate, now)

    fmt.Println("Physical:", calc.Physical())
    fmt.Println("Emotional:", calc.Emotional())
    fmt.Println("Intellectual:", calc.Intellectual())

    fmt.Println("Physical shortcut:", biorhythm.Physical(birthDate, now))
}
```

## Notes

- The calculation is based on elapsed days since birth and follows the standard sine-wave model used by biorhythm calculators.
