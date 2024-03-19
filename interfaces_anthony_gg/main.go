package main

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"
)

type AccountNotifier interface {
	NotifyAccountCreated(context.Context, Account) error
}

type Account struct {
	Username string
	Email    string
}

type AccountHandler struct {
	AccountNotifier AccountNotifier
}

type SimpleAccountNotifier struct {
}

func (n SimpleAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("account created by the simple notify", "username", account.Username, "email", account.Email)
	return nil
}

type BetterAccountNotifier struct {
}

func (n BetterAccountNotifier) NotifyAccountCreated(ctx context.Context, account Account) error {
	slog.Info("account created by the better notify", "username", account.Username, "email", account.Email)
	return nil

}

func (h *AccountHandler) handleCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account Account
	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		slog.Error("failed to decode the response body", "err", err)
		return
	}

	if err := h.AccountNotifier.NotifyAccountCreated(r.Context(), account); err != nil {
		slog.Error("failed to notify account created", "err", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(account)
}

func main() {
	mux := http.NewServeMux()

	accountHandler := &AccountHandler{
		AccountNotifier: BetterAccountNotifier{},
	}
	mux.HandleFunc("POST /account", accountHandler.handleCreateAccount)
	http.ListenAndServe(":3000", mux)
}
