# QValues

QValues is a parser for [Quality values](https://developer.mozilla.org/en-US/docs/Glossary/Quality_values). It's commonly used in HTTP Header like `Accept`, `Accept-Language`, `Accept-Encoding` etc.

The quality value this package provided is multiplied by 1000, for avoid float point sort.

The parser can also returned sorted quality values(bigger quality first), if two item has same quality, the one that occurs first takes precedence.

I write this package just for my own convenience, but if you have and advice, issues and PRs are welcome.

## Install

```bash
go get github.com/7sDream/qvalue
```

## Example

```go
package qvalue_test

import (
    "github.com/7sDream/qvalue"
)

var qvParser = qvalue.NewOptions()

func main() {
    acceptEncoding := "gzip, deflate;q=0.8, br;q=0.7, identity;q=0.5, *;q=0.1"
    // If you do not care quality, but want deal then in origin order, use `Parse` instead
    qvs, _ := qvParser.ParseAndSort(acceptEncoding)
    for _, qv := range qvs {
        // For example, we can only deal with gzip and identity
        switch qv.Value {
        case "gzip":
            {
                /* ... */
                break
            }
        case "*":
        case "identity":
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
