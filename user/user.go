package user

import (
	"context"

	"github.com/wutipong/mangaweb3-backend/ent"
	"github.com/wutipong/mangaweb3-backend/ent/user"
)

const (
	DEFAULT_EMAIL = "default@example.com"
)

func GetUser(ctx context.Context, client *ent.Client, email string) (u *ent.User, err error) {
	if email == "" {
		email = DEFAULT_EMAIL
	}

	u, err = client.User.Query().Where(
		user.Email(email),
	).Only(ctx)

	if !ent.IsNotFound(err) {
		return
	}

	u, err = client.User.Create().SetEmail(email).Save(ctx)

	return
}
