package sprites

import (
	"github.com/hajimehoshi/ebiten"
	"math"
)

type Point struct {
	x float64
	y float64
}

type ScaledImage struct {
	img            *ebiten.Image
	scaleX, scaleY float64
}

type Dimensions struct {
	width  int
	height int
}

type Movement struct {
	speed         float64
	direction     float64 // expressed in degrees
	gravity       float64
	downwardSpeed float64
	orientation   float64 // expressed in degrees
	rotationSpeed float64
}

type Sprite struct {
	Point
	Dimensions
	Movement
	ScaledImage
}

func (s *Sprite) Center() (float64, float64) {
	return s.x + float64(s.width)/2, s.y + float64(s.height)/2
}

func (s *Sprite) SetLocation(x_loc, y_loc float64) {
	s.x = x_loc
	s.y = y_loc
}

func (s *Sprite) SetSize(w, h int) {
	s.width = w
	s.height = h
	s.updateImageScale()
}

func (s *Sprite) SetSpeed(speed float64) {
	s.speed = speed
}

func (s *Sprite) SetDirection(direction float64) {
	s.direction = direction
}

func (s *Sprite) SetGravity(gravity float64) {
	s.gravity = gravity
}

func (s *Sprite) SetRotationSpeed(rspeed float64) {
	s.rotationSpeed = rspeed
}

func (s *Sprite) applyGravity() {
	s.downwardSpeed += s.gravity
	s.y += s.downwardSpeed
}
func (s *Sprite) Move() {
	s.x += s.speed * math.Sin(s.direction)
	s.y += s.speed * math.Cos(s.direction)
	s.applyGravity()
	s.orientation += s.rotationSpeed
}

func (s *Sprite) Stop() {
	s.speed = 0
	s.direction = 0
	s.gravity = 0
	s.rotationSpeed = 0
}

func (s *Sprite) updateImageScale() {
	s.scaleX = float64(s.width) / float64(s.img.Bounds().Dx())
	s.scaleY = float64(s.height) / float64(s.img.Bounds().Dy())
}

func (s *Sprite) SetImage(img *ebiten.Image) {
	s.img = img
	s.updateImageScale()
}

func (s *Sprite) DrawingParameters() (*ebiten.Image, *ebiten.DrawImageOptions) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(float64(s.orientation))
	op.GeoM.Scale(s.scaleX, s.scaleY)
	op.GeoM.Translate(float64(s.x), float64(s.y))

	return s.img, op
}

func NewSprite(x, y float64, width, height int, img *ebiten.Image) *Sprite {
	s := Sprite{}
	s.SetLocation(x, y)
	s.img = img
	s.SetSize(width, height)

	return &s
}

func (s *Sprite) Reset(x, y float64) {
	s.Stop()
	s.x = x
	s.y = y
	s.orientation = 0
	s.gravity = 0
}

func (s *Sprite) IsInside(x, y float64) bool {
	return x >= s.x && x <= s.x+float64(s.width) && y >= s.y && y < s.y+float64(s.height)
}

func IsTouching(s1 *Sprite, s2 *Sprite) bool {
	return s1.IsInside(s2.Center()) || s2.IsInside(s1.Center())
}

type CollisionDetection interface {
	Center() (float64, float64)
	IsInside(x, y float64) bool
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
	SetLocation(float64, float64)
	SetSize(float64, float64)
	Reset(float64, float64)
}
