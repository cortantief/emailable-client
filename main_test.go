package emailableclient

import (
	"os"
	"testing"
)

var emailable Emailable

func init() {
	emailable = NewEmailable(os.Getenv("EMAILABLE_API_KEY"))
}

func TestUniqueEmail(t *testing.T) {
	const email = "test@test.fr"
	v, err := emailable.VerifyEmail(email)
	if err != nil {
		t.Fatal(err)
	}
	if email != v.Email {
		t.Fail()
		return
	}
}

func TestBatchEmail(t *testing.T) {
	var emails = []string{
		"test@test.fr",
		"test2@test.fr",
	}
	v, err := emailable.VerifyBatch(emails)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := emailable.GetBatchStatus(v.Id); err != nil {
		t.Fatal(err)
	}
}
