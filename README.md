# testing

A PHPUnit-style test harness for [togo](https://github.com/togo-framework/togo):
HTTP request helpers + fluent assertions + an in-memory SQLite helper. Import as
`togotest`. The CLI `togo make:test` and `togo make:e2e` (Playwright, Dusk-style)
generate tests that use it.

```bash
togo install togo-framework/testing
```

```go
import togotest "github.com/togo-framework/testing"

func TestList(t *testing.T) {
  r := togotest.Do(t, handler, "GET", "/api/posts", nil)
  r.Status(t, 200).JSON(t, &out)
}
```
