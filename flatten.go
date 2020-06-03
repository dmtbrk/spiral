package spiral

import "errors"

// FlattenInt flattens the provided rectangular array (slice of slices) s unwrapping it in the clockwise direction.
// If a non-rectangular or nil slice is provided, an error is returned. An empty slice unwraps to
// an empty one.
func FlattenInt(s [][]int) ([]int, error) {
	// check if nil
	if s == nil {
		return nil, errors.New("cannot flatten an empty slice")
	}

	n := len(s)

	// check if not rectangular
	for _, row := range s {
		if len(row) != n {
			return nil, errors.New("cannot flatten a non-rectangular slice")
		}
	}

	flat := make([]int, n*n)

	i, j := 0, 0               // row and column indices of s
	startIdx, endIdx := 0, n-1 // lower and upper indices for the current loop of a spiral
	for fi := range flat {
		flat[fi] = s[i][j]

		if j == startIdx && i == startIdx+1 { // got to the end of loop, go to inner one
			startIdx += 1
			endIdx -= 1
		}
		if i == startIdx && j < endIdx { // i is at the start of a line, moving rigth
			j += 1
		} else if j == endIdx && i < endIdx { // j got to the end of a line, moving down
			i += 1
		} else if i == endIdx && j > startIdx { // i got to the end of a line, moving left
			j -= 1
		} else if j == startIdx && i > startIdx { // j got to the start of a line, moving up
			i -= 1
		}
	}

	return flat, nil
}
