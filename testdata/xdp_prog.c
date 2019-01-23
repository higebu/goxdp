#include <linux/bpf.h>
#include <linux/in.h>
#include <linux/if_ether.h>

#define SEC(NAME) __attribute__((section(NAME), used))

SEC("xdp")
int prog(struct xdp_md *ctx)
{
	void *data_end = (void *)(long)ctx->data_end;
	void *data = (void *)(long)ctx->data;
	struct ethhdr *eth = data;
	unsigned long nh_off;

	nh_off = sizeof(*eth);
	if (data + nh_off > data_end)
		return XDP_DROP;

	return XDP_PASS;
}

char _license[] SEC("license") = "MIT";
