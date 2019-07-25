package main

import (
	"github.com/gorilla/sessions"
)
import "encoding/json"

type state struct {
	Stage  int
	Attack int
	Money  int
	Name   string
}

func (s *state) getState() (int, int, int, string) {
	return s.Stage, s.Attack, s.Money, s.Name
}

func (s *state) setState(stage, attack, money int, name string) {
	s.Stage = stage
	s.Attack = attack
	s.Money = money
	s.Name = name
}

func loadFromSession(se *sessions.Session) state {
	s := state{}
	stage := se.Values["stage"].(int)
	name := se.Values["name"].(string)
	money := se.Values["money"].(int)
	attack := se.Values["attack"].(int)
	s.setState(stage, attack, money, name)
	return s
}

func loadFromSessionToJson(se *sessions.Session) ([]byte, error) {
	s := loadFromSession(se)
	j, err := json.Marshal(s)
	if err != nil {
		return []byte{}, err
	}
	return j, nil
}

// Msg

type SendMsg struct {
	From    string
	Content string
	Options []string
}

func (s *SendMsg) LoadSendToJson() ([]byte, error) {
	j, err := json.Marshal(s)
	if err != nil {
		return []byte{}, err
	}
	return j, nil
}

// Response

type RespMsg struct {
	From    string
	Content string
	Options []string
}

func ReadRespMsg(b []byte) (RespMsg, error) {
	var r RespMsg

	err := json.Unmarshal(b, r)
	if err != nil {
		return r, err
	}
	return r, nil
}
