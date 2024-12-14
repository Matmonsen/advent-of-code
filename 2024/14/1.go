package main

import (
	"advent-of-code/2024/14/utils"
	"fmt"
	"strconv"
)

const BoundsX = 101
const BoundsY = 103
const ElapsedSeconds = 100

func NewPos(current int, velocity int, max int) int {
	newVal := current + velocity
	n := newVal
	if newVal < 0 {
		n = max + newVal
	} else if newVal >= max {
		n = newVal - max
	}
	return n
}

func MoveRobot(robot *utils.Robot) {
	robot.Location.X = NewPos(robot.Location.X, robot.Velocity.X, BoundsX)
	robot.Location.Y = NewPos(robot.Location.Y, robot.Velocity.Y, BoundsY)
}

func PrintRobots(robots []*utils.Robot) {
	bathroom := make([][]string, BoundsY)

	for i := 0; i < BoundsY; i++ {
		row := make([]string, BoundsX)
		for i := 0; i < BoundsX; i++ {
			row[i] = "."
		}
		bathroom[i] = row
	}

	for _, robot := range robots {
		if bathroom[robot.Location.Y][robot.Location.X] == "." {
			bathroom[robot.Location.Y][robot.Location.X] = fmt.Sprintf("%d", 1)
		} else {
			n, _ := strconv.Atoi(bathroom[robot.Location.Y][robot.Location.X])
			bathroom[robot.Location.Y][robot.Location.X] = fmt.Sprintf("%d", n+1)
		}

	}

	for _, line := range bathroom {
		fmt.Println(line)
	}
	fmt.Println()
}

func CountRobotsInQuadrant(robots []*utils.Robot, quadrant utils.Quadrant) int {
	robotCount := 0
	for _, robot := range robots {
		isWithingBoundsX := robot.Location.X <= quadrant.Max.X && robot.Location.X >= quadrant.Min.X
		isWithingBoundsY := robot.Location.Y <= quadrant.Max.Y && robot.Location.Y >= quadrant.Min.Y

		if isWithingBoundsX && isWithingBoundsY {
			robotCount++
		}
	}

	return robotCount
}

func FindSafetyFactor(robots []*utils.Robot) int {
	halfBoundsY := BoundsY / 2
	halfBoundsX := BoundsX / 2
	maxBoundsY := BoundsY
	maxBoundsX := BoundsX

	q1 := CountRobotsInQuadrant(robots, utils.Quadrant{Min: utils.Point{X: 0, Y: 0}, Max: utils.Point{X: halfBoundsX - 1, Y: halfBoundsY - 1}})
	q2 := CountRobotsInQuadrant(robots, utils.Quadrant{Min: utils.Point{X: halfBoundsX + 1, Y: 0}, Max: utils.Point{X: maxBoundsX - 1, Y: halfBoundsY - 1}})
	q3 := CountRobotsInQuadrant(robots, utils.Quadrant{Min: utils.Point{X: 0, Y: halfBoundsY + 1}, Max: utils.Point{X: halfBoundsX - 1, Y: maxBoundsY - 1}})
	q4 := CountRobotsInQuadrant(robots, utils.Quadrant{Min: utils.Point{X: halfBoundsX + 1, Y: halfBoundsY + 1}, Max: utils.Point{X: maxBoundsX - 1, Y: maxBoundsY - 1}})

	return q1 * q2 * q3 * q4
}

func main() {
	robots := utils.GetInput()
	for i := 0; i < ElapsedSeconds; i++ {
		for _, robot := range robots {
			MoveRobot(robot)
		}
	}
	fmt.Println(FindSafetyFactor(robots))
}
