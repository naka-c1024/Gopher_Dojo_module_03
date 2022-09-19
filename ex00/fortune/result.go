package fortune

func RtnHealth(rdmNum int) string {
	switch rdmNum {
	case 0:
		return "Very good. It will improve with exercise."
	case 1, 2:
		return "Good."
	case 3:
		return "No problems."
	case 4:
		return "No problems. However, do not binge eat or drink."
	case 5:
		return "Try to make sure your body does not become cold."
	case 6:
		return "Your worries will continue. Sometimes, you just need to rest your mind."
	default:
		return "??"
	}
}

func RtnWish(rdmNum int) string {
	switch rdmNum {
	case 0:
		return "It will definitely come true."
	case 1:
		return "It shall come true without a hitch."
	case 2:
		return "Teamwork is the key to success. If you work together, it shall come true."
	case 3:
		return "If you take action, it will come true."
	case 4:
		return "It will take time, but it will happen."
	case 5:
		return "If you desire it, it will come true. However, this will come with much difficulty."
	case 6:
		return "It will be difficult for it to come true."
	default:
		return "??"
	}
}

func RtnFortune(rdmNum int) string {
	switch rdmNum {
	case 0:
		return "Dai-kichi"
	case 1:
		return "Kichi"
	case 2:
		return "Chuu-kichi"
	case 3:
		return "Sho-kichi"
	case 4:
		return "Sue-kichi"
	case 5:
		return "Kyo"
	case 6:
		return "Dai-kyo"
	default:
		return "??"
	}
}
