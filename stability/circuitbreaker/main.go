package main

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog"
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

type Service interface {
	Get() (string, error)
}

// UnstableServiceConf is the configuration structure for the UnstableService.
type UnstableServiceConf struct {
	// StopAfter is the number of calls that will happen before stopping.
	StopAfter int
	// StopDuring is the duration of the "fails".
	StopDuring time.Duration
}

// UnstableService represents a remote service that will respond erratically.
// It will reply to a certain number of time, then break for a certain time.
type UnstableService struct {
	conf        UnstableServiceConf
	nbCall      int
	brokenUntil time.Time
}

func (s *UnstableService) Get() (string, error) {
	logger := log.With().Str("source", "unstable service").Logger()

	if s.nbCall > s.conf.StopAfter {
		// Service broken but restore after a certain time.
		if time.Now().After(s.brokenUntil) {
			s.nbCall = 0
			logger.Debug().
				Msg("service restored")
		} else {
			logger.Debug().
				Msg("still broken")
			return "", errors.New("500 Internal Server Error (from UnstableService, still broken)")
		}
	}

	s.nbCall++

	// Breaks if too many calls.
	if s.nbCall > s.conf.StopAfter {
		s.brokenUntil = time.Now().Add(s.conf.StopDuring)
		logger.Debug().
			Time("broken-until", s.brokenUntil).
			Msg("too many calls")
		return "", errors.New("500 Internal Server Error (from UnstableService, just broken)")
	}

	return fmt.Sprintf("Pong! (nb of call without an error: %v)", s.nbCall), nil
}

func NewUnstableService(conf UnstableServiceConf) *UnstableService {
	return &UnstableService{
		conf: conf,
	}
}

type CircuitBreakerConf struct {
	// ErrorMax represents the number of consecutive failures from the service
	// before triggering the circuit breaker.
	ErrorMax int
	// Timeout represents the number of seconds to wait
	// before a new request to the service.
	Timeout time.Duration
}

// CircuitBreaker will act like a middleware or an adapter for the UnstableService.
type CircuitBreaker struct {
	conf        CircuitBreakerConf
	s           Service
	lastCall    time.Time
	errorsCount int
}

func (c *CircuitBreaker) Get() (string, error) {
	logger := log.With().Str("source", "circuit breaker").Logger()

	if c.errorsCount >= c.conf.ErrorMax {
		// If the circuit breaker is open, check if the time for a new test is passed.
		if time.Now().After(c.lastCall.Add(c.conf.Timeout)) {
			c.errorsCount = 0
			logger.Debug().
				Msg("timeout, let's try again")
		} else {
			logger.Debug().
				Msg("too many errors, open circuit, no more call to the service before some time")
			return "", errors.New("503 Service Unavailable - Please try later, go grab a coffee (from CircuitBreaker)")
		}
	}

	msg, err := c.s.Get()
	c.lastCall = time.Now()
	if err != nil {
		c.errorsCount++
		return msg, err
	}
	c.errorsCount = 0
	return msg, nil
}

func NewCircuitBreaker(conf CircuitBreakerConf, s Service) *CircuitBreaker {
	return &CircuitBreaker{conf: conf, s: s}
}

type ConsumerConf struct {
	// NbCall is the number of calls to	the	UnstableService before stopping.
	NbCall int
	// Period is the duration between calls to the UnstableService.
	Period time.Duration
}

type Consumer struct {
	conf        ConsumerConf
	errorsCount int
}

func (c *Consumer) Consume(s Service) {
	for i := 0; i < c.conf.NbCall; i++ {
		v, err := s.Get()
		if err != nil {
			log.Err(err).Msg("")
		} else {
			log.Info().Msg(v)
		}
		time.Sleep(c.conf.Period)
	}
}

func NewConsumer(conf ConsumerConf) Consumer {
	return Consumer{
		conf: conf,
	}
}

func main() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	log.Info().Msg("Circuit Breaker Demo")

	var s Service
	s = NewUnstableService(UnstableServiceConf{StopAfter: 4, StopDuring: 12 * time.Second})
	s = NewCircuitBreaker(CircuitBreakerConf{ErrorMax: 3, Timeout: 5 * time.Second}, s)
	c := NewConsumer(ConsumerConf{NbCall: 20, Period: 1 * time.Second})

	c.Consume(s)
}
