package ewallet_integration_test

import (
	"encoding/json"
	. "github.com/onsi/gomega"
	"io"
	"net/http"
	"testing"
)

func TestEwallet(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("Getting products", func(t *testing.T) {
		httpResponse, err := http.Get("http://localhost:8080/products")
		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(httpResponse).Should(
			HaveHTTPStatus(http.StatusOK))
		body, err := io.ReadAll(httpResponse.Body)
		g.Expect(err).Should(
			Not(HaveOccurred()))
		var data []interface{}
		err = json.Unmarshal(body, &data)
		g.Expect(err).Should(
			Not(HaveOccurred()))
		g.Expect(data).Should(
			ContainElement(And(
				HaveKeyWithValue("id", "ff5d8560-c82e-44f0-9947-930852ca5ecc"),
				HaveKeyWithValue("priceInCents", BeEquivalentTo(4390)),
				HaveKeyWithValue("title", "Suporte para Notebook"),
				HaveKeyWithValue("description", "Ideal para notebook ou tablet de qualquer tamanho"),
				HaveKeyWithValue("discount", And(
					HaveKeyWithValue("percentage", BeEquivalentTo(0)),
					HaveKeyWithValue("valueInCents", BeEquivalentTo(0)))))))
	})
}
