package executor

import (
	"fmt"
	"os/exec"

	"github.com/AtomJon/Powershell-REST-Server/resource"
)

func ExecutePowershell(resource resource.Resource) (string, error) {
	cmd := exec.Command("powershell.exe", "-c", resource.Content);

	result, err := cmd.CombinedOutput();
	if (err != nil) {
		return "", fmt.Errorf("cannot execute powershell:\nError Pipe:%s\nError:%v", string(result), err);
	}

	return string(result), nil;
}