# gopool


## Design

TaskPool 和 WorkerPool 全局都是只有一个，一个Pool对应多个Worker和多个Task。每个Worker都关联上P，然后加锁从P上获取task。加锁范围很小，不涉及task具体执行，所以锁竞争还好。 然后worker去执行他自身关联的Pool上的task

## Introduction

`gopool` is a high-performance goroutine pool which aims to reuse goroutines and limit the number of goroutines.

It is an alternative to the `go` keyword.

## Features

- High Performance
- Auto-recovering Panics
- Limit Goroutine Numbers
- Reuse Goroutine Stack

## QuickStart

Just replace your `go func(){...}` with `gopool.Go(func(){...})`.

old:
```go
go func() {
	// do your job
}()
```

new:
```go
gopool.Go(func(){
	/// do your job
})
```
