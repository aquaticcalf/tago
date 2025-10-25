module example/build

go 1.25.3

require (
	aqclf.xyz/tago v0.0.0
	aqclf.xyz/tago/css v0.0.0
)

replace (
	aqclf.xyz/tago => ../..
	aqclf.xyz/tago/css => ../../css
)
