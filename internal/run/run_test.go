package run

import (
	"testing"
)

func TestRunCmdNoArgs(t *testing.T) {
	err := RunCmd([]string{})
	if err == nil {
		t.Error("RunCmd with no args should return error")
	}
}

func TestRunCmdValidCommand(t *testing.T) {
	err := RunCmd([]string{"echo", "test"})
	if err != nil {
		t.Errorf("RunCmd with valid command failed: %v", err)
	}
}

func TestRunCmdInvalidCommand(t *testing.T) {
	err := RunCmd([]string{"nonexistent-command-xyz-123"})
	if err == nil {
		t.Error("RunCmd with invalid command should return error")
	}
}

func TestRunCmdWithArgs(t *testing.T) {
	err := RunCmd([]string{"echo", "hello", "world"})
	if err != nil {
		t.Errorf("RunCmd with multiple args failed: %v", err)
	}
}
