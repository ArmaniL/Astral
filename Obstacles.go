package main ;

import (
"image/color"
"github.com/hajimehoshi/ebiten"
"github.com/hajimehoshi/ebiten/ebitenutil"
)



type element interface{

Update()
Draw(image *ebiten.Image)

}



//sideways x shaped,teleport
type Teleporter struct{
   x float64
	y float64
  health float64
}


//shootsss at space ship
type Shooter struct{
  x float64
  y float64
  health float64
  size float64
  bb   Bullet;
}




func (s  *Shooter)Draw(screen *ebiten.Image) {
ebitenutil.DrawLine(screen,s.x,s.y,s.x+s.size,s.y,color.RGBA{0,0,0,0});
ebitenutil.DrawLine(screen,s.x+2*s.size/5,s.y,s.x+s.size/2,s.y+20,color.RGBA{0,0,0,0});
ebitenutil.DrawLine(screen,s.x+s.size/2,s.y+20,s.x+3*s.size/5,s.y,color.RGBA{0,0,0,0});
}




func (s *Shooter)Update(p Ship){
s.y++;
if p.x>s.x{
//if player is right and shooter is left
  s.x++;
  }else if p.x<s.x{
s.x--;
  }



}













func (t *Teleporter)Draw() {

}




func (t *Teleporter)Update(){


}
