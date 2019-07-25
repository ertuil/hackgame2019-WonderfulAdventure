package main

const (
	failstate = 11
	succstate = 10
	initstate = 0

	s0f0 = "蒙达鲁克硫斯伯古比奇巴勒城门卫"
	s0c0 = "站住！进门收钱！一人5元，童叟无欺！不进的走开，要进的速速报上姓名！"
	s0p0 = "我是达拉崩吧斑得贝迪卜多比鲁翁，我从千里之外，前来帮助国王救回公主。"
	s0p1 = "我是宇宙无敌帅炸裂二狗蛋王小二，我来嘲笑国王和他被恶龙捉走的心爱的公主。"
	s0a0 = "达拉崩吧斑得贝迪卜多比鲁翁"
	s0a1 = "宇宙无敌帅炸裂二狗蛋王小二"

	s1f1 = "国王"
	s1c1 = "愁啊，愁啊！爱女米娅莫拉苏娜丹妮谢莉红竟被那恶龙带走了。"
	s1f2 = "国王"
	s1c2 = "若是有哪位勇士能够打败恶龙，大大滴有赏赐"
	s1p0 = "我达拉崩吧斑得贝迪卜多比鲁翁，愿意跨过山与大海，去会会那恶龙昆图库塔卡提考特苏瓦西拉松。"
	s1p1 = "我达拉崩吧斑得贝迪卜多比鲁翁，十二丧母，十六丧父亲，如今肥宅一个，若能救下公主，还请国王将公主嫁给我。"
	s1a0 = "好！！若你成功救得公主，我赐你黄金万两。"
	s1a1 = "大胆！拖下去斩了！"

	s2f0 = "蒙达鲁克硫斯伯古比奇巴勒城门卫"
	s2c0 = "大胆刁民，竟敢在蒙达鲁克硫斯伯古比奇巴勒城门前撒野！让我试试你的身手，若是倒在我布拉多尔次拉西德乐的枪下，也算是你的荣幸！看招！"
)

var stateTrans [][]int

func stateTransInit() {
	stateTrans = make([][]int, 11)
	stateTrans[0] = []int{1, 2}
	stateTrans[1] = []int{3, 11}
}

func GetStateTrans(from, opt int) int {
	if from > len(stateTrans) {
		return failstate
	}
	if opt > len(stateTrans[from]) {
		return failstate
	}
	return stateTrans[from][opt]
}
