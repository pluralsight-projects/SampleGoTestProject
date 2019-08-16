package main

import (
  "log"
  "github.com/hajimehoshi/ebiten"
  "github.com/hajimehoshi/ebiten/ebitenutil"
)

func updated(screen *ebiten.Image) error {
  if ebiten.IsDrawingSkipped() {
    return nil
  }
  ebitenutil.DebugPrint(screen, "Hello, World!")
  return nil
}

func main() {
  if err := ebiten.Run(updated, 320, 240, 2, "Hello, World!"); err != nil {
    log.Fatal(err)
  }
}