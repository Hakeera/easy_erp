package logs

import (
	"log"

	"go.uber.org/zap"
)

var Logger *zap.Logger

func InitLogger() {
	var err error
	Logger, err = zap.NewProduction() // Usa o modo de produção (pode mudar para NewDevelopment para mais detalhes)
	if err != nil {
		log.Fatalf("Erro ao iniciar logger: %v", err)
	}
	defer Logger.Sync() // Garante que os logs sejam descarregados corretamente ao encerrar
}

