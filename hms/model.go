package hms

// InAppPurchaseData json model. Used when requesting In-App order and / or subscription verification.
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5
type InAppPurchaseData struct {
	// App ID.
	ApplicationID int64 `json:"applicationId"`

	// For consumables or non-consumables, the value is always false.
	// For subscriptions, possible values are:
	// true: A subscription is in active state and will be automatically renewed on the next renewal date.
	// false: A user has canceled the subscription. The user can access the subscribed content
	//        before the next renewal date but will be unable to access the content after the date
	//        unless the user enables automatic renewal. If a grace period is provided,
	//        this value remains true for the subscription as long as it is still in the grace period.
	//        The next settlement date is automatically extended every day until the grace period ends
	//        or the user changes the payment method.
	AutoRenewing bool `json:"autoRenewing"`

	// Order ID, which uniquely identifies a transaction and is generated by the Huawei IAP server during payment.
	OrderID string `json:"orderId"`

	// Product type. Possible values are:
	//    0: consumable
	//    1: non-consumable
	//    2: subscription
	Kind int64 `json:"kind"`

	// App package name.
	PackageName string `json:"packageName,omitempty"`

	// Product ID. Each product must have a unique ID, which is maintained in the PMS or passed when the app initiates a purchase.
	ProductID string `json:"productId"`

	// Product name.
	ProductName string `json:"productName,omitempty"`

	// Timestamp of the purchase time, which is the number of milliseconds from 00:00:00 on January 1, 1970 to the purchase time.
	PurchaseTime int64 `json:"purchaseTime,omitempty"`

	// Transaction status. Possible values are:
	//    -1: initialized
	//    0: purchased
	//    1: canceled
	//    2: refunded
	PurchaseState int64 `json:"purchaseState"`

	// Information stored on the merchant side, which is passed by the app when the payment API is called.
	DeveloperPayload string `json:"developerPayload,omitempty"`

	// Challenge defined when an app initiates a consumption request.
	// The challenge uniquely identifies the consumption request and exists only for one-off products.
	DeveloperChallenge string `json:"developerChallenge,omitempty"`

	// Consumption status, which exists only for one-off products. Possible values are:
	//    0: not consumed
	//    1: consumed
	ConsumptionState ConsumptionState `json:"consumptionState,omitempty"`

	// Purchase token, which uniquely identifies the mapping between a product and a user.
	// It is generated by the Huawei IAP server when the payment is complete.
	// NOTE:
	// * This parameter uniquely identifies the mapping between a product and a user.
	//   It does not change when the subscription is renewed.
	// * Currently, the value contains 92 characters and its length may be expanded.
	//   If the value needs to be stored, you are advised to reserve 128 characters.
	PurchaseToken string `json:"purchaseToken"`

	// Purchase type. Possible values are:
	//    0: in the sandbox
	//    1: in the promotion period (currently unsupported)
	// This parameter is not returned during formal purchase.
	// To avoid default value issues. check if PurchaseType != nil first, then read *PurchaseType.
	PurchaseType *int64 `json:"purchaseType,omitempty"`

	// Currency. The value must comply with the ISO 4217 standard.
	Currency string `json:"currency,omitempty"`

	// Value after the actual price of a product is multiplied by 100. The actual price is accurate to two decimal places.
	// For example, if the value of this parameter is 501, the actual product price is 5.01.
	Price int64 `json:"price,omitempty"`

	// Country or region code, which is used to identify a country or region. The value must comply with the ISO 3166 standard.
	Country string `json:"country,omitempty"`

	// Payment method. Possible values are:
	//    0: HUAWEI Point
	//    3: credit card
	//    4: Alipay
	//    6: carrier billing
	//    13: PayPal
	//    16: debit card
	//    17: WeChat Pay
	//    19: gift card
	//    20: balance
	//    21: HUAWEI Point card
	//    24: WorldPay
	//    31: Huawei Pay
	//    32: Ant Credit Pay
	//    200: M-Pesa
	PayType string `json:"payType,omitempty"`

	// Transaction order ID.
	PayOrderID string `json:"payOrderId,omitempty"`

	// Account type. Possible values are:
	//    0: HUAWEI ID
	//    1: AppTouch user account
	AccountFlag int64 `json:"accountFlag,omitempty"`

	// ===== The following parameters are returned only in the subscription scenario. =====

	// Order ID generated by the Huawei IAP server during fee deduction for the previous renewal.
	// The parameter value is the same as that of orderId when a subscription is purchased for the first time.
	LastOrderID string `json:"lastOrderId"`

	// ID of the subscription group to which a subscription belongs.
	ProductGroup string `json:"productGroup,omitempty"`

	// Timestamp of the purchase time, which is the number of milliseconds from 00:00:00 on January 1, 1970 to the purchase time.
	//
	// If the purchase is not complete, this parameter is left empty.
	// PurchaseTime int64 `json:"purchaseTime,omitempty"`

	// Timestamp of the first fee deduction time, which is the number of
	// milliseconds from 00:00:00 on January 1, 1970 to the first successful fee deduction time.
	OriPurchaseTime int64 `json:"oriPurchaseTime,omitempty"`

	// Subscription ID.
	//
	// NOTE: This parameter uniquely identifies the mapping between a product and a user. It does not change when the subscription is renewed.
	SubscriptionID string `json:"subscriptionId,omitempty"`

	// Original subscription ID. If this value is not empty, the current subscription is switched from another one.
	// It can be linked to the original subscription information.
	OriSubscriptionID string `json:"oriSubscriptionId,omitempty"`

	// Purchase quantity.
	Quantity int64 `json:"quantity,omitempty"`

	// Days of a paid subscription, excluding the free trial period and promotion period.
	DaysLasted int64 `json:"daysLasted,omitempty"`

	// Number of successful renewal periods. How many time the subscription has been renewed, excludes promotion.
	// If the value is 0 or does not exist, no renewal has been performed successfully.
	NumOfPeriods int64 `json:"numOfPeriods,omitempty"`

	// Number of successful renewal periods with promotion.
	NumOfDiscount int64 `json:"numOfDiscount,omitempty"`

	// Timestamp of the subscription expiration time. In milliseconds.
	//
	// For an automatic renewal receipt where the fee has been deducted successfully, this time indicates the renewal date or expiration date.
	// If the value is a past time, the subscription has expired.
	ExpirationDate int64 `json:"expirationDate,omitempty"`

	// Reason why a subscription expires. Possible values are:
	//    1: canceled by a user
	//    2: product being unavailable
	//    3: abnormal user signing information
	//    4: billing error
	//    5: price increase disagreed with by a user
	//    6: unknown error
	// If there are multiple exceptions, a smaller number indicates a higher priority (1 > 2 > 3...).
	ExpirationIntent int64 `json:"expirationIntent,omitempty"`

	// Indicates whether the system is still trying to renew an expired subscription. Possible values are:
	//    0: no
	//    1: yes
	RetryFlag int64 `json:"retryFlag,omitempty"`

	// Indicates whether a subscription is in a renewal period with promotion. Possible values are:
	//    0: no
	//    1: yes
	IntroductoryFlag int64 `json:"introductoryFlag,omitempty"`

	// Indicates whether a subscription is in a free trial period. Possible values are:
	//    0: no
	//    1: yes
	TrialFlag int64 `json:"trialFlag,omitempty"`

	// Time when a subscription is revoked. A refund is made and the service is unavailable immediately.
	// This value is returned when a user:
	//    a. makes a complaint and revokes a subscription through the customer service personnel; or
	//    b. performs subscription upgrade or crossgrade that immediately takes effect and revokes the previous subscription.
	// Note: If a receipt is revoked, it is deemed that the purchase is not complete.
	CancelTime int64 `json:"cancelTime,omitempty"`

	// Cause of subscription cancellation. Possible values are:
	//    0: other causes. For example, a user mistakenly purchases a subscription and has to cancel it.
	//    1: A user encounters a problem within the app and cancels the subscription.
	//    2: A user performs subscription upgrade or crossgrade.
	// Note: If this parameter is left empty but the cancelTime parameter has a value,
	// the cancellation is caused by an operation such as upgrade.
	CancelReason int64 `json:"cancelReason,omitempty"`

	// App information. This parameter is reserved.
	AppInfo string `json:"appInfo,omitempty"`

	// Indicates whether a user has disabled the subscription notification function. Possible values are:
	//    0: no
	//    1: yes
	// If the user disables the subscription notification function, no subscription notification will be sent to this user.
	NotifyClosed int64 `json:"notifyClosed,omitempty"`

	// Renewal status. Possible values are:
	//    1: The subscription renewal is normal.
	//    0: The user cancels subscription renewal.
	// For auto-renewable subscriptions, this parameter is valid for both valid and expired subscriptions. However, it does not represent users' subscription status. Generally, when the value is 0, the app can provide other subscription options for the user, for example, recommending another subscription with a lower level in the same group. The value 0 indicates that a user proactively cancels the subscription.
	RenewStatus int64 `json:"renewStatus,omitempty"`

	// User opinion on the price increase of a product. Possible values are:
	//    1: The user has agreed with the price increase.
	//    0: The user does not take any action. After the subscription expires, it becomes invalid.
	PriceConsentStatus int64 `json:"priceConsentStatus,omitempty"`

	// Price used upon the next renewal. It is provided as a reference for users when the priceConsentStatus parameter is returned.
	RenewPrice int64 `json:"renewPrice,omitempty"`

	// true: A user has been charged for a product, the product has not expired, and no refund has been made. In this case, you can provide services for the user.
	// false: The purchase of a product is not finished, the product has expired, or a refund has been made for the product after its purchase.
	//
	// NOTE
	// If a user has canceled a subscription, the subIsvalid parameter value is still true until the subscription expires.
	SubIsValid bool `json:"subIsvalid,omitempty"`

	// Indicates whether to postpone the settlement date.
	//    1: yes
	DeferFlag int64 `json:"deferFlag,omitempty"`

	// Subscription cancellation initiator. Possible values are:
	//    0: user
	//    1: developer
	//    2: Huawei
	CancelWay int64 `json:"cancelWay,omitempty"`

	// Timestamp (milliseconds in UTC) when you set a subscription renewal to be stopped in the future.
	// The subscription is still valid within the validity period, but the renewal will be stopped in the future. No refund is required.
	// NOTE:
	// cancelWay and cancellationTime are displayed when a subscription renewal stops (no refund is involved).
	CancellationTime int64 `json:"cancellationTime,omitempty"`

	// Number of days for retaining a subscription relationship after the subscription is canceled.
	CancelledSubKeepDays int64 `json:"cancelledSubKeepDays,omitempty"`

	// Confirmation status. Possible values are:
	//    0: not confirmed
	//    1: confirmed
	// If this parameter is left empty, no confirmation is required.
	Confirmed int64 `json:"confirmed,omitempty"`

	// Timestamp (milliseconds in UTC) when a paused subscription is resumed.
	ResumeTime int64 `json:"resumeTime,omitempty"`

	// Cancellation reason. Possible values are:
	//    0: others
	//    1: too high fee
	//    2: technical problem, for example, product not provided
	//    5: in the blocklist because of fraud
	//    7: subscription switchover
	//    9: service being rarely used and not required
	//    10: other better apps
	SurveyReason int64 `json:"surveyReason,omitempty"`

	// When the value of surveyReason is 0, this parameter is used to collect the cancellation reason entered by users.
	SurveyDetails string `json:"surveyDetails,omitempty"`

	// UTC timestamp (in milliseconds) when the grace period of a subscription ends.
	GraceExpirationTime int64 `json:"graceExpirationTime,omitempty"`
}

