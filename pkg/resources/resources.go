package resources

import (
	"sort"
	"strings"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/discovery"
)

// CategoryMap maps a list of resources to a category name
type CategoryMap map[string][]schema.GroupVersionResource

func (o CategoryMap) addCategory(
	category string,
	gvr schema.GroupVersionResource,
) {
	if _, ok := o[category]; !ok {
		o[category] = []schema.GroupVersionResource{}
	}
	o[category] = append(o[category], gvr)
}

func (o CategoryMap) String() string {
	var sb strings.Builder
	categories := make([]string, 0, len(o))
	for k := range o {
		categories = append(categories, k)
	}
	sort.Strings(categories)
	for _, category := range categories {
		sb.WriteString(category + ":\n")

		resources := o[category]
		sort.Slice(resources, func(i, j int) bool {
			r1, r2 := resources[i], resources[j]
			if r1.Group != r2.Group {
				return r1.Group < r2.Group
			}
			if r1.Version != r2.Version {
				return r1.Version < r2.Version
			}
			return r1.Resource < r2.Resource
		})
		for _, resource := range resources {
			sb.WriteString("  " +
				resource.Resource +
				" (" +
				resource.GroupVersion().String() +
				")\n",
			)
		}
	}
	return sb.String()
}

func GetResourceCategories(
	discoveryClient discovery.DiscoveryInterface,
) (CategoryMap, error) {

	categoryMap := CategoryMap{}

	resourceLists, err := discoveryClient.ServerPreferredResources()
	if err != nil {
		return nil, err
	}

	for _, resourceList := range resourceLists {
		gv, err := schema.ParseGroupVersion(resourceList.GroupVersion)
		if err != nil {
			return nil, err
		}
		for _, resource := range resourceList.APIResources {
			categories := resource.Categories
			for _, category := range categories {
				categoryMap.addCategory(
					category,
					gv.WithResource(resource.Name),
				)
			}
		}
	}
	return categoryMap, nil
}
