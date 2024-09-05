package main

// Launch the program and execute the appropriate code
func main() {
	if len(route) > 1 {
		switch route[1] {
		case "-h", "--help":
			help()
		case "-i", "--implement":
			driver()
			wpac = compiler(access.Wpac)
			prem = compiler(access.Prem)
			dev = compiler(access.Dev)

			if len(dev) > 0 {
				flag = "-d"
			}
			if len(prem) > 0 {
				flag = "-p"
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
