package main

import (
	"image"
	"image/png"
	"os"
	"strconv"
	"time"

	"github.com/kbinani/screenshot"
)

func main() {
	for {

		//每30秒截屏一次 截取的是windows桌面上面显示的页面

		time.Sleep(time.Second * 30)
		screen()
	}
}
func screen() {
	// 使用 GetDisplayBounds获取指定屏幕显示范围，全屏截图
	bounds := screenshot.GetDisplayBounds(0)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		panic(err)
	}
	// 拼接图片名
	t := time.Now().Unix()
	tt := strconv.Itoa(int(t)) + ".png"
	SavePath := "C:\\windows\\temp\\" + tt
	save(img, SavePath)
}

// save *image.RGBA to filePath with PNG format.
func save(img *image.RGBA, filePath string) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	png.Encode(file, img)
}
