package main

import (
	"strconv"
	"strings"

	"github.com/lxbot/lxlib/v2"
	"github.com/lxbot/lxlib/v2/common"
)

func main() {
	script, msgCh := lxlib.NewScript()

	for {
		msg := <-*msgCh

		if msg == nil {
			continue
		}

		next, err := msg.Copy()
		if err != nil {
			common.ErrorLog("message copy error:", err)
			continue
		}
		next.ResetContents()

		stateKey := "kip-counter_state_" + msg.Room.ID
		state := script.GetStorage(stateKey)
		if state == nil {
			state = ""
		}

		countKey := "kip-counter_count_" + msg.Room.ID
		countVal := script.GetStorage(countKey)
		shouldSetCount := false
		count := 0.0
		if countVal != nil {
			count = countVal.(float64)
		}

		for _, c := range msg.Contents {
			text := strings.ToLower(c.Text)
			if strings.HasPrefix(c.Text, "!kip") {
				commands := strings.Split(text, " ")
				if len(commands) < 2 {
					next.Reply().AddContent("コマンドが不明です。\n`!kip [enable|disable|show]`")
					script.SendMessage(next)
					continue
				}
				switch commands[1] {
				case "enable":
					script.SetStorage(stateKey, "enabled")
					next.Reply().AddContent("このルームの kipカウンター を有効化しました。")
					script.SendMessage(next)
					continue
				case "disable":
					script.SetStorage(stateKey, "disabled")
					next.Reply().AddContent("このルームの kipカウンター を無効化しました。")
					script.SendMessage(next)
					continue
				case "show":
					if state == "enabled" {
						next.Reply().AddContent(strconv.FormatFloat(count, 'f', -1, 64) + " ポイントです。")
						script.SendMessage(next)
					} else {
						next.Reply().AddContent("このルームでは kipカウンター が有効化されていません。\n有効化するには `!kip enable` と入力してください。")
						script.SendMessage(next)
					}
					continue
				default:
					next.Reply().AddContent("コマンドが不明です。\n`!kip [enable|disable|show]`")
					script.SendMessage(next)
					continue
				}
			} else if strings.Contains(text, "kip") || strings.Contains(text, "ｋｉｐ") || strings.Contains(text, "きもいえす") || strings.Contains(text, "きもies") {
				switch state {
				case "enabled":
					count++
					shouldSetCount = true
					continue
				case "":
					next.Reply().AddContent("このルームでは kipカウンター が有効化されていません。\n有効化するには `!kip enable` と入力してください。\nこの通知を無効にするには `!kip disable` と入力してください。")
					script.SendMessage(next)
					continue
				}
			}
		}
		if shouldSetCount {
			script.SetStorage(countKey, count)
		}
	}
}
