package repository

import (
	"context"
	"time"
)

func (r *AppRepository) UpsertRefreshToken(ctx context.Context, userId int64, refreshToken string, expires_at time.Time) error {
	_, err := r.DB.Exec(
		ctx,
		`INSERT INTO refresh_tokens (user_id, token, expires_at)
		VALUES($1, $2, $3)
		ON CONFLICT(user_id)
		DO UPDATE SET
		token=EXCLUDED.token, 
		expires_at=EXCLUDED.expires_at;
		`,
		userId, refreshToken, expires_at,
	)

	return err
}
