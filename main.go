package main

import (
	"log"

	"github.com/Yinzhongkan399/GoServerPS/baserun"
)

func main() {
	log.Println("Starting BaseRun()")
	if err := baserun.BaseRun(); err != nil {
		log.Fatalf("BaseRun failed: %v", err)
	}
	log.Println("BaseRun completed")

	log.Println("Running ReadBTFandGetItsMember()")
	funcs, err := baserun.ReadBTFandGetItsMember()
	if err != nil {
		log.Fatalf("ReadBTFandGetItsMember failed: %v", err)
	}
	log.Printf("ReadBTFandGetItsMember returned %d entries", len(funcs))

	log.Println("Running TranslateJSON()")
	if err := baserun.TranslateJSON(); err != nil {
		log.Fatalf("TranslateJSON failed: %v", err)
	}
	log.Println("TranslateJSON completed")

	log.Println("All steps finished successfully")
}
