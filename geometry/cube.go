package geometry

import (
	"errors"
)

//CubeVolume returns de volume of a cube.
func CubeVolume(n int) (int, error) {
	if n != 0 {
		return n * n * n, nil
	} else {
		return 0, errors.New("Zero length edge is not allowed")
	}
}
