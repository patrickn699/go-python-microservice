package convert

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os")

func Convert_file(img string, height,width,percent,watermark int, filename string){

	fmt.Println("Converting file: ", img)

	opend, err := os.Open(img)
	check_error(err)

	_, format, err := image.DecodeConfig(opend)
	check_error(err)

	if format == "jpeg" || format == "jpg" {
		fmt.Println("File is jpeg")
		opend.Seek(0, 0)
		img, err := jpeg.Decode(opend)
		check_error(err)
		my_sub_image := img.(interface {
			SubImage(r image.Rectangle) image.Image
		}).SubImage(image.Rect(0, 0, 10, 10))
	
		output_file, outputErr := os.Create("output.jpeg")
		if outputErr != nil {
			fmt.Println(outputErr)
		}
		jpeg.Encode(output_file, my_sub_image, nil)
	
	} else if format == "png" {
		fmt.Println("File is png")
		opend.Seek(0, 0)
		img, err := png.Decode(opend)
		check_error(err)
		my_sub_image := img.(interface {
			SubImage(r image.Rectangle) image.Image
		}).SubImage(image.Rect(0, 0, 10, 10))
	
		output_file, outputErr := os.Create("output.png")
		if outputErr != nil {
			fmt.Println(outputErr)
		}
		png.Encode(output_file, my_sub_image)
	}

}


func check_error(err error){
	if err != nil {
		fmt.Println(err)
	}
}