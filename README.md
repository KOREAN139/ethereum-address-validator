# ethav (ethereum-address-validator)
Simple address validator for **ETH**

Checks format and checksum of address

# Installation
1. Download and install:
```
go get -u github.com/KOREAN139/ethereum-address-validator
```
2. Import it in your code:
```go
import "github.com/KOREAN139/ethereum-address-validator"
```

# Quick Start

```go
package main

import "github.com/KOREAN139/ethereum-address-validator"

func main() {
    // ethav.Validate(ETH address) will return error if address is invalid
    if err := ethav.Validate("0xdbF03B407c01E7cD3CBea99509d93f8DDDC8C6FB"); err != nil {
        // do error handling
    }
    // do your work
}
```
