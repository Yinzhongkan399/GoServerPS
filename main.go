package main

import (
	"log"
)

func main() {
	log.Println("Starting BaseRun()")
	if err := BaseRun(); err != nil {
		log.Fatalf("BaseRun failed: %v", err)
	}
	log.Println("BaseRun completed")

	log.Println("Running ReadBTFandGetItsMember()")
	funcs, err := ReadBTFandGetItsMember()
	if err != nil {
		log.Fatalf("ReadBTFandGetItsMember failed: %v", err)
	}
	log.Printf("ReadBTFandGetItsMember returned %d entries", len(funcs))

	log.Println("Running TranslateJSON()")
	if err := TranslateJSON(); err != nil {
		log.Fatalf("TranslateJSON failed: %v", err)
	}
	log.Println("TranslateJSON completed")

	log.Println("All steps finished successfully")
}
