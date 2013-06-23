# unum
--
    import "github.com/go-utils/unum"

Go programming helpers for common maths needs; plus vectors, matrices and
quaternions.

## Usage

```go
const (
	PiDiv180  = math.Pi / 180
	PiDiv360  = math.Pi / 360
	PiHalfDiv = 0.5 / math.Pi
)
```

```go
var (
	Infinity    = math.Inf(1)
	NegInfinity = math.Inf(-1)

	Epsilon32 = Nextafter32(1, Infinity) - 1
	Epsilon64 = math.Nextafter(1, Infinity) - 1
)
```

#### func  AllEqual

```go
func AllEqual(test float64, vals ...float64) bool
```
Returns whether all `vals` equal `test`.

#### func  Clamp

```go
func Clamp(val, c0, c1 float64) float64
```
Clamps `val` between `c0` and `c1`.

#### func  DegToRad

```go
func DegToRad(degrees float64) float64
```
Converts the specified `degrees` to radians.

#### func  Din1

```go
func Din1(val, max float64) float64
```
Returns the "normalized ratio" of `val` to `max`. Example: for `max = 900` and
`val = 300`, returns `0.33333..`.

#### func  Fin1

```go
func Fin1(val, max float32) float32
```
Returns the "normalized ratio" of `val` to `max`. Example: for `max = 900` and
`val = 300`, returns `0.33333..`.

#### func  IsEven

```go
func IsEven(val int) bool
```
Returns whether `val` is even.

#### func  IsInt

```go
func IsInt(val float64) bool
```
Returns whether `val` represents an integer.

#### func  IsMod0

```go
func IsMod0(v, m int) bool
```
Returns whether `math.Mod(v, m)` is 0.

#### func  Mat3Identities

```go
func Mat3Identities(mats ...*Mat3)
```
Calls the `Identity` method on all specified `mats`.

#### func  Mat4Identities

```go
func Mat4Identities(mats ...*Mat4)
```
Calls the `Identity` method on all specified `mats`.

#### func  Mini

```go
func Mini(v1, v2 int) int
```
Returns the smaller of two `int` values.

#### func  Mix

```go
func Mix(x, y, a float64) float64
```
Returns `x` if `a` is 0, or `y` if `a` is 1, or else a corresponding mix of both
if `a` is between 0 and 1.

#### func  Nextafter32

```go
func Nextafter32(x, y float64) (r float64)
```
https://groups.google.com/d/topic/golang-nuts/dVtKN8QLUNM/discussion

#### func  RadToDeg

```go
func RadToDeg(radians float64) float64
```
Converts the specified `radians` to degrees.

#### func  Round

```go
func Round(v float64) (fint float64)
```
Returns (the equivalent of) `math.Ceil(v)` if fraction >= 0.5, otherwise returns
(the equivalent of) `math.Floor(v)`.

#### func  Saturate

```go
func Saturate(v float64) float64
```
Clamps `v` between 0 and 1.

#### func  Sign

```go
func Sign(v float64) (sign float64)
```
Returns -1 if `v` is negative, 1 if `v` is positive, or 0 if `v` is zero.

#### func  Step

```go
func Step(edge, x float64) (step int)
```
Returns 0 if `x` less-than `edge`, otherwise returns 1.

#### type Mat3

```go
type Mat3 [9]float64
```

Represents a 3x3 matrix.

```go
var (
	//	The 3x3 identity matrix.
	Mat3Identity Mat3
)
```

#### func  NewMat3Identity

```go
func NewMat3Identity() (mat *Mat3)
```
Returns a new 3x3 identity matrix.

#### func (*Mat3) Identity

```go
func (me *Mat3) Identity()
```
Sets this 3x3 matrix to `Mat3Identity`.

#### func (*Mat3) Transpose

```go
func (me *Mat3) Transpose()
```
Transposes this 3x3 matrix.

#### type Mat4

