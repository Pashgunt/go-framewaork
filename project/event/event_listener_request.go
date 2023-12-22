package event

import "project/dto"

type Subscriber chan *dto.ApiRequestDto

type Publisher struct {
	subscribers []Subscriber
}

func (p *Publisher) RequestSubscribe() Subscriber {
	ch := make(chan *dto.ApiRequestDto, 1)
	p.subscribers = append(p.subscribers, ch)
	return ch
}

func (p *Publisher) RequestPublish(data *dto.ApiRequestDto) {
	for _, subscriber := range p.subscribers {
		subscriber <- data
	}
}
