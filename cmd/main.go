package main

import (
	"fmt"
	. "gocourse05/core/funcs"
	_ "gocourse05/core/log"
	logger "gocourse05/core/log"
	"gocourse05/pkg/clinic"
	"gocourse05/pkg/patient"
	"log"
)

func main() {
	p1 := patient.Patient{`100`, `Frank`, 32, `A+`, []string{`Visit 1`, `Visit 2`, `Visit 2`}}
	p2 := patient.VIPPatient{patient.Patient{`101`, `Sam`, 27, `B-`, []string{`Visit 21`, `Visit 22`}}}
	c1 := &clinic.Clinic{}

	// Просте використання інтерфейсу
	if err := c1.RegisterPatient(p1); err != nil {
		log.Fatal(err)
	}
	if err := c1.RegisterPatient(p2); err != nil {
		log.Fatal(err)
	}

	// Поліморфізм з Інтерфейсами
	PerformHealthCheck(c1, p1.GetId())

	// Type Assertions та Type Switches
	PrintDetails(p1, p2, c1)

	// Інтерфейси як Контракти
	if err := ProcessPatient(p1); err != nil {
		log.Fatal(err)
	}
	if err := ProcessPatient(p2); err != nil {
		log.Println(err)
	}

	logger.AddClinic(`some clinic`, *c1)
	fmt.Printf(`%#v`, logger.Clinics())
}
