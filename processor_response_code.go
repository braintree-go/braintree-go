package braintree

import "strconv"

type ProcessorResponseCode int

const (
	ProcessorResponseCodeDoNotHonor                                                                                     ProcessorResponseCode = 2000
	ProcessorResponseCodeInsufficientFunds                                                                              ProcessorResponseCode = 2001
	ProcessorResponseCodeLimitExceeded                                                                                  ProcessorResponseCode = 2002
	ProcessorResponseCodeCardholdersActivityLimitExceeded                                                               ProcessorResponseCode = 2003
	ProcessorResponseCodeExpiredCard                                                                                    ProcessorResponseCode = 2004
	ProcessorResponseCodeInvalidCreditCardNumber                                                                        ProcessorResponseCode = 2005
	ProcessorResponseCodeInvalidExpirationDate                                                                          ProcessorResponseCode = 2006
	ProcessorResponseCodeNoAccount                                                                                      ProcessorResponseCode = 2007
	ProcessorResponseCodeCardAccountLengthError                                                                         ProcessorResponseCode = 2008
	ProcessorResponseCodeNoSuchIssuer                                                                                   ProcessorResponseCode = 2009
	ProcessorResponseCodeCardIssuerDeclinedCVV                                                                          ProcessorResponseCode = 2010
	ProcessorResponseCodeVoiceAuthorizationRequired                                                                     ProcessorResponseCode = 2011
	ProcessorResponseCodeProcessorDeclinedPossibleLostCard                                                              ProcessorResponseCode = 2012
	ProcessorResponseCodeProcessorDeclinedPossibleStolenCard                                                            ProcessorResponseCode = 2013
	ProcessorResponseCodeProcessorDeclinedFraudSuspected                                                                ProcessorResponseCode = 2014
	ProcessorResponseCodeTransactionNotAllowed                                                                          ProcessorResponseCode = 2015
	ProcessorResponseCodeDuplicateTransaction                                                                           ProcessorResponseCode = 2016
	ProcessorResponseCodeCardholderStoppedBilling                                                                       ProcessorResponseCode = 2017
	ProcessorResponseCodeCardholderStoppedAllBilling                                                                    ProcessorResponseCode = 2018
	ProcessorResponseCodeInvalidTransaction                                                                             ProcessorResponseCode = 2019
	ProcessorResponseCodeViolation                                                                                      ProcessorResponseCode = 2020
	ProcessorResponseCodeSecurityViolation                                                                              ProcessorResponseCode = 2021
	ProcessorResponseCodeDeclinedUpdatedCardholderAvailable                                                             ProcessorResponseCode = 2022
	ProcessorResponseCodeProcessorDoesNotSupportThisFeature                                                             ProcessorResponseCode = 2023
	ProcessorResponseCodeCardTypeNotEnabled                                                                             ProcessorResponseCode = 2024
	ProcessorResponseCodeSetUpErrorMerchant                                                                             ProcessorResponseCode = 2025
	ProcessorResponseCodeInvalidMerchantID                                                                              ProcessorResponseCode = 2026
	ProcessorResponseCodeSetUpErrorAmount                                                                               ProcessorResponseCode = 2027
	ProcessorResponseCodeSetUpErrorHierarchy                                                                            ProcessorResponseCode = 2028
	ProcessorResponseCodeSetUpErrorCard                                                                                 ProcessorResponseCode = 2029
	ProcessorResponseCodeSetUpErrorTerminal                                                                             ProcessorResponseCode = 2030
	ProcessorResponseCodeEncryptionError                                                                                ProcessorResponseCode = 2031
	ProcessorResponseCodeSurchargeNotPermitted                                                                          ProcessorResponseCode = 2032
	ProcessorResponseCodeInconsistentData                                                                               ProcessorResponseCode = 2033
	ProcessorResponseCodeNoActionTaken                                                                                  ProcessorResponseCode = 2034
	ProcessorResponseCodePartialApprovalForAmountInGroupIIIVersion                                                      ProcessorResponseCode = 2035
	ProcessorResponseCodeAuthorizationCouldNotBeFound                                                                   ProcessorResponseCode = 2036
	ProcessorResponseCodeAlreadyReversed                                                                                ProcessorResponseCode = 2037
	ProcessorResponseCodeProcessorDeclined                                                                              ProcessorResponseCode = 2038
	ProcessorResponseCodeInvalidAuthorizationCode                                                                       ProcessorResponseCode = 2039
	ProcessorResponseCodeInvalidStore                                                                                   ProcessorResponseCode = 2040
	ProcessorResponseCodeDeclinedCallForApproval                                                                        ProcessorResponseCode = 2041
	ProcessorResponseCodeInvalidClientID                                                                                ProcessorResponseCode = 2042
	ProcessorResponseCodeErrorDoNotRetryCallIssuer                                                                      ProcessorResponseCode = 2043
	ProcessorResponseCodeDeclinedCallIssuer                                                                             ProcessorResponseCode = 2044
	ProcessorResponseCodeInvalidMerchantNumber                                                                          ProcessorResponseCode = 2045
	ProcessorResponseCodeDeclined                                                                                       ProcessorResponseCode = 2046
	ProcessorResponseCodeCallIssuerPickUpCard                                                                           ProcessorResponseCode = 2047
	ProcessorResponseCodeInvalidAmount                                                                                  ProcessorResponseCode = 2048
	ProcessorResponseCodeInvalidSKUNumber                                                                               ProcessorResponseCode = 2049
	ProcessorResponseCodeInvalidCreditPlan                                                                              ProcessorResponseCode = 2050
	ProcessorResponseCodeCreditCardNumberDoesNotMatchMethodOfPayment                                                    ProcessorResponseCode = 2051
	ProcessorResponseCodeCardReportedAsLostOrStolen                                                                     ProcessorResponseCode = 2053
	ProcessorResponseCodeReversalAmountDoesNotMatchAuthorizationAmount                                                  ProcessorResponseCode = 2054
	ProcessorResponseCodeInvalidTransactionDivisionNumber                                                               ProcessorResponseCode = 2055
	ProcessorResponseCodeTransactionAmountExceedsTheTransactionDivisionLimit                                            ProcessorResponseCode = 2056
	ProcessorResponseCodeIssuerOrCardholderHasPutRestrictionOnCard                                                      ProcessorResponseCode = 2057
	ProcessorResponseCodeMerchantNotMastercardSecureCodeEnabled                                                         ProcessorResponseCode = 2058
	ProcessorResponseCodeAddressVerificationFailed                                                                      ProcessorResponseCode = 2059
	ProcessorResponseCodeAddressVerificationAndCardSecurityCodeFailed                                                   ProcessorResponseCode = 2060
	ProcessorResponseCodeInvalidTransactionData                                                                         ProcessorResponseCode = 2061
	ProcessorResponseCodeInvalidTaxAmount                                                                               ProcessorResponseCode = 2062
	ProcessorResponseCodePayPalBusinessAccountPreferenceResultedInTransactionFailing                                    ProcessorResponseCode = 2063
	ProcessorResponseCodeInvalidCurrencyCode                                                                            ProcessorResponseCode = 2064
	ProcessorResponseCodeRefundTimeLimitExceeded                                                                        ProcessorResponseCode = 2065
	ProcessorResponseCodePayPalBusinessAccountRestricted                                                                ProcessorResponseCode = 2066
	ProcessorResponseCodeAuthorizationExpired                                                                           ProcessorResponseCode = 2067
	ProcessorResponseCodePayPalBusinessAccountLockedOrClosed                                                            ProcessorResponseCode = 2068
	ProcessorResponseCodePayPalBlockingDuplicateOrderIDs                                                                ProcessorResponseCode = 2069
	ProcessorResponseCodePayPalBuyerRevokedPreApprovedPaymentAuthorization                                              ProcessorResponseCode = 2070
	ProcessorResponseCodePayPalPayeeAccountInvalidOrDoesNotHaveConfirmedEmail                                           ProcessorResponseCode = 2071
	ProcessorResponseCodePayPalPayeeEmailIncorrectlyFormatted                                                           ProcessorResponseCode = 2072
	ProcessorResponseCodePayPalValidationError                                                                          ProcessorResponseCode = 2073
	ProcessorResponseCodeFundingInstrumentInThePayPalAccountWasDeclinedByTheProcessorOrBankOrItCantBeUsedForThisPayment ProcessorResponseCode = 2074
	ProcessorResponseCodePayerAccountIsLockedOrClosed                                                                   ProcessorResponseCode = 2075
	ProcessorResponseCodePayerCannotPayForThisTransactionWithPayPal                                                     ProcessorResponseCode = 2076
	ProcessorResponseCodeTransactionRefusedDueToPayPalRiskModel                                                         ProcessorResponseCode = 2077
	ProcessorResponseCodePayPalMerchantAccountConfigurationError                                                        ProcessorResponseCode = 2079
	ProcessorResponseCodePayPalPendingPaymentsAreNotSupported                                                           ProcessorResponseCode = 2081
	ProcessorResponseCodePayPalDomesticTransactionRequired                                                              ProcessorResponseCode = 2082
	ProcessorResponseCodePayPalPhoneNumberRequired                                                                      ProcessorResponseCode = 2083
	ProcessorResponseCodePayPalTaxInfoRequired                                                                          ProcessorResponseCode = 2084
	ProcessorResponseCodePayPalPayeeBlockedTransaction                                                                  ProcessorResponseCode = 2085
	ProcessorResponseCodePayPalTransactionLimitExceeded                                                                 ProcessorResponseCode = 2086
	ProcessorResponseCodePayPalReferenceTransactionsNotEnabledForYourAccount                                            ProcessorResponseCode = 2087
	ProcessorResponseCodeCurrencyNotEnabledForYourPayPalSellerAccount                                                   ProcessorResponseCode = 2088
	ProcessorResponseCodePayPalPayeeEmailPermissionDeniedForThisRequest                                                 ProcessorResponseCode = 2089
	ProcessorResponseCodePayPalAccountNotConfiguredToRefundMoreThanSettledAmount                                        ProcessorResponseCode = 2090
	ProcessorResponseCodeCurrencyOfThisTransactionMustMatchCurrencyOfYourPayPalAccount                                  ProcessorResponseCode = 2091
	ProcessorResponseCodeNoDataFoundTryAnotherVerificationMethod                                                        ProcessorResponseCode = 2092
	ProcessorResponseCodePayPalPaymentMethodIsInvalid                                                                   ProcessorResponseCode = 2093
	ProcessorResponseCodePayPalPaymentHasAlreadyBeenCompleted                                                           ProcessorResponseCode = 2094
	ProcessorResponseCodePayPalRefundIsNotAllowedAfterPartialRefund                                                     ProcessorResponseCode = 2095
	ProcessorResponseCodePayPalBuyerAccountCantBeSameAsSellerAccount                                                    ProcessorResponseCode = 2096
	ProcessorResponseCodePayPalAuthorizationAmountLimitExceeded                                                         ProcessorResponseCode = 2097
	ProcessorResponseCodePayPalAuthorizationCountLimitExceeded                                                          ProcessorResponseCode = 2098
	ProcessorResponseCodeCardholderAuthenticationRequired                                                               ProcessorResponseCode = 2099
	ProcessorResponseCodePayPalChannelInitiatedBillingNotEnabledForYourAccount                                          ProcessorResponseCode = 2100
	ProcessorResponseCodeProcessorNetworkUnavailableTryAgain                                                            ProcessorResponseCode = 3000
)

func (rc ProcessorResponseCode) Int() int {
	return int(rc)
}

// UnmarshalText fills the response code with the integer value if the text contains one in string form. If the text is zero length, the response code's value is unchanged but unmarshaling is successful.
func (rc *ProcessorResponseCode) UnmarshalText(text []byte) error {
	if len(text) == 0 {
		return nil
	}

	n, err := strconv.Atoi(string(text))
	if err != nil {
		return err
	}

	*rc = ProcessorResponseCode(n)

	return nil
}

// MarshalText returns a string in bytes of the number, or nil in the case it is zero.
func (rc ProcessorResponseCode) MarshalText() ([]byte, error) {
	if rc == 0 {
		return nil, nil
	}
	return []byte(strconv.Itoa(int(rc))), nil
}