```go
type Mat4 [16]float64
```

Represents a 4x4 column-major matrix.

```go
var (
	//	The 4x4 identity matrix.
	Mat4Identity Mat4
)
```

#### func  NewMat4Add

```go
func NewMat4Add(a, b *Mat4) (mat *Mat4)
```
Returns a new `*Mat4` representing the result of adding `a` to `b`.

#### func  NewMat4Frustum

```go
func NewMat4Frustum(left, right, bottom, top, near, far float64) (mat *Mat4)
```
Returns a new `*Mat4` representing the specified frustum.

#### func  NewMat4Identity

```go
func NewMat4Identity() (mat *Mat4)
```
Returns a new `*Mat4` representing the identity matrix.

#### func  NewMat4Lookat

```go
func NewMat4Lookat(eyePos, lookTarget, upVec *Vec3) (mat *Mat4)
```
Returns a new `*Mat4` representing the "look-at matrix" computed from the
specified vectors.

#### func  NewMat4Mult1

```go
func NewMat4Mult1(m *Mat4, v float64) (mat *Mat4)
```
Returns a new `*Mat4` representing the result of multiplying all values in `m`
with `v`.

#### func  NewMat4Mult4

```go
func NewMat4Mult4(one, two *Mat4) (mat *Mat4)
```
Returns a new `*Mat4` that represents the result of multiplying `one` with
`two`.

#### func  NewMat4MultN

```go
func NewMat4MultN(mats ...*Mat4) (mat *Mat4)
```
Returns a new `*Mat4` that represents the result of multiplying all specified
`mats` with one another.

#### func  NewMat4Orient

```go
func NewMat4Orient(lookTarget, worldUp *Vec3) (mat *Mat4)
```
Returns a new `*Mat4` representing the "orientation matrix" computed from the
specified vectors.

#### func  NewMat4Perspective

```go
func NewMat4Perspective(fovY, aspect, near, far float64) (mat *Mat4)
```
Returns a new `*Mat4` that represents the specified perspective-projection
matrix.

#### func  NewMat4RotationX

```go
func NewMat4RotationX(rad float64) (mat *Mat4)
```
Returns a new `*Mat4` that represents a rotation of `rad` radians around the X
axis.

#### func  NewMat4RotationY

```go
func NewMat4RotationY(rad float64) (mat *Mat4)
```
Returns a new `*Mat4` that represents a rotation of `rad` radians around the Y
axis.

#### func  NewMat4RotationZ

```go
func NewMat4RotationZ(rad float64) (mat *Mat4)
```
Returns a new `*Mat4` that represents a rotation of `rad` radians around the Z
axis.

#### func  NewMat4Scaling

```go
func NewMat4Scaling(vec *Vec3) (mat *Mat4)
```
Returns a new `*Mat4` that represents a transformation of "scale by `vec`".

#### func  NewMat4Sub

```go
func NewMat4Sub(a, b *Mat4) (mat *Mat4)
```
Returns a new `*Mat4` that represents `a` minus `b`.

#### func  NewMat4Translation

```go
func NewMat4Translation(vec *Vec3) (mat *Mat4)
```
Returns a new `*Mat4` that represents a transformation of "translate by `vec`".

#### func (*Mat4) Abs

```go
func (me *Mat4) Abs() (abs *Mat4)
```
Returns a new `*Mat4` with each cell representing the `math.Abs` value of the
respective corresponding cell in `me`.

#### func (*Mat4) Add

```go
func (me *Mat4) Add(mat *Mat4)
```
Adds `mat` to `me`.

#### func (*Mat4) Clear

```go
func (me *Mat4) Clear()
```
Zeroes all cells in `me`.

#### func (*Mat4) Clone

```go
func (me *Mat4) Clone() (mat *Mat4)
```
Returns a new `*Mat` containing a copy of `me`.

#### func (*Mat4) CopyFrom

