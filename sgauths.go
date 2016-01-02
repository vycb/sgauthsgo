package main

import (
	//"bytes"
	//"fmt"
	//"log"
	"os/exec"
	"time"
)

func main() {
	run := func() *exec.Cmd {
		sgauth := exec.Command("/home/Progs/sgauth", "/home/Progs/sgauth.conf")
		err := sgauth.Start()
		if err != nil {
			panic(err)
		}
		return sgauth
	}

	sgauth := run()

	//var out bytes.Buffer

	for  {
		ping := exec.Command("ping", "-c 1", "google.com")
		//ping.Stdout = &out
		err := ping.Run()
		if err != nil {
			//log.Printf("Err:%q",err)
			sgauth.Process.Kill()
			sgauth.Process.Wait()
			sgauth = run()
		}
		//fmt.Printf("std.Out: %s", out.String())
		//fmt.Println("---------------")

		//out.Reset()
		time.Sleep(25 *time.Second)
	}

	defer sgauth.Process.Kill()
}