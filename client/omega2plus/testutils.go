package omega2plus

import (
	"os/exec"
	"strings"
	"testing"
)

func cmdExec(t *testing.T, cmds []string) {
	t.Helper()
	for _, cmd := range cmds {
		t.Logf("[INFO] %s", cmd)
		args := strings.Split(cmd, " ")

		if err := exec.Command(args[0], args[1:]...).Run(); err != nil {
			t.Logf("[WARN] %s %v", cmd, err)
		}
	}
}
