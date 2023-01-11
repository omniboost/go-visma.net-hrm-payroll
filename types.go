package vismanet

type Employees []Employee

type Employee struct {
	SalutationID               string `csv:"salutationid"`
	Salutation                 string `csv:"salutation"`
	Initials                   string `csv:"initials"`
	FirstName                  string `csv:"firstname"`
	Nickname                   string `csv:"nickname"`
	NameUsageEn                string `csv:"nameusage_en"`
	NameUsage                  string `csv:"nameusage"`
	Namecode                   string `csv:"namecode"`
	BirthName                  string `csv:"birthname"`
	PrefixBirthName            string `csv:"prefixbirthname"`
	FormattedName              string `csv:"formattedname"`
	DateOfBirth                Date   `csv:"dateofbirth"`
	Gender                     string `csv:"gender"`
	NationalityEn              string `csv:"nationality_en"`
	Nationality                string `csv:"nationality"`
	CorrAddressEn              string `csv:"corraddress_en"`
	CorrAddress                string `csv:"corraddress"`
	TitleBeforeEn              string `csv:"titlebefore_en"`
	TitleBefore                string `csv:"titlebefore"`
	TitleAfterEn               string `csv:"titleafter_en"`
	TitleAfter                 string `csv:"titleafter"`
	LanguageEn                 string `csv:"language_en"`
	Language                   string `csv:"language"`
	PartnerName                string `csv:"partnername"`
	PrefixPartnerName          string `csv:"prefixpartnername"`
	MaritalStatusID            string `csv:"maritalstatusid"`
	MaritalStatus              string `csv:"maritalstatus"`
	OfficialDate               Date   `csv:"officialdate"`
	EmployeeGroupID            string `csv:"employeegroupid"`
	EmployeeGroupNameEn        string `csv:"employeegroupname_en"`
	EmployeeGroupName          string `csv:"employeegroupname"`
	IsManager                  string `csv:"ismanager"`
	CompanyID                  string `csv:"companyid"`
	CompanyName                string `csv:"companyname"`
	BoardNumber                string `csv:"boardnumber"`
	SicknessStreet             string `csv:"sicknessstreet"`
	SicknessHouseNumber        string `csv:"sicknesshousenumber"`
	SicknessHouseNumberExt     string `csv:"sicknesshousenumberext"`
	SicknessZipcode            string `csv:"sicknesszipcode"`
	SicknessCity               string `csv:"sicknesscity"`
	SicknessCountryEn          string `csv:"sicknesscountry_en"`
	SicknessCountry            string `csv:"sicknesscountry"`
	PostalStreet               string `csv:"postalstreet"`
	PostalHouseNumber          string `csv:"postalhousenumber"`
	PostalHouseNumberExt       string `csv:"postalhousenumberext"`
	PostalZipCode              string `csv:"postalzipcode"`
	PostalCity                 string `csv:"postalcity"`
	PostalCountryEn            string `csv:"postalcountry_en"`
	PostalCountry              string `csv:"postalcountry"`
	HomeStreet                 string `csv:"homestreet"`
	HomeHouseNumber            string `csv:"homehousenumber"`
	HomeHouseNumberext         string `csv:"homehousenumberext"`
	HomeZipCode                string `csv:"homezipcode"`
	HomeCity                   string `csv:"homecity"`
	HomeCountryEn              string `csv:"homecountry_en"`
	HomeCountry                string `csv:"homecountry"`
	OtherAddressesEn           string `csv:"otheraddresses_en"`
	OtherAddresses             string `csv:"otheraddresses"`
	HomeEmailAddress           string `csv:"homeemailaddress"`
	BusinessEmailAddress       string `csv:"businessemailaddress"`
	HomewebEmailAddress        string `csv:"homewebemailaddress"`
	Businesswebemailaddress    string `csv:"businesswebemailaddress"`
	OtherEmailAddressesEn      string `csv:"otheremailaddresses_en"`
	OtherEmailaddresses        string `csv:"otheremailaddresses"`
	HomePhoneNumber            string `csv:"homephonenumber"`
	OfficeMobilePhoneNumber    string `csv:"officemobilephonenumber"`
	HomeFaxPhoneNumber         string `csv:"homefaxphonenumber"`
	PrivateMobilePhoneNumber   string `csv:"privatemobilephonenumber"`
	OfficetelPhoneNumber       string `csv:"officetelphonenumber"`
	OfficeextPhoneNumber       string `csv:"officeextphonenumber"`
	OfficefaxPhoneNumber       string `csv:"officefaxphonenumber"`
	SecretaryPhoneNumber       string `csv:"secretaryphonenumber"`
	SicknessAddressPhoneNumber string `csv:"sicknessaddressphonenumber"`
	EmergencyPhoneNumber       string `csv:"emergencyphonenumber"`
	OtherPhoneNumbersEn        string `csv:"otherphonenumbers_en"`
	OtherPhoneNumbers          string `csv:"otherphonenumbers"`
	ODPOrgID                   string `csv:"odporgid"`
	EmployeeID                 string `csv:"employeeid"`
	ExistenceStartUTC          string `csv:"existencestartutc"`
}
