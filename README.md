# Gravity

<img align="right" alt="Gravity" src="https://github.com/ekkx/gravity/assets/77382767/214a6ea8-f6c8-493f-b53f-dd030e884803" width="300">

Golang API wrapper for [Gravity](https://gravity.place/).

### Installation

```sh-session
go get github.com/ekkx/gravity
```

⚠️ This project is currently under construction. ⚠️

### Usage

Import the package into your project.

```go
import "github.com/ekkx/gravity"
```

Construct a new Gravity client.

```go
g, err := gravity.New("your_email", "your_password")
```

See Documentation and Examples below for more detailed information.

## Examples

Fetching feed contents after logging in with your email.

```go
package main

import "github.com/ekkx/gravity"

func main() {
  g, err := gravity.New("your_email", "your_password")
  if err != nil {
    return
  }

  // Home timeline
  feed, err := g.Feed.Popular()

  // Add comment to a feed
  resp, err := g.Feed.AddComment(&gravity.AddCommentParams{
    SID: ""
    Content: ""
  })
}

```

## Documentation

Coming soon!

<p align="center">
  <img alt="Gravity" src="https://github.com/ekkx/gravity/assets/77382767/245ee6b8-c4e8-48cf-aaaf-0c9a11a21929">
</p>
