package main

import (
	"gocv.io/x/gocv"
	"image"
)

func main(){


	frame := gocv.NewMat()
	defer frame.Close()

	hsv := gocv.NewMat()
	defer hsv.Close()

	mask1 := gocv.NewMat()
	defer mask1.Close()

	mask2 := gocv.NewMat()
	defer mask2.Close()

	kernel := gocv.NewMatWithSizeFromScalar(gocv.Scalar{Val1: 1}, 3, 3, gocv.MatTypeCV32F)
	defer kernel.Close()

	res := gocv.NewMat()
	defer res.Close()
	imgname := "Red.jpg"

	frame = gocv.IMRead(imgname, gocv.IMReadColor)


	window := gocv.NewWindow("Red Detection")
	defer window.Close()


	gocv.CvtColor(frame, &hsv, gocv.ColorBGRToHSV)

	gocv.InRangeWithScalar(hsv, gocv.Scalar{Val1: 0 , Val2: 100, Val3: 100}, gocv.Scalar{Val1: 10, Val2: 255, Val3: 255}, &mask1)
	gocv.InRangeWithScalar(hsv, gocv.Scalar{Val1: 160, Val2: 100, Val3: 100}, gocv.Scalar{Val1: 180, Val2: 255, Val3: 255}, &mask2)

	gocv.Add(mask1, mask2, &mask1)

	gocv.MorphologyEx(mask1, &mask1, gocv.MorphOpen, kernel)
	gocv.MorphologyEx(mask1, &mask1, gocv.MorphDilate, kernel)
	gocv.MorphologyEx(mask1, &mask1, gocv.MorphClose, kernel)

	gocv.BitwiseNot(mask1, &mask2)

	gocv.BitwiseAndWithMask(frame, frame, &res, mask1)


	img:=res
	pts1 := make([]gocv.Point2f, 3)
	pts1[0] = gocv.Point2f{ X: 83, Y: 90}
	pts1[1] = gocv.Point2f{ X: 447, Y: 90}
	pts1[2] = gocv.Point2f{ X: 83, Y: 472}

	pts2 := make([]gocv.Point2f, 3)
	pts2[0] = gocv.Point2f{ X: 83, Y: 90}
	pts2[1] = gocv.Point2f{ X: 447, Y: 90}
	pts2[2] = gocv.Point2f{ X: 150, Y: 472}
	test := gocv.GetAffineTransform2f(pts1,pts2)
	warp_dst := gocv.NewMatWithSize(img.Size()[0], img.Size()[1], img.Type())
	gocv.WarpAffine(img, &warp_dst, test, image.Point{
		X: img.Size()[0],
		Y: img.Size()[1],
	})
	window.IMShow(img)
	gray := gocv.NewMat()
	blurred := gocv.NewMat()
	canny:=gocv.NewMat()
	gocv.CvtColor(img,&gray, gocv.ColorRGBAToGray)
	gocv.GaussianBlur(gray,&blurred,image.Pt(75, 75),0,0,gocv.BorderDefault)
	gocv.Canny(blurred,&canny,120,255)

	//https://stackoverflow.com/questions/57125879/improve-rectangle-contour-detection-in-image-using-opencv
	// bunu dönüştürüyordum kanka devam edeceksen buradan etsene


	window.WaitKey(0)




	//green := color.RGBA{0, 255, 0, 0}



		}


	/*window := gocv.NewWindow("First")
	window2 :=gocv.NewWindow("Final")
	imgname :="grid.jpg"
	img := gocv.IMRead(imgname, gocv.IMReadColor)
	gocv.Circle(&img , image.Pt(83,90) , 5, color.RGBA{255, 0, 0,1} , -1)
	gocv.Circle(&img , image.Pt(447,90) , 5, color.RGBA{255, 0, 0,1} , -1)
	gocv.Circle(&img , image.Pt(83,472) , 5, color.RGBA{255, 0, 0,1} , -1)

	pts1 := make([]gocv.Point2f, 3)
	pts1[0] = gocv.Point2f{ X: 83, Y: 90}
	pts1[1] = gocv.Point2f{ X: 447, Y: 90}
	pts1[2] = gocv.Point2f{ X: 83, Y: 472}

	pts2 := make([]gocv.Point2f, 3)
	pts2[0] = gocv.Point2f{ X: 0, Y: 90}
	pts2[1] = gocv.Point2f{ X: 447, Y: 90}
	pts2[2] = gocv.Point2f{ X: 83, Y: 472}
	test := gocv.GetAffineTransform2f(pts1,pts2)
	warp_dst := gocv.NewMatWithSize(img.Size()[0], img.Size()[1], img.Type())
	gocv.WarpAffine(img, &warp_dst, test, image.Point{
		X: img.Size()[0],
		Y: img.Size()[1],
	})

	window.IMShow(img)
	window2.IMShow(warp_dst)
	window.WaitKey(0)*/
