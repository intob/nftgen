package gen

import (
	"github.com/intob/thinggen/util"
)

func RandomTraitMapping(traits map[string]Trait) (map[string]Variant, error) {
	m := make(map[string]Variant)
	for traitName, trait := range traits {
		keys := make([]string, 0, len(trait))
		for k := range trait {
			keys = append(keys, k)
		}
		v, err := util.RandInt(len(keys))
		if err != nil {
			return nil, err
		}
		m[traitName] = trait[keys[v]]
	}
	return m, nil
}

func CountPossibleMappings(traits map[string]Trait) (int, error) {
	n := 1
	for _, t := range traits {
		n *= len(t)
	}
	return n, nil
}
