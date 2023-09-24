package main



type data struct{
  length, breath,radius float64

}



func dCon(l,b,r float64)*data{
  return &data{
    length:l,
    breath:b,
    radius:r,
  }
}
