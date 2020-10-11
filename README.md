# logctxr

logctxr is a helper package for the [logr minimal logging API](https://github.com/go-logr/logr) that enables a Logger
instance to be embedded in and retrieved from a Go context. 

Since the core of a program only needs to import `github.com/iand/logctxr` it reduces the dependency on a specific
logger implementation. Typically a program will create a root logger at startup and embed it into a context that is
passed through the rest of the program. Submodules of the program can derive new loggers using the 
standard `logr` interface.

Example usage:

```Go
package main

import (
    "github.com/iand/logctxr"
    "github.com/iand/logfmtr"
)


func init() {
    // Ensure that logctxr knows how to make new logger instances
    // if it needs to
    logctxr.NewLogger = logfmtr.New
}

func main() {
    // Create a root logger
    root := logfmtr.New().WithName("root").V(2)

    // Embed the logger in a context
    loggerCtx := logctxr.NewContext(context.Background(), root)

    // Pass the context to the other function
    other(loggerCtx)

}


// A function that uses a context
func other(ctx context.Context) {
    // Retrieve the logger from the context
    logger := logctxr.FromContext(ctx)
    logger.Info("the sun is shining")
}
```

## Author

* [Ian Davis](http://github.com/iand) - <http://iandavis.com/>

## License

This is free and unencumbered software released into the public domain. Anyone is free to 
copy, modify, publish, use, compile, sell, or distribute this software, either in source 
code form or as a compiled binary, for any purpose, commercial or non-commercial, and by 
any means. For more information, see <http://unlicense.org/> or the 
accompanying [`UNLICENSE`](UNLICENSE) file.
