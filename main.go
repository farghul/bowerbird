package main

// Launch the program and execute the appropriate code
func main() {
	if len(route) > 1 {
		switch route[1] {
		case "-h", "--help":
			help()
		case "-i", "--implement":
			serialize()
			wpac = compiler(access.WPac)
			prem = compiler(access.Prem)
			dev = compiler(access.Dev)

			if len(dev) > 0 {
				flag = "-d"
				sift(dev)
			}
			if len(prem) > 0 {
				flag = "-p"
				sift(prem)
			}
			if len(wpac) > 0 {
				flag = "-w"
				rightplace()
				prepare()
				wpackagist()
			}
		case "-v", "--version":
			version()
		default:
			alert("Unknown flag detected -")
			help()
		}
	}
}
