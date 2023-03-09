# env

[![GoDoc](https://godoc.org/github.com/pkg-id/env?status.svg)](https://godoc.org/github.com/pkg-id/env)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/pkg-id/env/master/LICENSE)

env is a simple Go package that makes it easy to retrieve environment variables and provide fallback values if the specified key is not present. It also offers parsing of environment variables to various data types.

> This package is load or parse the dot env file. To load the dot env file, use [godotenv](https://github.com/joho/godotenv) package.

## Installation

```bash
go get github.com/pkg-id/env
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/pkg-id/env"
	"time"
)

func main() {
	// Get the value of the "FOO" environment variable or return "defaultFoo" if it doesn't exist.
	foo := env.String("FOO", "defaultFoo")
	fmt.Println(foo)

	// Get the value of the "BAR" environment variable as an integer or return 42 if it doesn't exist.
	bar := env.Int("BAR", 42)
	fmt.Println(bar)

	// Get the value of the "BAZ" environment variable as a float64 or return 3.14 if it doesn't exist.
	baz := env.Float64("BAZ", 3.14)
	fmt.Println(baz)

	// Get the value of the "QUX" environment variable as a boolean or return false if it doesn't exist.
	qux := env.Bool("QUX", false)
	fmt.Println(qux)

	// Get the value of the "MY_DURATION" environment variable as a time.Duration or return 10 seconds if it doesn't exist.
	duration := env.Duration("MY_DURATION", 10*time.Second)
	fmt.Println(duration)

	// Get the value of the "MY_LIST" environment variable as a list of integers or return []int{1, 2, 3} if it doesn't exist.
	list := env.List("MY_LIST", env.Parsers.Int(), []int{1, 2, 3})
	fmt.Println(list)
}

```


## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

