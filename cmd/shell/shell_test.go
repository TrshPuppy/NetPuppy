package shell

import (
	"testing"
)

func TestCBShellGetter(t *testing.T) {
	var fakeShellGetter TestShellGetter

	//expected := reflect.TypeOf(ShellInterface)
	actual, err := fakeShellGetter.GetConnectBackInitiatedShell()

	if err != nil {
		t.Errorf("Test Shell Getter returned error: %v\n", err)
	}

	_, success := actual.(*TestShell)
	if !success {
		t.Errorf("Test Shell Getter Failed - Expected: *TestShell, Actual: %v\n", actual)
	}
}
