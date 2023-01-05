package main

const (
	serverTypeFtp serverType = "ftp"
	serverTypeWebdav = "webdav"
	serverTypeSsh = "ssh"
)

type serverType string
