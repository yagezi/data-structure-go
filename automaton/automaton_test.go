/*
*@Time        : 2022/02/14
*@File        : automaton_test.go
*@Description :
 */

package acautomaton

import "testing"

func Test_AcamCreate(t *testing.T) {
	ac := NewAutomaton()
	ac.Add("abe")
	ac.Add("abc")
	ac.Add("abd")
	ac.Add("be")
	ac.Add("bc")
	ac.Add("cf")
	ac.GenFailurePointer()
	dfsAC(ac.root)
}

func Test_AcamMatch(t *testing.T) {
	ac := NewAutomaton()
	ac.Add("dine")
	ac.Add("nine")
	ac.Add("fine")
	ac.GenFailurePointer()
	ac.Match("ten little soldiers went out to dine, one choked himself and then there were nine")
}
