package utils

func RemoveItemFromSlice2[T interface{}](list []T, itemToRemove T, comparator func(itemA, itemB T) bool) []T {
	for i, item := range list {
		if comparator(item, itemToRemove) {
			list[i] = list[len(list)-1]
			return list[:len(list)-1]
		}
	}
	return list
}
