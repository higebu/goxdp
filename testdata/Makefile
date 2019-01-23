LLVM_PREFIX ?= /usr/bin
CLANG ?= $(LLVM_PREFIX)/clang
LLC ?= $(LLVM_PREFIX)/llc
INCLUDEFLAGS = -I/usr/include

all: xdp_prog.elf

clean:
	-$(RM) *.elf

%.elf : %.o
	$(LLC) -march=bpf -filetype=obj -o $@ $<

%.o : %.c
	$(CLANG) $(INCLUDEFLAGS) -target bpf -O2 -emit-llvm -g \
		-Wall -Werror \
		-Wno-unused-variable \
		-Wno-compare-distinct-pointer-types \
		-Wno-implicit-function-declaration \
		-c $< -o $@
