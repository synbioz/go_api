package models

import (
	"github.com/synbioz/go_api/config"
	"log"
	"time"
)

type Car struct {
	Id           int       `json:"id"`
	Manufacturer string    `json:"manufacturer"`
	Design       string    `json:"design"`
	Style        string    `json:"style"`
	Doors        uint8     `json:"doors"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type Cars []Car

func NewCar(c *Car) {
	if c == nil {
		log.Fatal(c)
	}
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	err := config.Db().QueryRow("INSERT INTO cars (manufacturer, design, style, doors, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6) RETURNING id;", c.Manufacturer, c.Design, c.Style, c.Doors, c.CreatedAt, c.UpdatedAt).Scan(&c.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func FindCarById(id int) *Car {
	var car Car

	row := config.Db().QueryRow("SELECT * FROM cars WHERE id = $1;", id)
	err := row.Scan(&car.Id, &car.Manufacturer, &car.Design, &car.Style, &car.Doors, &car.CreatedAt, &car.UpdatedAt)

	if err != nil {
		log.Fatal(err)
	}

	return &car
}

func AllCars() *Cars {
	var cars Cars

	rows, err := config.Db().Query("SELECT * FROM cars")

	if err != nil {
		log.Fatal(err)
	}

	// Close rows after all readed
	defer rows.Close()

	for rows.Next() {
		var c Car

		err := rows.Scan(&c.Id, &c.Manufacturer, &c.Design, &c.Style, &c.Doors, &c.CreatedAt, &c.UpdatedAt)

		if err != nil {
			log.Fatal(err)
		}

		cars = append(cars, c)
	}

	return &cars
}

func UpdateCar(car *Car) {
	car.UpdatedAt = time.Now()

	stmt, err := config.Db().Prepare("UPDATE cars SET manufacturer=$1, design=$2, style=$3, doors=$4, updated_at=$5 WHERE id=$6;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(car.Manufacturer, car.Design, car.Style, car.Doors, car.UpdatedAt, car.Id)

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteCarById(id int) error {
	stmt, err := config.Db().Prepare("DELETE FROM cars WHERE id=$1;")

	if err != nil {
		log.Fatal(err)
	}

	_, err = stmt.Exec(id)

	return err
}
