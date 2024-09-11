package main

// Launch the program and execute the appropriate code
func main() {
	if len(route) > 1 {
		switch route[1] {
		case "-h", "--help":
			help()
		case "-c", "--core":
			serialize()
			core = compiler(access.Core)
			if len(core) > 0 {
				flag = "-w"
				rightplace()
				prepare()
				packagist(core)
			} else {
				journal("Checked for WordPress core updates, none found.")
			}
		case "-d", "--developer":
			serialize()
			dev = compiler(access.Dev)
			if len(dev) > 0 {
				flag = "-d"
				sift(dev)
			} else {
				journal("Checked for Developer plugin updates, none found.")
			}
		case "-p", "--premium":
			serialize()
			prem = compiler(access.Prem)
			if len(prem) > 0 {
				flag = "-p"
				sift(prem)
			} else {
				journal("Checked for Premium plugin updates, none found.")
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
				journal("Checked for WPackagist plugin updates, none found.")
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
