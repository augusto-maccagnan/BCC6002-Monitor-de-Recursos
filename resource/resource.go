package resource

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
)

func GetResources()  (map[string][]string, error) {
	var cpuCount int
	var fields []string
	var resources = make(map[string][]string)

	cpuFrequency := getCpuUsage(&cpuCount)
	cpuMaxFreq := getCpuMaxFreq()
	cpuMinFreq := getCpuMinFreq()
	memInfo := getMemInfo()

	for i := range cpuFrequency {
		fields = append(fields, "CPU " + strconv.Itoa(i) + " frequency: " + cpuFrequency[i])
	}
	resources["CPU Frequency"] = fields
	fields = nil
	fields = append(fields, "CPU Max Frequency: " + cpuMaxFreq)
	resources["CPU Max Frequency"] = fields
	fields = nil
	fields = append(fields, "CPU Min Frequency: " + cpuMinFreq)
	resources["CPU Min Frequency"] = fields
	fields = nil
	fields = append(fields, "Number of Physical Cores: " + strconv.Itoa(cpuCount))
	resources["Number of Physical Cores"] = fields
	
	fields = nil
	fields = append(fields, "Memory " + memInfo[0])
	resources["Total Memory"] = fields
	fields = nil
	fields = append(fields, "Memory " + memInfo[1])
	resources["Free Memory"] = fields
	fields = nil
	fields = append(fields, "Memory " + memInfo[2])
	resources["Available Memory"] = fields

	return resources, nil
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
	return cpuFrequency
}

func getMemInfo() []string {
	var memInfo []string

	cmd := exec.Command("cat", "/proc/meminfo")
	temp, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	var found bool = true
	var after string = string(temp)

	for found {
		_, after, found = strings.Cut(after, "Mem")
		if(!found) {
			break
		}
		memInfo = append(memInfo, after[:strings.Index(after, "\n")])
	}

	return memInfo
}