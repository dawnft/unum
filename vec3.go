package unum

import (
	"math"
)

//	Represents a 3-dimensional vector.
type Vec3 struct {
	X, Y, Z float64
}

//	Returns the `math.Max()` of the `math.Abs()` values of all 3 components in `me`.
func (me *Vec3) AbsMax() float64 {
	return math.Max(math.Abs(me.X), math.Max(math.Abs(me.Y), math.Abs(me.Z)))
}

//	Adds `vec` to `me` in-place.
func (me *Vec3) Add(vec *Vec3) {
	me.X, me.Y, me.Z = me.X+vec.X, me.Y+vec.Y, me.Z+vec.Z
}

//	Returns the sum of `me` and `vec`.
func (me *Vec3) AddTo(vec *Vec3) *Vec3 {
	return &Vec3{me.X + vec.X, me.Y + vec.Y, me.Z + vec.Z}
}

//	Adds `val` to all 3 components of `me`.
func (me *Vec3) Add1(val float64) {
	me.X, me.Y, me.Z = me.X+val, me.Y+val, me.Z+val
}

//	Adds the specified 3 components to the respective components in `me`.
func (me *Vec3) Add3(x, y, z float64) {
	me.X, me.Y, me.Z = me.X+x, me.Y+y, me.Z+z
}

//	Returns whether all 3 components in `me` equal `val`.
func (me *Vec3) AllEqual(val float64) bool {
	return (me.X == val) && (me.Y == val) && (me.Z == val)
}

//	Returns whether all 3 components in `me` are greater than or equal to their respective component counterparts in `vec`.
func (me *Vec3) AllGreaterOrEqual(vec *Vec3) bool {
	return (me.X >= vec.X) && (me.Y >= vec.Y) && (me.Z >= vec.Z)
}

//	Returns whether all 3 components in `me` are greater than (or equal to) `min`, and also less than `max`.
func (me *Vec3) AllInRange(min, max float64) bool {
	return (me.X >= min) && (me.X < max) && (me.Y >= min) && (me.Y < max) && (me.Z >= min) && (me.Z < max)
}

//	Returns whether all 3 components in `me` are greater than `min`, and also less than `max`.
func (me *Vec3) AllInside(min, max *Vec3) bool {
	return (me.X > min.X) && (me.X < max.X) && (me.Y > min.Y) && (me.Y < max.Y) && (me.Z > min.Z) && (me.Z < max.Z)
}

//	Returns whether all 3 components in `me` are less than their respective component counterparts in `vec`.
func (me *Vec3) AllLessOrEqual(vec *Vec3) bool {
	return (me.X <= vec.X) && (me.Y <= vec.Y) && (me.Z <= vec.Z)
}

//	Zeroes all 3 components in `me`.
func (me *Vec3) Clear() {
	me.X, me.Y, me.Z = 0, 0, 0
}

//	Returns a new `*Vec3` that represents the cross-product of `me` and `vec`.
func (me *Vec3) Cross(vec *Vec3) *Vec3 {
	return &Vec3{(me.Y * vec.Z) - (me.Z * vec.Y), (me.Z * vec.X) - (me.X * vec.Z), (me.X * vec.Y) - (me.Y * vec.X)}
}

//	Returns a new `*Vec` that represents the cross-product of `me` and `vec`, normalized.
func (me *Vec3) CrossNormalized(vec *Vec3) (r *Vec3) {
	r = me.Cross(vec) // should be inlined.
	//	&Vec3{(me.Y * vec.Z) - (me.Z * vec.Y), (me.Z * vec.X) - (me.X * vec.Z), (me.X * vec.Y) - (me.Y * vec.X)}
	r.Normalize()
	return
}

//	Returns the distance of `me` from `vec`.
func (me *Vec3) DistanceFrom(vec *Vec3) float64 {
	return math.Sqrt(me.SubDot(vec))
}

//	Returns the distance of `me` from "the center" (zero).
func (me *Vec3) DistanceFromZero() float64 {
	return math.Sqrt(me.Dot(me))
}

