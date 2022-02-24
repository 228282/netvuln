package v1

import (
	"fmt"

	api "netvuln/api"

	_ "github.com/Ullaakut/nmap"
)

func CheckVuln(req api.CheckVulnRequest) api.CheckVulnResponse {
	fmt.Println("hi!")
	return api.CheckVulnResponse{}
}
