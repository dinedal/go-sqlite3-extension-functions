# go-sqlite3-extension-functions

## Install

```bash
mkdir ./build
cd ./build
cmake ..
make && make install
```

## What

This is the same file contributed by Liam Healy on 2010-02-06 15:45:07 at [https://www.sqlite.org/contrib?orderby=date](https://www.sqlite.org/contrib?orderby=date)

All this does is use `CMake` to create a cross-platform build that can be used in [go-sqlite3](https://github.com/mattn/go-sqlite3)

## Usage

Use like so:

```go
package main

import (
	"database/sql"

	sqlite3 "github.com/mattn/go-sqlite3"
	"github.com/dinedal/go-sqlite3-extension-functions/go"
)

func Main() {
	db, err := sql.Open("sqlite3-extension-functions", ":memory:")
}
```

Or use the code in [extension-functions.go](github.com/dinedal/go-sqlite3-extension-functions/go/blob/master/go/extension-functions.go) to create your own driver hook.