```go
func (me *Mat4) CopyFrom(mat *Mat4)
```
Copies all cells from `mat` to `me`.

#### func (*Mat4) CopyTo

```go
func (me *Mat4) CopyTo(mat *Mat4)
```
Copies all cells from `me` to `mat`.

#### func (*Mat4) Frustum

```go
func (me *Mat4) Frustum(left, right, bottom, top, near, far float64)
```
Sets `me` to represent the specified frustum.

#### func (*Mat4) Identity

```go
func (me *Mat4) Identity()
```
Copies all cells from `Mat4Identity` to `me`.

#### func (*Mat4) Lookat

```go
func (me *Mat4) Lookat(eyePos, lookTarget, upVec *Vec3)
```
Sets `me` to the "look-at matrix" computed from the specified vectors.

#### func (*Mat4) Mult1

```go
func (me *Mat4) Mult1(v float64)
```
Multiplies all cells in `me` with `v`.

#### func (*Mat4) Orient

```go
func (me *Mat4) Orient(lookTarget, worldUp *Vec3)
```
Sets `me` to the "orientation matrix" computed from the specified vectors.

#### func (*Mat4) Perspective

```go
func (me *Mat4) Perspective(fovYDeg, a, n, f float64) (fovYRadHalf float64)
```
Sets `me` to the specified perspective-projection matrix.

`fovYRad` -- vertical field-of-view angle in radians. `a` -- aspect ratio. `n`
-- near-plane. `f` -- far-plane.

#### func (*Mat4) RotationX

```go
func (me *Mat4) RotationX(rad float64)
```
Sets `me` to a rotation matrix representing "rotate `rad` radians around the X
axis".

#### func (*Mat4) RotationY

```go
func (me *Mat4) RotationY(rad float64)
```
Sets `me` to a rotation matrix representing "rotate `rad` radians around the Y
axis".

#### func (*Mat4) RotationZ

```go
func (me *Mat4) RotationZ(rad float64)
```
Sets `me` to a rotation matrix representing "rotate `rad` radians around the Z
axis".

#### func (*Mat4) Scaling

```go
func (me *Mat4) Scaling(vec *Vec3)
```
Sets `me` to a transformation matrix representing "scale by `vec`"

#### func (*Mat4) SetFromMult4

```go
func (me *Mat4) SetFromMult4(one, two *Mat4)
```
Sets `me` to the result of multiplying `one` times `two`.

#### func (*Mat4) SetFromMultN

```go
func (me *Mat4) SetFromMultN(mats ...*Mat4)
```
Sets `me` to the result of multiplying all the specified `mats` with one
another.

#### func (*Mat4) SetFromTransposeOf

```go
func (me *Mat4) SetFromTransposeOf(mat *Mat4)
```
Sets `me` to the transpose of `mat`.

#### func (*Mat4) Sub

```go
func (me *Mat4) Sub(mat *Mat4)
```
Subtracts `mat` from `me`.

#### func (*Mat4) ToInverseMat3

```go
func (me *Mat4) ToInverseMat3(mat *Mat3)
```
Sets `mat` to the inverse of `me`.

#### func (*Mat4) Translation

```go
func (me *Mat4) Translation(vec *Vec3)
```
Sets `me` to a transformation matrix representing "translate by `vec`"

#### func (*Mat4) Transposed

```go
func (me *Mat4) Transposed() (mat *Mat4)
```
Returns the transpose of `me`.

#### type Vec2

```go
type Vec2 struct{ X, Y float64 }
```

A 2-dimensional vector.

#### func (*Vec2) Div

```go
func (me *Vec2) Div(vec *Vec2) *Vec2
```
Returns a new `*Vec2` that is the result of dividing `me` by `vec`.

#### func (*Vec2) Dot

```go
func (me *Vec2) Dot(vec *Vec2) float64
```
Returns the dot product of `me` and `vec`.

#### func (*Vec2) Length

