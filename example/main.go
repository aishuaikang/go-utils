package main

import (
	"log"
	"time"

	"github.com/aishuaikang/go-utils"
)

func main() {
	log.SetFlags(log.Ltime | log.Lmsgprefix | log.Llongfile | log.Lshortfile)

	var hwnd int32

	// 启动线程监控窗口
	go func() {
		for {
			newHwnd, err := utils.GetHwndByTitle("v2rayN - V3.29 - 2020/12/05")
			if err != nil {
				log.Println(err)
				break
			}
			if hwnd != newHwnd {
				hwnd = newHwnd

				winRect, err := utils.GetWindowRect(hwnd)
				if err != nil {
					log.Println(err)
					break
				}

				winCenterX := winRect.Width / 2
				winCenterY := winRect.Height / 2

				err = utils.CaptureInit(0, winCenterX-320, winCenterY-320, 640, 640)
				if err != nil {
					log.Println(err)
					break
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()

	defer utils.CaptureRelease()

	// 计算屏幕中心点
	width, height, err := utils.GetScreenResolution()
	if err != nil {
		log.Println(err)
		return
	}
	centerX := int(width / 2)
	centerY := int(height / 2)

	// 初始化算法（可以重复调用更新参数）
	utils.SpeedInit(80, 0.1)
	utils.PidInit(0.1, 0.002, 0.003)
	utils.LinearInit(0.1)
	utils.MagneticInit(0.3)

	// for循环模拟推理
	for {
		// 判断窗口是否存在
		if hwnd <= 0 || utils.IsWindowIconic(hwnd) || !utils.IsWindowVisible(hwnd) {
			log.Println("窗口不存在或最小化或不可见")
			time.Sleep(1 * time.Second)
			continue
		}

		// 截图
		var data []byte

		startTime := time.Now()
		utils.CaptureBitmap(&data)
		captureTime := time.Since(startTime)

		// 获取当前鼠标位置
		currentX, currentY, err := utils.GetCursorPos()
		if err != nil {
			log.Println(err)
			return
		}

		// 调用算法
		outputX, outputY, err := utils.Compute(utils.Pid, currentX, currentY, float64(centerX), float64(centerY))
		if err != nil {
			log.Println(err)
			return
		}

		// 鼠标移动
		utils.MoveMouse(outputX, outputY)

		// 计算信息
		var fps int64

		cm := captureTime.Milliseconds()
		if cm == 0 {
			fps = 99999999999999999
		} else {
			fps = 1000 / cm
		}

		log.Printf("当前位置：(%2d,%2d) - 截图耗时：%5v - 帧数(每秒)：%d\n", int(outputX), int(outputY), captureTime, fps)
		utils.BitmapSaveBMP("test.bmp")
	}

}
