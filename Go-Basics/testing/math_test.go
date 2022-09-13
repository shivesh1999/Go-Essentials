package math

import "testing"

func TestAbs(t *testing.T) {
	// t.log("log printed here")
	// t.Fail()    // will show a test case failed in the results
	// t.FailNow() // t.fail() + safely exit without continuing
	// t.Error() // t.log() + t.Fail()
	// t.Fatal() // t.log() + t.failNow()
	// if Abs(-1) < 0 {
	// 	t.E("negative value found in abs()")
	// }
	// if Abs(0) < 0 {
	// 	t.Error("negative value found in abs()")
	// }
	// if Abs(1) < 0 {
	// 	t.Error("negative value found in abs()")
	// }
}

func TestAdd(t *testing.T) {

	got := Add(4, 6)
	want := 10

	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}
