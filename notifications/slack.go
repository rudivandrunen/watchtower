package notifications

import (
	"github.com/johntdyer/slackrus"
	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

const (
	slackType = "slack"
)

type slackTypeNotifier struct {
	slackrus.SlackrusHook
}

func newSlackNotifier(c *cli.Context) typeNotifier {
	n := &slackTypeNotifier{
		SlackrusHook: slackrus.SlackrusHook{
			HookURL:        c.GlobalString("notification-slack-hook-url"),
			Username:       c.GlobalString("notification-slack-identifier"),
			AcceptedLevels: slackrus.LevelThreshold(logrus.DebugLevel),
		},
	}

	log.AddHook(n)

	return n
}

func (s *slackTypeNotifier) StartNotification() {}

func (s *slackTypeNotifier) SendNotification() {}
