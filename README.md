# Go: gontext

## INSTALL

```bash
go get github.com/hidori/go-gontext@latest
```

## USAGE

CODE:

```go
package main

import (
 "context"
 "fmt"
 "time"

 "github.com/hidori/go-gontext"
)

func main() {
 production()
 test()
 production()
 test()
}

func production() {
 ctx := context.Background()
 ctx = contexttime.SetDefaultTime(ctx)

 fmt.Println(contexttime.GetTime(ctx).Now())
}

var _ = contexttime.Time((*MockTime)(nil))

type MockTime struct{}

func (t *MockTime) Now() time.Time {
 return time.Date(2025, 2, 21, 12, 34, 56, 789, time.UTC)
}

func test() {
 ctx := context.Background()
 ctx = contexttime.SetTime(ctx, &MockTime{})

 fmt.Println(contexttime.GetTime(ctx).Now())
}
```

OUTPUT:

```text
2025-02-26 01:36:57.317319169 +0900 JST m=+0.000065680
2025-02-21 12:34:56.000000789 +0000 UTC
2025-02-26 01:36:57.317429996 +0900 JST m=+0.000176489
2025-02-21 12:34:56.000000789 +0000 UTC
```
