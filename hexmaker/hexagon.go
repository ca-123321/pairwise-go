package hexmaker

import (
  "github.com/fogleman/gg"
  "github.com/golang/freetype/truetype"
  "golang.org/x/image/font/gofont/goregular"
  "math"
  "strconv"
)

func MakeHexagon() {
  const S = 1024
  dc := gg.NewContext(S, S)
  radius := 300.0

  // Set up font
  font, err := truetype.Parse(goregular.TTF)
  if err != nil {
    panic("")
  }
  face := truetype.NewFace(font, &truetype.Options{
    Size: 40,
  })
  dc.SetFontFace(face)

  // Draw hexagon
  dc.DrawRegularPolygon(6, S/2, S/2, radius, 0) // sides, x, y, radius, rotation
  dc.SetHexColor("016d9a")
  dc.Fill()

  // Draw circles and text
  for i := 0; i < 6; i++ {
    dc.Push()
    dc.SetHexColor("e42828")
    angle := gg.Radians(float64(i*60 - 30)) // -30 for edge-alignment
    x := S/2 + 200*math.Cos(angle)
    y := S/2 + 200*math.Sin(angle)
    dc.RotateAbout(angle, x, y)
    dc.DrawCircle(x, y, 50)
    dc.Fill()
    dc.SetRGB(1, 1, 1)
    dc.RotateAbout(gg.Radians(270), x, y)
    text := strconv.Itoa(i+1)
    dc.DrawStringAnchored(text, x, y, 0.5, 0.5)
    dc.Pop()
  }
  dc.SavePNG("out.png")
}

