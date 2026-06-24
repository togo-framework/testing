<!-- togo-brand -->
<p align="center">
  <img src=".github/assets/togo-mark.svg" width="96" alt="togo" />
</p>
<h1 align="center">testing</h1>
<p align="center"><sub>part of the <a href="https://github.com/togo-framework">togo-framework</a> — the full-stack Go + React framework</sub></p>

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


---

## 💎 Premium sponsors

togo is proudly sponsored by **ID8 Media** and **One Studio**.

<p align="center">
  <a href="https://id8media.com"><img src=".github/assets/id8media.svg" height="44" alt="ID8 Media" /></a>
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
  <a href="https://one-studio.co"><img src=".github/assets/one-studio.jpeg" height="44" alt="One Studio" /></a>
</p>
