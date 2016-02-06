package util

import (
	"../config"
)

func TypeChecker(param string) bool {
	var OssTypes = config.GetOssTypes()

	var isContain = false
	for _, value := range OssTypes {
		if param == value {
			isContain = true
		}
	}
	return isContain
}
