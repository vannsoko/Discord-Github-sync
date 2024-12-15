package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"encoding/json"
	"io"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	_ "net/http/pprof"
)

type GithubWebhook struct {
	Action	string	`json:"action"`
	Discussion Discussion `json:"discussion"`
	Comment struct {Body string `json:"body"`} `json:"comment"`
	Repository struct {Name string `json:"name"`} `json:"repository"`	
	Sender Sender `json:"sender"`
	Changes struct {Body struct {From string `json:"from"`} `json:"body"`} `json:"changes"`
}
type Discussion struct {
	ID int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
	Category Category `json:"category"`
	User User `json:"user"`
}
type Category struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}
type User struct {
	Login string `json:"login"`
	Type string `json:"type"`
}
type Sender struct {
	Login string `json:"login"`
	Type string `json:"type"`
	ID int `json:"id"`
}

const DEBUG bool = true

func main() {
	if !DEBUG {
		log.SetOutput(io.Discard)
	}
    // Load environment variables
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found")
    }

    // Initialize Discord Bot
    discordToken := os.Getenv("DISCORD_TOKEN")



    discord.InitBot(discordToken)
    defer discord.CloseBot()

    // Setup Gin Router
    router := gin.Default()

	router.POST("/github-webhook", handleGithubWebhook)


    // Start Gin server
    go func() {
        if err := router.Run(":8080"); err != nil {
            log.Fatalln("Failed to start server:", err)
        }
    }()
	
    // Wait for termination signal
    stop := make(chan os.Signal, 1)
    signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
    <-stop

    log.Println("Shutting down gracefully")
}

func handleGithubWebhook(c *gin.Context) {
	body, err := io.ReadAll(c.Request.Body)	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Unable to read request body",
			"details": err.Error(),
		})
		return
	}
	var githubwebhook GithubWebhook
	err = json.Unmarshal(body, &githubwebhook)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid JSON format",
			"details": err.Error(),
		})
		return
	}

	switch githubwebhook.Action {
	case "created":
		
	case "edited":
		
	case "deleted":
		 
	}
}

// TODO using sender ID ???
