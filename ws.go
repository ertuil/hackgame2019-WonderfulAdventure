package main

import (
	"github.com/gorilla/websocket"
	"strconv"
	"time"
)

// 初始状态
func Stage_0(c *websocket.Conn, se *state) error {

	j, err := MsgInitJson(s0f0, s0c0, []string{s0p0, s0p1}, *se)
	if err != nil {
		return err
	}
	_ = c.WriteMessage(websocket.TextMessage, j)

	_, ret, err := c.ReadMessage()
	if err != nil {
		return err
	}
	opt, err := strconv.Atoi(string(ret))
	if err != nil {
		return err
	}

	if opt == 0 {
		se.Name = s0a0
		se.Money = se.Money - 5
	} else {
		se.Name = s0a1
	}

	se.Stage = GetStateTrans(se.Stage, opt)

	return nil
}

// 国王
func Stage_1(c *websocket.Conn, se *state) error {

	j, err := MsgInitJson(s1f1, s1c1, []string{}, *se)
	if err != nil {
		return err
	}
	_ = c.WriteMessage(websocket.TextMessage, j)

	time.Sleep(300 * time.Microsecond)

	j, err = MsgInitJson(s1f2, s1c2, []string{s1p0, s1p1}, *se)
	if err != nil {
		return err
	}
	_ = c.WriteMessage(websocket.TextMessage, j)

	_, ret, err := c.ReadMessage()
	if err != nil {
		return err
	}
	opt, err := strconv.Atoi(string(ret))
	if err != nil {
		return err
	}

	if opt == 0 {
		j, err = MsgInitJson(s1f2, s1a0, []string{}, *se)
		if err != nil {
			return err
		}
		_ = c.WriteMessage(websocket.TextMessage, j)
	} else {
		j, err = MsgInitJson(s1f2, s1a1, []string{}, *se)
		if err != nil {
			return err
		}
		_ = c.WriteMessage(websocket.TextMessage, j)
	}

	se.Stage = GetStateTrans(se.Stage, opt)
	return nil
}

// 门口失败
func Stage_2(c *websocket.Conn, se *state) error {
	j, err := MsgInitJson(s2f0, s2c0, []string{}, *se)

	if err != nil {
		return err
	}

	c.WriteMessage(websocket.TextMessage, j)
	se.Stage = failstate
	return nil
}

// 门口失败
func Stage_3(c *websocket.Conn, se *state) error {

	se.Stage = succstate
	return nil
}
