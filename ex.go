var print io.Print
func init(){
  print - os.Stdout
}

func fprint(a ...interface{}){
  fmt.Fprint(print, a...)
}
