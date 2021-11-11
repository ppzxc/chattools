package utils

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/metadata"
)

func ContextValueExtractor(ctx context.Context, fields logrus.Fields) logrus.Fields {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if sessionId := md.Get("session.id"); len(sessionId) > 0 {
			fields["session.id"] = sessionId[0]
		}

		if sessionUserId := md.Get("session.user.id"); len(sessionUserId) > 0 {
			fields["session.user.id"] = sessionUserId[0]
		}

		if browserId := md.Get("browser.id"); len(browserId) > 0 {
			fields["session.browser.id"] = browserId[0]
		}

		if deviceId := md.Get("device.id"); len(deviceId) > 0 {
			fields["session.device.id"] = deviceId[0]
		}

		if transactionId := md.Get("transaction.id"); len(transactionId) > 0 {
			fields["transaction.id"] = transactionId[0]
		}

		if uuid := md.Get("uuid"); len(uuid) > 0 {
			fields["uuid"] = uuid[0]
		}
	} else {
		fields["session.id"] = ctx.Value("session.id")
		fields["session.user.id"] = ctx.Value("session.user.id")
		fields["session.browser.id"] = ctx.Value("session.browser.id")
		fields["session.device.id"] = ctx.Value("session.device.id")
		fields["transaction.id"] = ctx.Value("transaction.id")
		fields["uuid"] = ctx.Value("uuid")
	}

	return fields
}
