package sprites

import (
	"github.com/hajimehoshi/ebiten"
	"math"
)

type Point struct {
	X float64
	Y float64
}

type ScaledImage struct {
	Img            *ebiten.Image
	ScaleX, ScaleY float64
}

type Dimensions struct {
	Width  int
	Height int
}

type Movement struct {
	Speed         float64
	Direction     float64 // expressed in degrees
	Gravity       float64
	downwardSpeed float64
	Orientation   float64 // expressed in degrees
	rotationSpeed float64
}

type Sprite struct {
	Point
	Dimensions
	Movement
	ScaledImage
}

func (s *Sprite) Center() (float64, float64) {
	return s.X + float64(s.Width)/2, s.Y + float64(s.Height)/2
}

func (s *Sprite) SetLocation(x_loc, y_loc float64) {
	s.X = x_loc
	s.Y = y_loc
}

func (s *Sprite) SetSize(w, h int) {
	s.Width = w
	s.Height = h
	s.updateImageScale()
}

func (s *Sprite) SetSpeed(speed float64) {
	s.Speed = speed
	s.Movement.rotationSpeed = s.Speed / 100
}

func (s *Sprite) SetDirection(direction float64) {
	s.Direction = direction
}

func (s *Sprite) SetGravity(gravity float64) {
	s.Gravity = gravity
}

func (s *Sprite) SetRotationSpeed(rspeed float64) {
	s.rotationSpeed = rspeed
}

func (s *Sprite) applyGravity() {
	s.downwardSpeed += s.Gravity
	s.Y += s.downwardSpeed
}
func (s *Sprite) Move() {
	s.X += s.Speed * math.Cos(s.Direction*math.Pi/180)
	s.Y -= s.Speed * math.Sin(s.Direction*math.Pi/180)
	s.applyGravity()
	s.Orientation += s.rotationSpeed
}

func (s *Sprite) Stop() {
	s.Speed = 0
	s.Direction = 0
	s.rotationSpeed = 0
}

func (s *Sprite) updateImageScale() {
	if s.Img == nil {
		return
	}
	s.ScaleX = float64(s.Width) / float64(s.Img.Bounds().Dx())
	s.ScaleY = float64(s.Height) / float64(s.Img.Bounds().Dy())
}

func (s *Sprite) SetImage(img *ebiten.Image) {
	s.Img = img
	s.updateImageScale()
}

func (s *Sprite) DrawingParameters() (*ebiten.Image, *ebiten.DrawImageOptions) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Rotate(float64(s.Orientation))
	op.GeoM.Scale(s.ScaleX, s.ScaleY)
	op.GeoM.Translate(float64(s.X), float64(s.Y))

	return s.Img, op
}

func NewSprite(x, y float64, Width, Height int) *Sprite {
	s := Sprite{}
	s.Img = nil
	s.SetLocation(x, y)
	s.SetSize(Width, Height)

	return &s
}

func (s *Sprite) Reset(x, y float64) {
	s.Stop()
	s.X = x
	s.Y = y
	s.Movement.downwardSpeed = 0
	s.Orientation = 0
}

func (s *Sprite) IsInside(x, y float64) bool {
	return x >= s.X && x <= s.X+float64(s.Width) && y >= s.Y && y < s.Y+float64(s.Height)
}

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
	SetLocation(float64, float64)
	SetSize(float64, float64)
	Reset(float64, float64)
}
