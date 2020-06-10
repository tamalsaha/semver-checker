package main

import (
	"fmt"
	"github.com/Masterminds/semver"
	"sort"
)

func main() {
	raw := []string{"v2020.6.10-alpha.0", "v2020.6.10-v.0", "v2020.6.10-v.10", "v2020.6.10-v.2", "1.0", "v2020.7.1", "2", "0.4.2"}
	vs := make([]*semver.Version, len(raw))
	for i, r := range raw {
		v, err := semver.NewVersion(r)
		if err != nil {
			panic(fmt.Errorf("error parsing version: %s", err))
		}

		vs[i] = v
	}

	sort.Sort(semver.Collection(vs))

	fmt.Println(vs)
}
