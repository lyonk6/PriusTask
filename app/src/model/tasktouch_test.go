package model

import (
	"fmt"
	"math/rand"
	"testing"
)

/**
 * Test saveTaskTouch. Start by creating a tasktouch object and submitting
 * it to the database. Then Clean up by removing it from the database.
 */
func TestSaveTaskTouch(t *testing.T) {
	fmt.Println("TestSaveTaskTouch")
	tasktouch := &TaskTouch{}
	tasktouch.TouchType = "CREATED"
	tasktouch.Latitude = rand.Float64() * 31
	tasktouch.Longitude = rand.Float64() * 29

}

func TestPostTaskTouch(t *testing.T) {
	fmt.Println("TestPostTaskTouch")
}
