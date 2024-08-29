package main

// Launch the program and execute the appropriate code
func main() {
	switch route {
	case "-v", "--version":
		version()
	case "-h", "--help":
		about()
	case "-r", "--release":
		flag = "-r"
		changedir()
		prepare()
		released()
	default:
		driver()
		free = compiler("wpackagist")
		paid = compiler("premium")
		dev = compiler("bcgov")

		if len(free) > 0 {
			flag = "-p"
			changedir()
			prepare()
			wpackagist()
		}
		if len(paid) > 0 {
			flag = "-s"
			premium()
		}
		if len(dev) > 0 {
			flag = "-d"
		}
	}
}
