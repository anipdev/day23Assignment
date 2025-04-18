package utils

import (
	"fmt"
)

func GenerateImageFilename(productID, productName string) string {
	return fmt.Sprintf("%s-%s.jpg", productID, Slugify(productName))
}
