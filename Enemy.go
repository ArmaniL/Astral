package main ;

import (
"image/color"
"github.com/hajimehoshi/ebiten"
"github.com/hajimehoshi/ebiten/ebitenutil"
)






// x shpaed very basic
type Enemy struct{
    health float64
    x float64
    y float64
    size float64
    isvisible bool
    ismoving bool
}
func (e  *Enemy )Draw(screen *ebiten.Image){
  if e.isvisible{
	ebitenutil.DrawLine(screen,e.x,e.y,e.x+e.size,e.y+e.size,color.RGBA{0,0,0,255});
    ebitenutil.DrawLine(screen,e.x,e.y+e.size,e.x+e.size,e.y,color.RGBA{0,0,0,255});
}
}


func (e  *Enemy)Update(){
if e.isvisible && e.isvisible{
    if e.y>=screenHeight{
    	e.Reset();
    }
	e.y++;
}
}
func(e *Enemy)Reset(){
e.isvisible=true;
e.ismoving=true;
e.x=random(float64(screenWidth-30),0)
e.y=-1*random(float64(0.8*screenWidth),0)
}


func (e *Enemy)Collision(b Bullet){





halfway:=b.x+b.length/2;

if  (halfway>=e.x  &&  halfway<=e.x+e.size) && ( b.y>= e.y && b.y<= e.y+e.size) {

b.isFired=false;
count++;
      e.health-=1;
      if e.health<=0{
         e.isvisible=false;
         e.Reset();
}


}}