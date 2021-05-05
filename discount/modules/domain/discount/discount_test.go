package discount

import (
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
	"testing"
)

func TestDiscount(t *testing.T) {
	g := NewGomegaWithT(t)
	t.Run("Do not add percentage calculates same value", func(t *testing.T) {
		maxPercentage := 10 * Percent
		valueInCents := 100
		sut := NewDiscount(valueInCents, maxPercentage)

		g.Expect(sut).Should(
			MatchFields(IgnoreExtras, Fields{
				"Percentage":   BeEquivalentTo(0 * Percent),
				"ValueInCents": BeEquivalentTo(valueInCents),
			}))
	})
	t.Run("Adding zero percentage calculates same value", func(t *testing.T) {
		maxPercentage := 10 * Percent
		valueInCents := 100
		sut := NewDiscount(valueInCents, maxPercentage)

		sut.AddByPercentage(0 * Percent)

		g.Expect(sut).Should(
			MatchFields(IgnoreExtras, Fields{
				"Percentage":   BeEquivalentTo(0 * Percent),
				"ValueInCents": BeEquivalentTo(valueInCents),
			}))
	})
	t.Run("Adding percentage less than max percentage calculates normally", func(t *testing.T) {
		maxPercentage := 10 * Percent
		valueInCents := 100
		sut := NewDiscount(valueInCents, maxPercentage)

		sut.AddByPercentage(9 * Percent)

		g.Expect(sut).Should(
			MatchFields(IgnoreExtras, Fields{
				"Percentage":   BeEquivalentTo(9 * Percent),
				"ValueInCents": BeEquivalentTo(91),
			}))
	})
	t.Run("Adding percentage same as max percentage calculates normally", func(t *testing.T) {
		maxPercentage := 10 * Percent
		valueInCents := 100
		sut := NewDiscount(valueInCents, maxPercentage)

		sut.AddByPercentage(maxPercentage)

		g.Expect(sut).Should(
			MatchFields(IgnoreExtras, Fields{
				"Percentage":   BeEquivalentTo(maxPercentage),
				"ValueInCents": BeEquivalentTo(90),
			}))
	})
	t.Run("Adding percentage more than max percentage calculates based on max percentage", func(t *testing.T) {
		maxPercentage := 10 * Percent
		valueInCents := 100
		sut := NewDiscount(valueInCents, maxPercentage)

		sut.AddByPercentage(11 * Percent)

		g.Expect(sut).Should(
			MatchFields(IgnoreExtras, Fields{
				"Percentage":   BeEquivalentTo(maxPercentage),
				"ValueInCents": BeEquivalentTo(90),
			}))
	})
	t.Run("Adding percentages less than max percentage calculates normally", func(t *testing.T) {
		maxPercentage := 10 * Percent
		valueInCents := 100
		sut := NewDiscount(valueInCents, maxPercentage)

		sut.AddByPercentage(3 * Percent)
		sut.AddByPercentage(4 * Percent)

		g.Expect(sut).Should(
			MatchFields(IgnoreExtras, Fields{
				"Percentage":   BeEquivalentTo(7 * Percent),
				"ValueInCents": BeEquivalentTo(93),
			}))
	})
	t.Run("Adding percentages more than max percentage calculates based on max percentage", func(t *testing.T) {
		maxPercentage := 10 * Percent
		valueInCents := 100
		sut := NewDiscount(valueInCents, maxPercentage)

		sut.AddByPercentage(7 * Percent)
		sut.AddByPercentage(4 * Percent)

		g.Expect(sut).Should(
			MatchFields(IgnoreExtras, Fields{
				"Percentage":   BeEquivalentTo(maxPercentage),
				"ValueInCents": BeEquivalentTo(90),
			}))
	})
}
