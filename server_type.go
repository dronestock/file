package main

const (
	// nolint: staticcheck
	serverTypeFtp    serverType = "ftp"
	serverTypeWebdav            = "webdav"
	serverTypeSsh               = "ssh"
)

type serverType string
