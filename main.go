package main

import (
  "fmt"
  "os"
  "time"
  "net/http"
  "encoding/json"

  "github.com/codegangsta/cli"
  "github.com/gin-gonic/gin"
)

func main() {
  var httpPort string
  var wsPort string
  var initPeers string

  app := cli.NewApp()
  app.Name = "go-naivechain"
  app.Usage = "naivechain written in Go"
  app.Version = "0.0.1"

  app.Flags = []cli.Flag {
    cli.StringFlag {
      Name:  "httpPort",
      Value: "8000",
      Usage: "HTTP port",
      Destination: &httpPort,
    },
    cli.StringFlag {
      Name:  "wsPort",
      Value: "8001",
      Usage: "WebSocket port",
      Destination: &wsPort,
    },
    cli.StringFlag {
      Name:  "peers",
      Value: "",
      Usage: "Initial peers",
      Destination: &initPeers,
    },
  }

  app.Action = func(c *cli.Context) error {
    fmt.Println("HTTP Port:", httpPort)
    fmt.Println("WebSocket Port:", wsPort)
    fmt.Println("Initial Peers:", initPeers)

    // server run!
    runServer()

    return nil
  }

  app.Run(os.Args)
}

// Block
type Block struct {
  Index        int
  PreviousHash string
  Timestamp    time.Time
  Data         string
  Hash         string
}

func runServer() {
  genesisBlock := Block {
    Index:        0,
    PreviousHash: "0",
    Timestamp:    time.Now(),
    Data:         "Genesis Block",
    Hash:         "816534932c2b7154836da6afc367695e6337db8a921823784c14378abed4f7d7",
  }

  blockChain := []Block {
    genesisBlock,
  }

  fmt.Println(blockChain)

  r := gin.Default()

  r.GET("/ping", func(c *gin.Context) {
    c.String(http.StatusOK, "Pong!")
  })

  r.GET("/blocks", func(c *gin.Context) {
    b, _ := json.Marshal(blockChain)
    c.String(http.StatusOK, string(b))
  })

  r.POST("/mineBlock", func(c *gin.Context) {
    // Block 作成

    // blockChain に追加

    // broadcastする

    c.String(http.StatusOK, "Added block!")
  })

  r.GET("/peers", func(c *gin.Context) {
    // TODO
  })

  r.POST("/addPeer", func(c *gin.Context) {
    // TODO
  })

  r.Run()
}
