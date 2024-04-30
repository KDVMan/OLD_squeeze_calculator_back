package services_websocket

import (
	"backend/core/models"
	"backend/enums"
	"backend/services/exchange_limit"
	"backend/services/symbol"
	"encoding/json"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type ConnectionModel struct {
	ws               *websocket.Conn
	websocketService *WebsocketService
	sendChan         chan *core_models.BroadcastChannelModel
}

type WebsocketService struct {
	symbolService        *services_symbol.SymbolService
	exchangeLimitService *services_exchange_limit.ExchangeLimitService
	connections          map[*ConnectionModel]bool
	registerChan         chan *ConnectionModel
	unregisterChan       chan *ConnectionModel
	BroadcastChan        chan *core_models.BroadcastChannelModel
	ProgressChan         chan *core_models.ProgressChannelModel
	lock                 sync.Mutex
}

func New(symbolService *services_symbol.SymbolService, exchangeLimitService *services_exchange_limit.ExchangeLimitService) *WebsocketService {
	return &WebsocketService{
		symbolService:        symbolService,
		exchangeLimitService: exchangeLimitService,
		connections:          make(map[*ConnectionModel]bool),
		registerChan:         make(chan *ConnectionModel),
		unregisterChan:       make(chan *ConnectionModel),
		BroadcastChan:        make(chan *core_models.BroadcastChannelModel),
		ProgressChan:         make(chan *core_models.ProgressChannelModel),
	}
}

func HandleConnections(websocketService *WebsocketService, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)

	if err != nil {
		log.Println("Websocket upgrade error:", err)
		return
	}

	connection := &ConnectionModel{
		ws:               conn,
		websocketService: websocketService,
		sendChan:         make(chan *core_models.BroadcastChannelModel, 256),
	}

	websocketService.registerChan <- connection

	go connection.write()
	go connection.read()
}

func (websocketService *WebsocketService) Start() {
	for {
		select {
		case connection := <-websocketService.registerChan:
			websocketService.lock.Lock()
			websocketService.connections[connection] = true
			websocketService.lock.Unlock()

			log.Println("[Start] Connection Registered")

			go websocketService.broadcastSymbols()
			go websocketService.broadcastExchangeLimits()
		case connection := <-websocketService.unregisterChan:
			websocketService.lock.Lock()

			if _, ok := websocketService.connections[connection]; ok {
				delete(websocketService.connections, connection)
				close(connection.sendChan)

				log.Println("[Start] Connection Unregistered")
			}

			websocketService.lock.Unlock()
		case broadcast := <-websocketService.BroadcastChan:
			websocketService.broadcast(broadcast)

		case progress := <-websocketService.ProgressChan:
			// log.Println("[Start] Progress: ", progress)

			broadcastModel := &core_models.BroadcastChannelModel{
				Event: enums.WebsocketEventCalculator,
				Data:  progress,
			}

			websocketService.broadcast(broadcastModel)
		}
	}
}

func (websocketService *WebsocketService) broadcast(broadcastModel *core_models.BroadcastChannelModel) {
	websocketService.lock.Lock()
	defer websocketService.lock.Unlock()

	for connection := range websocketService.connections {
		select {
		case connection.sendChan <- broadcastModel:
			// log.Println("Message sent successfully:", broadcastModel)
		default:
			log.Println("Failed to send message:", broadcastModel)
		}
	}
}

func (connection *ConnectionModel) read() {
	defer func() {
		connection.websocketService.unregisterChan <- connection

		if err := connection.ws.Close(); err != nil {
			log.Println("Error closing websocket connection", err)
			return
		}
	}()

	for {
		_, _, err := connection.ws.ReadMessage()

		if err != nil {
			log.Println("Error reading websocket message:", err)
			break
		}
	}
}

func (connection *ConnectionModel) write() {
	for {
		broadcastModel, ok := <-connection.sendChan

		if !ok {
			if err := connection.ws.WriteMessage(websocket.CloseMessage, []byte{}); err != nil {
				log.Println("Error writing close message:", err)
				return
			}

			return
		}

		message, err := json.Marshal(broadcastModel)

		if err != nil {
			log.Println("Error marshaling message:", err)
			continue
		}

		if err := connection.ws.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("Error writing text message:", err)
			return
		}
	}
}
