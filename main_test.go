package main

import "testing"

func TestValidateEmail(t *testing.T) {
	testsPass := []string{"test@test.de", "otto@company.co.uk"}
	testsFail := []string{"test@@test.de", "@test.de", "test.de", "test@o.k", "test@ok"}

	for _, v := range testsPass {
		if ValidateEmail(v) == false {
			t.Errorf("%s should be passing", v)
		}
	}

	for _, v := range testsFail {
		if ValidateEmail(v) == true {
			t.Errorf("%s should be failing", v)
		}
	}
}
