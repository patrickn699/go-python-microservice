package convert

import ("fmt"
        "io"
	    "log"
		"github.com/sunshineplan/imgconv")


func Convert_file(image string,height,width,percent int,watermark string) {

	fmt.Println("Hello World, from convert package")

	src, err := imgconv.Open(image)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}

	// Resize the image to width = 200px preserving the aspect ratio.
	mark := imgconv.Resize(src, imgconv.ResizeOption{Width: 200})

	// Add random watermark set opacity = 128.
	dst := imgconv.Watermark(src, imgconv.WatermarkOption{Mark: mark, Opacity: 128, Random: true})

	// Write the resulting image as TIFF.
	err = imgconv.Write(io.Discard, dst, imgconv.FormatOption{Format: imgconv.TIFF})
	if err != nil {
		log.Fatalf("failed to write image: %v", err)
	}
}
