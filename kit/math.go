package kit

import (
	"fmt"
	"math"
	"math/cmplx"
)

func QuadraticRoots(a, b, c float64) (float64, float64, error) {
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return 0, 0, fmt.Errorf("complex roots")
	}

	root1 := (-b + math.Sqrt(discriminant)) / (2 * a)
	root2 := (-b - math.Sqrt(discriminant)) / (2 * a)
	return root1, root2, nil
}

func QuadraticRootsComplex(a, b, c float64) (complex128, complex128) {
	discriminant := cmplx.Sqrt(complex(b*b-4*a*c, 0))
	root1 := (-complex(b, 0) + discriminant) / complex(2*a, 0)
	root2 := (-complex(b, 0) - discriminant) / complex(2*a, 0)
	return root1, root2
}
