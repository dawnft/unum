package unum

import (
	"math"
)

//	Represents an arbitrary 4-dimensional vector or a Quaternion.
type Vec4 struct {
	//	X, Y, Z
	Vec3

	//	The 4th vector component.
	W float64
}

//	Returns a new `*Vec4` containing a copy of `me`.
func (me *Vec4) Clone() (q *Vec4) {
	*q = *me
	return
}

//	Negates the `X`, `Y`, `Z` components in `me`, but not `W`.
func (me *Vec4) Conjugate() {
	me.X, me.Y, me.Z = -me.X, -me.Y, -me.Z
}

//	Returns a new `*Vec4` that represents `me` conjugated.
func (me *Vec4) Conjugated() (v *Vec4) {
	v = new(Vec4)
	v.X, v.Y, v.Z, v.W = -me.X, -me.Y, -me.Z, me.W
	return
}

//	Returns the 4D vector length of `me`.
func (me *Vec4) Length() float64 {
	return (me.X * me.X) + (me.Y * me.Y) + (me.Z * me.Z) + (me.W * me.W)
}

//	Returns the 4D vector magnitude of `me`.
func (me *Vec4) Magnitude() float64 {
	return math.Sqrt(me.Length())
}

//	Normalizes `me` according to `me.Magnitude()`.
func (me *Vec4) Normalize() {
	me.NormalizeFrom(me.Magnitude())
}

//	Normalizes `me` according to the specified `magnitude`.
func (me *Vec4) NormalizeFrom(magnitude float64) {
	if magnitude == 0 {
		me.X, me.Y, me.Z, me.W = 0, 0, 0, 0
	} else {
		magnitude = 1 / magnitude
		me.X, me.Y, me.Z, me.W = me.X*magnitude, me.Y*magnitude, me.Z*magnitude, me.W*magnitude
	}
}

//	Returns a new `*Vec4` that represents `me` normalized according to `me.Magnitude()`.
func (me *Vec4) Normalized() *Vec4 {
	var q Vec4
	if mag := me.Magnitude(); mag != 0 {
		mag = 1 / mag
		q.X, q.Y, q.Z, q.W = me.X*mag, me.Y*mag, me.Z*mag, me.W*mag
	}
	return &q
}

//	Sets `me` to `c` conjugated.
func (me *Vec4) SetFromConjugated(c *Vec4) {
	me.X, me.Y, me.Z, me.W = -c.X, -c.Y, -c.Z, c.W
}

//	Applies various 4D vector component computations of `l` and `r` to `me`, as needed by the `Vec3.RotateRad()` method.
func (me *Vec4) SetFromMult(l, r *Vec4) {
	me.W = (l.W * r.W) - (l.X * r.X) - (l.Y * r.Y) - (l.Z * r.Z)
	me.X = (l.X * r.W) + (l.W * r.X) + (l.Y * r.Z) - (l.Z * r.Y)
	me.Y = (l.Y * r.W) + (l.W * r.Y) + (l.Z * r.X) - (l.X * r.Z)
	me.Z = (l.Z * r.W) + (l.W * r.Z) + (l.X * r.Y) - (l.Y * r.X)
}

//	Scales all 4 vector components in `me` by factor `v`.
func (me *Vec4) Scale(v float64) {
	me.X, me.Y, me.Z, me.W = me.X*v, me.Y*v, me.Z*v, me.W*v
}

//	Sets `me` to the result of multiplying the specified `*Mat4` with the specified `*Vec3`.
func (me *Vec4) MultMat4Vec3(mat *Mat4, vec *Vec3) {
	me.X = (mat[0] * vec.X) + (mat[4] * vec.Y) + (mat[8] * vec.Z) + (mat[12] * 1)
	me.Y = (mat[1] * vec.X) + (mat[5] * vec.Y) + (mat[9] * vec.Z) + (mat[13] * 1)
	me.Z = (mat[2] * vec.X) + (mat[6] * vec.Y) + (mat[10] * vec.Z) + (mat[14] * 1)
	me.W = (mat[3] * vec.X) + (mat[7] * vec.Y) + (mat[11] * vec.Z) + (mat[15] * 1)
}

//	Sets `me` to the result of multiplying the specified `*Mat4` with the specified `*Vec4`.
func (me *Vec4) MultMat4Vec4(mat *Mat4, vec *Vec4) {
	me.X = (mat[0] * vec.X) + (mat[4] * vec.Y) + (mat[8] * vec.Z) + (mat[12] * vec.W)
	me.Y = (mat[1] * vec.X) + (mat[5] * vec.Y) + (mat[9] * vec.Z) + (mat[13] * vec.W)
	me.Z = (mat[2] * vec.X) + (mat[6] * vec.Y) + (mat[10] * vec.Z) + (mat[14] * vec.W)
	me.W = (mat[3] * vec.X) + (mat[7] * vec.Y) + (mat[11] * vec.Z) + (mat[15] * vec.W)
}

//	Applies various 4D vector component computations of `q` and `v` to `me`, as needed by the `Vec3.RotateRad()` method.
func (me *Vec4) SetFromMult3(q *Vec4, v *Vec3) {
	me.W = -(q.X * v.X) - (q.Y * v.Y) - (q.Z * v.Z)
	me.X = (q.W * v.X) + (q.Y * v.Z) - (q.Z * v.Y)
	me.Y = (q.W * v.Y) + (q.Z * v.X) - (q.X * v.Z)
	me.Z = (q.W * v.Z) + (q.X * v.Y) - (q.Y * v.X)
}

//	Sets `me` to the result of multiplying the specified `*Mat4` with `me`.
func (me *Vec4) SetFromMultMat4(mat *Mat4) {
	me.MultMat4Vec4(mat, me.Clone())
}

//	Returns a human-readable (imprecise) `string` representation of `me`.
func (me *Vec4) String() string {
	return strf("{X:%1.2f Y:%1.2f Z:%1.2f W:%1.2f}", me.X, me.Y, me.Z, me.W)
}
