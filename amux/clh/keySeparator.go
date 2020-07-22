package main

import "os"

type keyArgSlice struct {
	keyMap map[string][]string
}

func KeyArgs() keyArgSlice {
	allArgs := os.Args
	kaSl := keyArgSlice{}
	kaSl.keyMap = make(map[string][]string)
	var tempSl []string
	for i := range allArgs {
		if string(allArgs[i][0]) != "-" {
			tempSl = append(tempSl, allArgs[i])
			continue
		}
		kaSl.keyMap[allArgs[i]] = tempSl
		tempSl = nil
	}

	return kaSl
}
