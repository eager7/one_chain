contract:=MultiSigWalletWithDailyLimit
pkg:=msw

default:go

sol:
	solc --abi ${contract}.sol | awk '/JSON ABI/{x=1;next}x' > ${pkg}.json
	solc --bin ${contract}.sol | awk '/Binary:/{x=1;next}x' > ${pkg}.bin

go:
	abigen --bin=${pkg}.bin --abi=${pkg}.json --pkg=${pkg} --out=${pkg}.go

clean:
	rm ${pkg}.json
	rm ${pkg}.bin
	rm ${pkg}.go