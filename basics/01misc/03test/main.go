package main

const (
	_   = 1 << (10 * iota)
	kiB // 1024
	miB // 1048576
	giB // 1073741824
	tiB // 1099511627776 (превышает 1 << 32)
	piB // 1125899906842624 EiB // 1152921504606846976
	ziB // 1180591620717411303424 (превышает 1 << 64)
	yiB // 1208925819614629174706176
)

const (
	_   = 1 << (10 * iota)
	kB // 1024
	mB // 1048576
	gB // 1073741824
	tB // 1099511627776 (превышает 1 << 32)
	pB // 1125899906842624 EiB // 1152921504606846976
	zB // 1180591620717411303424 (превышает 1 << 64)
	yB // 1208925819614629174706176
)

func main() {

}