```go
func (me *Vec2) Length() float64
```
Returns the 2D vector length of `me`.

#### func (*Vec2) Magnitude

```go
func (me *Vec2) Magnitude() float64
```
Returns the 2D vector magnitude of `me`.

#### func (*Vec2) Mult

```go
func (me *Vec2) Mult(vec *Vec2) *Vec2
```
Returns a new `*Vec2` that is the result of multiplying `me` with `vec`.

#### func (*Vec2) Normalize

```go
func (me *Vec2) Normalize()
```
Normalizes `me` in-place.

#### func (*Vec2) Normalized

```go
func (me *Vec2) Normalized() *Vec2
```
Returns a new `*Vec2` that is the normalized representation of `me`.

#### func (*Vec2) NormalizedScaled

```go
func (me *Vec2) NormalizedScaled(factor float64) (vec *Vec2)
```
Returns a new `*Vec2` that is the normalized representation of `me` scaled by
`factor`.

#### func (*Vec2) Scale

```go
func (me *Vec2) Scale(factor float64)
```
Multiplies all components in `me` with `factor`.

#### func (*Vec2) Scaled

```go
func (me *Vec2) Scaled(factor float64) *Vec2
```
Returns a new `*Vec2` that represents `me` scaled by `factor`.

#### func (*Vec2) String

```go
func (me *Vec2) String() string
```
Returns a human-readable (imprecise) `string` representation of `me`.

#### func (*Vec2) Sub

```go
func (me *Vec2) Sub(vec *Vec2) *Vec2
```
Returns a new `*Vec2` that represents `me` minus `vec`.

#### type Vec3

```go
type Vec3 struct {
	X, Y, Z float64
}
```

Represents a 3-dimensional vector.

#### func (*Vec3) AbsMax

```go
func (me *Vec3) AbsMax() float64
```
Returns the `math.Max` of the `math.Abs` values of all 3 components in `me`.

#### func (*Vec3) Add

```go
func (me *Vec3) Add(vec *Vec3)
```
Adds `vec` to `me` in-place.

#### func (*Vec3) Add1

```go
func (me *Vec3) Add1(val float64)
```
Adds `val` to all 3 components of `me`.

#### func (*Vec3) Add3

```go
func (me *Vec3) Add3(x, y, z float64)
```
Adds the specified 3 components to the respective components in `me`.

#### func (*Vec3) AddTo

```go
func (me *Vec3) AddTo(vec *Vec3) *Vec3
```
Returns the sum of `me` and `vec`.

#### func (*Vec3) AllEqual

```go
func (me *Vec3) AllEqual(val float64) bool
```
Returns whether all 3 components in `me` equal `val`.

#### func (*Vec3) AllGreaterOrEqual

```go
func (me *Vec3) AllGreaterOrEqual(vec *Vec3) bool
```
Returns whether all 3 components in `me` are greater than or equal to their
respective component counterparts in `vec`.

#### func (*Vec3) AllInRange

```go
func (me *Vec3) AllInRange(min, max float64) bool
```
Returns whether all 3 components in `me` are greater than (or equal to) `min`,
and also less than `max`.

#### func (*Vec3) AllInside

```go
func (me *Vec3) AllInside(min, max *Vec3) bool
```
Returns whether all 3 components in `me` are greater than `min`, and also less
than `max`.

#### func (*Vec3) AllLessOrEqual

```go
func (me *Vec3) AllLessOrEqual(vec *Vec3) bool
```
Returns whether all 3 components in `me` are less than their respective
component counterparts in `vec`.

#### func (*Vec3) Clamp

```go
func (me *Vec3) Clamp(min, max *Vec3)
```
Clamps each component in `me` between the respective corresponding counter-part
component in `min` and `max`.

#### func (*Vec3) Clear

```go
func (me *Vec3) Clear()
```
Zeroes all 3 components in `me`.

#### func (*Vec3) Cross