//	Returns a new `*Vec3` that represents `me` divided by `vec`.
func (me *Vec3) Div(vec *Vec3) *Vec3 {
	return &Vec3{me.X / vec.X, me.Y / vec.Y, me.Z / vec.Z}
}

//	Returns a new `*Vec3` that represents all 3 components in `me`, each divided by `val`.
func (me *Vec3) Div1(val float64) *Vec3 {
	return &Vec3{me.X / val, me.Y / val, me.Z / val}
}

//	Returns the dot-product of `me` and `vec`.
func (me *Vec3) Dot(vec *Vec3) float64 {
	return (me.X * vec.X) + (me.Y * vec.Y) + (me.Z * vec.Z)
}

//	Returns the dot-product of `me` and (`vec1` minus `vec2`).
func (me *Vec3) DotSub(vec1, vec2 *Vec3) float64 {
	return (me.X * (vec1.X - vec2.X)) + (me.Y * (vec1.Y - vec2.Y)) + (me.Z * (vec1.Z - vec2.Z))
}

/*
//	Returns whether `me` equals `vec`.
func (me *Vec3) Equals(vec *Vec3) bool {
	// return *me == *vec?
	return (me.X == vec.X) && (me.Y == vec.Y) && (me.Z == vec.Z)
}
*/

//	Returns a new `*Vec3` representing the inverse of `me`.
func (me *Vec3) Inv() *Vec3 {
	return &Vec3{1 / me.X, 1 / me.Y, 1 / me.Z}
}

//	Returns the 3D vector length of `me`.
func (me *Vec3) Length() float64 {
	return (me.X * me.X) + (me.Y * me.Y) + (me.Z * me.Z)
}

//	Returns the 3D vector magnitude of `me`.
func (me *Vec3) Magnitude() float64 {
	return math.Sqrt(me.Length())
}

//	Sets each component in `me` to the corresponding component in `vec` only if the former is infinity.
func (me *Vec3) MakeFinite(vec *Vec3) {
	if math.IsInf(me.X, 0) {
		me.X = vec.X
	}
	if math.IsInf(me.Y, 0) {
		me.Y = vec.Y
	}
	if math.IsInf(me.Z, 0) {
		me.Z = vec.Z
	}
}

//	Returns the "manhattan distance" of `me` from `vec`.
func (me *Vec3) ManhattanDistanceFrom(vec *Vec3) float64 {
	return math.Abs(vec.X-me.X) + math.Abs(vec.Y-me.Y) + math.Abs(vec.Z-me.Z)
}

//	Returns the largest of the 3 components in `me`.
func (me *Vec3) Max() float64 {
	return math.Max(me.X, math.Max(me.Y, me.Z))
}

//	Returns the smallest of the 3 components in `me`.
func (me *Vec3) Min() float64 {
	return math.Min(me.X, math.Min(me.Y, me.Z))
}

//	Clamps each component in `me` between the respective corresponding counter-part component in `min` and `max`.
func (me *Vec3) Clamp(min, max *Vec3) {
	if me.X < min.X {
		min.X = me.X
	}
	if me.X > max.X {
		max.X = me.X
	}
	if me.Y < min.Y {
		min.Y = me.Y
	}
	if me.Y > max.Y {
		max.Y = me.Y
	}
	if me.Z < min.Z {
		min.Z = me.Z
	}
	if me.Z > max.Z {
		max.Z = me.Z
	}
}

//	Returns a new `*Vec3` that represents `me` multiplied with `vec`.
func (me *Vec3) Mult(vec *Vec3) *Vec3 {
	return &Vec3{me.X * vec.X, me.Y * vec.Y, me.Z * vec.Z}
}

//	Returns a new `*Vec3` with each component in `me` multiplied by the respective corresponding specified factor.
func (me *Vec3) Times(x, y, z float64) *Vec3 {
	return &Vec3{me.X * x, me.Y * y, me.Z * z}
}

//	Reverses the signs of all 3 vector components in `me`.
func (me *Vec3) Negate() {
	me.X, me.Y, me.Z = -me.X, -me.Y, -me.Z
}

