package modules

type MobilePushServiceLinks struct {
	UserService struct {
		GetCouldIdList string
	}
}

type SmsServiceLinks struct {
	UserService struct {
		GetPhoneNumberList string
	}
}

type EmailServiceLinks struct {
	UserService struct {
		GetEmailList string
	}
}

type UserServiceLinks struct {
	SmsService struct {
		SendSms string
	}
}

type AnyBackendServiceLinks struct {
	ConfigService struct {
		UpdateRemoteConfig string
	}
}

type MdmAdapterServiceLinks struct {
	MdmService struct {
		HandleRecord      string
		HandleRecordBatch string
	}
}

type MdmServiceLinks struct {
	MdmNotifierService struct {
		BroadcastNotification string
	}
}

type MdmNotifyLinks struct {
	MdmConverterService struct {
		ConvertToSudirBatchList string
		ConvertToFindBatchList  string
		ConvertAnyBatchList     string
		ConvertErlBatchList     string
	}
}

type MdmApiServiceLinks struct {
	MdmService struct {
		GetRecord           string
		GetAttributes       string
		GetRecordsList      string
		GetByExternalIdList string
	}
	MdmConverterService struct {
		ConvertToSudir               string
		ConvertToFind                string
		ConvertSearchRequestForSudir string
		ConvertSearchRequestForFind  string

		FilterAttributes    string
		FilterSearchRequest string
	}
	MdmSearchService struct {
		Search       string
		SearchIdList string
	}
	MdmAsyncApiService struct {
		CreateJob              string
		GetJobStatus           string
		GetInternalStoreResult string
	}
}

type MdmDumperService struct {
	MdmConverterService struct {
		ConvertToSudirBatchList    string
		ConvertToPdpSudirBatchList string
		ConvertToFindBatchList     string
		ConvertAnyBatchList        string
		ConvertErlBatchList        string
		FilterBatchList            string
	}
	MdmSearchService struct {
		Count string
	}
}

type MdmAsyncApiService struct {
	MdmSearchService struct {
		SearchIdWithScroll   string
		PreferredSlicesCount string
	}
}
