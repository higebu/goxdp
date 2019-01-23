# goxdp

A example command line tool to attach/detach an XDP program.

## Install

```
go get -u github.com/higebu/goxdp
```

## Usage

```
sudo goxdp attach --device lo --object ./testdata/xdp_prog.elf
sudo goxdp detach --device lo
```
