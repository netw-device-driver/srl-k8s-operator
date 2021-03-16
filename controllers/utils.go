package controllers

import "time"

const (
	targetNotFoundRetryDelay = time.Second * 60
)

// StringInList returns a boolean indicating whether strToSearch is a
// member of the string slice passed as the first argument.
func StringInList(list []string, strToSearch string) bool {
	for _, item := range list {
		if item == strToSearch {
			return true
		}
	}
	return false
}

// FilterStringFromList produces a new string slice that does not
// include the strToFilter argument.
func FilterStringFromList(list []string, strToFilter string) (newList []string) {
	for _, item := range list {
		if item != strToFilter {
			newList = append(newList, item)
		}
	}
	return
}

func removeString(slice []string, s string) (result []string) {
	for _, v := range slice {
		if v != s {
			result = append(result, s)
		}
	}
	return result
}
