# Coin Toss

Coin Toss is a single-purpose package that will randomly return a 0 or 1. It comes with an executable, `toss` which will write the result of the flip as, again, a 0 or 1. Perfect when you want to randomly fail a build script step with a non-zero return value. Prank your friends and co-workers.

## How to Get

If you're new to Go, and especially Go on Windows, I've [written about what I've learned so far](https://hyr.mn/go-structure-windows/). 

You can download the source using `go get`. From the command line, run

```
> go get github.com/hyrmn/cointoss/cmd/toss
```

This will download the entire repository and place it within your `%GOPATH%\src` directory

Next, if you want, you can have Go compile the program and copy it to the `%GOPATH%\bin` directory

```
> go install github.com/hyrmn/cointoss/cmd/toss
```

Now you'll have an executable named `toss` (`toss.exe` on Windows).

## How to Use

```
> toss
```

## Runtime Considerations

This uses Golang's `rand` package. As such, it is not suitable for cryptographic applications. But, then, if you're using a random 0 or 1 with a 50/50 distribution over enough iterations, you're probably not thinking of cryptographic applications 😜

## Wait, What?

Ok, this isn't really a serious project. Although the fact that only returns a 0 or 1 to stdout does make it suitable to chain with other CLI tools. 

## What's Next

I'm going to include a web server and Docker deployment target. Why? Because.
