package ltools

// Xor reports true when exactly one of its operands is true.
func Xor(operands ...bool) bool {
	trueCount := 0
	for _, operand := range operands {
		if operand {
			trueCount++
		}
		if trueCount > 1 {
			break
		}
	}
	if trueCount == 1 {
		return true
	}
	return false
}
