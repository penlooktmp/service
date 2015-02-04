// Copyright 2014 Penlook Development Team. All rights reserved.
// Use of this source code is governed by
// license that can be found in the LICENSE file.
// Author : Loi Nguyen <loint@penlook.com>

package rabit

import (
	"github.com/streadway/amqp"
)

type Rabit struct {
	Name       string
	Server     string
	Connection *amqp.Connection
	Channel    *amqp.Channel
}

func (rabit Rabit) Connect() (Rabit, error) {

	connection, err := amqp.Dial(rabit.Server)

	if err != nil {
		return rabit, err
	}

	rabit.Connection = connection

	channel, err := rabit.Connection.Channel()

	if err != nil {
		return rabit, err
	}

	rabit.Channel = channel

	return rabit, nil
}

func (rabit Rabit) Sample() {

	for i := 0; i < 10; i++ {
		rabit.Channel.Publish(
			"example",
			"hello.world",
			false,
			false,
			amqp.Publishing{
				Headers:         amqp.Table{},
				ContentType:     "text/plain",
				ContentEncoding: "UTF-8",
				Body:            []byte("Hi"),
				DeliveryMode:    amqp.Transient,
				Priority:        0,
			},
		)
	}

}
