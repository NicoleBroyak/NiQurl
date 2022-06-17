# urlhandler
--
    import "github.com/nicolebroyak/niqurl/tools/urlhandler"


## Usage

#### func  ShortenURLError

```go
func ShortenURLError() error
```

#### type NiqURL

```go
type NiqURL struct {
	*url.URL
	LongURL  string
	ShortURL string
	UserName string
}
```

Alias for *url.URL type

#### func  StringToNiqURL

```go
func StringToNiqURL(input string) (*NiqURL, error)
```

#### func (*NiqURL) GenerateShortURLPath

```go
func (NiqURL *NiqURL) GenerateShortURLPath(urlLen int)
```

#### func (*NiqURL) IfEmptySchemeAddHTTPS

```go
func (NiqURL *NiqURL) IfEmptySchemeAddHTTPS() *NiqURL
```
