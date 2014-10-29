package tic_tac_toe_test

import (
	"bytes"
	. "github.com/MikeDanaher/go-tic-tac-toe/tic-tac-toe"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"reflect"
)

var _ = Describe("Console Setup", func() {
	var (
		mockReader bytes.Buffer
		mockWriter bytes.Buffer
	)

	It("Sets up a new human vs human console game", func() {
		SetResponses(&mockReader, []string{"1", "x", "o", "1"})
		config := NewConfig3x3(&mockReader, &mockWriter)
		setup := NewConsoleSetup(config)
		Expect(reflect.TypeOf(setup.Player1)).To(Equal(reflect.TypeOf(new(HumanPlayer))))
		Expect(reflect.TypeOf(setup.Player2)).To(Equal(reflect.TypeOf(new(HumanPlayer))))
	})

	It("Sets up a new human vs computer console game, computer goes first", func() {
		SetResponses(&mockReader, []string{"2", "x", "o", "2"})
		config := NewConfig3x3(&mockReader, &mockWriter)
		setup := NewConsoleSetup(config)
		Expect(reflect.TypeOf(setup.Player1)).To(Equal(reflect.TypeOf(new(ComputerPlayer))))
		Expect(reflect.TypeOf(setup.Player2)).To(Equal(reflect.TypeOf(new(HumanPlayer))))
	})

	It("Sets up a new computer vs computer console game", func() {
		SetResponses(&mockReader, []string{"3", "x", "o", "1"})
		config := NewConfig3x3(&mockReader, &mockWriter)
		setup := NewConsoleSetup(config)
		Expect(reflect.TypeOf(setup.Player1)).To(Equal(reflect.TypeOf(new(ComputerPlayer))))
		Expect(reflect.TypeOf(setup.Player2)).To(Equal(reflect.TypeOf(new(ComputerPlayer))))
	})

})
