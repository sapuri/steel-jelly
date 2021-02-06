# steel-jelly
A Go library and CLI tool for porn sites

## Installation
### Library
```
go get github.com/sapuri/steel-jelly/steeljelly
```

### CLI
```
go get github.com/sapuri/steel-jelly/cmd/steeljelly
```

(optional) To run unit tests:
```
make generate
make test
```

## Examples
### Library
```go
package main

import (
	"fmt"
	"log"

	"github.com/sapuri/steel-jelly/steeljelly"
)

func main() {
	const videoURL = "https://jp.pornhub.com/view_video.php?viewkey=ph5f756e8a650b3"

	client := steeljelly.NewClient()
	res, err := client.GetThumbnailURLs(steeljelly.SiteTypePornhub, videoURL)
	if err != nil {
		log.Fatal(err)
	}

	// https://ei.phncdn.com/videos/202010/01/356624402/original/(m=eaAaGwObaaaa)(mh=3zxiu3wi3w-_5ZlG)1.jpg
	fmt.Println(res[0])
}
```

### CLI
Usage
```
steeljelly

NAME:
   steeljelly - A CLI tool for porn sites

USAGE:
   steeljelly [global options] command [command options] [arguments...]

VERSION:
   0.0.1

COMMANDS:
   eroterest  エロタレスト
   pornhub    Pornhub
   help, h    Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

(ex.) Retrieve thumbnail URLs of Pornhub video
```bash
steeljelly pornhub get-thumbnails --url https://jp.pornhub.com/view_video.php\?viewkey\=ph5f756e8a650b3
# https://ei.phncdn.com/videos/202010/01/356624402/original/(m=eaAaGwObaaaa)(mh=3zxiu3wi3w-_5ZlG)1.jpg
# ...
```