```go
func (me *Vec3) Cross(vec *Vec3) *Vec3
```
Returns a new `*Vec3` that represents the cross-product of `me` and `vec`.

#### func (*Vec3) CrossNormalized

```go
func (me *Vec3) CrossNormalized(vec *Vec3) (r *Vec3)
```
Returns a new `*Vec` that represents the cross-product of `me` and `vec`,
normalized.

#### func (*Vec3) DistanceFrom

```go
func (me *Vec3) DistanceFrom(vec *Vec3) float64
```
Returns the distance of `me` from `vec`.

#### func (*Vec3) DistanceFromZero

```go
func (me *Vec3) DistanceFromZero() float64
```
Returns the distance of `me` from "the center" (zero).

#### func (*Vec3) Div

```go
func (me *Vec3) Div(vec *Vec3) *Vec3
```
Returns a new `*Vec3` that represents `me` divided by `vec`.

#### func (*Vec3) Div1

```go
func (me *Vec3) Div1(val float64) *Vec3
```
Returns a new `*Vec3` that represents all 3 components in `me`, each divided by
`val`.

#### func (*Vec3) Dot

```go
func (me *Vec3) Dot(vec *Vec3) float64
```
Returns the dot-product of `me` and `vec`.

#### func (*Vec3) DotSub

```go
func (me *Vec3) DotSub(vec1, vec2 *Vec3) float64
```
Returns the dot-product of `me` and (`vec1` minus `vec2`).

#### func (*Vec3) Inv

```go
func (me *Vec3) Inv() *Vec3
```
Returns a new `*Vec3` representing the inverse of `me`.

#### func (*Vec3) Length

```go
func (me *Vec3) Length() float64
```
Returns the 3D vector length of `me`.

#### func (*Vec3) Magnitude

```go
func (me *Vec3) Magnitude() float64
```
Returns the 3D vector magnitude of `me`.

#### func (*Vec3) MakeFinite

```go
func (me *Vec3) MakeFinite(vec *Vec3)
```
Sets each component in `me` to the corresponding component in `vec` only if the
former is infinity.

#### func (*Vec3) ManhattanDistanceFrom

```go
func (me *Vec3) ManhattanDistanceFrom(vec *Vec3) float64
```
Returns the "manhattan distance" of `me` from `vec`.

#### func (*Vec3) Max

```go
func (me *Vec3) Max() float64
```
Returns the largest of the 3 components in `me`.

#### func (*Vec3) Min

```go
func (me *Vec3) Min() float64
```
Returns the smallest of the 3 components in `me`.

#### func (*Vec3) Mult

```go
func (me *Vec3) Mult(vec *Vec3) *Vec3
```
Returns a new `*Vec3` that represents `me` multiplied with `vec`.

#### func (*Vec3) Negate

```go
func (me *Vec3) Negate()
```
Reverses the signs of all 3 vector components in `me`.

#### func (*Vec3) Negated

```go
func (me *Vec3) Negated() *Vec3
```
Returns a new `*Vec` with each component representing the negative (sign
inverted) corresponding component in `me`.

#### func (*Vec3) Normalize

```go
func (me *Vec3) Normalize()
```
Normalizes `me` in-place.

#### func (*Vec3) Normalized

```go
func (me *Vec3) Normalized() (vec *Vec3)
```
Returns a new `*Vec3` that represents `me`, normalized.

#### func (*Vec3) NormalizedScaled

```go
func (me *Vec3) NormalizedScaled(factor float64) (vec *Vec3)
```
Returns a new `*Vec3` that represents `me` normalized, then scaled by `factor`.

#### func (*Vec3) RotateDeg

```go
func (me *Vec3) RotateDeg(angleDeg float64, axis *Vec3)
```
Rotates `me` `angleDeg` degrees around the specified `axis`.

#### func (*Vec3) RotateRad

