package auth

import (
	"Derzhavnaya/internal/db"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
)

type RateLimiter struct {
	attempts map[string]int
	mu       sync.Mutex
}

func NewRateLimiter() *RateLimiter {
	return &RateLimiter{attempts: make(map[string]int)}
}

// GetIP достает реальный IP пользователя с учетом Traefik
func GetIP(r *http.Request) string {
	if ip := r.Header.Get("X-Real-Ip"); ip != "" {
		return ip
	}
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return strings.Split(ip, ",")[0]
	}
	host, _, _ := net.SplitHostPort(r.RemoteAddr)
	return host
}

// Allow проверяет, не превысил ли IP лимит (например, 5 попыток)
func (rl *RateLimiter) Allow(ip string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()

	if rl.attempts[ip] >= 5 {
		return false
	}
	rl.attempts[ip]++

	// Очистим счетчик через 15 минут (упрощенно)
	go func() {
		time.Sleep(15 * time.Minute)
		rl.mu.Lock()
		delete(rl.attempts, ip)
		rl.mu.Unlock()
	}()

	return true
}

func FilterMenuByRole(items []db.WebMenuItem, userRole string) []db.WebMenuItem {
	var filtered []db.WebMenuItem
	for _, item := range items {
		if !item.Role.Valid || item.Role.String == "" {
			filtered = append(filtered, item)
			continue
		}
		if strings.EqualFold(item.Role.String, userRole) || strings.EqualFold(userRole, "admin") {
			filtered = append(filtered, item)
		}
	}
	return filtered
}
