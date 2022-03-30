package image

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"io"
	"net/http"
	"os"
)

// 下载文件到本地
func downloadFile(URL, fileName string) error {
	//Get the response bytes from the url
	response, err := http.Get(URL)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}
	//Create a empty file
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	//Write the bytes to the fiel
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return err
	}

	return nil
}

func AddMask(origin, mask, target string) error {
	//图片，网上随便找了一张
	img_file, err := os.Open(origin)
	if err != nil {
		fmt.Println("打开图片出错")
		return err
	}
	defer img_file.Close()
	img, err := jpeg.Decode(img_file)
	if err != nil {
		fmt.Println("把图片解码为结构体时出错")
		return err
	}

	//水印,用的是我自己支付宝的二维码
	wmb_file, err := os.Open(mask)
	if err != nil {
		fmt.Println("打开水印图片出错")
		return err
	}
	defer wmb_file.Close()
	wmb_img, err := png.Decode(wmb_file)
	if err != nil {
		fmt.Println("把水印图片解码为结构体时出错")
		return err
	}

	//把水印写在右下角，并向0坐标偏移10个像素
	offset := image.Pt(img.Bounds().Dx()-wmb_img.Bounds().Dx()-10, img.Bounds().Dy()-wmb_img.Bounds().Dy()-10)
	b := img.Bounds()
	//根据b画布的大小新建一个新图像
	m := image.NewRGBA(b)

	//image.ZP代表Point结构体，目标的源点，即(0,0)
	//draw.Src源图像透过遮罩后，替换掉目标图像
	//draw.Over源图像透过遮罩后，覆盖在目标图像上（类似图层）
	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, wmb_img.Bounds().Add(offset), wmb_img, image.ZP, draw.Over)

	//生成新图片new.jpg,并设置图片质量
	imgw, err := os.Create(target)
	if err != nil {
		return err
	}
	jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
	defer imgw.Close()

	fmt.Println("添加水印图片结束请查看")
	return nil
}

func AddWatermarkByNet(url, mask, target string) error {
	//Get the response bytes from the url
	response, err := http.Get(url)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return errors.New("Received non 200 response code")
	}

	fileBuffer := new(bytes.Buffer)
	_, err = io.Copy(fileBuffer, response.Body)
	if err != nil {
		fmt.Println("图片下载时出错")
		return err
	}

	img, err := jpeg.Decode(fileBuffer)
	if err != nil {
		fmt.Println("把图片解码为结构体时出错")
		return err
	}

	//水印,用的是我自己支付宝的二维码
	wmb_file, err := os.Open(mask)
	if err != nil {
		fmt.Println("打开水印图片出错")
		return err
	}
	defer wmb_file.Close()
	wmb_img, err := png.Decode(wmb_file)
	if err != nil {
		fmt.Println("把水印图片解码为结构体时出错")
		return err
	}

	//把水印写在右下角，并向0坐标偏移10个像素
	offset := image.Pt(img.Bounds().Dx()-wmb_img.Bounds().Dx()-10, img.Bounds().Dy()-wmb_img.Bounds().Dy()-10)
	b := img.Bounds()
	//根据b画布的大小新建一个新图像
	m := image.NewRGBA(b)

	//image.ZP代表Point结构体，目标的源点，即(0,0)
	//draw.Src源图像透过遮罩后，替换掉目标图像
	//draw.Over源图像透过遮罩后，覆盖在目标图像上（类似图层）
	draw.Draw(m, b, img, image.ZP, draw.Src)
	draw.Draw(m, wmb_img.Bounds().Add(offset), wmb_img, image.ZP, draw.Over)

	//生成新图片new.jpg,并设置图片质量
	imgw, err := os.Create(target)
	if err != nil {
		return err
	}
	jpeg.Encode(imgw, m, &jpeg.Options{jpeg.DefaultQuality})
	defer imgw.Close()

	fmt.Println("添加水印图片结束请查看")
	return nil
}
