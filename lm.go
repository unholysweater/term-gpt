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
    // Create a new buffered reader from standard input
    reader := bufio.NewReader(os.Stdin)
    
    // Create initial thread instance outside of the prompt loop
    myThread := thread.New()

    for {
        // Prompt
        fmt.Print("Ask: ")
        userInput, err := reader.ReadString('\n')
        if err != nil {
            fmt.Println("Error reading input: ", err)
            continue // Skip to next iteration of the loop
        }

        // Trim newline char from the input
        userInput = userInput[:len(userInput)-1]

        // Continue until exit is entered
        if userInput == "exit" {
            break
        }
        
        // Add message to the already initialized thread
	    myThread.AddMessage(
		    thread.NewUserMessage().AddContent(
			    thread.NewTextContent(userInput),
		    ),
	    )

        // Generate OpenAI response
        err = openai.New().Generate(context.Background(), myThread)
	    if err != nil {
            fmt.Println("Error generating response: ", err)
            continue // Skip to next iteration of the loop
	    }

        // Print response
	    fmt.Println(myThread)
    }
}