// CanceledPurchaseList response from query canceled or refunded purchase list
//
// Document: https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/api-cancel-or-refund-record-0000001050746117-V5
type CanceledPurchaseList struct {
	// Result code. Possible values are:
	//    0: success
	//    Other values: failure. For details about the result codes, please refer to https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-error-code-0000001050166248-V5
	ResponseCode string `json:"responseCode"`

	// Response description.
	ResponseMessage string `json:"responseMessage,omitempty"`

	// List of canceled or refunded purchase information, in JSON strings. Each string indicates a purchase record.
	//
	// For details about the purchase information format, please refer to CanceledPurchase{}
	CancelledPurchaseList string `json:"cancelledPurchaseList,omitempty"`

	// Token to query data on the next page. If a value is returned, pass it in the next query request to query data on the next page.
	ContinuationToken string `json:"continuationToken,omitempty"`
}

// CanceledPurchase individual canceled purchase information, for CanceledPurchaseList.
//
// Document: https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/api-cancel-or-refund-record-0000001050746117-V5
type CanceledPurchase struct {
	// Unique order ID of a subscription or subscription renewal.
	OrderID string `json:"orderId"`

	// Product ID.
	ProductID string `json:"productId"`

	// Purchase token.
	PurchaseToken string `json:"purchaseToken"`

	// Purchase timestamp in UTC.
	PurchaseTime int64 `json:"purchaseTime"`

	// Cancellation or refund timestamp in UTC.
	CancelledTime int64 `json:"cancelledTime"`

	// Cancellation or refund initiator. Possible values are:
	//    0: user
	//    1: developer
	//    2: Huawei
	CancelledSource int64 `json:"cancelledSource"`

	// Cancellation or refund reason. Possible values are:
	//    0: others
	//    1: user repentance
	//    2: product not provided
	//    3: abnormal app service
	//    4: accidental purchase
	//    5: fraud
	//    6: chargeback
	//    7: upgrade or downgrade
	//    8: user service area change
	CancelledReason int64 `json:"cancelledReason"`
}

