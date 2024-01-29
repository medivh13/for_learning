package circuitbreaker

import (
	"time"

	"github.com/sony/gobreaker"
)

func NewCircuitBreakerInstance() *gobreaker.CircuitBreaker {
	st := gobreaker.Settings{
		Name:        "integrationCircuitBreaker",
		MaxRequests: 20,
		Interval:    2 * time.Second,
		Timeout:     40 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 3
		},
	}

	cb := gobreaker.NewCircuitBreaker(st)
	return cb
}

// MaxRequests:
// Deskripsi: Jumlah maksimum permintaan yang dapat dilayani
// sebelum Circuit Breaker memutus sirkuit. Jika jumlah permintaan melebihi batas ini,
// Circuit Breaker dapat memutus sirkuit untuk sementara waktu.

// Interval:
// Deskripsi: Waktu interval antara dua pengukuran untuk memeriksa apakah
// sirkuit harus dibuka atau ditutup kembali.
// Interval ini berfungsi sebagai jendela waktu di mana kegagalan diukur untuk menentukan
// apakah sirkuit harus dibuka atau tidak.

// Timeout:
// Deskripsi: Waktu maksimum yang diizinkan untuk menunggu
// permintaan sebelum dianggap sebagai kegagalan.
// Jika waktu yang dihabiskan untuk menyelesaikan permintaan melebihi batas waktu ini, permintaan dianggap gagal.

// ReadyToTrip:
// Deskripsi: Ini adalah fungsi yang menentukan apakah sirkuit harus dibuka atau tidak.
// Fungsi ini menerima parameter counts yang berisi informasi tentang jumlah kegagalan berturut-turut,
// dan kemudian mengembalikan nilai boolean. Jika nilai yang dikembalikan adalah true, sirkuit akan dibuka.
