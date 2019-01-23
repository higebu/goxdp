package cmd

type Options struct {
	Device  string
	Object  string
	Section string
}

func NewOptions() *Options {
	return &Options{}
}