// Below are all constants that can be used to determine different parameter states in IAP Purchase Data

// Constants for InAppPurchaseData.Kind product type
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5
const (
	InAppPurchaseDataKindConsumable    int64 = 0
	InAppPurchaseDataKindNonConsumable int64 = 1
	InAppPurchaseDataKindSubscription  int64 = 2
)

// Constants for InAppPurchaseData.PurchaseState
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5
const (
	InAppPurchaseDataPurchaseStateInitialized int64 = -1
	InAppPurchaseDataPurchaseStatePurchased   int64 = 0
	InAppPurchaseDataPurchaseStateCanceled    int64 = 1
	InAppPurchaseDataPurchaseStateRefunded    int64 = 2
)

// Constants for InAppPurchaseData.ConsumptionState
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5
type ConsumptionState int

const (
	InAppPurchaseDataConsumptionStateNotConsumed ConsumptionState = 0
	InAppPurchaseDataConsumptionStateConsumed    ConsumptionState = 1
)

// Constants for InAppPurchaseData.PurchaseType indicates if the product is purchased through sandbox environment or promotion.
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5
//
// Note: You should always check InAppPurchaseData.PurchaseType != nil first.
const (
	InAppPurchaseDataPurchaseTypeSandbox   int64 = 0
	InAppPurchaseDataPurchaseTypePromotion int64 = 1
)

