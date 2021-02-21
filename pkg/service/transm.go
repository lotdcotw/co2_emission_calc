package service

// emission values will be loaded with etcd or static ini file
// if not, 1 will be used for tests
var emissions = map[string]float64{
	"carsmall_diesel":        1,
	"carsmall_petrol":        1,
	"carsmall_pluginhybrid":  1,
	"carsmall_electric":      1,
	"carmedium_diesel":       1,
	"carmedium_petrol":       1,
	"carmedium_pluginhybrid": 1,
	"carmedium_electric":     1,
	"carlarge_diesel":        1,
	"carlarge_petrol":        1,
	"carlarge_pluginhybrid":  1,
	"carlarge_electric":      1,
	"bus_generic":            1,
	"train_generic":          1,
}

var transmMap = map[string]string{
	"carsmall_diesel":       "small-diesel-car",
	"carsmall_petrol":       "small-petrol-car",
	"carsmall_pluginhybrid": "small-plugin-hybrid-car",
	"carsmall_electric":     "small-electric-car",

	"carmedium_diesel":       "medium-diesel-car",
	"carmedium_petrol":       "medium-petrol-car",
	"carmedium_pluginhybrid": "medium-plugin-hybrid-car",
	"carmedium_electric":     "medium-electric-car",

	"carlarge_diesel":       "large-diesel-car",
	"carlarge_petrol":       "large-petrol-car",
	"carlarge_pluginhybrid": "large-plugin-hybrid-car",
	"carlarge_electric":     "large-electric-car",

	"bus_generic": "bus",

	"train_generic": "train",
}

func keyConversion(param string) string {
	for k, v := range transmMap {
		if v == param {
			return k
		}
	}
	return ""
}
