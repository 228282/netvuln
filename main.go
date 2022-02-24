package netvuln

import (
	"fmt"

	api "netvuln/api"
	v1 "netvuln/internal/v1"
)

func main() {
	fmt.Println("ololo")

	a := v1.CheckVuln(api.CheckVulnRequest{})
	fmt.Println(a.Results)
}
