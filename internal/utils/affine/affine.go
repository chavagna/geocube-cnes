// Package to handle 2D affine transformations, following GDAL affine convention
package affine

// Affine follows the GDAL transform convention
type Affine [6]float64

func NewAffine(a, b, c, d, e, f float64) *Affine {
	res := Affine([6]float64{a, b, c, d, e, f})
	return &res
}

// Translation creates a translation transform from (offx, offy)
func Translation(offx, offy float64) *Affine {
	return NewAffine(offx, 1.0, 0, offy, 0, 1.0)
}

// Scale creates a scale transform from (scalex, scaley)
func Scale(scalex, scaley float64) *Affine {
	return NewAffine(0, scalex, 0, 0, 0, scaley)
}

// Rx returns the X resolution
func (a *Affine) Rx() float64 {
	return float64(a[1])
}

// Ry returns the Y resolution
func (a *Affine) Ry() float64 {
	return float64(a[5])
}

// IsInvertible returns true if the transformation is invertible
func (a *Affine) IsInvertible() bool {
	return a[1]*a[5] != a[2]*a[4] // det != 0
}

// Inverse creates the inverse of the affine transform.
// Inverse panics if it is not inversible
func (a *Affine) Inverse() *Affine {
	idet := 1.0 / (a[1]*a[5] - a[2]*a[4])
	res := Affine([6]float64{0, a[5] * idet, -a[2] * idet, 0, -a[4] * idet, a[1] * idet})
	res[0], res[3] = res.Transform(-a[0], -a[3])
	return &res
}

// Multiply merges the two affines transforms into one.
func (a *Affine) Multiply(b *Affine) *Affine {
	return NewAffine(
		a[1]*b[0]+a[2]*b[3]+a[0],
		a[1]*b[1]+a[2]*b[4],
		a[1]*b[2]+a[2]*b[5],
		a[4]*b[0]+a[5]*b[3]+a[3],
		a[4]*b[1]+a[5]*b[4],
		a[4]*b[2]+a[5]*b[5],
	)
}

// Transform applies the affine transform to the point (x, y)
func (a *Affine) Transform(x float64, y float64) (float64, float64) {
	return a[0] + a[1]*x + a[2]*y, a[3] + a[4]*x + a[5]*y
}
