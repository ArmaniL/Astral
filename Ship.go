package main;
import (
"github.com/hajimehoshi/ebiten/ebitenutil"
"image/color"
"github.com/hajimehoshi/ebiten"
)

type Ship struct{
     x float64
     y float64
     speed float64
     health float64
     score float64
}



func (s *Ship)Update(){

 if ebiten.IsKeyPressed(ebiten.KeyLeft) && s.x>0 {
                s.x-=s.speed
        }

 if ebiten.IsKeyPressed(ebiten.KeyRight) && (s.x+size)<screenWidth {
                s.x+=s.speed
        }


}


func (s *Ship)Draw(  screen *ebiten.Image){
ebitenutil.DrawLine(screen,s.x,s.y+size,s.x+size/2,s.y,color.RGBA{1,1,1,255})
 ebitenutil.DrawLine(screen,  s.x+size/2,s.y, s.x+size, s.y+size , color.RGBA{1,1,1,255})
ebitenutil.DrawLine(screen, s.x+size, s.y+size, s.x, s.y+size , color.RGBA{1,1,1,255})
}
