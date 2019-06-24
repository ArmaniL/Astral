package main;
import (


"image/color"
"github.com/hajimehoshi/ebiten"
"github.com/hajimehoshi/ebiten/ebitenutil"
)

type Bullet struct{
  x  float64
  y  float64
  length float64
  width float64
  isFired  bool
  ship Ship
  acc float64
}

func (b *Bullet)Update(){

if ebiten.IsKeyPressed(ebiten.KeySpace){
    b.isFired=true;
}

if(b.isFired){
if(b.y<=0){
b.isFired=false
b.x=b.ship.x+2
b.y=b.ship.y-20;
 }else{
b.y-=b.acc
}
}else{
b.x=b.ship.x+2

}

}


func (b *Bullet)Draw(screen *ebiten.Image){
if(b.isFired){
    ebitenutil.DrawRect(screen,b.x,b.y,b.length,b.width,color.RGBA{1,1,1,255})
}
}
