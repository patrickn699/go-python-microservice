package convert

import (
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"io")

func Convert_file(img string, height,width,percent,watermark int, filename string){

	fmt.Println("Converting file: ", img)
	
	fil, er := os.Open("./converted")
	fmt.Println(fil)
	check_error(er)
	//err := os.Mkdir("./converted", os.ModePerm)
	//check_error(err)

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
		}).SubImage(image.Rect(height, width, percent, watermark))
	
		output_file, outputErr := os.Create("convert/op.jpeg")
		check_error(outputErr)
		jpeg.Encode(output_file, my_sub_image, nil)
	
	} else if format == "png" {
		fmt.Println("File is png")
		opend.Seek(0, 0)
		img, err := png.Decode(opend)
		check_error(err)
		my_sub_image := img.(interface {
			SubImage(r image.Rectangle) image.Image
		}).SubImage(image.Rect(height, width, percent, watermark))
	
		output_file, outputErr := os.Create("output.png")
		check_error(outputErr)
		png.Encode(output_file, my_sub_image)
		by,e := io.Copy(fil, output_file)
		check_error(e)
		fmt.Println("Bytes written: ", by)
	}

}


func check_error(err error){
	if err != nil {
		fmt.Println(err)
	}
}