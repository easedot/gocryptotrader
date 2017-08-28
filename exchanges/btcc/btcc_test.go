package btcc

import (
	"testing"
	"time"

	"github.com/thrasher-/gocryptotrader/config"
)

// Please supply your own APIkeys here to do better tests
const (
	apiKey    = ""
	apiSecret = ""
)

var b BTCC

func TestSetDefaults(t *testing.T) {
	b.SetDefaults()
}

func TestSetup(t *testing.T) {
	conf := config.ExchangeConfig{
		Enabled: true,
	}
	b.Setup(conf)

	conf = config.ExchangeConfig{
		Enabled:   false,
		APIKey:    apiKey,
		APISecret: apiSecret,
	}
	b.Setup(conf)
}

func TestGetFee(t *testing.T) {
	if b.GetFee() != 0 {
		t.Error("Test failed - GetFee() error")
	}
}

func TestGetTicker(t *testing.T) {
	_, err := b.GetTicker("ltccny")
	if err != nil {
		t.Error("Test failed - GetTicker() error", err)
	}
}

func TestGetTradesLast24h(t *testing.T) {
	_, err := b.GetTradesLast24h("ltccny")
	if err != nil {
		t.Error("Test failed - GetTradesLast24h() error", err)
	}
}

func TestGetTradeHistory(t *testing.T) {
	_, err := b.GetTradeHistory("ltccny", 0, 0, time.Time{})
	if err != nil {
		t.Error("Test failed - GetTradeHistory() error", err)
	}
}

func TestGetOrderBook(t *testing.T) {
	_, err := b.GetOrderBook("ltccny", 100)
	if err != nil {
		t.Error("Test failed - GetOrderBook() error", err)
	}
	_, err = b.GetOrderBook("ltccny", 0)
	if err != nil {
		t.Error("Test failed - GetOrderBook() error", err)
	}
}

func TestGetAccountInfo(t *testing.T) {
	err := b.GetAccountInfo("")
	if err == nil {
		t.Error("Test failed - GetAccountInfo() error", err)
	}
}
