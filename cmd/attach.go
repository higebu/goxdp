// Copyright © 2019 Yuya Kusakabe <yuya.kusakabe@gmail.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"os"

	"github.com/newtools/ebpf"
	"github.com/spf13/cobra"
	"github.com/vishvananda/netlink"
)

// attachCmd represents the attach command
func NewAttachCmd(opt *Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "attach",
		Short: "Attach the XDP program to the device.",
		Long:  "Equivalent to: ip link set dev DEVICE xdp object FILE section NAME",
		Run: func(cmd *cobra.Command, args []string) {
			attach(cmd, opt)
		},
	}

	cmd.Flags().StringVarP(&opt.Object, "object", "o", "", "elf object file path")
	cmd.MarkFlagRequired("object")
	cmd.Flags().StringVarP(&opt.Section, "section", "s", "prog", "prog name")

	return cmd
}

func attach(cmd *cobra.Command, opt *Options) {
	f, err := os.Open(opt.Object)
	if err != nil {
		cmd.Println(err)
		return
	}
	spec, err := ebpf.LoadCollectionSpecFromReader(f)
	if err != nil {
		cmd.Println(err)
		return
	}
	coll, err := ebpf.NewCollection(spec)
	if err != nil {
		cmd.Println(err)
		return
	}
	prog, ok := coll.Programs[opt.Section]
	if !ok {
		cmd.Printf("%s not found in object\n", opt.Section)
		return
	}

	link, err := netlink.LinkByName(opt.Device)
	if err != nil {
		cmd.Printf("failed to get device %s: %s\n", opt.Device, err)
		return
	}
	if err := netlink.LinkSetXdpFd(link, prog.FD()); err != nil {
		cmd.Println(err)
		return
	}
	cmd.Println("attached")
}
