package locking

import (
	"time"
)

// Lock ticket.
//
// If the ticket is the current lock holder, it will have its lease timeout set. If not, it will have its acquisition
// timeout set. Note that once a ticket is dereferenced by the manager, these rules no longer hold.
type Ticket interface {
	// Ticket ID.
	//
	// Identifies the specific locking attempt or lease.
	Id() int64

	// Acquired.
	//
	// Channel that will eventually emit the state of the acquisition attempt of the ticket.
	Acquired() <-chan bool
}

// Lock ticket implementation.
type ticketImpl struct {
	// Lease ID.
	id int64

	// First lease timeout upon acquisition.
	firstLeaseTimeout time.Duration

	// Acquisition notification channel.
	acquiredChan chan bool

	// Acquisition timeout as a monotonic timestamp.
	acquireTimeoutAt time.Duration

	// Lease timeout as a monotonic timestamp.
	leaseTimeoutAt time.Duration
}

func (t *ticketImpl) Id() int64 {
	return t.id
}

func (t *ticketImpl) Acquired() <-chan bool {
	return t.acquiredChan
}
