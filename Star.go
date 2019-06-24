package main;
import (

    "image/color"
  "github.com/hajimehoshi/ebiten"
  "github.com/hajimehoshi/ebiten/ebitenutil"
)
type Star struct{

     x float64
     y float64
     size float64
     speed float64


}







func (star *Star)Update() {

if (star.y<screenHeight){

    star.y+=star.speed

}else{

star.x=random(float64(screenWidth),0)
star.y=0;
 }

}



func NewStar()Star{

    return Star{random(float64(screenWidth),0),random(float64(-2),-20),random(1,0.5),random(6,1) }
}

func (star *Star)Draw( screen *ebiten.Image){
ebitenutil.DrawRect(screen,star.x,star.y,star.size,star.size,color.RGBA{2,2,2,254});
ebitenutil.DrawLine(screen,star.x+star.size/4,star.y,star.x+star.size/2,star.y-(star.size+4),color.RGBA{2,2,2,100});
}
