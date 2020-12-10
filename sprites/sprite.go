package sprites

import (
	"github.com/hajimehoshi/ebiten"
	"math"
)

type Point struct {
	x float64
	y float64
}

func (s *Point) Location() (float64, float64) {
	return s.x, s.y
}

func (s *Point) SetLocation(x_loc, y_loc float64) {
	s.x = x_loc
	s.y = y_loc
}

type ScaledImage struct {
	img            *ebiten.Image
	scaleX, scaleY float64
}

func (s *ScaledImage) UpdateImageScale(w, h int) {
	if s.img == nil {
		return
	}
	bounds := s.img.Bounds()
	s.scaleX = float64(w) / float64(bounds.Dx())
	s.scaleY = float64(h) / float64(bounds.Dy())
}

func (s *ScaledImage) Image() *ebiten.Image {
	return s.img
}

func (s *ScaledImage) SetImage(img *ebiten.Image) {
	s.img = img
}

func (s *ScaledImage) Scales() (float64, float64) {
	return s.scaleX, s.scaleY
}

type Dimensions struct {
	width  int
	height int
}

func (s *Dimensions) Size() (int, int) {
	return s.width, s.height
}
func (s *Dimensions) SetSize(w, h int) {
	s.width = w
	s.height = h
}

type Movement struct {
	speed         float64
	direction     float64 // expressed in degrees
	gravity       float64
	downwardSpeed float64
	orientation   float64 // expressed in degrees
	rotationSpeed float64
}

func (s *Movement) Speed() float64 {
	return s.speed
}

func (s *Movement) SetSpeed(speed float64) {
	s.speed = speed
}

func (s *Movement) Direction() float64 {
	return s.direction
}

func (s *Movement) SetDirection(direction float64) {
	s.direction = direction
}

func (s *Movement) Gravity() float64 {
	return s.gravity
}

func (s *Movement) SetGravity(gravity float64) {
	s.gravity = gravity
}

func (s *Movement) Rotationspeed() float64 {
	return s.rotationSpeed
}

func (s *Movement) SetRotationSpeed(rspeed float64) {
	s.rotationSpeed = rspeed
}

func (s *Movement) ApplyGravity() {
	s.downwardSpeed += s.gravity
}

func (s *Movement) Stop() {
	s.SetSpeed(0)
	s.SetDirection(0)
	s.SetGravity(0)
	s.SetOrientation(0)
	s.SetRotationSpeed(0)
}

type Sprite struct {
	Point
	Dimensions
	Movement
	ScaledImage
}

func (s *Sprite) SetImage(img *ebiten.Image) {
	s.ScaledImage.SetImage(img)
	s.ScaledImage.UpdateImageScale(s.Size())
}

func (s *Sprite) applyGravity() {
	s.Movement.ApplyGravity()
	s.downwardSpeed += s.Gravity()
	s.Point.y += s.downwardSpeed
}

func (s *Sprite) Move() {
	s.Point.x += s.Movement.speed * math.Cos(s.Movement.direction*math.Pi/180)
	s.Point.y -= s.Movement.speed * math.Sin(s.Movement.direction*math.Pi/180)
	s.applyGravity()
	s.orientation += s.rotationSpeed
}

func (s *Sprite) Center() (float64, float64) {
	return s.x + float64(s.Dimensions.width)/2, s.y + float64(s.Dimensions.height)/2
}

func (s *Sprite) Size() (int, int) {
	return s.Dimensions.width, s.Dimensions.height
}

func (s *Sprite) SetSize(w, h int) {
	s.Dimensions.SetSize(w, h)
	s.ScaledImage.UpdateImageScale(s.Size())
}

func (s *Sprite) DrawingParameters() (*ebiten.Image, *ebiten.DrawImageOptions) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(float64(s.Movement.orientation))
	op.GeoM.Scale(s.ScaledImage.Scales())
	x, y := s.Size()
	op.GeoM.Translate(float64(x), float64(y))

	return s.ScaledImage.Image(), op
}

func NewSprite(x, y float64, Width, Height int) *Sprite {
	s := Sprite{}
	s.ScaledImage = ScaledImage{nil, 0, 0}
	s.SetLocation(x, y)
	s.SetSize(Width, Height)

	return &s
}

func (s *Sprite) Reset(x, y float64) {
	s.Stop()
	s.SetLocation(x, y)
	s.Orientation = 0
	s.Gravity = 0
}

func (s *Sprite) IsInside(x, y float64) bool {
	return x >= s.x && x <= s.x+float64(s.Width) && y >= s.Y && y < s.Y+float64(s.Height)
}

func (s *Sprite) Align(d Dimensions)

type CollisionDetection interface {
	Center() (float64, float64)
	IsInside(x, y float64) bool
}

func IsTouching(s1 CollisionDetection, s2 CollisionDetection) bool {
	return s1.IsInside(s2.Center()) || s2.IsInside(s1.Center())
}

type Drawable interface {
	SetImage(*ebiten.Image)
	DrawingParameters() (*ebiten.Image, *ebiten.DrawImageOptions)
}

type Moveable interface {
	SetSpeed(float64)
	SetAngle(float64)
	SetRotationSpeed(float64)
	SetGravity(float64)
	Move()
	Stop()
}

type Playable interface {
	Location() (float64, float64)
	Size() (int, int)
	SetLocation(float64, float64)
	SetSize(float64, float64)
	Reset(float64, float64)
}
