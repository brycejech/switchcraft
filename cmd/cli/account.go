package cli

import (
	"fmt"
	"log"
	"switchcraft/core"
	"switchcraft/types"
	"time"

	"github.com/spf13/cobra"
)

func registerAccountModule(core *core.Core) {
	var accountCmd = &cobra.Command{
		Use:   "account",
		Short: "SwitchCraft CLI account module",
	}
	rootCmd.AddCommand(accountCmd)

	/* ------------------------ */
	/* === GET ACCOUNTS CMD === */
	/* ------------------------ */
	var getAccountsTenantID int64
	var getAccountsCmd = &cobra.Command{
		Use:   "getMany",
		Short: "Get multiple accounts",
		Run: func(cmd *cobra.Command, args []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			var tenantID *int64
			if cmd.Flags().Changed("tenantId") {
				tenantID = &getAccountsTenantID
			}

			accounts, err := core.AccountGetMany(opCtx, tenantID)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(accounts)
		},
	}
	getAccountsCmd.Flags().Int64Var(&getAccountsTenantID, "tenantId", 0, "account.createdBy")
	getAccountsCmd.MarkFlagRequired("tenantId")
	accountCmd.AddCommand(getAccountsCmd)

	/* -------------------------- */
	/* === CREATE ACCOUNT CMD === */
	/* -------------------------- */
	createAccountCmdArgs := struct {
		TenantID  int64
		FirstName string
		LastName  string
		Email     string
		Username  string
	}{}
	var createAccountCmd = &cobra.Command{
		Use:   "create",
		Short: "Create new account",
		Run: func(_ *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			args := core.NewAccountCreateArgs(
				createAccountCmdArgs.TenantID,
				createAccountCmdArgs.FirstName,
				createAccountCmdArgs.LastName,
				createAccountCmdArgs.Email,
				createAccountCmdArgs.Username,
				nil,
			)

			account, err := core.AccountCreate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(account)
		},
	}

	createAccountCmd.Flags().Int64Var(&createAccountCmdArgs.TenantID, "tenantId", 0, "account.tenantId")
	createAccountCmd.Flags().StringVar(&createAccountCmdArgs.FirstName, "firstName", "", "account.firstName")
	createAccountCmd.MarkFlagRequired("firstName")
	createAccountCmd.Flags().StringVar(&createAccountCmdArgs.LastName, "lastName", "", "account.lastName")
	createAccountCmd.MarkFlagRequired("lastName")
	createAccountCmd.Flags().StringVar(&createAccountCmdArgs.Email, "email", "", "account.email")
	createAccountCmd.MarkFlagRequired("email")
	createAccountCmd.Flags().StringVar(&createAccountCmdArgs.Username, "username", "", "account.username")
	createAccountCmd.MarkFlagRequired("username")
	accountCmd.AddCommand(createAccountCmd)

	/* ----------------------- */
	/* === GET ACCOUNT CMD === */
	/* ----------------------- */
	var getAccountTenantID int64
	var getAccountID int64
	var getAccountUUID string
	var getAccountUsername string
	var getAccountCmd = &cobra.Command{
		Use:   "getOne",
		Short: "Get an account by id, uuid, or username",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			var (
				tenantID *int64
				id       *int64
				uuid     *string
				username *string
			)
			if cmd.Flags().Changed("tenantId") {
				tenantID = &getAccountTenantID
			}
			if cmd.Flags().Changed("id") {
				id = &getAccountID
			}
			if cmd.Flags().Changed("uuid") {
				uuid = &getAccountUUID
			}
			if cmd.Flags().Changed("username") {
				username = &getAccountUsername
			}

			args := core.NewAccountGetOneArgs(
				tenantID,
				id,
				uuid,
				username,
			)

			account, err := core.AccountGetOne(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(account)
		},
	}
	getAccountCmd.Flags().Int64Var(&getAccountTenantID, "tenantId", 0, "account.tenantId")
	getAccountCmd.Flags().Int64Var(&getAccountID, "id", 0, "account.id")
	getAccountCmd.Flags().StringVar(&getAccountUUID, "uuid", "", "account.uuid")
	getAccountCmd.Flags().StringVar(&getAccountUsername, "username", "", "account.username")
	accountCmd.AddCommand(getAccountCmd)

	/* -------------------------- */
	/* === UPDATE ACCOUNT CMD === */
	/* -------------------------- */
	updateAccountCmdArgs := struct {
		TenantID  int64
		ID        int64
		FirstName string
		LastName  string
		Email     string
		Username  string
	}{}
	var updateAccountCmd = &cobra.Command{
		Use:   "update",
		Short: "Update an existing account",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			var tenantID *int64
			if cmd.Flags().Changed("tenantId") {
				tenantID = &updateAccountCmdArgs.TenantID
			}
			args := core.NewAccountUpdateArgs(
				tenantID,
				updateAccountCmdArgs.ID,
				updateAccountCmdArgs.FirstName,
				updateAccountCmdArgs.LastName,
				updateAccountCmdArgs.Email,
				updateAccountCmdArgs.Username,
			)

			account, err := core.AccountUpdate(opCtx, args)
			if err != nil {
				log.Fatal(err)
			}

			printJSON(account)
		},
	}

	updateAccountCmd.Flags().Int64Var(&updateAccountCmdArgs.TenantID, "tenantId", 0, "account.tenantId")
	updateAccountCmd.Flags().Int64Var(&updateAccountCmdArgs.ID, "id", 0, "account.id")
	updateAccountCmd.MarkFlagRequired("id")
	updateAccountCmd.Flags().StringVar(&updateAccountCmdArgs.FirstName, "firstName", "", "account.firstName")
	updateAccountCmd.MarkFlagRequired("firstName")
	updateAccountCmd.Flags().StringVar(&updateAccountCmdArgs.LastName, "lastName", "", "account.lastName")
	updateAccountCmd.MarkFlagRequired("lastName")
	updateAccountCmd.Flags().StringVar(&updateAccountCmdArgs.Email, "email", "", "account.email")
	updateAccountCmd.MarkFlagRequired("email")
	updateAccountCmd.Flags().StringVar(&updateAccountCmdArgs.Username, "username", "", "account.username")
	updateAccountCmd.MarkFlagRequired("username")
	accountCmd.AddCommand(updateAccountCmd)

	/* -------------------------- */
	/* === DELETE ACCOUNT CMD === */
	/* -------------------------- */
	var deleteAccountTenantID int64
	var deleteAccountID int64
	var deleteAccountCmd = &cobra.Command{
		Use:   "delete",
		Short: "Delete an account",
		Run: func(cmd *cobra.Command, _ []string) {
			authAccount := mustAuthn(core)
			opCtx := types.NewOperationCtx(baseCtx, "", time.Now(), *authAccount)

			var tenantID *int64
			if cmd.Flags().Changed("tenantId") {
				tenantID = &deleteAccountTenantID
			}
			if err := core.AccountDelete(opCtx, tenantID, deleteAccountID); err != nil {
				log.Fatal(err)
			}
			fmt.Printf("Account '%v' deleted successfully\n", deleteAccountID)
		},
	}

	deleteAccountCmd.Flags().Int64Var(&deleteAccountID, "id", 0, "account.id")
	deleteAccountCmd.MarkFlagRequired("id")
	accountCmd.AddCommand(deleteAccountCmd)
}
