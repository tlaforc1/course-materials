package scanner

import (
	"testing"
)

// THESE TESTS ARE LIKELY TO FAIL IF YOU DO NOT CHANGE HOW the worker connects (e.g., you should use DialTimeout)
func TestOpenPort(t *testing.T){
	testVar := 23
	want := 0
    open, _ := PortScanner(testVar) // Currently function returns only number of open ports
	if testVar >= 22{
		want = 1
	}	// default value when passing in 1024 TO scanme; also only works because currently PortScanner only returns 
	      //consider what would happen if you parameterize the portscanner address and ports to scan

    if open != want {
        t.Errorf("got %d, wanted %d", open, want)
    }
}

func TestTotalPortsScanned(t *testing.T){
	// THIS TEST WILL FAIL - YOU MUST MODIFY THE OUTPUT OF PortScanner()
	testVar := 20
    open, closed := PortScanner(testVar) // Currently function returns only number of open ports
    want := testVar // default value; consider what would happen if you parameterize the portscanner ports to scan

    if open+closed != want {
        t.Errorf("got %d, wanted %d", open+closed, want)
    }
}
