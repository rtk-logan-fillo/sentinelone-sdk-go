module github.com/rtk-logan-fillo/sentinelone-sdk-go

// needed for running examples and to play nicely with vscode
replace github.com/rtk-logan-fillo/sentinelone-sdk-go/sentinelone => ./sentinelone

go 1.18

require github.com/rtk-logan-fillo/sentinelone-sdk-go/sentinelone v0.0.0-00010101000000-000000000000

require (
	github.com/golang/protobuf v1.4.2 // indirect
	golang.org/x/net v0.0.0-20200822124328-c89045814202 // indirect
	golang.org/x/oauth2 v0.0.0-20210323180902-22b0adad7558 // indirect
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/protobuf v1.25.0 // indirect
)
