# QValues

[Document](https://pkg.go.dev/github.com/7sDream/qvalue)

QValues is a parser for [quality values](https://developer.mozilla.org/en-US/docs/Glossary/Quality_values). It's commonly used in HTTP Headers like `Accept`, `Accept-Language`, `Accept-Encoding` etc.

The quality value this package provided is multiplied by 1000, to avoid floating point sorting.

The parser can also return sorted quality values(bigger quality first), if two item has same quality, the one that occurs first takes precedence.

I write this package just for my own convenience, but if you have any suggestion, issues and PRs are both welcome.

## Install

```bash
go get github.com/7sDream/qvalue
```

## Example

```go
package main

import (
    "github.com/7sDream/qvalue"
)

var qvParser = qvalue.NewOptions()

func main() {
    acceptEncoding := "gzip, deflate;q=0.8, br;q=0.7, identity;q=0.5, *;q=0.1"
    // If you do not care quality, just want deal them in origin order, use `Parse` instead
    qvs, _ := qvParser.ParseAndSort(acceptEncoding)
    for _, qv := range qvs {
        // For example, we can only deal with gzip and identity
        switch qv.Value {
        case "gzip":
            {
                /* ... */
                break
            }
        case "*", "identity":
            {
                /* ... */
                break
            }
        }
    }
}
```

## License

MIT.