```go
func (me *Vec3) RotateRad(angleRad float64, axis *Vec3)
```
Rotates `me` `angleRad` radians around the specified `axis`.

#### func (*Vec3) Scale

```go
func (me *Vec3) Scale(factor float64)
```
Scales `me` by `factor`.

#### func (*Vec3) ScaleAdd

```go
func (me *Vec3) ScaleAdd(factor, add *Vec3)
```
Scales `me` by `factor`, then adds `add`.

#### func (*Vec3) Scaled

```go
func (me *Vec3) Scaled(factor float64) *Vec3
```
Returns a new `*Vec3` that represents `me` scaled by `factor`.

#### func (*Vec3) ScaledAdded

```go
func (me *Vec3) ScaledAdded(factor float64, add *Vec3) *Vec3
```
Returns a new `*Vec3` that represents `me` scaled by `factor`, then `add` added.

#### func (*Vec3) Set

```go
func (me *Vec3) Set(x, y, z float64)
```
Sets all 3 vector components in `me` to the corresponding respective specified
value.

#### func (*Vec3) SetFromAdd

```go
func (me *Vec3) SetFromAdd(vec1, vec2 *Vec3)
```
Sets `me` to the result of adding `vec1` and `vec2`.

#### func (*Vec3) SetFromAddAdd

```go
func (me *Vec3) SetFromAddAdd(a, b, c *Vec3)
```
`me = a + b + c`

#### func (*Vec3) SetFromAddMult

```go
func (me *Vec3) SetFromAddMult(add, mul1, mul2 *Vec3)
```
`me = add + (mul1 * mul2)`

#### func (*Vec3) SetFromAddScaled

```go
func (me *Vec3) SetFromAddScaled(vec1, vec2 *Vec3, mul float64)
```
`me = vec2 + vec2 * mul`

#### func (*Vec3) SetFromAddSub

```go
func (me *Vec3) SetFromAddSub(a, b, c *Vec3)
```
`me = a + b - c`

#### func (*Vec3) SetFromCos

```go
func (me *Vec3) SetFromCos(vec *Vec3)
```
Sets each vector component in `me` to the `math.Cos` of the respective
corresponding component in `vec`.

#### func (*Vec3) SetFromCross

```go
func (me *Vec3) SetFromCross(vec *Vec3)
```
Sets `me` to the cross-product of `me` and `vec`.

#### func (*Vec3) SetFromCrossOf

```go
func (me *Vec3) SetFromCrossOf(one, two *Vec3)
```
Sets `me` to the cross-product of `one` and `two`.

#### func (*Vec3) SetFromDegToRad

```go
func (me *Vec3) SetFromDegToRad(deg *Vec3)
```
Sets each vector component in `me` to the radian equivalent of the degree angle
stored in the respective corresponding component of `vec`.

#### func (*Vec3) SetFromEpsilon32

```go
func (me *Vec3) SetFromEpsilon32()
```
Sets each vector component in `me` to `Epsilon32` if it is 0 or greater but
smaller than `Epsilon32`.

#### func (*Vec3) SetFromEpsilon64

```go
func (me *Vec3) SetFromEpsilon64()
```
Sets each vector component in `me` to `Epsilon64` if it is 0 or greater but
smaller than `Epsilon64`.

#### func (*Vec3) SetFromInv

```go
func (me *Vec3) SetFromInv(vec *Vec3)
```
Sets `me` to the inverse of `vec`.

#### func (*Vec3) SetFromMult

```go
func (me *Vec3) SetFromMult(v1, v2 *Vec3)
```
`me = v1 * v2`

#### func (*Vec3) SetFromNegated

```go
func (me *Vec3) SetFromNegated(vec *Vec3)
```
`me = -vec`

#### func (*Vec3) SetFromNormalized

```go
func (me *Vec3) SetFromNormalized(vec *Vec3)
```
Sets `me` to `vec` normalized.