//	Returns a new `*Vec` with each component representing the negative (sign inverted) corresponding component in `me`.
func (me *Vec3) Negated() *Vec3 {
	return &Vec3{-me.X, -me.Y, -me.Z}
}

//	Normalizes `me` in-place.
func (me *Vec3) Normalize() {
	if tmpVal := me.Magnitude(); tmpVal == 0 {
		me.X, me.Y, me.Z = 0, 0, 0
	} else {
		me.Scale(1 / tmpVal)
	}
}

//	Returns a new `*Vec3` that represents `me`, normalized.
func (me *Vec3) Normalized() (vec *Vec3) {
	vec = &Vec3{}
	vec.SetFromNormalized(me)
	// if tmpVal := me.Magnitude(); tmpVal != 0 {
	// 	vec.SetFromMult1(me, 1/tmpVal)
	// }
	return
}

//	Returns a new `*Vec3` that represents `me` normalized, then scaled by `factor`.
func (me *Vec3) NormalizedScaled(factor float64) (vec *Vec3) {
	vec = me.Normalized()
	vec.Scale(factor)
	return
}

//	Rotates `me` `angleDeg` degrees around the specified `axis`.
func (me *Vec3) RotateDeg(angleDeg float64, axis *Vec3) {
	me.RotateRad(DegToRad(angleDeg/2), axis)
}

//	Rotates `me` `angleRad` radians around the specified `axis`.
func (me *Vec3) RotateRad(angleRad float64, axis *Vec3) {
	var tmpQ, tmpQw, tmpQr, tmpQc Vec4
	sin := math.Sin(angleRad)
	tmpQr.X, tmpQr.Y, tmpQr.Z, tmpQr.W = axis.X*sin, axis.Y*sin, axis.Z*sin, math.Cos(angleRad)
	tmpQc.SetFromConjugated(&tmpQr)
	tmpQ.SetFromMult3(&tmpQr, me)
	tmpQw.SetFromMult(&tmpQ, &tmpQc)
	me.X, me.Y, me.Z = tmpQw.X, tmpQw.Y, tmpQw.Z
}

//	Scales `me` by `factor`.
func (me *Vec3) Scale(factor float64) {
	me.X, me.Y, me.Z = me.X*factor, me.Y*factor, me.Z*factor
}

//	Scales `me` by `factor`, then adds `add`.
func (me *Vec3) ScaleAdd(factor, add *Vec3) {
	me.X, me.Y, me.Z = (me.X*factor.X)+add.X, (me.Y*factor.Y)+add.Y, (me.Z*factor.Z)+add.Z
}

//	Returns a new `*Vec3` that represents `me` scaled by `factor`.
func (me *Vec3) Scaled(factor float64) *Vec3 {
	return &Vec3{me.X * factor, me.Y * factor, me.Z * factor}
}

//	Returns a new `*Vec3` that represents `me` scaled by `factor`, then `add` added.
func (me *Vec3) ScaledAdded(factor float64, add *Vec3) *Vec3 {
	return &Vec3{(me.X * factor) + add.X, (me.Y * factor) + add.Y, (me.Z * factor) + add.Z}
}

//	Sets all 3 vector components in `me` to the corresponding respective specified value.
func (me *Vec3) Set(x, y, z float64) {
	me.X, me.Y, me.Z = x, y, z
}

/*
//	Sets each vector component in `me` to the corresponding respective component in `vec`.
func (me *Vec3) SetFrom(vec *Vec3) {
	me.X, me.Y, me.Z = vec.X, vec.Y, vec.Z
}
*/

//	Sets `me` to the result of adding `vec1` and `vec2`.
func (me *Vec3) SetFromAdd(vec1, vec2 *Vec3) {
	me.X, me.Y, me.Z = vec1.X+vec2.X, vec1.Y+vec2.Y, vec1.Z+vec2.Z
}

//	`me = a + b + c`
func (me *Vec3) SetFromAddAdd(a, b, c *Vec3) {
	me.X, me.Y, me.Z = a.X+b.X+c.X, a.Y+b.Y+c.Y, a.Z+b.Z+c.Z
}

