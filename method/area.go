package main



func(data *data)area()float64{

  total := (data.breath * data.length)

  return total
} 


func (data *data)perimeter ()float64{
  return 2*(data.length + data.breath)
}
