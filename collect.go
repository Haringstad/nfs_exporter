package main

import (
	"os/exec"

        log "github.com/sirupsen/logrus"

)

func execCommand(execpath, address string) (string, bool) {
	params := []string{"-e", address}

	out, err := exec.Command(execpath, params...).Output()
	if err != nil {
		log.Errorf("Get NFS storage path failed: %v", err)
		return string(out), false
	}

	return string(out), true
}
