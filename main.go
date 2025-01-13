package main

// Launch the program and execute the appropriate code
func main() {
	if len(route) > 1 {
		switch route[1] {
		case "-c", "--core":
			serialize()
			core = compiler(access.Core)
			if len(core) > 0 {
				flag = "-w"
				rightplace()
				prepare()
				packagist(core)
			} else {
				journal("No WordPress core update tickets to process.")
			}
		case "-h", "--help":
			help()
		case "-p", "--premium":
			serialize()
			prem = compiler(access.Prem)
			if len(prem) > 0 {
				flag = "-p"
				sift(prem)
			} else {
				journal("No Premium plugin update tickets to process.")
			}
		case "-w", "--wpackagist":
			serialize()
			wpac = compiler(access.WPac)
			if len(wpac) > 0 {
				flag = "-w"
				rightplace()
				prepare()
				packagist(wpac)
			} else {
				journal("No WPackagist plugin update tickets to process.")
			}
		case "-v", "--version":
			version()
		default:
			alert("Unknown flag detected -")
			help()
		}
	} else {
		alert("No flag detected -")
	}
}
