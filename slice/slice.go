package slice

func AllocateVar(b int) []byte {
	return []byte{}
}

func AllocateMake(b int) []byte {
	return make([]byte, b)
}

func Reslice(slc []byte, lidx int, uidx int) []byte {
	s := CopySlice(slc)
	return s[lidx:uidx]
}

func CopySlice(slc []byte) []byte {
	s := AllocateMake(len(slc))
	copy(s, slc)
	return s
}
