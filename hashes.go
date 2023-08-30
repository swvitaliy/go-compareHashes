package compareHashes

func djb2(a []byte) uint64 {
	var hash uint64 = 5381
	for _, b := range a {
		hash = ((hash << 5) + hash) + uint64(b)
	}
	return hash
}

func sdbm(a []byte) uint64 {
	hash := uint64(0)
	for _, c := range a {
		hash = uint64(c) + (hash << 6) + (hash << 16) - hash
	}
	return hash
}
