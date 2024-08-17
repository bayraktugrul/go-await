# go-await [![](https://img.shields.io/badge/License-MIT-yellow.svg)](https://img.shields.io/badge/License-MIT-yellow.svg)

go-await lets you use to synchronize async operations. It is inspired from popular Java library 'awaitility [https://github.com/awaitility/awaitility]'

## Install

```bash
go get github.com/bayraktugrul/go-await
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

## Credits

* [Tuğrul Bayrak](https://github.com/bayraktugrul)
