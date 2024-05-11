package models

func RefreshTokenUserIDKey(jwtID string) string {
	return "refresh_token:" + jwtID + ":user_id"
}