// Constants for InAppPurchaseData.PayType payment methods when buying the product.
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5#EN-US_TOPIC_0000001050986133__section135412662210
const (
	InAppPurchaseDataPayTypeHuaweiPoint     string = "0"
	InAppPurchaseDataPayTypeCreditCard      string = "3"
	InAppPurchaseDataPayTypeAlipay          string = "4"
	InAppPurchaseDataPayTypeCarrier         string = "6"
	InAppPurchaseDataPayTypePayPal          string = "13"
	InAppPurchaseDataPayTypeDebitCard       string = "16"
	InAppPurchaseDataPayTypeWeChatPay       string = "17"
	InAppPurchaseDataPayTypeGiftCard        string = "19"
	InAppPurchaseDataPayTypeBalance         string = "20"
	InAppPurchaseDataPayTypeHuaweiPointCard string = "21"
	InAppPurchaseDataPayTypeWorldPay        string = "24"
	InAppPurchaseDataPayTypeHuaweiPay       string = "31"
	InAppPurchaseDataPayTypeAntCreditPay    string = "32"
	InAppPurchaseDataPayTypeMPesa           string = "200" // M-Pesa
)

// Constants for InAppPurchaseData.AccountFlag Account type.
// See https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/api-common-statement-0000001050986127-V5#EN-US_TOPIC_0000001050986127__section1741234185817 for details.
const (
	InAppPurchaseDataAccountFlagHuaweiID int64 = 0
	InAppPurchaseDataAccountFlagAppTouch int64 = 1
)

// Constants for InAppPurchaseData.ExpirationIntent Reasons why a subscription expires.
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5
const (
	InAppPurchaseDataExpirationIntentCanceledByUser         int64 = 1
	InAppPurchaseDataExpirationIntentProductUnavaliable     int64 = 2
	InAppPurchaseDataExpirationIntentAbnormalUserSigning    int64 = 3
	InAppPurchaseDataExpirationIntentBillingError           int64 = 4
	InAppPurchaseDataExpirationIntentPriceIncreaseDisagreed int64 = 5
	InAppPurchaseDataExpirationIntentUnknownError           int64 = 6
)

