package main

// func (cfg *config) mathCommandCallback(bot *tgbotapi.BotAPI, _ *tgbotapi.Update, msg *tgbotapi.MessageConfig) error {
// 	const defaultErrormsg = "could not calculate: %v"
// 	u := tgbotapi.NewUpdate(1)
// 	u.Timeout = 60

// 	updates := bot.GetUpdatesChan(u)
// 	step := 1
// 	operations := []string{
// 		"addition",
// 		"subtraction",
// 		"multiplication",
// 		"division",
// 	}
// 	var firstNum, secondNum int

// 	for up := range updates {
// 		if up.Message == nil {
// 			return errors.New("invalid message")
// 		}

// 		switch step {
// 		case 1:
// 			mathMsg := tgbotapi.NewMessage(up.Message.From.ID, "Choose the first number")
// 			_, err := bot.Send(mathMsg)
// 			if err != nil {
// 				return fmt.Errorf(defaultErrormsg, err)
// 			}
// 			step = 2
// 		case 2:
// 			num, err := strconv.Atoi(up.Message.Text)
// 			if err != nil {
// 				return errors.New("please put a valid number")
// 			}
// 			firstNum = num

// 			mathMsg := tgbotapi.NewMessage(up.Message.From.ID, "Choose the second number")
// 			_, err = bot.Send(mathMsg)
// 			if err != nil {
// 				return fmt.Errorf(defaultErrormsg, err)
// 			}
// 			step = 3
// 		case 3:
// 			num, err := strconv.Atoi(up.Message.Text)
// 			if err != nil {
// 				return errors.New("please put a valid number")
// 			}
// 			secondNum = num

// 			mathMsg := tgbotapi.NewMessage(up.Message.From.ID, "Choose the operation")
// 			for _, opr := range operations {
// 				mathMsg.Text += "\n" + opr
// 			}
// 			_, err = bot.Send(mathMsg)
// 			if err != nil {
// 				return fmt.Errorf(defaultErrormsg, err)
// 			}
// 			step = 4
// 		case 4:
// 			var result int
// 			switch up.Message.Text {
// 			case operations[0]:
// 				result = firstNum + secondNum
// 			case operations[1]:
// 				result = firstNum - secondNum
// 			case operations[2]:
// 				result = firstNum * secondNum
// 			case operations[3]:
// 				if secondNum == 0 {
// 					return errors.New("division by zero is not allowed")
// 				}
// 				result = firstNum / secondNum
// 			}
// 			bot.Send(tgbotapi.NewMessage(up.Message.From.ID, fmt.Sprintf("The result of the operation is %v", result)))
// 			return nil
// 		}
// 	}
// 	return nil
// }
