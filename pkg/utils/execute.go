package utils

import (
	"fmt"
	"time"

	vegeta "github.com/tsenart/vegeta/lib"
)

//refer https://github.com/zendesk/zopim-vegeta
//func ExecuteLoad(url string, method string, headers string, noOfReqPerSec int, totalDurationInSec int) {
func ExecuteLoad() {
	fmt.Println("Start load")
	rate := vegeta.Rate{Freq: 10, Per: time.Second}
	duration := 4 * time.Second

	targeter := vegeta.NewStaticTargeter(vegeta.Target{Method: "GET", URL: "https://reqres.in/api/users?page=2"})

	attacker := vegeta.NewAttacker()

	var metrics vegeta.Metrics

	for res := range attacker.Attack(targeter, rate, duration, "myapi") {
		metrics.Add(res)
		fmt.Println(res.Code)
	}
	metrics.Close()

	fmt.Printf("metrics at latencies at 99th percentile:%s \n", metrics.Latencies.P99)
}
