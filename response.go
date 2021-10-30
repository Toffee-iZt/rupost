package rupost

import "time"

// Error struct.
type Error struct {
	Code        int    `json:"code"`
	Description string `json:"description"`
}

// Response struct.
type Response struct {
	AttorneyServiceIsEnabled     bool        `json:"attorneyServiceIsEnabled"`
	UserTrackingItemID           interface{} `json:"userTrackingItemId"`
	UserTitle                    interface{} `json:"userTitle"`
	ItemAddedDate                interface{} `json:"itemAddedDate"`
	DeleteDate                   interface{} `json:"deleteDate"`
	LastOperationViewed          bool        `json:"lastOperationViewed"`
	Deleted                      bool        `json:"deleted"`
	AutoAdded                    bool        `json:"autoAdded"`
	LastOperationViewedTimestamp interface{} `json:"lastOperationViewedTimestamp"`
	Payable                      bool        `json:"payable"`
	PaymentStatus                interface{} `json:"paymentStatus"`
	PaymentSystem                interface{} `json:"paymentSystem"`
	DelayInsuranceInfo           struct {
		PaymentStatus  string      `json:"paymentStatus"`
		ProductInfo    interface{} `json:"productInfo"`
		DescriptionURL interface{} `json:"descriptionUrl"`
	} `json:"delayInsuranceInfo"`
	RefusalInsuranceInfo struct {
		PaymentStatus  string      `json:"paymentStatus"`
		ProductInfo    interface{} `json:"productInfo"`
		DescriptionURL interface{} `json:"descriptionUrl"`
	} `json:"refusalInsuranceInfo"`
	PaymentDate  interface{} `json:"paymentDate"`
	ParcelStatus interface{} `json:"parcelStatus"`
	Euv          bool        `json:"euv"`
	EuvStatus    string      `json:"euvStatus"`
	Amount       interface{} `json:"amount"`
	Formless     bool        `json:"formless"`
	TrackingItem struct {
		LastOperationDateTime                int64         `json:"lastOperationDateTime"`
		AcceptanceOperationDateTime          int64         `json:"acceptanceOperationDateTime"`
		TpoCustomPayment                     interface{}   `json:"tpoCustomPayment"`
		CustomsPayments                      []interface{} `json:"customsPayments"`
		DestinationCountryName               string        `json:"destinationCountryName"`
		DestinationCountryNameGenitiveCase   string        `json:"destinationCountryNameGenitiveCase"`
		DestinationCityName                  string        `json:"destinationCityName"`
		OriginCountryName                    string        `json:"originCountryName"`
		OriginCityName                       interface{}   `json:"originCityName"`
		MailRank                             interface{}   `json:"mailRank"`
		MailCtg                              int           `json:"mailCtg"`
		PostMark                             interface{}   `json:"postMark"`
		Insurance                            interface{}   `json:"insurance"`
		InsuranceMoney                       interface{}   `json:"insuranceMoney"`
		IsDestinationInInternationalTracking bool          `json:"isDestinationInInternationalTracking"`
		IsOriginInInternationalTracking      bool          `json:"isOriginInInternationalTracking"`
		HiddenHistoryList                    []interface{} `json:"hiddenHistoryList"`
		FuturePathList                       []struct {
			Date                      interface{} `json:"date"`
			HumanStatus               string      `json:"humanStatus"`
			HumanOperationStatus      interface{} `json:"humanOperationStatus"`
			OperationType             int         `json:"operationType"`
			OperationAttr             int         `json:"operationAttr"`
			CountryID                 int         `json:"countryId"`
			Index                     interface{} `json:"index"`
			CityName                  interface{} `json:"cityName"`
			CountryName               string      `json:"countryName"`
			CountryNameGenitiveCase   string      `json:"countryNameGenitiveCase"`
			CountryCustomName         string      `json:"countryCustomName"`
			IsInInternationalTracking interface{} `json:"isInInternationalTracking"`
			Description               interface{} `json:"description"`
			Weight                    interface{} `json:"weight"`
		} `json:"futurePathList"`
		CashOnDeliveryEventsList []interface{} `json:"cashOnDeliveryEventsList"`
		ShipmentTripInfo         struct {
			Acceptance struct {
				Date          time.Time   `json:"date"`
				OperationType int         `json:"operationType"`
				OperationAttr int         `json:"operationAttr"`
				Index         string      `json:"index"`
				IndexTo       interface{} `json:"indexTo"`
			} `json:"acceptance"`
			CustomsPassing                  interface{} `json:"customsPassing"`
			Arrived                         interface{} `json:"arrived"`
			HiddenInternalHistoryRecord     interface{} `json:"hiddenInternalHistoryRecord"`
			ExpectedDeliveryDays            interface{} `json:"expectedDeliveryDays"`
			StartDeliveryDate               interface{} `json:"startDeliveryDate"`
			ExpectedDeliveryDate            interface{} `json:"expectedDeliveryDate"`
			IsExpectedDeliveryDateAvailable bool        `json:"isExpectedDeliveryDateAvailable"`
		} `json:"shipmentTripInfo"`
		Sender                  string      `json:"sender"`
		Recipient               string      `json:"recipient"`
		Weight                  int         `json:"weight"`
		StorageTime             int         `json:"storageTime"`
		Title                   string      `json:"title"`
		LiferayWebContentID     interface{} `json:"liferayWebContentId"`
		TrackingHistoryItemList []struct {
			Date                      time.Time   `json:"date"`
			HumanStatus               string      `json:"humanStatus"`
			HumanOperationStatus      interface{} `json:"humanOperationStatus"`
			OperationType             int         `json:"operationType"`
			OperationAttr             int         `json:"operationAttr"`
			CountryID                 int         `json:"countryId"`
			Index                     string      `json:"index"`
			CityName                  interface{} `json:"cityName"`
			CountryName               string      `json:"countryName"`
			CountryNameGenitiveCase   string      `json:"countryNameGenitiveCase"`
			CountryCustomName         string      `json:"countryCustomName"`
			IsInInternationalTracking bool        `json:"isInInternationalTracking"`
			Description               string      `json:"description"`
			Weight                    int         `json:"weight"`
		} `json:"trackingHistoryItemList"`
		GlobalStatus                 string      `json:"globalStatus"`
		MailType                     string      `json:"mailType"`
		MailTypeCode                 int         `json:"mailTypeCode"`
		CountryFromCode              int         `json:"countryFromCode"`
		CountryToCode                int         `json:"countryToCode"`
		CustomDuty                   interface{} `json:"customDuty"`
		CashOnDelivery               interface{} `json:"cashOnDelivery"`
		IndexFrom                    interface{} `json:"indexFrom"`
		IndexTo                      string      `json:"indexTo"`
		CanBeOrdered                 bool        `json:"canBeOrdered"`
		CanBePickedUp                bool        `json:"canBePickedUp"`
		DeliveryOrderDate            interface{} `json:"deliveryOrderDate"`
		CommonStatus                 string      `json:"commonStatus"`
		FirstOperationDate           int64       `json:"firstOperationDate"`
		LastOperationDate            int64       `json:"lastOperationDate"`
		LastOperationTimezoneOffset  int         `json:"lastOperationTimezoneOffset"`
		ComplexCode                  interface{} `json:"complexCode"`
		ComplexType                  interface{} `json:"complexType"`
		ComplexDeliveryMethod        interface{} `json:"complexDeliveryMethod"`
		Barcode                      string      `json:"barcode"`
		BarcodeImage                 interface{} `json:"barcodeImage"`
		NotificationBarcode          interface{} `json:"notificationBarcode"`
		NotificationTitle            interface{} `json:"notificationTitle"`
		SourceBarcode                interface{} `json:"sourceBarcode"`
		SourceTitle                  interface{} `json:"sourceTitle"`
		EndStorageDate               interface{} `json:"endStorageDate"`
		LastDayInOps                 interface{} `json:"lastDayInOps"`
		HasBeenGiven                 interface{} `json:"hasBeenGiven"`
		LastOperationAttr            int         `json:"lastOperationAttr"`
		LastOperationType            int         `json:"lastOperationType"`
		ID                           interface{} `json:"id"`
		PostmarkText                 interface{} `json:"postmarkText"`
		MailTypeText                 string      `json:"mailTypeText"`
		CustomsPaymentStatus         interface{} `json:"customsPaymentStatus"`
		CustomsOperatorDuty          interface{} `json:"customsOperatorDuty"`
		CustomsOperatorDutyFee       interface{} `json:"customsOperatorDutyFee"`
		CustomsOperatorDutyFeeOnline interface{} `json:"customsOperatorDutyFeeOnline"`
		CustomsPayment               interface{} `json:"customsPayment"`
		MailCtgText                  string      `json:"mailCtgText"`
		MailRankText                 interface{} `json:"mailRankText"`
		ReturnRate                   interface{} `json:"returnRate"`
		RedispatchRate               interface{} `json:"redispatchRate"`
	} `json:"trackingItem"`
	OfficeReservationURL string      `json:"officeReservationUrl"`
	AlienOrder           interface{} `json:"alienOrder"`
	OfficeSummary        interface{} `json:"officeSummary"`
	PostmanDeliveryInfo  interface{} `json:"postmanDeliveryInfo"`
	CourierDeliveryInfo  interface{} `json:"courierDeliveryInfo"`
	FormF22Params        struct {
		MailTypeText   string      `json:"MailTypeText"`
		SenderAddress  string      `json:"senderAddress"`
		RecipientIndex string      `json:"RecipientIndex"`
		WeightGr       int         `json:"WeightGr"`
		SendingType    string      `json:"SendingType"`
		Ordered        bool        `json:"Ordered"`
		MailCtgText    string      `json:"MailCtgText"`
		State          interface{} `json:"state"`
		PostID         string      `json:"PostId"`
		Other          bool        `json:"Other"`
	} `json:"formF22Params"`
}
