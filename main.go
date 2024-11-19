//go:build linux
// +build linux

package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"
)


func main(){


	var cpuFrequency []string
	var cpuCount int
	for true {
		ClearScreen()
		cpuFrequency = getCpuUsage(&cpuCount)
		for i := range cpuFrequency {
			fmt.Println(cpuFrequency[i])
		}

		fmt.Printf("CPU Max Freq: %s", getCpuMaxFreq())
		fmt.Printf("CPU Min Freq: %s", getCpuMinFreq())
		fmt.Printf("Number of Physical Cores: %d\n", cpuCount)
		time.Sleep(1 * time.Second)
	}
}

func getCpuMaxFreq() string {
	cmd := exec.Command("cat", "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_max_freq")
	cpuMaxFreq, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(cpuMaxFreq)
}

func getCpuMinFreq() string {
	cmd := exec.Command("cat", "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_min_freq")
	cpuMinFreq, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	return string(cpuMinFreq)
}

func getCpuUsage(Number *int) []string {
	cmd := exec.Command("cat", "/proc/cpuinfo")
	osOutput, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	text := string(osOutput)

	// before, cpu0Frequency, found := strings.Cut(cpuFrequency, "\n")
	var cpuFrequency []string
	var found bool = true
	var i, index int
	var temp string
	for (found) { // Search for next line
		index = strings.Index(text, "cpu MHz")
		if(index == -1){
			break
		}
		text = text[index:]
		temp, text, found = strings.Cut(text, "\n")
		cpuFrequency = append(cpuFrequency, temp)
		i++
	}
	*Number = i
	// fmt.Println(i)
	return cpuFrequency
}


func ClearScreen() {
    if runtime.GOOS == "windows" {
        cmd := exec.Command("cmd", "/c", "cls")
        cmd.Stdout = os.Stdout
        cmd.Run()
    } else {
        cmd := exec.Command("clear")
        cmd.Stdout = os.Stdout
        cmd.Run()
    }
}