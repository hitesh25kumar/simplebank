package gapi

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/hitesh25kumar/db/util"
	"github.com/hitesh25kumar/pb"
	db "github.com/techschool/simplebank/db/sqlc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {

	user, err := server.store.GetUser(ctx, req.GetUsername())
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "No user found: %s", err)
		}
		return nil, status.Errorf(codes.Internal, "Failed to find user: %s", err)
	}

	err = util.CheckPassword(req.Password, user.HashedPassword)

	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "Incorrect password: %s", err)
	}

	accessToken, _, err := server.tokenMaker.CreateToken(user.Username, server.config.AccessTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create access token: %s", err)
	}

	refreshToken, refreshPayload, err := server.tokenMaker.CreateToken(user.Username, server.config.RefreshTokenDuration)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create refresh token: %s", err)
	}
	session, err := server.store.CreateSession(ctx, db.CreateSessionParams{
		ID:           refreshPayload.ID,
		Username:     user.Username,
		RefreshToken: refreshToken,
		UserAgent:    "",
		ClientIp:     "",
		IsBlocked:    false,
		ExpiresAt:    refreshPayload.ExpiredAt,
	})

	rsp := *pb.LoginUserResponse{
		AccessToken: accessToken,
		User:        newUserResponse(user),
	}
	ctx.JSON(http.StatusOK, rsp)

	return nil, status.Errorf(codes.Unimplemented, "method LoginUser not implemented")
}
