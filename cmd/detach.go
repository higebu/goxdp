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
	"github.com/spf13/cobra"
	"github.com/vishvananda/netlink"
)

func NewDetachCmd(opt *Options) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "detach",
		Short: "Detach the XDP program from the device.",
		Long:  "Equivalent to: ip link set dev DEVICE xdp off",
		Run: func(cmd *cobra.Command, args []string) {
			detach(cmd, opt)
		},
	}
	return cmd
}

func detach(cmd *cobra.Command, opt *Options) {
	link, err := netlink.LinkByName(opt.Device)
	if err != nil {
		cmd.Printf("failed to get device %s: %s\n", opt.Device, err)
		return
	}
	if err := netlink.LinkSetXdpFd(link, -1); err != nil {
		cmd.Println(err)
		return
	}
	cmd.Println("detached")
}
