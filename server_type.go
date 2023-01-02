package main

const (
	serverTypeFtp serverType = iota
	serverTypeWebdav
)

type serverType uint8
