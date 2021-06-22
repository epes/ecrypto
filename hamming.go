package ecrypto

import "math"

func HammingDistance(a []byte, b []byte) int {
	if len(a) != len(b) {
		return 0
	}

	var h int

	for i := 0; i < len(a); i++ {
		h += HammingDistanceB(a[i], b[i])
	}

	return h
}

func HammingDistanceB(a byte, b byte) int {
	return CountTrueBits(a ^ b)
}

// HammingBlockNormalized returns the normalized hamming distance between
// k-sized blocks of the input byte slice.
func HammingBlockNormalized(b []byte, k int) float64 {
	// nothing to compare it to
	if k < 2 {
		return math.MaxFloat64
	}

	// input not big enough to get comparable k-sized blocks
	if k*2 > len(b) {
		return math.MaxFloat64
	}

	// -1 because there are n-1 comparisons between n items
	comparisons := len(b)/k - 1

	var distance float64
	fk := float64(k)

	for i := 0; i < comparisons; i++ {
		distance += float64(HammingDistance(b[i*k:i*k+k], b[i*k+k:i*k+2*k])) / fk
	}

	return distance / float64(comparisons)
}

// HammingNLowestBlockSizesInRange takes the HammingBlockNormalized values for
// input b for all values (k) between [min,max] and returns the n lowest valued sizes.
// Result is length min(max-min+1, n).
func HammingNLowestBlockSizesInRange(b []byte, n int, min int, max int) []int {
	all := make([]float64, max-min+1)

	for i := 0; i <= max-min; i++ {
		all[i] = HammingBlockNormalized(b, min+i)
	}

	var k []int

	if len(all) < n {
		k = make([]int, 0, len(all))
	} else {
		k = make([]int, 0, n)
	}

	// O(n^2) - let's hope the range isn't too big!
	// TODO: heap if this is a bottleneck. find min(n, len(all)) lowest values
	for i := 0; i < cap(k); i++ {
		lowestIndex := 0
		lowestDistance := math.MaxFloat64

		for ki, kdistance := range all {
			if kdistance < lowestDistance {
				lowestDistance = kdistance
				lowestIndex = ki
			}
		}

		all[lowestIndex] = math.MaxFloat64
		k = append(k, min+lowestIndex)
	}

	return k
}
