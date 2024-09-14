package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/getlantern/systray"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/host"

)

var logFile *os.File

func main() {

	var err error
	logFile, err = os.OpenFile("gouge.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		
		log.SetOutput(os.Stdout)
	} else {
		defer logFile.Close()
		log.SetOutput(logFile)
	}
	defer func() {
		if logFile != nil {
			logFile.Close()
		}
		cleanUp()
	}()
	log.SetOutput(logFile)


	systray.Run(onReady, onExit)

}

func onReady() {
	
	icon, err := getIcon()
	if err != nil {
		log.Println("Error loading icon:", err)

	} else {
		if icon == nil || len(icon) == 0 {
			log.Println("Icon is empty or nil")
		} else {
			log.Printf("Setting icon, size: %d bytes", len(icon))
			systray.SetIcon(icon)
		}
	}

	systray.SetTitle("GOuge")
	systray.SetTooltip("GOuge")

	mCPU := systray.AddMenuItem("CPU: -", "CPU Usage")
	mRAM := systray.AddMenuItem("RAM: -", "RAM Usage")
	mDisk := systray.AddMenuItem("Disk: -", "Disk Usage")
	mUptime := systray.AddMenuItem("Uptime: -", "System Uptime")

	systray.AddSeparator()
	mQuit := systray.AddMenuItem("Terminate", "Exit")

	go func() {
		for {
			updateSystemMetrics(mCPU, mRAM, mDisk, mUptime)
			time.Sleep(2 * time.Second)
		}
	}()

	go func() {
		<-mQuit.ClickedCh
		systray.Quit()
	}()
}

func updateSystemMetrics(mCPU, mRAM, mDisk, mUptime *systray.MenuItem) {
	cpuUsage, err := cpu.Percent(time.Second, false)
	if err != nil {
		log.Println("Error getting CPU usage:", err)
		return
	}
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		log.Println("Error getting memory info:", err)
		return
	}
	diskInfo, err := disk.Usage("/")
	if err != nil {
		log.Println("Error getting disk info:", err)
		return
	}


	uptimeSeconds, err := host.Uptime()
	if err != nil {
		log.Println("Error getting system uptime:", err)
		return
	}
	uptimeDuration := time.Duration(uptimeSeconds) * time.Second
	uptimeString := fmt.Sprintf("%d d, %d h, %d m",
		int(uptimeDuration.Hours())/24,
		int(uptimeDuration.Hours())%24,
		int(uptimeDuration.Minutes())%60)

	mCPU.SetTitle(fmt.Sprintf("CPU: %.1f%%", cpuUsage[0]))
	mRAM.SetTitle(fmt.Sprintf("RAM: %.1f%%", memInfo.UsedPercent))
	mDisk.SetTitle(fmt.Sprintf("Disk: %.1f%%", diskInfo.UsedPercent))
	mUptime.SetTitle(fmt.Sprintf("Uptime: %s", uptimeString))
}

func onExit() {
	fmt.Println("onExit called")
}

func getIcon() ([]byte, error) {
    possiblePaths := []string{
        "gouge.png",
        "Resources/gouge.png",
        "../Resources/gouge.png",
        "gouge.icns",
        "Resources/gouge.icns",
        "../Resources/gouge.icns",
    }

    execPath, err := os.Executable()
    if err == nil {
        execDir := filepath.Dir(execPath)
        possiblePaths = append(possiblePaths,
            filepath.Join(execDir, "gouge.png"),
            filepath.Join(execDir, "Resources", "gouge.png"),
            filepath.Join(execDir, "..", "Resources", "gouge.png"),
            filepath.Join(execDir, "gouge.icns"),
            filepath.Join(execDir, "Resources", "gouge.icns"),
            filepath.Join(execDir, "..", "Resources", "gouge.icns"),
        )
    }

    for _, path := range possiblePaths {
        log.Printf("Trying to load icon from: %s", path)
        icon, err := os.ReadFile(path)
        if err == nil {
            log.Printf("Icon found at: %s, size: %d bytes", path, len(icon))
            return icon, nil
        } else {
            log.Printf("Failed to load icon from %s: %v", path, err)
        }
    }

    return nil, fmt.Errorf("icon file not found in any of the expected locations")
}

func cleanUp() {
	fmt.Println("Application is closing and cleaning up...")
	os.Remove("gouge.log") 
}
