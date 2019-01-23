package cmd

import (
	"bytes"
	"fmt"
	"strings"
	"testing"

	"github.com/vishvananda/netlink"
)

func setup(t *testing.T) func() {
	link := &netlink.Veth{LinkAttrs: netlink.LinkAttrs{Name: "test"}, PeerName: "testpeer"}
	if err := netlink.LinkAdd(link); err != nil {
		t.Fatal(err)
	}
	return func() {
		if err := netlink.LinkDel(link); err != nil {
			t.Fatal(err)
		}
	}
}

func TestAttachDetach(t *testing.T) {
	teardown := setup(t)
	defer teardown()
	cases := []struct {
		name    string
		command string
		want    string
	}{
		{name: "attach", command: "goxdp attach --device test --object ../testdata/xdp_prog.elf", want: "attached\n"},
		{name: "detach", command: "goxdp detach --device test", want: "detached\n"},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			buf := new(bytes.Buffer)
			opt := NewOptions()
			cmd := NewRootCmd(opt)
			cmd.SetOutput(buf)
			cmdArgs := strings.Split(c.command, " ")
			fmt.Printf("cmdArgs %+v\n", cmdArgs)
			cmd.SetArgs(cmdArgs[1:])
			cmd.Execute()

			get := buf.String()
			if c.want != get {
				t.Errorf("unexpected response: want:%+v, get:%+v", c.want, get)
			}
		})
	}
}
