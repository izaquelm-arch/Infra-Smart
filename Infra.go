package main

import (
	"fmt"
	"sync"
	"time"
)

type Semaforo struct {
	Tempo time.Duration // Tempo de permanência em emergência
	Color string
	Mutex sync.Mutex
}

type EstadoSinal struct {
	Cor     string
	Duracao time.Duration
}

func main() {
	semaforo := &Semaforo{
		Tempo: 5 * time.Second,
		Color: "Verde",
	}

	chPrioridade := make(chan string, 5)

	// Goroutine em segundo plano gerando emergências
	go func() {
		time.Sleep(3 * time.Second)
		chPrioridade <- "Ambulância SAMU 192"

		time.Sleep(10 * time.Second)
		chPrioridade <- "Carro de Bombeiros"
	}()

	// Ciclo normal das cores
	ciclo := []EstadoSinal{
		{Cor: "Verde", Duracao: 4 * time.Second},
		{Cor: "Amarelo", Duracao: 2 * time.Second},
		{Cor: "Vermelho", Duracao: 4 * time.Second},
	}

	fmt.Println("=== 🚥 Iniciando Sinalização de Trânsito Inteligente ===")

	i := 0
	for {
		estadoAtual := ciclo[i%len(ciclo)]

		semaforo.Mutex.Lock()
		semaforo.Color = estadoAtual.Cor
		fmt.Printf("\n🚥 [CICLO NORMAL] Cor alterada para: %s (Duração: %v)\n", semaforo.Color, estadoAtual.Duracao)
		semaforo.Mutex.Unlock()

		timer := time.NewTimer(estadoAtual.Duracao)

		// Verifica se há emergência ANTES de esperar o timer
		select {
		case msg := <-chPrioridade:
			timer.Stop()
			executarPrioridade(semaforo, msg)

		default:
			// Se não há emergência imediata, aguarda o timer OU uma emergência futura
			select {
			case msg := <-chPrioridade:
				timer.Stop()
				executarPrioridade(semaforo, msg)

			case <-timer.C:
				// Tempo normal encerrado com sucesso
			}
		}

		i++
	}
}

func executarPrioridade(s *Semaforo, msg string) {
	fmt.Printf("\n🚨 [ALERTA DE EMERGÊNCIA] Veículo detectado: %s\n", msg)

	s.Mutex.Lock()
	corAnterior := s.Color
	s.Color = "Vermelho"
	fmt.Printf("⛔ Sinal forçado para: %s (Garantindo travamento do cruzamento)\n", s.Color)
	s.Mutex.Unlock()

	// Aguarda o tempo de passagem do veículo de emergência
	time.Sleep(s.Tempo)

	s.Mutex.Lock()
	s.Color = corAnterior
	fmt.Printf("✅ Emergência finalizada. Estado restaurado para: %s\n", s.Color)
	s.Mutex.Unlock()
}