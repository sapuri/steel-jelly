# steel-jelly

## Installation
```
go get github.com/sapuri/steel-jelly/steeljelly
```

(optional) To run unit tests:
```
make generate
make test
```

## Examples
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
