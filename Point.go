package main;
import (
"math"
)


type Point struct{

x float64
y float64

}


func(p *Point) Translate(xb float64, yb float64 ){
p.x+=xb
p.y+=yb
}




func(p *Point)  Rotate(degrees float64,fromx float64,fromy float64){

xb:=p.x-fromx;
yb:=p.y-fromy;
p.x=(xb)*math.Cos(degrees)-(yb)*math.Sin(degrees)
p.y=(xb)*math.Sin(degrees)+(yb)*math.Cos(degrees)
p.x+=fromx;
p.y+=fromy;


}
