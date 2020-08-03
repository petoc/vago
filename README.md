# ISO Go

Simple ISO validators for Go.

## Standards

[ISO-3166-1](https://github.com/petoc/isogo/tree/master/iso31661) (alpha2, alpha3, numeric)<br>
[ISO-3166-2](https://github.com/petoc/isogo/tree/master/iso31662)<br>
[ISO-3166-3](https://github.com/petoc/isogo/tree/master/iso31663)<br>
[ISO-4217](https://github.com/petoc/isogo/tree/master/iso4217)

## Usage

```go
import "github.com/petoc/isogo/iso31661"
import "github.com/petoc/isogo/iso31662"
import "github.com/petoc/isogo/iso31663"
import "github.com/petoc/isogo/4217"
```

```go
iso31661.IsAlpha2("12")   // false
iso31661.IsAlpha2("AU")   // true
iso31661.IsAlpha3("123")  // false
iso31661.IsAlpha3("AUS")  // true
iso31661.IsNumeric("AUS") // false
iso31661.IsNumeric("004") // true
iso31661.Is("12")         // false
iso31661.Is("AU")         // true
iso31661.Is("XYZ")        // false
iso31661.Is("AUS")        // true
iso31661.Is("AUST")       // false

iso31662.Is("12-34")  // false
iso31662.Is("AU-NSW") // true

iso31663.Is("1234") // false
iso31663.Is("BQAQ") // true

iso4217.Is("123") // false
iso4217.Is("AUD") // true
```

## Sources

ISO-3166-1 (https://github.com/lukes/ISO-3166-Countries-with-Regional-Codes/blob/master/all/all.json)<br>
ISO-3166-2 (https://github.com/olahol/iso-3166-2.json/blob/master/iso-3166-2.json)<br>
ISO-3166-3 (https://github.com/haliaeetus/iso-3166/blob/master/data/iso_3166-3.json)<br>
ISO-4217 (https://github.com/umpirsky/currency-list/blob/master/data/en/currency.json)<br>

## License

Licensed under MIT license.
