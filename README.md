# LogX

A concurrent log framework in Golang

Each logging statement is executed concurrently in a goroutine. So if we have write to file or stdout operations, main thread is not blocked. 

Custom formatters are supported. See logx/formatters/JSONFormatterFn for example.

## Features in pipeline

* Signals (Channels)

* Custom Writers - can even be used for pushing logs to an external system

* More Custom Formatters OOTB

## History

### 0.20.4 

* Channels and Signals

(Inspired by Michael Van Sickle's codecamp (https://www.youtube.com/watch?v=YS4e4q9oBaU&t=21910s))

* Caller stack support

### 0.20.3 (Initial Release)

Initial version does not implement signals. It uses sync.WaitGroup{} only

## Packages Overview

### logx

logx

# Examples

Refer tests

TODO: Add more tests

# Usage 

TODO

## Install package

```shell
go get https://github.com/mayukh42/logx
```

## Use Package 'lambda'

```shell
import github.com/mayukh42/logx/logx
```
