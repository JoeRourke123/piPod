package logger

import (
	"context"
	"fmt"
	"github.com/jomei/notionapi"
	"os"
)

func Info(ctx context.Context, message string, tags ...string) {
	logEvent(ctx, message, "Info", tags...)
}

func Error(ctx context.Context, message string, err error, tags ...string) {
	logEvent(ctx, message+err.Error(), "Error", tags...)
}

func Warn(ctx context.Context, message string, tags ...string) {
	logEvent(ctx, message, "Warn", tags...)
}

func logEvent(ctx context.Context, message string, severity string, tags ...string) {
	properties := make(map[string]notionapi.Property)
	properties["Message"] = notionapi.TitleProperty{Title: []notionapi.RichText{{Text: &notionapi.Text{Content: message}}}}
	properties["Level"] = notionapi.SelectProperty{Select: notionapi.Option{Name: severity}}
	properties["Tags"] = fetchTags(tags)

	fmt.Println("[", severity, "]", message, "\n", tags, "\n----")

	//createRequest := notionapi.PageCreateRequest{
	//	Properties: properties,
	//	Parent: notionapi.Parent{
	//		Type:       notionapi.ParentTypeDatabaseID,
	//		DatabaseID: dbId,
	//	},
	//}

	//_, err := client.Page.Create(ctx, &createRequest)

	//if err != nil {
	//	log.Println("could not log to notion: ", err)
	//}
}

const (
	dbId = "63e5236180444d56945aad561d42903e"
)

var (
	client = notionapi.NewClient(notionapi.Token(os.Getenv("NOTION_SECRET")))
)
