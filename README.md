# 🚥 Smart Traffic Light with Emergency Priority em Go

Um sistema autônomo e concorrente de controle de semáforo de trânsito inteligente desenvolvido em **Go (Golang)**, capaz de gerenciar o ciclo normal de sinalização e interrompê-lo instantaneamente ao detectar veículos de emergência (ambulâncias, bombeiros, viaturas) em tempo real.

---

## 📌 Sobre o Projeto

Em engenharia de tráfego e cidades inteligentes (*Smart Cities*), sistemas de sinalização precisam garantir tanto o fluxo contínuo de veículos quanto a prioridade absoluta para serviços de urgência. 

Este projeto demonstra a implementação de um semáforo reativo utilizando o modelo de concorrência nativo do Go. O sistema altera as cores do sinal de forma temporizada e escuta um canal de eventos assíncronos. Ao receber uma notificação de emergência, o ciclo normal é suspenso imediatamente e o cruzamento é travado em **Vermelho** até a passagem do veículo prioritário.

---

## 🚀 Diferenciais Técnicos & Conceitos Aplicados

- **Padrão *Double Select* (Prioridade Estrita)**: Por padrão, a estrutura `select` do Go sorteia aleatoriamente entre `cases` que estejam prontos ao mesmo tempo. Para evitar que o temporizador normal ignore uma emergência simultânea, utiliza-se um `select` aninhado com `default` para checar o canal de prioridade primeiro.
- **Gerenciamento de Recursos com `time.Timer`**: Em vez de travar a thread com `time.Sleep`, o ciclo utiliza `time.NewTimer`. Isso permite abortar e limpar a contagem de tempo pendente com `timer.Stop()` no exato milissegundo em que um evento de emergência é ingerido.
- **Goroutines Assíncronas**: A simulação de eventos de tráfego e veículos de emergência roda em segundo plano sem bloquear o motor do semáforo.
- **Proteção de Estado (`sync.Mutex`)**: Garantia de consistência do estado e cor atual do semáforo durante transições de emergência.

---

## 🛠️ Como Executar

### Pré-requisitos
- [Go](https://golang.org/doc/install) instalado (versão 1.18 ou superior).

### Passo a Passo

1. Clone o repositório:
   ```bash
   git clone [https://github.com/izaquelm-arch/Infra-Smart.git](https://github.com/izaquelm-arch/Infra-Smart)