//	`me = add + (mul1 * mul2)`
func (me *Vec3) SetFromAddMult(add, mul1, mul2 *Vec3) {
	me.X, me.Y, me.Z = add.X+(mul1.X*mul2.X), add.Y+(mul1.Y*mul2.Y), add.Z+(mul1.Z*mul2.Z)
}

//	`me = vec2 + vec2 * mul`
func (me *Vec3) SetFromAddScaled(vec1, vec2 *Vec3, mul float64) {
	me.X, me.Y, me.Z = vec1.X+vec2.X*mul, vec1.Y+vec2.Y*mul, vec1.Z+vec2.Z*mul
}

//	`me = a + b - c`
func (me *Vec3) SetFromAddSub(a, b, c *Vec3) {
	me.X, me.Y, me.Z = a.X+b.X-c.X, a.Y+b.Y-c.Y, a.Z+b.Z-c.Z
}

//	Sets each vector component in `me` to the `math.Cos()` of the respective corresponding component in `vec`.
func (me *Vec3) SetFromCos(vec *Vec3) {
	me.X, me.Y, me.Z = math.Cos(vec.X), math.Cos(vec.Y), math.Cos(vec.Z)
}

//	Sets `me` to the cross-product of `me` and `vec`.
func (me *Vec3) SetFromCross(vec *Vec3) {
	me.X, me.Y, me.Z = (me.Y*vec.Z)-(me.Z*vec.Y), (me.Z*vec.X)-(me.X*vec.Z), (me.X*vec.Y)-(me.Y*vec.X)
}

//	Sets `me` to the cross-product of `one` and `two`.
func (me *Vec3) SetFromCrossOf(one, two *Vec3) {
	me.X, me.Y, me.Z = (one.Y*two.Z)-(one.Z*two.Y), (one.Z*two.X)-(one.X*two.Z), (one.X*two.Y)-(one.Y*two.X)
}

//	Sets each vector component in `me` to the radian equivalent of the degree angle stored in the respective corresponding component of `vec`.
func (me *Vec3) SetFromDegToRad(deg *Vec3) {
	me.X, me.Y, me.Z = DegToRad(deg.X), DegToRad(deg.Y), DegToRad(deg.Z)
}

//	Sets each vector component in `me` to `Epsilon32` if it is 0 or greater but smaller than `Epsilon32`.
func (me *Vec3) SetFromEpsilon32() {
	if (me.X >= 0) && (me.X < Epsilon32) {
		me.X = Epsilon32
	}
	if (me.Y >= 0) && (me.Y < Epsilon32) {
		me.Y = Epsilon32
	}
	if (me.Z >= 0) && (me.Z < Epsilon32) {
		me.Z = Epsilon32
	}
}

//	Sets each vector component in `me` to `Epsilon64` if it is 0 or greater but smaller than `Epsilon64`.
func (me *Vec3) SetFromEpsilon64() {
	if (me.X >= 0) && (me.X < Epsilon64) {
		me.X = Epsilon64
	}
	if (me.Y >= 0) && (me.Y < Epsilon64) {
		me.Y = Epsilon64
	}
	if (me.Z >= 0) && (me.Z < Epsilon64) {
		me.Z = Epsilon64
	}
}

//	Sets `me` to the inverse of `vec`.
func (me *Vec3) SetFromInv(vec *Vec3) {
	me.X, me.Y, me.Z = 1/vec.X, 1/vec.Y, 1/vec.Z
}

//	`me = v1 * v2`
func (me *Vec3) SetFromMult(v1, v2 *Vec3) {
	me.X, me.Y, me.Z = v1.X*v2.X, v1.Y*v2.Y, v1.Z*v2.Z
}

//	`me = vec * mul`
func (me *Vec3) SetFromScaled(vec *Vec3, mul float64) {
	me.X, me.Y, me.Z = vec.X*mul, vec.Y*mul, vec.Z*mul
}

//	`me = (vec1 - vec2) * mul`
func (me *Vec3) SetFromScaledSub(vec1, vec2 *Vec3, mul float64) {
	me.X, me.Y, me.Z = (vec1.X-vec2.X)*mul, (vec1.Y-vec2.Y)*mul, (vec1.Z-vec2.Z)*mul
}