#### func (*Vec3) SetFromRotation

```go
func (me *Vec3) SetFromRotation(pos, rotCos, rotSin *Vec3)
```
Sets `me` to `pos` rotated as expressed in `rotCos` and `rotSin`.

#### func (*Vec3) SetFromScaled

```go
func (me *Vec3) SetFromScaled(vec *Vec3, mul float64)
```
`me = vec * mul`

#### func (*Vec3) SetFromScaledSub

```go
func (me *Vec3) SetFromScaledSub(vec1, vec2 *Vec3, mul float64)
```
`me = (vec1 - vec2) * mul`

#### func (*Vec3) SetFromSin

```go
func (me *Vec3) SetFromSin(vec *Vec3)
```
Sets each vector component in `me` to the `math.Sin` of the respective
corresponding component in `vec`.

#### func (*Vec3) SetFromStep

```go
func (me *Vec3) SetFromStep(edge float64, vec, v0, v1 *Vec3)
```
Component-wise, set `me` to `v0` if it is less than `edge`, else `v1`.

#### func (*Vec3) SetFromSub

```go
func (me *Vec3) SetFromSub(vec1, vec2 *Vec3)
```
`me = vec1 - vec2`.

#### func (*Vec3) SetFromSubAdd

```go
func (me *Vec3) SetFromSubAdd(a, b, c *Vec3)
```
`me = a - b + c`

#### func (*Vec3) SetFromSubMult

```go
func (me *Vec3) SetFromSubMult(sub1, sub2, mul *Vec3)
```
`me = (sub1 - sub2) * mul`

#### func (*Vec3) SetFromSubScaled

```go
func (me *Vec3) SetFromSubScaled(v1, v2 *Vec3, v2Scale float64)
```
`me = v1 - v2 * v2Scale`

#### func (*Vec3) SetFromSubSub

```go
func (me *Vec3) SetFromSubSub(a, b, c *Vec3)
```
`me = a - b - c`

#### func (*Vec3) SetToMax

```go
func (me *Vec3) SetToMax()
```
Sets all 3 vector components in `me` to `math.MaxFloat64`.

#### func (*Vec3) SetToMin

```go
func (me *Vec3) SetToMin()
```
Sets all 3 vector components in `me` to `-math.MaxFloat64`.

#### func (*Vec3) Sign

```go
func (me *Vec3) Sign() *Vec3
```
Returns a new `*Vec3` with each vector component indicating the sign (-1, 1 or
0) of the respective corresponding component in `me`.

#### func (*Vec3) String

```go
func (me *Vec3) String() string
```
Returns a human-readable (imprecise) `string` representation of `me`.

#### func (*Vec3) Sub

```go
func (me *Vec3) Sub(vec *Vec3) *Vec3
```
Returns a new `*Vec3` that represents `me` minus `vec`.

#### func (*Vec3) SubDivMult

```go
func (me *Vec3) SubDivMult(sub, div, mul *Vec3) *Vec3
```
Returns a new `*Vec3` that represents `((me - sub) / div) * mul`.

#### func (*Vec3) SubDot

```go
func (me *Vec3) SubDot(vec *Vec3) float64
```
Returns the dot-product of `me` minus `vec`.

#### func (*Vec3) SubFloorDivMult

```go
func (me *Vec3) SubFloorDivMult(floorDiv, mul float64) *Vec3
```
Returns a new `*Vec3` that represents `mul * math.Floor(me / floorDiv)`.

#### func (*Vec3) SubFrom

```go
func (me *Vec3) SubFrom(val float64) *Vec3
```
Returns a new `*Vec3` that represents `val` minus `me`.

#### func (*Vec3) SubScaled

```go
func (me *Vec3) SubScaled(vec *Vec3, val float64) *Vec3
```
Returns a new `*Vec3` that represents `(me - vec) * val`.

#### func (*Vec3) SubVec

