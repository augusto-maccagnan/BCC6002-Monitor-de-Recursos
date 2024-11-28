package main

import (
	"encoding/json"
	"fmt"
	"main/model"
	"main/resource"
	"net/http"
	"strconv"
	"strings"
)

func main() {

 // Create a new request multiplexer
 // Take incoming requests and dispatch them to the matching handlers
 mux := http.NewServeMux()

//  r := 

 // Register the routes and handlers
 mux.Handle("/", &homeHandler{})
 mux.Handle("/resources", &resourcesHandler{})

 // Run the server
 fmt.Println("Starting server on :8080")
 http.ListenAndServe(":8080", mux)
}

type homeHandler struct{}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
//  w.Write([]byte(getResources()))
    w.Write([]byte("This is my home page"))
}

type resourcesHandler struct{}

func (h *resourcesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    enableCors(&w)
    switch {
    case r.Method == http.MethodGet:
        h.listResources(w, r)
        return
    }
}

func (h *resourcesHandler) listResources(w http.ResponseWriter, r *http.Request){
	dataText, err := resource.GetResources()
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	jsonBytes, err := convertData(dataText)
	if err != nil {
		InternalServerErrorHandler(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func InternalServerErrorHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusInternalServerError)
    w.Write([]byte("500 Internal Server Error"))
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusNotFound)
    w.Write([]byte("404 Not Found"))
}

// Enable any origin to access the API
func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	}

func convertData(dataMap map[string][]string) ([]byte, error) {
	var data []model.Resources
	var CpuParams []model.CpuParams
	var CpuCores []model.CpuCores
	var Memory []model.Memory
	var Disk []model.Disk
	var TotalFreq int

	tmp := dataMap["CPU Max Frequency"]
	_, after, _ := strings.Cut(tmp[0], ":")
	idx := strings.Index(after, ":")
	after = after[idx+2:strings.Index(after, "\n")]
	MaxFreq, _ := strconv.Atoi(after[:len(after)-3])
	// after = after[:len(after)-3] + "." + after[len(after)-3:]
	// CpuParams = append(CpuParams, model.CpuParams{Name: before, Frequency: MaxFreq})

	tmp = dataMap["CPU Min Frequency"]
	_, after, _ = strings.Cut(tmp[0], ":")
	idx = strings.Index(after, ":")
	after = after[idx+2:strings.Index(after, "\n")]
	MinFreq, _ := strconv.Atoi(after[:len(after)-3])
	// after = after[:len(after)-3] + "." + after[len(after)-3:]
	
	tmp = dataMap["Number of Physical Cores"]
	_ , after, _ = strings.Cut(tmp[0], ":")
	CoreCount, err := strconv.Atoi(after[1:])
	if(err != nil){
		return nil, err
	}

	tmp = dataMap["CPU Frequency"]
	for i := range tmp{
		before, after, _ := strings.Cut(tmp[i], ":")
		idx := strings.Index(after, ":")
		after = after[idx+2:]
		after = after[:strings.Index(after, ".")-1] + after[strings.Index(after, ".")+1:]
		after = after[:len(after)-2]
		TempFreq, _ := strconv.Atoi(after)
		CpuCores = append(CpuCores, model.CpuCores{Name: before, Frequency: TempFreq, Percentage: calculatePercentage(TempFreq, MaxFreq, MinFreq)})
		TotalFreq += TempFreq
	}

	CpuParams = append(CpuParams, model.CpuParams{MaxFrequency: MaxFreq, MinFrequency: MinFreq, CoreNumber: CoreCount, TotalUse: float64(TotalFreq/CoreCount), Percentage: calculatePercentage((TotalFreq/CoreCount), MaxFreq, MinFreq)})

	tmp = dataMap["Total Memory"]
	_, after, _ = strings.Cut(tmp[0], ":")
	after = after[:strings.Index(after, "k")-1]
	TotalMem, err := strconv.ParseFloat(after[strings.LastIndex(after, " ")+1:], 64)
	if(err != nil){
		return nil, err
	}

	tmp = dataMap["Free Memory"]
	_, after, _ = strings.Cut(tmp[0], ":")
	after = after[:strings.Index(after, "k")-1]
	FreeMem, err := strconv.ParseFloat(after[strings.LastIndex(after, " ")+1:], 64)
	if(err != nil){
		return nil, err
	}

	tmp = dataMap["Available Memory"]
	_ , after, _ = strings.Cut(tmp[0], ":")
	after = after[:strings.Index(after, "k")-1]
	AvailableMem, err := strconv.ParseFloat(after[strings.LastIndex(after, " ")+1:], 64)
	if(err != nil){
		return nil, err
	}
	UsedMem := TotalMem - AvailableMem
	percentage, err := strconv.ParseFloat(fmt.Sprintf("%.2f", ((UsedMem/TotalMem)*100)), 64)
	if(err != nil){
		return nil, err
	}
	Memory = append(Memory, model.Memory{Total: TotalMem, Available: AvailableMem, Free: FreeMem, Used: UsedMem, Percentage: percentage})

	name := dataMap["Disk Name"]
	total := dataMap["Disk Total"]
	usage := dataMap["Disk Usage"]
	free := dataMap["Disk Free"]
	for i := range total {
		_, after, _ := strings.Cut(name[i], ":")
		Name := after[strings.LastIndex(after, " ")+1:]

		_, after, _ = strings.Cut(total[i], ":")
		TotalDisk, err := strconv.Atoi(after[strings.LastIndex(after, " ")+1:])
		if err != nil {
			return nil, err
		}
		_, after, _ = strings.Cut(usage[i], ":")
		UsedDisk, err := strconv.Atoi(after[strings.LastIndex(after, " ")+1:])
		if err != nil {
			return nil, err
		}
		_, after, _ = strings.Cut(free[i], ":")
		FreeDisk, err := strconv.Atoi(after[strings.LastIndex(after, " ")+1:])
		if err != nil {
			return nil, err
		}
		Disk = append(Disk, model.Disk{Name : Name, Total: TotalDisk, Used: UsedDisk, Free: FreeDisk, Percentage: calculatePercentage(UsedDisk, TotalDisk, 0)})
	}
	data = append(data, model.Resources{Cpu: CpuParams, CpuCores: CpuCores, Memory: Memory, Disk: Disk})

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		err = fmt.Errorf("error marshalling data: %v", err)
		return nil, err
	}
	return jsonBytes, nil
}

func calculatePercentage(frequency int, maxFreq int, minFreq int) float64 {
	percentage, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", float64(frequency-minFreq)/float64(maxFreq-minFreq)*100), 64)
	return percentage
}