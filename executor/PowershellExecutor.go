package executor

import (
	"fmt"
	"os/exec"

	"github.com/AtomJon/subscriptrestserver/resource"
)

func ExecutePowershell(resource resource.Resource) (string, error) {
	cmd := exec.Command("powershell.exe", "-c", resource.Content);

	result, err := cmd.CombinedOutput();
	if (err != nil) {
		return "", fmt.Errorf("powershell:\n\tError Pipe: %s\n\tError: %v", string(result), err);
	}

	return string(result), nil;
}