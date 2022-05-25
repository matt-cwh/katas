package main

import "testing"

func TestNextState_UnderPopulation(t *testing.T) {
	nextState := getNextState(1)
	if nextState {
		t.Errorf("expect next state to be false when live neighbours less than 2")
	}
}

func TestNextState_LivesOn(t *testing.T) {
	nextState := getNextState(2)
	if !nextState {
		t.Errorf("expect next state to be true when live neighbours is 2")
	}

	nextState = getNextState(3)
	if !nextState {
		t.Errorf("expect next state to be true when live neighbours is 3")
	}
}

func TestNextState_OverPopulation(t *testing.T) {
	nextState := getNextState(4)
	if nextState {
		t.Errorf("expect next state to be false when live cell has live neighbours of 4")
	}
}

func TestNextState_Reproduction(t *testing.T) {
	nextState := getNextState(3)
	if !nextState {
		t.Errorf("expect next state to be true when dead cell has live neighbours equal to 3")
	}

	nextState = getNextState(4)
	if nextState {
		t.Errorf("expect next state to be false when dead cell has live neighbours more than 3")
	}
}