//	`me = -vec`
func (me *Vec3) SetFromNegated(vec *Vec3) {
	me.X, me.Y, me.Z = -vec.X, -vec.Y, -vec.Z
}

//	Sets `me` to `vec` normalized.
func (me *Vec3) SetFromNormalized(vec *Vec3) {
	if tmpVal := vec.Magnitude(); tmpVal != 0 {
		me.SetFromScaled(vec, 1/tmpVal)
	}
}

//	Sets `me` to `pos` rotated as expressed in `rotCos` and `rotSin`.
func (me *Vec3) SetFromRotation(pos, rotCos, rotSin *Vec3) {
	tmpVal := ((pos.Y * rotSin.X) + (pos.Z * rotCos.X))
	me.X = (pos.X * rotCos.Y) + (tmpVal * rotSin.Y)
	me.Y = (pos.Y * rotCos.X) - (pos.Z * rotSin.X)
	me.Z = (-pos.X * rotSin.Y) + (tmpVal * rotCos.Y)
}

//	Sets each vector component in `me` to the `math.Sin()` of the respective corresponding component in `vec`.
func (me *Vec3) SetFromSin(vec *Vec3) {
	me.X, me.Y, me.Z = math.Sin(vec.X), math.Sin(vec.Y), math.Sin(vec.Z)
}

//	Component-wise, set `me` to `v0` if it is less than `edge`, else `v1`.
func (me *Vec3) SetFromStep(edge float64, vec, v0, v1 *Vec3) {
	if vec.X < edge {
		me.X = v0.X
	} else {
		me.X = v1.X
	}
	if vec.Y < edge {
		me.Y = v0.Y
	} else {
		me.Y = v1.Y
	}
	if vec.Z < edge {
		me.Z = v0.Z
	} else {
		me.Z = v1.Z
	}
}

//	`me = vec1 - vec2`.
func (me *Vec3) SetFromSub(vec1, vec2 *Vec3) {
	me.X, me.Y, me.Z = vec1.X-vec2.X, vec1.Y-vec2.Y, vec1.Z-vec2.Z
}

//	`me = a - b + c`
func (me *Vec3) SetFromSubAdd(a, b, c *Vec3) {
	me.X, me.Y, me.Z = a.X-b.X+c.X, a.Y-b.Y+c.Y, a.Z-b.Z+c.Z
}

//	`me = v1 - v2 * v2Scale`
func (me *Vec3) SetFromSubScaled(v1, v2 *Vec3, v2Scale float64) {
	me.X, me.Y, me.Z = v1.X-v2.X*v2Scale, v1.Y-v2.Y*v2Scale, v1.Z-v2.Z*v2Scale
}

//	`me = a - b - c`
func (me *Vec3) SetFromSubSub(a, b, c *Vec3) {
	me.X, me.Y, me.Z = a.X-b.X-c.X, a.Y-b.Y-c.Y, a.Z-b.Z-c.Z
}

//	`me = (sub1 - sub2) * mul`
func (me *Vec3) SetFromSubMult(sub1, sub2, mul *Vec3) {
	me.X, me.Y, me.Z = (sub1.X-sub2.X)*mul.X, (sub1.Y-sub2.Y)*mul.Y, (sub1.Z-sub2.Z)*mul.Z
}

//	Sets all 3 vector components in `me` to `math.MaxFloat64`.
func (me *Vec3) SetToMax() {
	me.X, me.Y, me.Z = math.MaxFloat64, math.MaxFloat64, math.MaxFloat64
}

//	Sets all 3 vector components in `me` to `-math.MaxFloat64`.
func (me *Vec3) SetToMin() {
	me.X, me.Y, me.Z = -math.MaxFloat64, -math.MaxFloat64, -math.MaxFloat64
}

//	Returns a new `*Vec3` with each vector component indicating the sign (-1, 1 or 0) of the respective corresponding component in `me`.
func (me *Vec3) Sign() *Vec3 {
	return &Vec3{Sign(me.X), Sign(me.Y), Sign(me.Z)}
}

