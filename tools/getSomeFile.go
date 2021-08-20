package tools

import (
	"fmt"
	"github.com/discoriver/building-tool-suite/pkg/file"
	"github.com/discoriver/building-tool-suite/pkg/log"
)

type GetSomeFileCommandFlag struct {
	FilePath string
}

func (cmd *GetSomeFileCommandFlag) GetSomeFile() {
	if bytes, err := file.GetFileBytesPtr(cmd.FilePath); err != nil {
		log.Error("GetSomeFile:: Error reading file bytes: %s", err)
	} else {
		fmt.Print(bytes)
	}
}
