package main;

import (
"github.com/hajimehoshi/ebiten"
)

const (

amountofenemies int=10
)
var (
    lmode=1;
	count int
 enemies [amountofenemies]Enemy;
 dancers [amountofenemies]Dancer;

	)
var layouts=[]func(screen *ebiten.Image){


//First Function
    func(screen *ebiten.Image){

  for  i := 0; i < amountofenemies; i++{

        enemies[i].isvisible=true;

        enemies[i].Update();
        enemies[i].Draw(screen);
        enemies[i].Collision(bullet);


    }
    if count==amountofenemies{
    	lmode=int(random(2,1))
    }

    }         ,

    //Second layout function
func(screen *ebiten.Image){
   for  i := 0; i < amountofenemies; i++{
         dancers[i].isvisible=true;
         dancers[i].Update();
         dancers[i].Draw(screen);
         dancers[i].Collision(bullet);
     }
     if count==amountofenemies{
     	lmode=int(random(2,1))
     }
     }}
