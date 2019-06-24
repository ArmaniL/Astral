package main ;

import (
"image/color"
"math/rand"
"log"
"github.com/hajimehoshi/ebiten"
"github.com/hajimehoshi/ebiten/ebitenutil"
  "github.com/hajimehoshi/ebiten/text"
  "github.com/golang/freetype/truetype"
  "golang.org/x/image/font"
  "io/ioutil"
)

const(
screenWidth=470
screenHeight=450
size=30
amountofstars int=10
)


func random(max,min float64)float64{
 return (rand.Float64()*max)+min

}

// ALl the variables
var (Player Ship;
 Stars [amountofstars]Star;
 mode int=0;
 fng  font.Face  ;
 v int;
 incr bool;
 bullet Bullet;
 )
func  update(screen *ebiten.Image) error {

if mode==0{  Startscreen(screen);
   }else  {
        ActualGame(screen);
   }

return nil
}


func Startscreen (screen *ebiten.Image)error{

if incr{  v+=1;
    if v>100{
        incr=false;
        } }else{
            v-=1;
            if v<40{
                incr=true;
            }
        }


if err:=screen.Fill(color.RGBA{255,255,255,255});err!=nil{
    log.Fatal(err)
}
for  i := 0; i < amountofstars; i++{

        Stars[i].Update();
        Stars[i].Draw(screen);

    }
 text.Draw(screen, "Astral", fng, screenWidth*0.4, screenHeight/3, color.Black)
text.Draw(screen, "Press SpaceBar", fng, screenWidth*0.3, screenHeight*0.8, color.RGBA{10,10,10,uint8(v)})
if ebiten.IsKeyPressed(ebiten.KeySpace)  {
                mode++
        }

    return nil


}


func ActualGame(screen *ebiten.Image)error{


if err:=screen.Fill(color.RGBA{255,255,255,255});err!=nil{
    log.Fatal(err)
}
for  i := 0; i < amountofstars; i++{

        Stars[i].Update();
        Stars[i].Draw(screen);

    }
Player.Draw(screen);
Player.Update();
bullet.Draw(screen)
bullet.x=Player.x+size/2
bullet.Update();
layouts[lmode-1](screen);
return nil;
}

 func main() {
     f, err := ebitenutil.OpenFile("res/georgiab.ttf")
        if err != nil {
                log.Fatal(err)
        }
        defer f.Close()

        b, err := ioutil.ReadAll(f)
        if err != nil {
                log.Fatal(err)
        }

        tt, err := truetype.Parse(b)
        if err != nil {
                log.Fatal(err)
        }

        const dpi = 72
     fng= truetype.NewFace(tt, &truetype.Options{
                Size:    24,
                DPI:     dpi,
                Hinting: font.HintingFull,
        })
    Player =Ship{float64(screenWidth/2),float64(screenHeight*0.8),3,10,0}
    bullet= Bullet{Player.x+(size/2),Player.y-20,4,14,false,Player,17}
    for  i := 0; i < amountofstars; i++{

        Stars[i]=NewStar();
    }

   for  i := 0; i < amountofenemies; i++{
        dancers[i]=Dancer{random(screenWidth,1),random(screenHeight*-1,-1),30.0,random(3,1),false,true,true}
        enemies[i]=Enemy{random(2,1),random(screenWidth,1),random(screenHeight*-1,-1),30.0,false,true}
    }

 	if err:=ebiten.Run(update, screenWidth, screenHeight, 1, "Astral");err != nil {
		log.Fatal(err)
	}
}
