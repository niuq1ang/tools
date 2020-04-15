package timestamp

import (
	"fmt"
	"testing"
	"time"
)

// func Test_createWokerID(t *testing.T) {
// 	pid := os.Getpid()
// 	workerIDByPid := createWorkerIDByPid()

// 	mac := getMacAddr()
// 	workerIDByMac := createWorkerIDByMac()

// 	fmt.Printf("pid: %x, mac: %s, len = %d \n", pid, mac, len(mac))
// 	fmt.Printf("workerIDByPid: %x, workerIDByMac: %x \n", workerIDByPid, workerIDByMac)
// }

func Test_Uuid(t *testing.T) {
	SetWorkerID(3)
	for i := 0; i < 10; i++ {
		// fmt.Printf("MicroSec = %b, UUID = %b\n", GetMicroSec(), Next())
		// fmt.Printf("MicroSec = %x, UUID = %x\n", GetMicroSec(), Next())
		// fmt.Printf("MicroSec = %x\n", getMaskedMicroSec())
	}
}

func Test_UuidPerfomance(t *testing.T) {

	SetWorkerID(3)
	num := 0
	go func() {
		for {
			// fmt.Printf("MicroSec = %x\n", Next())
			Next()
			num++
		}
	}()
	<-time.After(time.Second)
	fmt.Printf("Get %d stamp in second. \n", num)
}
