# Chat Application

Questa è un'applicazione di chat scritta in Go, che implementa un server e un client per la comunicazione in tempo reale.

## Struttura del Progetto

- **cmd/**
  - **client/**: Contiene il codice principale del client.
  - **server/**: Contiene il codice principale del server.
- **internal/**
  - **client/**: Contiene le utility e gli handler per il client.
  - **server/**: Contiene gli handler per il server.
  - **shared/**: Moduli condivisi tra client e server.
  - **types/**: Definizioni delle strutture dati utilizzate.
- **go.mod / go.sum**: File di gestione delle dipendenze.
- **.gitignore**: File di configurazione per escludere determinati file dal repository.

## Requisiti

- Go (Versione consigliata: 1.20+)

## Installazione e Avvio

1. Clona il repository:
   ```sh
   git clone <repository-url>
   cd Chat
   ```
2. Installa le dipendenze:
   ```sh
   go mod tidy
   ```
3. Avvia il server:
   ```sh
   go run cmd/server/main.go
   ```
4. Avvia il client (in un altro terminale):
   ```sh
   go run cmd/client/main.go
   ```

## Funzionalità

- Comunicazione in tempo reale tra client e server.
- Gestione delle connessioni multiple.
- Invio e ricezione di messaggi testuali.
- Struttura modulare per una facile estensibilità.

## Esempio di Utilizzo

Dopo aver avviato il server e il client, puoi inviare messaggi digitando direttamente nel terminale del client. Il server riceverà e inoltrerà i messaggi agli altri utenti connessi.