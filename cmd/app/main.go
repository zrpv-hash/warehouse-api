package main

import (
	"context"
	"os/signal"
	"syscall"

	"warehousesvc/internal/application/product/getall"
	"warehousesvc/internal/application/reserve/release"
	"warehousesvc/internal/application/reserve/reserve"
	"warehousesvc/internal/infrastructure/crosscutting/cleanenv"
	"warehousesvc/internal/infrastructure/crosscutting/logger"
	"warehousesvc/internal/infrastructure/crosscutting/pgclient"
	inventoryRepo "warehousesvc/internal/infrastructure/repository/inventory/pgsql"
	productRepo "warehousesvc/internal/infrastructure/repository/product/pgsql"
	reserveRepo "warehousesvc/internal/infrastructure/repository/reserve/pgsql"
	"warehousesvc/internal/infrastructure/tx/pgsqltx"
	httprest "warehousesvc/internal/interface/http_rest"
	"warehousesvc/internal/interface/http_rest/common"
	deleterelease "warehousesvc/internal/interface/http_rest/product/delete_release"
	postreserve "warehousesvc/internal/interface/http_rest/product/post_reserve"
	getproducts "warehousesvc/internal/interface/http_rest/warehouse/get_products"
)

// @title Warehouse API
// @version 1.0
// @description Warehouse API
// @termsOfService http://swagger.io/terms/
// @host localhost:8080
// @BasePath /
func main() {
	var handlers []common.Handler

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg, err := cleanenv.New()
	if err != nil {
		panic(err)
	}

	log := logger.New(cfg)

	pgc, err := pgclient.New(cfg)
	if err != nil {
		panic(err)
	}

	tx := pgsqltx.New(pgc)

	ir := inventoryRepo.New(pgc)
	rr := reserveRepo.New(pgc)
	pr := productRepo.New(pgc)

	resuc := reserve.New(tx, ir, rr)
	reluc := release.New(tx, ir, rr)
	guc := getall.New(pr)

	handlers = append(handlers, getproducts.New(guc), postreserve.New(resuc), deleterelease.New(reluc))

	httprest.New(ctx, cfg, log, handlers)
}
