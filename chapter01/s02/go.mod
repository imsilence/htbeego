module chapter01/s02

go 1.12

replace (
	golang.org/x/crypto => github.com/golang/crypto v0.0.0-20190426145343-a29dc8fdc734
	golang.org/x/net => github.com/golang/net v0.0.0-20190424112056-4829fb13d2c6
	golang.org/x/sync => github.com/golang/sync v0.0.0-20190423024810-112230192c58
	golang.org/x/sys => github.com/golang/sys v0.0.0-20190428183149-804c0c7841b5
	golang.org/x/text => github.com/golang/text v0.3.2
	golang.org/x/tools => github.com/golang/tools v0.0.0-20190428024724-550556f78a90
)

require github.com/astaxie/beego v1.11.1
