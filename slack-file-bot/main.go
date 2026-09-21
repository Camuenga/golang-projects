/*
package main

import (

	"fmt"
	"os"
	//"github.com/slack-io/slacker"
	"github.com/slack-go/slack"

)

func main() {

	api := slack.New(os.Getenv("SLACK_BOT_TOKEN"))
	channelArr := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"Effective-Modern-C++.pdf"}

	for i := 0; i < len(fileArr); i++ {
		params := slack.UploadFileParameters{
			Channels: channelArr,
			File:     fileArr[i],
			Filename: "Effective-Modern-C++.pdf",
			FileSize: 43,
		}
		file, err := api.UploadFile(params)
		if err != nil {
			fmt.Printf("%s \n", err)
			return
		}
		fmt.Printf("Name: %s \n", file.Title)
	}

}
*/
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/slack-go/slack"
)

func main() {

	token := os.Getenv("SLACK_BOT_TOKEN")
	channelid := []string{os.Getenv("CHANNEL_ID")}
	fileArr := []string{"test.txt"}

	if token == "" {
		fmt.Println("SLACK_BOT_TOKEN environment variable is required")
		os.Exit(1)
	}
	api := slack.New(token)

	ctx := context.Background()

	// Upload a file
	params := slack.UploadFileParameters{
		Channels: channelid,
		Filename: "test.txt",
		File:     fileArr[0],
		FileSize: 43,
	}
	file, err := api.UploadFileContext(ctx, params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}
	fmt.Printf("ID: %s, title: %s\n", file.ID, file.Title)

}
