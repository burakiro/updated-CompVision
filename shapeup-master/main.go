// shapes - openCV test for shape detection
package main

import (
	"fmt"
	"gocv.io/x/gocv"
	"image"
	"image/color"
)

func main() {
	window := gocv.NewWindow("Shapes")
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowAutosize)
	img := gocv.IMRead("Red.jpg", gocv.IMReadColor)
	defer img.Close()

	// load the image and resize it to a smaller factor so that
	// the shapes can be approximated better
	resized := gocv.NewMat()
	defer resized.Close()
	imgW := float64(img.Rows())
	imgH := float64(img.Cols())
	resizedW := float64(600)
	resizedH := float64(resizedW / imgW * imgH)
	gocv.Resize(img, &resized, image.Pt(int(resizedW), int(resizedH)), 0, 0, gocv.InterpolationArea)
	ratioX := imgW / resizedW
	ratioY := imgH / resizedH
	fmt.Printf("ratioX: %v, ratioY: %v, img.Rows: %v, img.Cols: %v, resized.Rows: %v, resized.Cols: %v\n", ratioX, ratioY, img.Rows(), img.Cols(), resized.Rows(), resized.Cols())
	window.IMShow(resized)
	gocv.WaitKey(0)

	// convert the resized image to grayscale
	gray := gocv.NewMat()
	defer gray.Close()
	gocv.CvtColor(resized, &gray, gocv.ColorBGRToGray)
	window.IMShow(gray)
	gocv.WaitKey(0)

	// blur it slightly
	blurred := gocv.NewMat()
	defer blurred.Close()
	gocv.GaussianBlur(gray, &blurred, image.Point{1, 1}, 0, 0, gocv.BorderConstant)
	window.IMShow(blurred)
	gocv.WaitKey(0)

	// canny
	canny := gocv.NewMat()
	defer canny.Close()
	gocv.Canny(blurred, &canny, 100, 200)
	// gocv.Canny(blurred, canny, 0, 50)
	window.IMShow(canny)
	gocv.WaitKey(0)

	// dilate canny output to remove potential holes between edge segments
	// dilate := gocv.NewMat()
	// defer dilate.Close()
	// kernel := gocv.NewMat()
	// defer kernel.Close()
	// gocv.Dilate(canny, dilate, kernel)
	// window.IMShow(dilate)
	// gocv.WaitKey(0)

	// and threshold it
	thresh := gocv.NewMat()
	defer thresh.Close()
	gocv.Threshold(canny, &thresh, 60, 255, gocv.ThresholdBinary)
	window.IMShow(thresh)
	gocv.WaitKey(0)
	


	// invert the image
	// inverted := gocv.NewMat()
	// defer inverted.Close()
	// gocv.Threshold(thresh, inverted, 60, 255, gocv.ThresholdBinaryInv)
	// window.IMShow(inverted)
	// gocv.WaitKey(0)

	// find contours in the thresholded image
	cnts := gocv.FindContours(thresh, gocv.RetrievalList, gocv.ChainApproxSimple)

	// loop over the contours
	//for _, c := range cnts {
		// compute the center of the contour, then detect the name of the
		// shape using only the contour
		/*
		   tmp := gocv.NewMat()
		   defer tmp.Close()
		   tmp.p = c
		   M := gocv.Moments(tmp, true)
		   cX := int((M["m10"] / M["m00"]) * ratioX)
		   cY := int((M["m01"] / M["m00"]) * ratioY)
		*/
		//shape, err := detect(c, resizedW*resizedH)
		//if err != nil {
		//	continue // ignore if shape wasn't detected from contour
		//}

		// multiply the contour (x, y)-coordinates by the resize ratio,
		// then draw the contours and the name of the shape on the image
		// for i := range c {
		// 	c[i].X = int(float64(c[i].X) * ratioX)
		// 	c[i].Y = int(float64(c[i].Y) * ratioY)
		// }
		gocv.DrawContours(&resized, cnts, -1, color.RGBA{0, 255, 0, 0}, 2)
		// gocv.putText(img, shape, image.Point{cX, cY}, gocv.FontHersheySimplex, 0.5, color.RGBA{255, 255, 255, 0}, 2)
		//gocv.PutText(&resized, shape, c[0], gocv.FontHersheySimplex, 0.5, color.RGBA{255, 255, 255, 0}, 2)
	//}

	// show image and wait for key to be pressed
	window.IMShow(resized)
	gocv.WaitKey(0)
}

func detect(pts []image.Point, imgArea float64) (string, error) {
	//  initialize the shape name and approximate the contour
	shape := "unidentified"
	peri := gocv.ArcLength(pts, true)
	approx := gocv.ApproxPolyDP(pts, 0.02*peri, true)


	switch len(approx) {
	// if the shape is a triangle, it will have 3 vertices
	case 3.0:
		shape = "triangle"

	// if the shape has 4 vertices, it is either a square or
	// a rectangle
	case 4.0:
		/*
		   // compute the bounding box of the contour and use the
		   // bounding box to compute the aspect ratio
		   (x, y, w, h) = gocv.boundingRect(approx)
		   ar = w / float(h)
		   // a square will have an aspect ratio that is approximately
		   // equal to one, otherwise, the shape is a rectangle
		   shape = "square" if ar >= 0.95 and ar <= 1.05 else "rectangle"
		*/
		shape = "box"

	// if the shape is a pentagon, it will have 5 vertices
	case 5.0:
		shape = "pentagon"

	// otherwise, we assume the shape is a circle
	default:
		shape = "circle"
	}

	// return the name of the shape
	return shape, nil
}