//go:build linux
// +build linux

package main

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)


func main(){
	var cmd *exec.Cmd
	// cmd := exec.Command("cat", "/proc/cpuinfo")
	// cpuFrequency, err := cmd.Output()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	cmd = exec.Command("cat", "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_max_freq")
	cpuMaxFreq, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	cmd = exec.Command("cat", "/sys/devices/system/cpu/cpu0/cpufreq/cpuinfo_min_freq")
	cpuMinFreq, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(cpuFrequency))
	var cpuFrequency [8]string
	cpuFrequency = getCpuInfo()
	for i := range cpuFrequency {
		fmt.Println(cpuFrequency[i])
	}
	// fmt.Println(getCpuInfo())
	fmt.Println(string(cpuMaxFreq))
	fmt.Println(string(cpuMinFreq))

}

func getCpuInfo() [8]string {
	cmd := exec.Command("cat", "/proc/cpuinfo")
	osOutput, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	text := string(osOutput)

	// before, cpu0Frequency, found := strings.Cut(cpuFrequency, "\n")
	var cpuFrequency [8]string
	var found bool = true
	var i, index int
	for (found) { // Search for next line
		index = strings.Index(text, "cpu MHz")
		if(index == -1){
			break
		}
		text = text[index:]
		cpuFrequency[i], text, found = strings.Cut(text, "\n")
		i++
	}
	return cpuFrequency
}