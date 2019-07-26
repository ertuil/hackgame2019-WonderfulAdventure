package main

import (
	"errors"
	"github.com/gorilla/websocket"
	"math/rand"
	"strconv"
	"time"
)

// 初始状态
func Stage0(c *websocket.Conn, se *state) error {

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
		se.Money = se.Money - doorprice
	} else {
		se.Name = s0a1
	}

	se.Stage = GetStateTrans(se.Stage, opt)

	return nil
}

// 国王
func Stage1(c *websocket.Conn, se *state) error {

	j, err := MsgInitJson(s1f1, s1c1, []string{}, *se)
	if err != nil {
		return err
	}
	_ = c.WriteMessage(websocket.TextMessage, j)

	time.Sleep(1000 * time.Microsecond)

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
func Stage2(c *websocket.Conn, se *state) error {
	j, err := MsgInitJson(s2f0, s2c0, []string{}, *se)

	if err != nil {
		return err
	}
	c.WriteMessage(websocket.TextMessage, j)

	se.Stage = failstate
	return errors.New("你被门卫击败了。")
}

// 	去哪里
func Stage3(c *websocket.Conn, se *state) error {

	j, err := MsgInitJson(s3f0, s3c0, []string{s3a0, s3a1, s3a2}, *se)
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

	se.Stage = GetStateTrans(se.Stage, opt)

	return nil
}

// 大市场 漏洞主体
func Stage4(c *websocket.Conn, se *state) error {

	j, err := MsgInitJson(s4f0, s4c0, []string{s4a0, s4a1, s4a2, s4a3, s4a4}, *se)
	if err != nil {
		return err
	}
	_ = c.WriteMessage(websocket.TextMessage, j)

	_, ret, err := c.ReadMessage()
	if err != nil {
		return err
	}
	opt, err := strconv.Atoi(string(ret))

	if err != nil || opt*price > se.Money {
		j, err := MsgInitJson(s4f1, s4c2, []string{}, *se)
		if err != nil {
			return err
		}
		_ = c.WriteMessage(websocket.TextMessage, j)
	} else {
		se.Money = se.Money - opt*price
		se.Attack = se.Attack + opt*exp

		j, err := MsgInitJson(s4f1, s4c1, []string{}, *se)
		if err != nil {
			return err
		}
		_ = c.WriteMessage(websocket.TextMessage, j)
	}

	se.Stage = 3
	return nil
}

// 打怪刷机
func Stage5(c *websocket.Conn, se *state) error {

	rand.Seed(time.Now().Unix())
	dat := 10 * rand.Intn(20)

	j, err := MsgInitJson(s3f0, "你遇到一只攻击力为"+strconv.Itoa(dat)+"的小怪兽", []string{}, *se)
	if err != nil {
		return err
	}
	_ = c.WriteMessage(websocket.TextMessage, j)

	if se.Attack >= dat {
		se.Attack += beatexp
		j, err := MsgInitJson(s3f0, "你击败了这支小怪兽，获得了经验："+strconv.Itoa(beatexp), []string{}, *se)
		if err != nil {
			return err
		}
		_ = c.WriteMessage(websocket.TextMessage, j)
	} else {
		return errors.New("战斗力不足，你被击败了。")
	}
	se.Stage = 3
	return nil
}

// 恶龙巢穴
func Stage6(c *websocket.Conn, se *state) error {

	// 第一句话
	j, err := MsgInitJson(s6f0, s6c0, []string{s6a0, s6a1}, *se)
	if err != nil {
		return err
	}
	_ = c.WriteMessage(websocket.TextMessage, j)

	_, _, err = c.ReadMessage()
	if err != nil {
		return err
	}

	// 随机攻击力

	dragonAttck := rand.Intn(10000)*10000 + 23333

	// 第二句话
	j, err = MsgInitJson(s6f0, s6c1+strconv.Itoa(dragonAttck)+s6c11, []string{}, *se)
	if err != nil {
		return err
	}
	_ = c.WriteMessage(websocket.TextMessage, j)

	// 结果判断
	if se.Attack >= dragonAttck {
		j, err = MsgInitJson(s6f1, s6c2, []string{}, *se)
		if err != nil {
			return err
		}
		_ = c.WriteMessage(websocket.TextMessage, j)
	} else {
		return errors.New(s6c3)
	}

	se.Stage = 7
	return nil
}

func Stage7(c *websocket.Conn, se *state) error {

	// 第一句话
	j, err := MsgInitJson(s7f0, s7c0, []string{s7a0, s7a1, s7a2}, *se)
	if err != nil {
		return err
	}
	_ = c.WriteMessage(websocket.TextMessage, j)

	_, ret, err := c.ReadMessage()
	if err != nil {
		return err
	}
	opt, err := strconv.Atoi(string(ret))

	if err != nil || opt != 2 {
		se.Stage = 11
	} else {
		se.Stage = succstate
	}

	j, err = MsgInitJson(s7f0, s7c1, []string{}, *se)
	if err != nil {
		return err
	}
	_ = c.WriteMessage(websocket.TextMessage, j)

	return nil
}
