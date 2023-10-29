package services

import (
	"onden-backend/config"
	"time"

	"github.com/livekit/protocol/auth"
)

var livekitConfig *config.LiveKitConfig;

func LiveKitInit(config *config.LiveKitConfig) {
	livekitConfig = config;
}

func GetLiveKitToken(name, room string) (string, error) {
	cansubscribe := true;
	at := auth.NewAccessToken(livekitConfig.APIKey, livekitConfig.APISecret);
	grant := &auth.VideoGrant{
		RoomJoin: true,
		Room: room,
		CanSubscribe: &cansubscribe,
	};

	at.AddGrant(grant).SetIdentity(name).SetValidFor(time.Hour * 24);

	result, err := at.ToJWT();
	if err != nil {
		return "", err;
	}

	return result, nil;
}