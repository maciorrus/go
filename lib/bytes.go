package lib

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

func BytesToKB(size int) float32 {
	return float32(size) / KB
}

func BytesToMB(size int) float32 {
	return float32(size) / MB
}

func BytesToGB(size int) float32 {
	return float32(size) / GB
}

func BytesToTB(size int) float32 {
	return float32(size) / TB
}

func BytesToPB(size int) float32 {
	return float32(size) / PB
}

func BytesToZB(size int) float32 {
	return float32(size) / ZB
}

func BytesToYB(size int) float32 {
	return float32(size) / YB
}
