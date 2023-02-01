package packageone

import "fmt"

var PackageVar = "This is the package level variable in packageOne"

func PrintMe(s1, s2 string) {
	fmt.Println(s1, s2, PackageVar)

}
