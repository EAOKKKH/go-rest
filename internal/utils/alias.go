package utils

import "math/rand"

func GetAlias(size int) string {
	const alNum = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	alias := make([]byte, size)
	for i := range alias {
		alias[i] = alNum[rand.Intn(len(alNum))]
	}

	return string(alias)
}
