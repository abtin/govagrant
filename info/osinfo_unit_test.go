// +build unit

package info

import (
	"testing"
)


func TestMd5sum(t *testing.T) {
	info := OSInfo{
			Platform: "Microsoft Windows 10 Pro",
			Family:   "Standalone Workstation",
			Version:  "10.0.19042 Build 19042",
		}
	want := "IUew7yjW55KX9AdGgNmbAg=="
	got := info.Md5sum()
	if  got != want {
		t.Errorf("\nwant:\t%q\ngot:\t%q\n", want, got)
	}
}
