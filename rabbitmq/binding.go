package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"
	log "github.com/sirupsen/logrus"
)

// DeclareBindings declares all bindings from a list of bindings
func (self *Declarator) DeclareBindings(bindings []Binding) {
	for _, binding := range bindings {
		self.Bind(binding)
	}
}

// Bind binds a queue or exchange to another exchange
func (self *Declarator) Bind(binding Binding) {
	if binding.DestinationType == "queue" {
		self.bindQueue(binding)
	}

	if binding.DestinationType == "exchange" {
		self.bindExchange(binding)
	}
}

func (self *Declarator) bindQueue(binding Binding) {
	err := self.conn.QueueBind(
		binding.Destination,
		binding.RoutingKey,
		binding.Source,
		binding.NoWait,
		amqp.Table{},
	)
	if err != nil {
		log.Error("[RabbitMQ] [Binding] Error when trying bind queue " + err.Error())
		return
	}

	log.Warn("[RabbitMQ] [Binding] Queue " + binding.Destination + " binded with exchange " + binding.Source)
}

func (self *Declarator) bindExchange(binding Binding) {
	err := self.conn.ExchangeBind(
		binding.Destination,
		binding.RoutingKey,
		binding.Source,
		binding.NoWait,
		amqp.Table{},
	)
	if err != nil {
		log.Error("[RabbitMQ] [Binding] Error when trying bind exchange " + err.Error())
		return
	}

	log.Warn("[RabbitMQ] [Binding] Exchange " + binding.Destination + " binded with exchange " + binding.Source)
}
