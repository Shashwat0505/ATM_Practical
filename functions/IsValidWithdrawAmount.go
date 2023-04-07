package functions

func IsValidWithdrawAmount(withdrawCash float64, balanceCash float64) bool {

	if balanceCash < withdrawCash {
		return false

	}

	if withdrawCash > 500 && int(withdrawCash)%500 == 0 {
		return true
	}

	withdrawtemp := int(withdrawCash)
	if withdrawtemp%500 != 0 {
		return false

	}

	if !(balanceCash-withdrawCash >= 500) {
		return false

	}

	return true

}
