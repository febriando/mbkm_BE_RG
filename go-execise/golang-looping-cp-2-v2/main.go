package main

import (
	"fmt"
)

// hello World => d_l_r_o_W o_l_l_e_H
func ReverseString(str string) string {
	var temp = len(str)
	var result string
	for i := temp - 1; i >= 0; i-- {
		if i == temp-1 {
			result += (string(str[i]))
		} else if (string(str[i+1])) == " " {
			result += (string(str[i]))
		} else if (string(str[i])) == " " {
			result += (string(str[i]))
		} else {
			result += ("_" + string(str[i]))
		}
	}

	return result // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	var lorem = "Lorem ipsum dolor sit amet consectetur adipiscing elit sed do eiusmod tempor incididunt ut labore et dolore magna aliqua Volutpat lacus laoreet non curabitur Amet risus nullam eget felis eget nunc Amet risus nullam eget felis eget Consectetur purus ut faucibus pulvinar elementum Suspendisse ultrices gravida dictum fusce ut placerat orci nulla pellentesque Eu nisl nunc mi ipsum Ac turpis egestas maecenas pharetra convallis Egestas tellus rutrum tellus pellentesque eu tincidunt tortor Nunc sed velit dignissim sodales ut Ut venenatis tellus in metus vulputate eu scelerisque felis Feugiat pretium nibh ipsum consequat nisl vel Porta non pulvinar neque laoreet suspendisse interdum consectetur Sed elementum tempus egestas sed sed risus Blandit volutpat maecenas volutpat blandit aliquam etiam erat Porta non pulvinar neque laoreet suspendisse interdum Vel fringilla est ullamcorper eget nulla facilisi etiam dignissim diam Nisl suscipit adipiscing bibendum est ultricies integer quis auctor Vitae tempus quam pellentesque nec nam aliquam sem Malesuada bibendum arcu vitae elementum curabitur vitae nunc sed Quam elementum pulvinar etiam non quam lacus suspendisse Nunc mattis enim ut tellus elementum sagittis vitae et Ultrices eros in cursus turpis massa tincidunt Morbi tincidunt ornare massa eget egestas purus viverra Massa ultricies mi quis hendrerit dolor magna eget Ut eu sem integer vitae justo eget magna fermentum iaculis Diam sit amet nisl suscipit adipiscing Convallis aenean et tortor at risus viverra Nulla posuere sollicitudin aliquam ultrices sagittis orci a Quam viverra orci sagittis eu volutpat odio facilisis mauris Libero volutpat sed cras ornare arcu dui vivamus Hendrerit dolor magna eget est lorem ipsum dolor sit amet Congue eu consequat ac felis donec et odio Etiam tempor orci eu lobortis elementum nibh tellus At auctor urna nunc id cursus Velit egestas dui id ornare arcu odio ut Blandit massa enim nec dui nunc mattis enim ut tellus Risus quis varius quam quisque Interdum velit laoreet id donec ultrices tincidunt arcu non sodales Vel risus commodo viverra maecenas"
	fmt.Println(ReverseString("Hello World"))
	fmt.Println(ReverseString("I am a student"))
	fmt.Println(ReverseString("I am a stranger"))
	fmt.Println(ReverseString("I am a student and Stranger And Person to Improvement"))
	fmt.Println(ReverseString(lorem))
}
