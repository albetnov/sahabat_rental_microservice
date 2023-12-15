package client

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"io"
	"net/http"
	"sahabatrental.com/stock_movement/v2/modules/stock_movement/models"
	"strconv"
	"time"
)

func NotifyStockCreated(movement models.StockMovement, earning models.Earning) error {
	clientUrl := viper.GetString("CLIENT_URL")

	if clientUrl == "" {
		return errors.New("no client url registered, aborting")
	}

	payload, err := json.Marshal(map[string]uint{
		"earning_id":        earning.ID,
		"stock_movement_id": movement.ID,
	})

	if err != nil {
		return err
	}

	// let's tolerate just 15 seconds
	timeout, _ := context.WithTimeout(context.Background(), time.Second*15)

	url := fmt.Sprintf("%s/api/webhooks/stock/created", clientUrl)

	client := &http.Client{}

	req, err := http.NewRequestWithContext(timeout, http.MethodPost, url, bytes.NewBuffer(payload))

	if err != nil {
		return err
	}

	timestamp := time.Now().Unix()
	sign := fmt.Sprintf("%d;%s", timestamp, string(payload))

	signed := hmac.New(sha256.New, []byte(viper.GetString("APP_KEY")))
	signed.Write([]byte(sign))

	req.Header.Add("Authorization", fmt.Sprintf("%x", signed.Sum(nil)))
	req.Header.Add("Timestamp", strconv.FormatInt(timestamp, 10))
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)

	if err != nil {
		return err
	}

	defer res.Body.Close()
	result, err := io.ReadAll(res.Body)

	if err != nil {
		return err
	}

	body := struct {
		Status string
	}{}

	if err := json.Unmarshal(result, &body); err != nil {
		return err
	}

	if res.StatusCode != 200 || body.Status != "ok" {
		return errors.New("something went wrong in target client")
	}

	return nil
}
