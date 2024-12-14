package utils

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Quadrant struct {
	Min Point
	Max Point
}

type Robot struct {
	Id       int
	Location Point
	Velocity Point
}

func (p *Point) PrintPoint() {
	fmt.Println("X: ", p.X, ", Y: ", p.Y)
}
func (r *Robot) PrintRobot() {
	fmt.Println("Id:", r.Id, ", Location:", r.Location, ", Velocity:", r.Velocity)
}

func GetInput() []*Robot {
	file, _ := os.Open("input.txt")
	defer file.Close()

	var robots []*Robot
	scanner := bufio.NewScanner(file)
	id := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		parts := strings.Split(line, " ")
		posStr := strings.TrimPrefix(parts[0], "p=")
		velStr := strings.TrimPrefix(parts[1], "v=")

		posNums := strings.Split(posStr, ",")
		velNums := strings.Split(velStr, ",")

		px, _ := strconv.Atoi(posNums[0])
		py, _ := strconv.Atoi(posNums[1])
		vx, _ := strconv.Atoi(velNums[0])
		vy, _ := strconv.Atoi(velNums[1])

		robots = append(robots, &Robot{
			Id:       id,
			Location: Point{X: px, Y: py},
			Velocity: Point{X: vx, Y: vy},
		})
		id++
	}

	return robots
}
