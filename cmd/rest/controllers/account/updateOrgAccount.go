package account

import (
	"net/http"
	"strconv"
	"switchcraft/cmd/rest/restutils"
	"switchcraft/types"
)

type updateOrgAccountArgs struct {
	ID        int64  `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Username  string `json:"username"`
}

func (c *accountController) UpdateOrgAccount(w http.ResponseWriter, r *http.Request) {
	orgSlug := r.PathValue("orgSlug")
	accountIDStr := r.PathValue("accountID")
	if orgSlug == "" || accountIDStr == "" {
		restutils.NotFound(w, r)
		return
	}

	body := &updateOrgAccountArgs{}
	if err := restutils.DecodeBody(r, body); err != nil {
		restutils.JSONParseError(w, r)
		return
	}

	var (
		accountID int64
		err       error
	)
	if accountID, err = strconv.ParseInt(accountIDStr, 10, 64); err != nil {
		restutils.BadRequest(w, r)
		return
	}

	account, err := c.core.AccountGetOne(r.Context(),
		c.core.NewAccountGetOneArgs(&orgSlug, &accountID, nil, nil),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	tracer, _ := r.Context().Value(types.CtxOperationTracer).(types.OperationTracer)

	if account.ID != body.ID {
		restutils.BadRequest(w, r)
		c.logger.Warn(tracer, "Account update ID mismatch detected", map[string]any{
			"user":            tracer.AuthAccount.Username,
			"requestBody":     body,
			"existingAccount": account,
		})
		return
	}

	updatedAccount, err := c.core.AccountUpdate(r.Context(),
		c.core.NewAccountUpdateArgs(
			&orgSlug,
			account.ID,
			account.IsInstanceAdmin,
			body.FirstName,
			body.LastName,
			body.Email,
			body.Username,
		),
	)
	if err != nil {
		restutils.HandleCoreErr(w, r, err)
		return
	}

	restutils.Render(w, r, http.StatusOK, updatedAccount)
}
