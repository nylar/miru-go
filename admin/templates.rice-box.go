package admin

import (
	"github.com/GeertJohan/go.rice/embedded"
	"time"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    `add.html`,
		FileModTime: time.Unix(1424610480, 0),
		Content:     string([]byte{0x3c, 0x21, 0x44, 0x4f, 0x43, 0x54, 0x59, 0x50, 0x45, 0x20, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x20, 0x6c, 0x61, 0x6e, 0x67, 0x3d, 0x22, 0x65, 0x6e, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x68, 0x65, 0x61, 0x64, 0x3e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x41, 0x64, 0x64, 0x20, 0x53, 0x69, 0x74, 0x65, 0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x22, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x22, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x63, 0x73, 0x73, 0x22, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2f, 0x63, 0x73, 0x73, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x73, 0x73, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6e, 0x61, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x75, 0x6c, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x77, 0x72, 0x61, 0x70, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x69, 0x3e, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x64, 0x64, 0x2f, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x3e, 0x41, 0x64, 0x64, 0x20, 0x53, 0x69, 0x74, 0x65, 0x3c, 0x2f, 0x61, 0x3e, 0x3c, 0x2f, 0x6c, 0x69, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x69, 0x3e, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x22, 0x3e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x3c, 0x2f, 0x61, 0x3e, 0x3c, 0x2f, 0x6c, 0x69, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x75, 0x6c, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x6e, 0x61, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x61, 0x6c, 0x65, 0x72, 0x74, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7b, 0x7b, 0x69, 0x66, 0x20, 0x2e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x7d, 0x7d, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x70, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x73, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x22, 0x3e, 0x53, 0x75, 0x63, 0x63, 0x65, 0x73, 0x73, 0x66, 0x75, 0x6c, 0x6c, 0x79, 0x20, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x65, 0x64, 0x3c, 0x62, 0x72, 0x20, 0x2f, 0x3e, 0x3c, 0x62, 0x72, 0x20, 0x2f, 0x3e, 0x7b, 0x7b, 0x2e, 0x55, 0x72, 0x6c, 0x7d, 0x7d, 0x3c, 0x2f, 0x70, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7b, 0x7b, 0x65, 0x6c, 0x73, 0x65, 0x7d, 0x7d, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x70, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x22, 0x3e, 0x46, 0x61, 0x69, 0x6c, 0x65, 0x64, 0x20, 0x74, 0x6f, 0x20, 0x63, 0x72, 0x61, 0x77, 0x6c, 0x3c, 0x62, 0x72, 0x20, 0x2f, 0x3e, 0x3c, 0x62, 0x72, 0x20, 0x2f, 0x3e, 0x7b, 0x7b, 0x2e, 0x55, 0x72, 0x6c, 0x7d, 0x7d, 0x3c, 0x2f, 0x70, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x7b, 0x7b, 0x65, 0x6e, 0x64, 0x7d, 0x7d, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}
	file3 := &embedded.EmbeddedFile{
		Filename:    `new.html`,
		FileModTime: time.Unix(1424561561, 0),
		Content:     string([]byte{0x3c, 0x21, 0x44, 0x4f, 0x43, 0x54, 0x59, 0x50, 0x45, 0x20, 0x68, 0x74, 0x6d, 0x6c, 0x3e, 0xa, 0x3c, 0x68, 0x74, 0x6d, 0x6c, 0x20, 0x6c, 0x61, 0x6e, 0x67, 0x3d, 0x22, 0x65, 0x6e, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x68, 0x65, 0x61, 0x64, 0x3e, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0x41, 0x64, 0x64, 0x20, 0x53, 0x69, 0x74, 0x65, 0x3c, 0x2f, 0x74, 0x69, 0x74, 0x6c, 0x65, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x69, 0x6e, 0x6b, 0x20, 0x72, 0x65, 0x6c, 0x3d, 0x22, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x73, 0x68, 0x65, 0x65, 0x74, 0x22, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x74, 0x65, 0x78, 0x74, 0x2f, 0x63, 0x73, 0x73, 0x22, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x73, 0x74, 0x61, 0x74, 0x69, 0x63, 0x2f, 0x63, 0x73, 0x73, 0x2f, 0x6d, 0x61, 0x69, 0x6e, 0x2e, 0x63, 0x73, 0x73, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x68, 0x65, 0x61, 0x64, 0x3e, 0xa, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6e, 0x61, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x75, 0x6c, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x77, 0x72, 0x61, 0x70, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x69, 0x3e, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x61, 0x64, 0x64, 0x2f, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x61, 0x63, 0x74, 0x69, 0x76, 0x65, 0x22, 0x3e, 0x41, 0x64, 0x64, 0x20, 0x53, 0x69, 0x74, 0x65, 0x3c, 0x2f, 0x61, 0x3e, 0x3c, 0x2f, 0x6c, 0x69, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x69, 0x3e, 0x3c, 0x61, 0x20, 0x68, 0x72, 0x65, 0x66, 0x3d, 0x22, 0x2f, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x71, 0x75, 0x65, 0x75, 0x65, 0x2f, 0x22, 0x3e, 0x51, 0x75, 0x65, 0x75, 0x65, 0x3c, 0x2f, 0x61, 0x3e, 0x3c, 0x2f, 0x6c, 0x69, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x75, 0x6c, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x6e, 0x61, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x66, 0x6f, 0x72, 0x6d, 0x20, 0x6d, 0x65, 0x74, 0x68, 0x6f, 0x64, 0x3d, 0x22, 0x50, 0x4f, 0x53, 0x54, 0x22, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x77, 0x72, 0x61, 0x70, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x20, 0x66, 0x6f, 0x72, 0x3d, 0x22, 0x69, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x3e, 0x53, 0x69, 0x74, 0x65, 0x20, 0x55, 0x52, 0x4c, 0x3c, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x75, 0x72, 0x6c, 0x22, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x75, 0x72, 0x6c, 0x22, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x69, 0x64, 0x5f, 0x75, 0x72, 0x6c, 0x22, 0x20, 0x70, 0x6c, 0x61, 0x63, 0x65, 0x68, 0x6f, 0x6c, 0x64, 0x65, 0x72, 0x3d, 0x22, 0x68, 0x74, 0x74, 0x70, 0x3a, 0x2f, 0x2f, 0x65, 0x78, 0x61, 0x6d, 0x70, 0x6c, 0x65, 0x2e, 0x63, 0x6f, 0x6d, 0x22, 0x20, 0x72, 0x65, 0x71, 0x75, 0x69, 0x72, 0x65, 0x64, 0x2f, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3e, 0xa, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x20, 0x66, 0x6f, 0x72, 0x3d, 0x22, 0x69, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x3e, 0x4c, 0x69, 0x6e, 0x6b, 0x73, 0x20, 0x74, 0x6f, 0x20, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x3c, 0x2f, 0x73, 0x74, 0x72, 0x6f, 0x6e, 0x67, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x22, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0x74, 0x72, 0x75, 0x65, 0x22, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x69, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x30, 0x22, 0x20, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x3d, 0x22, 0x63, 0x68, 0x65, 0x63, 0x6b, 0x65, 0x64, 0x22, 0x3e, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x6f, 0x6e, 0x6c, 0x79, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x20, 0x74, 0x79, 0x70, 0x65, 0x3d, 0x22, 0x72, 0x61, 0x64, 0x69, 0x6f, 0x22, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x3d, 0x22, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x22, 0x20, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3d, 0x22, 0x66, 0x61, 0x6c, 0x73, 0x65, 0x22, 0x20, 0x69, 0x64, 0x3d, 0x22, 0x69, 0x64, 0x5f, 0x6f, 0x6e, 0x6c, 0x79, 0x5f, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x31, 0x22, 0x3e, 0x20, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x20, 0x26, 0x61, 0x6d, 0x70, 0x3b, 0x20, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x64, 0x69, 0x76, 0x20, 0x63, 0x6c, 0x61, 0x73, 0x73, 0x3d, 0x22, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x74, 0x22, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x3e, 0x41, 0x64, 0x64, 0x20, 0x53, 0x69, 0x74, 0x65, 0x3c, 0x2f, 0x62, 0x75, 0x74, 0x74, 0x6f, 0x6e, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x64, 0x69, 0x76, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x66, 0x6f, 0x72, 0x6d, 0x3e, 0xa, 0x20, 0x20, 0x20, 0x20, 0x3c, 0x2f, 0x62, 0x6f, 0x64, 0x79, 0x3e, 0xa, 0x3c, 0x2f, 0x68, 0x74, 0x6d, 0x6c, 0x3e}), //++ TODO: optimize? (double allocation) or does compiler already optimize this?
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   ``,
		DirModTime: time.Unix(1424611205, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // add.html
			file3, // new.html

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`templates`, &embedded.EmbeddedBox{
		Name: `templates`,
		Time: time.Unix(1424611210, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"": dir1,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"add.html": file2,
			"new.html": file3,
		},
	})
}
