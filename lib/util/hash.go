package util

const primeRK = 16777619

func HashStr(str string) uint32 {
	hash := uint32(0)
	for i := 0; i < len(str); i++ {
		hash = hash*primeRK + uint32(str[i])
	}
	return hash
}
