module test

go 1.20

require github.com/HP-AKCL/Go/learn/module/akcl_mod_cal v0.0.0

require (
	golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c // indirect
	rsc.io/quote v1.5.2 // indirect
	rsc.io/sampler v1.3.0 // indirect
)

replace github.com/HP-AKCL/Go/learn/module/akcl_mod_cal => ./module/
