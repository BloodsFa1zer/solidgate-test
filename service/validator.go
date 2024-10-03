package service

import (
	"github.com/rs/zerolog/log"
	"solidgate-test/model"
	"solidgate-test/util"
	"time"
)

type ValidatorInterface interface {
	IsValidCard(card model.Card) *util.Error
}

type CardValidator struct{}

func NewCardValidator() *CardValidator {
	log.Info().Msg("New CardValidator instance created")
	return &CardValidator{}
}

func (cv *CardValidator) IsValidCard(card model.Card) *util.Error {
	log.Info().Str("cardNumber", card.Number).Int("expiryYear", card.ExpirationDate.Year).Int("expiryMonth", card.ExpirationDate.Month).Msg("Validating card")

	isDateValid, err := cv.isCardDateValid(card.ExpirationDate.Year, card.ExpirationDate.Month)
	isLuhnValid := cv.isCardLuhnAlgorithmValid(card.Number)

	if !isLuhnValid && !isDateValid {
		log.Warn().Err(util.InvalidCardNumberAndExpirationDateError()).Msg("Both Card Number and Expiry Date invalid")
		return util.InvalidCardNumberAndExpirationDateError()
	} else if !isLuhnValid {
		log.Warn().Err(util.InvalidCardNumberError()).Msg("Card Number invalid")
		return util.InvalidCardNumberError()
	} else if !isDateValid {
		log.Warn().Err(err).Msg("Expiry Date invalid")
		return err
	}

	log.Info().Msg("Card validation successful")
	return nil
}

func (cv *CardValidator) isCardLuhnAlgorithmValid(cardNumber string) bool {
	log.Debug().Str("cardNumber", cardNumber).Msg("Checking Luhn algorithm validity")

	if len(cardNumber) == 0 {
		log.Warn().Msg("Card number is empty")
		return false
	}

	total := 0
	isSecondDigit := false

	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit := int(cardNumber[i] - '0')

		if isSecondDigit {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		total += digit
		isSecondDigit = !isSecondDigit
	}

	isValid := total%10 == 0
	log.Debug().Bool("isValid", isValid).Msg("Luhn algorithm check result")
	return isValid
}

func (cv *CardValidator) isCardDateValid(cardExpirationYear, cardExpirationMonth int) (bool, *util.Error) {
	log.Debug().Int("expiryYear", cardExpirationYear).Int("expiryMonth", cardExpirationMonth).Msg("Checking card expiration date")

	currentYear, currentMonth, _ := time.Now().Date()

	if cardExpirationMonth < 1 || cardExpirationMonth > 12 {
		log.Warn().Err(util.InvalidExpirationDateError()).Msg("Month of expiry date is invalid")
		return false, util.InvalidExpirationDateError()
	}

	if cardExpirationYear < currentYear {
		log.Warn().Err(util.ExpiredCardError()).Msg("Year of expiry date is in the past")
		return false, util.ExpiredCardError()
	}

	if cardExpirationYear == currentYear && cardExpirationMonth < int(currentMonth) {
		log.Warn().Err(util.ExpiredCardError()).Msg("Card is expired")
		return false, util.ExpiredCardError()
	}

	log.Debug().Msg("Card expiration date is valid")
	return true, nil
}
