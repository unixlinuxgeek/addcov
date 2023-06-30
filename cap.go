// Capture 1 frame from video
// Захватываем из видео один кадр
// example: 
//   ./cap /var/tmp/in.mp4 /var/tmp/out5.jpg /var/tmp/out.mp4

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	capture()
	merge()
}

func merge() {
	args := os.Args

	if len(args) == 4 {
		app := "/usr/bin/ffmpeg"
		arg0 := "-i"
		arg1 := args[1]
		arg2 := "-i"
		arg3 := args[2]
		arg4 := "-c:v"
		arg5 := "copy"
		arg6 := "-c:a"
		arg7 := "aac"
		arg8 := args[3]

		cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8)
		err := cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func capture() {
	args := os.Args
	if len(args) == 4 {
		vidName := args[1]
		imgName := args[2]
		f, err := os.OpenFile(vidName, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 777)
		err = f.Close()
		if err != nil {
			log.Fatal(err)
		}

		app := "/usr/bin/ffmpeg"
		arg0 := "-y"
		arg1 := "-ss"
		arg2 := "00:01:00"
		arg3 := "-i"
		arg4 := vidName
		arg5 := "-vf"
		arg6 := "scale=200:200"

		arg7 := "-frames:v"
		arg8 := "1"
		arg9 := "-q:v"
		arg10 := "2"
		arg11 := imgName

		cmd := exec.Command(app, arg0, arg1, arg2, arg3, arg4, arg5, arg6, arg7, arg8, arg9, arg10, arg11)
		err = cmd.Run()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		fmt.Println("Введите название видеофайла и выходного изображения!!!")
	}
}
