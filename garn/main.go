package main

import (
	"fmt"
	"github.com/abkshvn/garn/resource_manager"
)

//#TODO(abkshvn):remove this and add tests
func main() {
	r := resource_manager.ResourceManager{"localhost", 8088}
	info, err := r.ClusterInformation()
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(info.State)
}
