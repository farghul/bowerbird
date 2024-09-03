package main

// Launch the program and execute the appropriate code
func main() {
	if len(route) > 1 {
		switch route[1] {
		case "-h", "--help":
			help()
		case "-u", "--update":
			driver()
			free = compiler(access.Free)
			paid = compiler(access.Paid)
			dev = compiler(access.Dev)

			if len(free) > 0 {
				flag = "-p"
				rightplace()
				prepare()
				wpackagist()
			}
			if len(paid) > 0 {
				flag = "-s"
			}
			if len(dev) > 0 {
				flag = "-d"
			}
		case "-v", "--version":
			version()
		default:
			alert("Unknown flag detected -")
			help()
		}
	}
}
