package gocurrency

import "testing"
import "log"

func TestCurrencyRate(t *testing.T) {
	f, e := CurrencyRate("TWD", "CNY")
	if e != nil {
		t.Fatal(e)
		return
	}

	t.Error(f)
}

func TestUnmarshal(t *testing.T) {
	s := `{
  "disclaimer": "Usage subject to terms: https://openexchangerates.org/terms",
  "license": "https://openexchangerates.org/license",
  "timestamp": 1509588000,
  "base": "USD",
  "rates": {
    "AED": 3.673097,
    "AFN": 68.6245,
    "ALL": 114.718761,
    "AMD": 482.586304,
    "ANG": 1.785544,
    "AOA": 165.9225,
    "ARS": 17.6055,
    "AUD": 1.296702,
    "AWG": 1.790501,
    "AZN": 1.6985,
    "BAM": 1.681092,
    "BBD": 2,
    "BDT": 83.195604,
    "BGN": 1.68304,
    "BHD": 0.37731,
    "BIF": 1752.45,
    "BMD": 1,
    "BND": 1.361297,
    "BOB": 6.927284,
    "BRL": 3.267317,
    "BSD": 1,
    "BTC": 0.000145454449,
    "BTN": 64.780658,
    "BWP": 10.557528,
    "BYN": 1.975378,
    "BZD": 2.010703,
    "CAD": 1.28385,
    "CDF": 1562.881563,
    "CHF": 0.999218,
    "CLF": 0.02359,
    "CLP": 636.1,
    "CNH": 6.59056,
    "CNY": 6.5878,
    "COP": 3065,
    "CRC": 569.935,
    "CUC": 1,
    "CUP": 25.5,
    "CVE": 95.1,
    "CZK": 21.919868,
    "DJF": 178.97,
    "DKK": 6.38178,
    "DOP": 47.861625,
    "DZD": 115.095,
    "EGP": 17.648,
    "ERN": 15.191062,
    "ETB": 27.286919,
    "EUR": 0.857576,
    "FJD": 2.076448,
    "FKP": 0.752286,
    "GBP": 0.752286,
    "GEL": 2.545329,
    "GGP": 0.752286,
    "GHS": 4.401408,
    "GIP": 0.752286,
    "GMD": 47.43,
    "GNF": 9028.05,
    "GTQ": 7.354849,
    "GYD": 207.840851,
    "HKD": 7.80155,
    "HNL": 23.632113,
    "HRK": 6.451,
    "HTG": 63.438864,
    "HUF": 266.459,
    "IDR": 13536.080047,
    "ILS": 3.50997,
    "IMP": 0.752286,
    "INR": 64.507,
    "IQD": 1165.409957,
    "IRR": 34576.946723,
    "ISK": 106,
    "JEP": 0.752286,
    "JMD": 127.147067,
    "JOD": 0.709001,
    "JPY": 113.8355,
    "KES": 103.675219,
    "KGS": 68.735479,
    "KHR": 4040.35,
    "KMF": 423.104935,
    "KPW": 900,
    "KRW": 1112.54,
    "KWD": 0.302199,
    "KYD": 0.833577,
    "KZT": 335.522793,
    "LAK": 8302.75,
    "LBP": 1516.45,
    "LKR": 153.767344,
    "LRD": 119.354136,
    "LSL": 14.145022,
    "LYD": 1.368013,
    "MAD": 9.492378,
    "MDL": 17.456,
    "MGA": 3173.55,
    "MKD": 52.805598,
    "MMK": 1354.9,
    "MNT": 2451.969147,
    "MOP": 8.039064,
    "MRO": 355.302795,
    "MUR": 34.2485,
    "MVR": 15.409873,
    "MWK": 726.664645,
    "MXN": 19.05825,
    "MYR": 4.227438,
    "MZN": 60.904857,
    "NAD": 14.145022,
    "NGN": 359.730849,
    "NIO": 30.784803,
    "NOK": 8.123,
    "NPR": 103.377151,
    "NZD": 1.443139,
    "OMR": 0.384856,
    "PAB": 1,
    "PEN": 3.249011,
    "PGK": 3.219138,
    "PHP": 51.495,
    "PKR": 105.335922,
    "PLN": 3.63233,
    "PYG": 5638.55,
    "QAR": 3.806363,
    "RON": 3.950701,
    "RSD": 101.985,
    "RUB": 58.24,
    "RWF": 854.315,
    "SAR": 3.75025,
    "SBD": 7.800511,
    "SCR": 13.410431,
    "SDG": 6.678746,
    "SEK": 8.376607,
    "SGD": 1.358047,
    "SHP": 0.752286,
    "SLL": 7643.328121,
    "SOS": 578.561731,
    "SRD": 7.448,
    "SSP": 130.2634,
    "STD": 20903.500286,
    "SVC": 8.752562,
    "SYP": 514.98999,
    "SZL": 14.142681,
    "THB": 33.079,
    "TJS": 8.81295,
    "TMT": 3.499986,
    "TND": 2.491311,
    "TOP": 2.268082,
    "TRY": 3.8102,
    "TTD": 6.732185,
    "TWD": 30.152,
    "TZS": 2248.9,
    "UAH": 26.888392,
    "UGX": 3653.2,
    "USD": 1,
    "UYU": 29.175769,
    "UZS": 8080.75,
    "VEF": 10.78085,
    "VND": 22711.919974,
    "VUV": 105.939999,
    "WST": 2.532807,
    "XAF": 562.533234,
    "XAG": 0.0581819,
    "XAU": 0.00078098,
    "XCD": 2.70255,
    "XDR": 0.711691,
    "XOF": 562.533234,
    "XPD": 0.00099308,
    "XPF": 102.336084,
    "XPT": 0.00107011,
    "YER": 250.306642,
    "ZAR": 14.006517,
    "ZMW": 10.028208,
    "ZWL": 322.355011
  }
}`
	log.Println(s)
	v, err := unmarshal([]byte(s))
	t.Error(v, err)
}

func TestCompute(t *testing.T) {
	v, _ := compute(6.5991, 30.1905)
	t.Error(v)
}
