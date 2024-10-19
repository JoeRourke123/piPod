package logger

import (
	"github.com/jomei/notionapi"
	"runtime"
)

const (
	DbWriteTag          = "db:write"
	DbReadTag           = "db.read"
	OperationFetchTag   = "api:op:fetch"
	ServiceOperationTag = "service:op"
	SocketOperationTag  = "socket:op"
)

func FromTag(from string) string {
	return "from:" + from
}

func ApiTag(service string, endpoint string) string {
	return "api:" + service + ":" + endpoint
}

func DeviceTag(deviceId string) string {
	return "device:" + deviceId
}

func UriTag(spotifyUri string) string {
	return "uri:" + spotifyUri
}

func fetchTags(tags []string) notionapi.MultiSelectProperty {
	tags = append(tags, "os:"+runtime.GOOS)
	tags = append(tags, "")
	return notionapi.MultiSelectProperty{MultiSelect: make([]notionapi.Option, 0)}
}