// Constants for InAppPurchaseData.RetryFlag Indicates whether the system still tries to renew an expired subscription.
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5
const (
	InAppPurchaseDataRetryFlagNo  int64 = 0
	InAppPurchaseDataRetryFlagYes int64 = 1
)

// Constants for InAppPurchaseData.IntroductoryFlag Indicates whether a subscription is in the renewal period with promotion or not.
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5
const (
	InAppPurchaseDataIntroductoryFlagNo  int64 = 0
	InAppPurchaseDataIntroductoryFlagYes int64 = 1
)

// Constants for InAppPurchaseData.TrialFlag Indicates whether a subscription is in the free trial period or not.
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5
const (
	InAppPurchaseDataTrialFlagNo  int64 = 0
	InAppPurchaseDataTrialFlagYes int64 = 1
)

// Constants for InAppPurchaseData.CancelReason Causes of subscription cancellation.
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5
//
// Note: You should check SubIsValid first.
const (
	InAppPurchaseDataCancelReasonOther               int64 = 0 // Other causes. For example, a user mistakenly purchases a subscription and has to cancel it.
	InAppPurchaseDataCancelReasonUserIssue           int64 = 1 // A user encounters a problem within the app and cancels the subscription.
	InAppPurchaseDataCancelReasonUpgradeOrCrossgrade int64 = 2 // A user performs subscription upgrade or crossgrade.
)

// Constants for InAppPurchaseData.NotifyClosed Indicates whether a user has disabled the subscription notification function.
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5
const (
	InAppPurchaseDataNotifyClosedNo  int64 = 0
	InAppPurchaseDataNotifyClosedYes int64 = 1
)

// Constants for InAppPurchaseData.RenewStatus Indecates whether the auto-renewal is canceled or not.
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5
const (
	InAppPurchaseDataRenewStatusCanceledByUser int64 = 0 // A user proactively canceled the subscription auto-renewal.
	InAppPurchaseDataRenewStatusNormal         int64 = 1 // The subscription will be auto-renewed.
)

// Constants for InAppPurchaseData.PriceConsentStatus User opinions collected when the product price increased.
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5
const (
	InAppPurchaseDataPriceConsentStatusNoResponse int64 = 0 // The user did not give any comfirmation. Subscription will be terminated without renewal after expires.
	InAppPurchaseDataPriceConsentStatusAgreed     int64 = 1 // The user has agreed with the price increase.
)

// Constants for InAppPurchaseData.DeferFlag Indicates whether to postpone the settlement date.
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5
const (
	InAppPurchaseDataDeferFlagYes int64 = 1
)

// Constants for InAppPurchaseData.CancelWay how does the subscription be canceled
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/server-data-model-0000001050986133-V5
//
// Note: You should check SubIsValid first.
const (
	InAppPurchaseDataCancelWayByUser      int64 = 0
	InAppPurchaseDataCancelWayByDeveloper int64 = 1
	InAppPurchaseDataCancelWayByHuawei    int64 = 2
)

// Constants for CanceledPurchase.CancelledSource
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/api-cancel-or-refund-record-0000001050746117-V5
const (
	CanceledPurchaseCancelledSourceByUser      int64 = 0
	CanceledPurchaseCancelledSourceByDeveloper int64 = 1
	CanceledPurchaseCancelledSourceByHuawei    int64 = 2
)

// Constants for CanceledPurchase.CancelledReason
// https://developer.huawei.com/consumer/en/doc/HMSCore-References-V5/api-cancel-or-refund-record-0000001050746117-V5
const (
	CanceledPurchaseCancelledReasonOther              int64 = 0
	CanceledPurchaseCancelledReasonUserRepentance     int64 = 1
	CanceledPurchaseCancelledReasonProductNotProvided int64 = 2
	CanceledPurchaseCancelledReasonAbnormal           int64 = 3
	CanceledPurchaseCancelledReasonAccidental         int64 = 4
	CanceledPurchaseCancelledReasonFraud              int64 = 5
	CanceledPurchaseCancelledReasonChargeback         int64 = 6
	CanceledPurchaseCancelledReasonUpgradeOrDowngrade int64 = 7
	CanceledPurchaseCancelledReasonServiceAreaChanged int64 = 8
)
