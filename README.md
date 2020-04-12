# LogX

A concurrent log framework in Golang

Each logging statement is executed concurrently in a goroutine. So if we have write to file or stdout operations, main thread is not blocked. 

Custom formatters are supported. See logx/formatters/JSONFormatterFn for example.

## Features in pipeline

Signals (Channels)

Custom Writers 

More Custom Formatters OOTB

## History

0.20.3 (Initial Release)

Initial (0.20.3) version does not implement signals. It uses sync.WaitGroup{}

## Packages Overview

### logx

logx

# Examples

Refer tests. 

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