```go
func (me *Vec3) SubVec(vec *Vec3)
```
Subtracts `vec` from `me`.

#### func (*Vec3) Times

```go
func (me *Vec3) Times(x, y, z float64) *Vec3
```
Returns a new `*Vec3` with each component in `me` multiplied by the respective
corresponding specified factor.

#### func (*Vec3) TransformCoord

```go
func (me *Vec3) TransformCoord(mat *Mat4)
```
Transform coordinate vector `me` according to the specified `*Mat4`.

#### func (*Vec3) TransformNormal

```go
func (me *Vec3) TransformNormal(mat *Mat4, absMat bool)
```
Transform normal vector `me` according to the specified `*Mat4`.

#### type Vec4

```go
type Vec4 struct {
	//	X, Y, Z
	Vec3

	//	The 4th vector component.
	W float64
}
```

Represents an arbitrary 4-dimensional vector or a Quaternion.

#### func (*Vec4) Clone

```go
func (me *Vec4) Clone() (q *Vec4)
```
Returns a new `*Vec4` containing a copy of `me`.

#### func (*Vec4) Conjugate

```go
func (me *Vec4) Conjugate()
```
Negates the `X`, `Y`, `Z` components in `me`, but not `W`.

#### func (*Vec4) Conjugated

```go
func (me *Vec4) Conjugated() (v *Vec4)
```
Returns a new `*Vec4` that represents `me` conjugated.

#### func (*Vec4) Length

```go
func (me *Vec4) Length() float64
```
Returns the 4D vector length of `me`.

#### func (*Vec4) Magnitude

```go
func (me *Vec4) Magnitude() float64
```
Returns the 4D vector magnitude of `me`.

#### func (*Vec4) MultMat4Vec3

```go
func (me *Vec4) MultMat4Vec3(mat *Mat4, vec *Vec3)
```
Sets `me` to the result of multiplying the specified `*Mat4` with the specified
`*Vec3`.

#### func (*Vec4) MultMat4Vec4

```go
func (me *Vec4) MultMat4Vec4(mat *Mat4, vec *Vec4)
```
Sets `me` to the result of multiplying the specified `*Mat4` with the specified
`*Vec4`.

#### func (*Vec4) Normalize

```go
func (me *Vec4) Normalize()
```
Normalizes `me` according to `me.Magnitude`.

#### func (*Vec4) NormalizeFrom

```go
func (me *Vec4) NormalizeFrom(magnitude float64)
```
Normalizes `me` according to the specified `magnitude`.

#### func (*Vec4) Normalized

```go
func (me *Vec4) Normalized() *Vec4
```
Returns a new `*Vec4` that represents `me` normalized according to
`me.Magnitude`.

#### func (*Vec4) Scale

```go
func (me *Vec4) Scale(v float64)
```
Scales all 4 vector components in `me` by factor `v`.

#### func (*Vec4) SetFromConjugated

```go
func (me *Vec4) SetFromConjugated(c *Vec4)
```
Sets `me` to `c` conjugated.

#### func (*Vec4) SetFromMult

```go
func (me *Vec4) SetFromMult(l, r *Vec4)
```
Applies various 4D vector component computations of `l` and `r` to `me`, as
needed by the `Vec3.RotateRad` method.

#### func (*Vec4) SetFromMult3

```go
func (me *Vec4) SetFromMult3(q *Vec4, v *Vec3)
```
Applies various 4D vector component computations of `q` and `v` to `me`, as
needed by the `Vec3.RotateRad` method.

#### func (*Vec4) SetFromMultMat4

```go
func (me *Vec4) SetFromMultMat4(mat *Mat4)
```
Sets `me` to the result of multiplying the specified `*Mat4` with `me`.

#### func (*Vec4) String

```go
func (me *Vec4) String() string
```
Returns a human-readable (imprecise) `string` representation of `me`.

--
**godocdown** http://github.com/robertkrimen/godocdown
