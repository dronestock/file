package main

const (
	serverTypeFtp serverType = iota
	serverTypeWebdav
	serverTypeSsh
)

type serverType uint8
