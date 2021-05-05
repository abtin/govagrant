// +build integration

package info

import (
	"reflect"
	"runtime"
	"testing"
)

func TestNewOsInfo_windows10(t *testing.T) {
	if runtime.GOOS != "windows" {
		t.Skip("Skipping TestNewOsInfo_windows10 on a non windows host")
		return
	}

	want := OSInfo{
		Platform: "Microsoft Windows 10 Pro",
		Family:   "Standalone Workstation",
		Version:  "10.0.16299 Build 16299",
	}
	got, err := NewOsInfo()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nwanted:\t%+v\ngot:\t%+v\n", want, got)
		t.Fail()
	}
}

func TestNewOsInfo_linux(t *testing.T) {
	if runtime.GOOS != "linux" {
		t.Skip("Skipping TestNewOsInfo_linux_ubuntu on a non linux host")
		return
	}
	want := OSInfo{
		Platform: "centos",
		Family:   "rhel",
		Version:  "8.3.2011",
	}

	got, err := NewOsInfo()
	if err != nil {
		t.Error(err)
		t.Fail()
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("\nwanted:\t%+v\ngot:\t%+v\n", want, got)
		t.Fail()
	}
}
