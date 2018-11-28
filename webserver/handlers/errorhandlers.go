package handlers

import (
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func handleGeneralGrpcError(env *Env, err error) error {
	switch status.Convert(err).Code() {
	case codes.OK:
		return nil
	case codes.Internal:
		env.Logger.Errorw("Internal error during grpc request",
			"err", status.Convert(err).Err(),
		)
		return StatusData{
			Code: http.StatusInternalServerError,
			Data: map[string]string{
				"error": "Internal server error",
			},
		}
	default:
		env.Logger.Errorw("Unknown error during grpc request",
			"err", status.Convert(err).Err(),
		)
		return StatusData{
			Code: http.StatusInternalServerError,
			Data: map[string]string{
				"error": "Unknown error",
			},
		}
	}
}

func handleGetOneUserGrpcError(env *Env, err error) error {
	switch status.Convert(err).Code() {
	case codes.OK:
		return nil
	case codes.NotFound:
		env.Logger.Infow("Requested change of data for non-existing user")
		return StatusData{
			Code: http.StatusNotFound,
			Data: map[string]string{
				"error": "User not found",
			},
		}
	case codes.Internal:
		env.Logger.Errorw("Internal error during grpc request",
			"grpc", "user",
			"err", status.Convert(err).Err(),
		)
		return StatusData{
			Code: http.StatusInternalServerError,
			Data: map[string]string{
				"error": "Internal server error",
			},
		}
	default:
		env.Logger.Errorw("Unknown error during grpc request",
			"grpc", "user",
			"err", status.Convert(err).Err(),
		)
		return StatusData{
			Code: http.StatusInternalServerError,
			Data: map[string]string{
				"error": "Unknown error",
			},
		}
	}
}
