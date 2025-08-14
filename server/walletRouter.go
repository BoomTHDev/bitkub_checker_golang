package server

import (
	"github.com/boomthdev/wld_check_bk/middleware"
	_walletController "github.com/boomthdev/wld_check_bk/pkg/wallet/controller"
	_walletRepository "github.com/boomthdev/wld_check_bk/pkg/wallet/repository"
	_walletService "github.com/boomthdev/wld_check_bk/pkg/wallet/service"
)

func (s *fiberServer) initWalletRouter() {
	walletRepository := _walletRepository.NewWalletRepositoryImpl(s.conf.Bitkub)
	walletService := _walletService.NewWalletService(walletRepository)
	walletController := _walletController.NewWalletController(walletService)

	walletRouter := s.app.Group("/api/v1/wallet")
	walletRouter.Get("/", middleware.ApiKeyAuth(), walletController.GetWallet)
}
