# Validators

Collection of simple validators for Go.

## Available validators

[ISO-3166-1](https://github.com/petoc/vago/tree/master/iso31661) (alpha2, alpha3, numeric)<br>
[ISO-3166-2](https://github.com/petoc/vago/tree/master/iso31662)<br>
[ISO-3166-3](https://github.com/petoc/vago/tree/master/iso31663)<br>
[ISO-4217](https://github.com/petoc/vago/tree/master/iso4217)<br>
[Time Zone (IANA)](https://github.com/petoc/vago/tree/master/tz)<br>

## Usage

### ISO-3166-1

```go
import "github.com/petoc/vago/iso31661"
```

```go
iso31661.IsAlpha2("AU")     // true
iso31661.NameAlpha2("AU")   // Australia
iso31661.IsAlpha3("AUS")    // true
iso31661.NameAlpha3("AUS")  // Australia
iso31661.IsNumeric("004")   // true
iso31661.NameNumeric("AUS") // Australia
iso31661.Is("AU")           // true
iso31661.Is("AUS")          // true
iso31661.Is("004")          // true
iso31661.Name("AU")         // Australia
iso31661.Name("AUS")        // Australia
iso31661.Name("004")        // Australia
```

### ISO-3166-2

```go
import "github.com/petoc/vago/iso31662"
```

```go
iso31662.Is("AU-NSW")   // true
iso31662.Name("AU-NSW") // New South Wales
```

### ISO-3166-3

```go
import "github.com/petoc/vago/iso31663"
```

```go
iso31663.Is("BQAQ")   // true
iso31663.Name("BQAQ") // British Antarctic Territory
```

### ISO-4217

```go
import "github.com/petoc/vago/iso4217"
```

```go
iso4217.Is("EUR")   // true
iso4217.Name("EUR") // Euro
```

### Time zone (IANA)

```go
import "github.com/petoc/vago/tz"
```

```go
tz.Is("Europe/Berlin") // true
tz.Names("DE")         // [Europe/Berlin, Europe/Busingen]
```

## Sources

ISO-3166-1 (https://github.com/lukes/ISO-3166-Countries-with-Regional-Codes/blob/master/all/all.json)<br>
ISO-3166-2 (https://github.com/olahol/iso-3166-2.json/blob/master/iso-3166-2.json)<br>
ISO-3166-3 (https://github.com/haliaeetus/iso-3166/blob/master/data/iso_3166-3.json)<br>
ISO-4217 (https://github.com/umpirsky/currency-list/blob/master/data/en/currency.json)<br>
Time Zone (IANA) (https://github.com/eggert/tz/blob/master/zone.tab)<br>

## License

Licensed under MIT license.
