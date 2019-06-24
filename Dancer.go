package main;
import (
  "image/color"
  "github.com/hajimehoshi/ebiten"
  "github.com/hajimehoshi/ebiten/ebitenutil"
)

//move around the screen randomly
type Dancer struct{
x float64
y float64
size float64
health float64
isvisible bool
ismoving bool
isright  bool
}


//Draws it like an hour glass
func (d  *Dancer)Draw(screen *ebiten.Image) {
ebitenutil.DrawLine(screen,d.x,d.y,d.x+d.size,d.y,color.RGBA{0,0,0,255});
ebitenutil.DrawLine(screen,d.x+d.size,d.y,d.x+d.size/2,d.y+d.size/2,color.RGBA{0,0,0,255});
ebitenutil.DrawLine(screen,d.x+d.size/2,d.y+d.size/2,d.x,d.y+d.size,color.RGBA{0,0,0,255});
ebitenutil.DrawLine(screen,d.x,d.y+d.size,d.x+d.size,d.y+d.size,color.RGBA{0,0,0,255});
ebitenutil.DrawLine(screen,d.x+d.size,d.y+d.size,d.x,d.y,color.RGBA{0,0,0,255});
}
//Handle Collision with bullet
func(d *Dancer)Collision(b Bullet){
  halfway:=b.x+b.length/2;
  if  (halfway>=d.x  &&  halfway<=d.x+d.size) && ( d.y>= d.y && d.y<= d.y+d.size) {
  b.isFired=false;
  count++;
        d.health-=1;
        if d.health<=0{
           d.isvisible=false;
           d.Reset();
  }
  }
}
func (d *Dancer)Reset(){
  d.isvisible=true;
  d.ismoving=true;
  d.isright=true;
  d.x=random(float64(screenWidth-30),0)
  d.y=-1*random(float64(0.8*screenWidth),0)
}
func (d *Dancer)Update(){

if d.isvisible && d.ismoving{
  d.y++;
if  d.x>=screenWidth{
d.isright=false;
}else{
  d.isright=true;
}

  if  d.isright{
    d.x++;
  }else{
d.x-=10;
  }
}}