//	Returns a human-readable (imprecise) `string` representation of `me`.
func (me *Vec3) String() string {
	return strf("{X:%1.2f Y:%1.2f Z:%1.2f}", me.X, me.Y, me.Z)
}

//	Returns a new `*Vec3` that represents `me` minus `vec`.
func (me *Vec3) Sub(vec *Vec3) *Vec3 {
	return &Vec3{me.X - vec.X, me.Y - vec.Y, me.Z - vec.Z}
}

//	Returns a new `*Vec3` that represents `((me - sub) / div) * mul`.
func (me *Vec3) SubDivMult(sub, div, mul *Vec3) *Vec3 {
	return &Vec3{((me.X - sub.X) / div.X) * mul.X, ((me.Y - sub.Y) / div.Y) * mul.Y, ((me.Z - sub.Z) / div.Z) * mul.Z}
}

//	Returns the dot-product of `me` minus `vec`.
func (me *Vec3) SubDot(vec *Vec3) float64 {
	return ((me.X - vec.X) * (me.X - vec.X)) + ((me.Y - vec.Y) * (me.Y - vec.Y)) + ((me.Z - vec.Z) * (me.Z - vec.Z))
}

//	Returns a new `*Vec3` that represents `mul * math.Floor(me / floorDiv)`.
func (me *Vec3) SubFloorDivMult(floorDiv, mul float64) *Vec3 {
	return me.Sub(&Vec3{math.Floor(me.X/floorDiv) * mul, math.Floor(me.Y/floorDiv) * mul, math.Floor(me.Z/floorDiv) * mul})
}

//	Returns a new `*Vec3` that represents `val` minus `me`.
func (me *Vec3) SubFrom(val float64) *Vec3 {
	return &Vec3{val - me.X, val - me.Y, val - me.Z}
}

//	Returns a new `*Vec3` that represents `(me - vec) * val`.
func (me *Vec3) SubScaled(vec *Vec3, val float64) *Vec3 {
	return &Vec3{(me.X - vec.X) * val, (me.Y - vec.Y) * val, (me.Z - vec.Z) * val}
}

//	Subtracts `vec` from `me`.
func (me *Vec3) SubVec(vec *Vec3) {
	me.X, me.Y, me.Z = me.X-vec.X, me.Y-vec.Y, me.Z-vec.Z
}

//	Transform coordinate vector `me` according to the specified `*Mat4`.
func (me *Vec3) TransformCoord(mat *Mat4) {
	var q Vec4
	q.MultMat4Vec3(mat, me)
	q.W = 1 / q.W
	me.X, me.Y, me.Z = q.X*q.W, q.Y*q.W, q.Z*q.W
}

//	Transform normal vector `me` according to the specified `*Mat4`.
func (me *Vec3) TransformNormal(mat *Mat4, absMat bool) {
	m11, m21, m31 := mat[0], mat[1], mat[2]
	m12, m22, m32 := mat[4], mat[5], mat[6]
	m13, m23, m33 := mat[8], mat[9], mat[10]
	if absMat {
		m11, m21, m31 = math.Abs(m11), math.Abs(m21), math.Abs(m31)
		m12, m22, m32 = math.Abs(m12), math.Abs(m22), math.Abs(m32)
		m13, m23, m33 = math.Abs(m13), math.Abs(m23), math.Abs(m33)
	}
	x := ((me.X * m11) + (me.Y * m21)) + (me.Z * m31)
	y := ((me.X * m12) + (me.Y * m22)) + (me.Z * m32)
	z := ((me.X * m13) + (me.Y * m23)) + (me.Z * m33)
	me.X, me.Y, me.Z = x, y, z
}

/*
func (me *Vec3) ToDegInts () []int {
	return []int { int(RadToDeg(me.X)), int(RadToDeg(me.Y)), int(RadToDeg(me.Z)) }
}

func (me *Vec3) ToInts () []int {
	return []int { int(me.X), int(me.Y), int(me.Z) }
}
*/
