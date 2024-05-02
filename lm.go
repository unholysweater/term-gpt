package main

import (
	"context"
	"fmt"
    "bufio"
    "os"

	"github.com/henomis/lingoose/llm/openai"
	"github.com/henomis/lingoose/thread"
)

func main() {
    // Create the thread on session start
    reader := bufio.NewReader(os.Stdin)
    myThread := thread.New()

    for {
        // Prompt for input
        fmt.Print("# ")
        userInput, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input: ", err)
            continue
        }

        // Trim newline char from the input
        userInput = userInput[:len(userInput)-1]

        // Continue until exit is entered (alternatively just use C-c)
        if userInput == "exit" {
            break
        }
       
        // Append message to thread object
	    myThread.AddMessage(
		    thread.NewUserMessage().AddContent(
			    thread.NewTextContent(userInput),
		    ),
	    )

        fmt.Println("-----------")

        // Generate OpenAI response, append to thread object
        err = openai.New().Generate(context.Background(), myThread)
	    if err != nil {
            fmt.Println("Error generating response: ", err)
            continue
	    }

	    fmt.Println(myThread)
    }
}
