# go-await [![GoDoc][doc-img]][doc] [![][workflow-badge]][workflow-actions] [![Release][release-badge]][release] [![][license-badge]][license]

go-await lets you use to synchronize async operations. It is inspired from popular Java library 'awaitility [https://github.com/awaitility/awaitility]'

## Install
With [Go's module support](https://go.dev/wiki/Modules#how-to-use-modules), `go [build|run|test]` automatically fetches the necessary dependencies when you add the import in your code:

```sh
import "github.com/bayraktugrul/go-await"
```

Alternatively, use `go get`:

```sh
go get -u github.com/bayraktugrul/go-await
```

## Examples

### Await with default parameters

```go
producer.publishMessage(orderCreatedMessage)

// waits for up to 1 second until repo.Exist returns true, using a default polling interval of 100ms
err := await.New().Await(func() bool {
    return repo.Exist(id)
})

if err != nil {}
```

### Await with custom parameters

```go

producer.publishMessage(orderCreatedMessage)

// waits for up to 2 seconds until repo.Exist returns true, using a polling interval of 500ms
err := await.New().PollInterval(500 * time.Millisecond).AtMost(2 * time.Second).Await(func() bool {
    return repo.Exist(id)
})

if err != nil {}
```

### Await with custom poll interval strategy

The default polling strategy is Fixed, returning the same interval for each iteration. You can customize this by setting a different poll interval strategy.
```go

producer.publishMessage(orderCreatedMessage)

// waits with double poll interval strategy 
// for up to 2 seconds until repo.Exist returns true, 
// using a polling interval of 500ms
err := await.New().
    PollStrategy(poll.Double).
    PollInterval(100 * time.Millisecond).
    AtMost(2 * time.Second).
    Await(func() bool {
        return repo.Exist(id) // any condition you want to wait for until it returns true
    })

if err != nil {}
```

## Credits
* [Tuğrul Bayrak](https://github.com/bayraktugrul)

[doc-img]: https://godoc.org/github.com/bayraktugrul/go-await?status.svg
[doc]: https://godoc.org/github.com/bayraktugrul/go-await

[workflow-actions]: https://github.com/bayraktugrul/go-await/actions
[workflow-badge]: https://github.com/bayraktugrul/go-await/workflows/build/badge.svg

[license]:https://github.com/bayraktugrul/go-await/blob/main/LICENSE
[license-badge]:https://img.shields.io/badge/License-MIT-blue.svg

[release]: https://github.com/bayraktugrul/go-await/releases
[release-badge]: https://img.shields.io/github/v/release/bayraktugrul/go-await.svg