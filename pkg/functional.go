package pkg

import "practiva/web/internal/domain"

func Any(slice []domain.Product, predicate func(domain.Product) bool) bool {
	for _, product := range slice {
		if predicate(product) {
			return true
		}
	}
	return false
}
