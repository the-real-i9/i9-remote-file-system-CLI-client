package rfssession

import "fmt"

func printWorkDir() {
	fmt.Printf("/%s%s\n", user.Username, workPath)
}
