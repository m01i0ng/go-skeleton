package db

import (
  "time"

  mqtt "github.com/eclipse/paho.mqtt.golang"
  "github.com/google/wire"
  "github.com/kataras/golog"
  "github.com/m01i0ng/go-skeleton/internal/config"
  "github.com/m01i0ng/go-skeleton/util"
)

var MqttProvider = wire.NewSet(NewMqtt)

func NewMqtt(c *config.Config) (mqtt.Client, error) {
  m := c.Mqtt

  opts := mqtt.NewClientOptions().
    SetClientID(util.RandSeq(8)).
    AddBroker(m.Broker).
    SetUsername(m.User).
    SetPassword(m.Password).
    SetCleanSession(true).
    SetAutoReconnect(true).
    SetMaxReconnectInterval(time.Second * 5)
  client := mqtt.NewClient(opts)

  if token := client.Connect(); token.Wait() && token.Error() != nil {
    golog.Error(token.Error())
    return nil, token.Error()
  }

  golog.Infof("Mqtt connected to: %s", m.Broker)

  return client, nil
}
