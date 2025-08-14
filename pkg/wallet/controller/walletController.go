package controller

import (
	_walletService "github.com/boomthdev/wld_check_bk/pkg/wallet/service"
	"github.com/gofiber/fiber/v2"
)

type WalletController struct {
	walletService _walletService.WalletService
}

func NewWalletController(walletService _walletService.WalletService) *WalletController {
	return &WalletController{walletService: walletService}
}

func (c *WalletController) GetWallet(ctx *fiber.Ctx) error {
	coin := ctx.Query("coin", "BTC")
	symbol := ctx.Query("symbol", "THB_BTC")
	resp, appErr := c.walletService.GetWallet(coin, symbol)
	if appErr != nil {
		return appErr
	}

	return ctx.Status(fiber.StatusOK).JSON(resp)
}
