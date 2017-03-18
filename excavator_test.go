package excavator

import "testing"

func TestGetHistoryWithInvalidShell(t *testing.T) {
    commands, err := GetHistory("invalidShell", 30)
    if err == nil && commands != nil {
        t.Error("Should have raised error for an invalid shell.")
    }
}

func TestGetHistoryWithBash(t *testing.T) {
    commands, err := GetHistory("bash", 30)
    if err !=  nil && commands == nil {
        t.Error("Should have returned comamnds for Bash shell.")
    }
}
