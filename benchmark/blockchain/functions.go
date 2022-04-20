package blockchain

const (
	FuncPing = "ping"

	FuncCreateUser     = "CreateUser"
	FuncCreateReviewer = "CreateReviewer"
	FuncCreateAdmin    = "CreateAdmin"
	FuncCreatePaper    = "CreatePaper"

	FuncUpdateUserByEmail                  = "UpdateUserByEmail"
	FuncUpdateUserName                     = "UpdateUserName"
	FuncUpdateUserPassword                 = "UpdateUserPassword"
	FuncUpdateUserIsReviewer               = "UpdateUserIsReviewer"
	FuncUpdateUserIsNotReviewer            = "UpdateUserIsNotReviewer"
	FuncUpdateUserIsAdmin                  = "UpdateUserIsAdmin"
	FuncUpdateUserIsNotAdmin               = "UpdateUserIsNotAdmin"
	FuncUpdateUserReviewing                = "UpdateUserReviewing"
	FuncUpdatePaperById                    = "UpdatePaperById"
	FuncUpdatePeerReviewByPaperAndReviewer = "UpdatePeerReviewByPaperAndReviewer"

	FuncRetrieveUsersByQuery                                   = "RetrieveUsersByQuery"
	FuncRetrieveUsers                                          = "RetrieveUsers"
	FuncRetrieveUsersSortByEmail                               = "RetrieveUsersSortByEmail"
	FuncRetrieveUsersSortByName                                = "RetrieveUsersSortByName"
	FuncRetrieveUsersByNameSortByEmail                         = "RetrieveUsersByNameSortByEmail"
	FuncRetrieveReviewersSortByEmail                           = "RetrieveReviewersSortByEmail"
	FuncRetrieveReviewersSortByName                            = "RetrieveReviewersSortByName"
	FuncRetrieveReviewersByPaperIdSortByEmail                  = "RetrieveReviewersByPaperIdSortByEmail"
	FuncRetrieveReviewersByPaperIdSortByName                   = "RetrieveReviewersByPaperIdSortByName"
	FuncRetrieveAdminsSortByEmail                              = "RetrieveAdminsSortByEmail"
	FuncRetrieveAdminsSortByName                               = "RetrieveAdminsSortByName"
	FuncRetrieveUserByEmail                                    = "RetrieveUserByEmail"
	FuncRetrievePapersByQuery                                  = "RetrievePapersByQuery"
	FuncRetrievePapers                                         = "RetrievePapers"
	FuncRetrievePapersSortByTitle                              = "RetrievePapersSortByTitle"
	FuncRetrievePapersSortByUploadTime                         = "RetrievePapersSortByUploadTime"
	FuncRetrieveAcceptedPapersSortByTitle                      = "RetrieveAcceptedPapersSortByTitle"
	FuncRetrieveAcceptedPapersSortByUploadTime                 = "RetrieveAcceptedPapersSortByUploadTime"
	FuncRetrieveAcceptedPapersSortByPublishTime                = "RetrieveAcceptedPapersSortByPublishTime"
	FuncRetrieveRejectedPapersSortByTitle                      = "RetrieveRejectedPapersSortByTitle"
	FuncRetrieveRejectedPapersSortByUploadTime                 = "RetrieveRejectedPapersSortByUploadTime"
	FuncRetrieveReviewingPapersSortByTitle                     = "RetrieveReviewingPapersSortByTitle"
	FuncRetrieveReviewingPapersSortByUploadTime                = "RetrieveReviewingPapersSortByUploadTime"
	FuncRetrieveAcceptedPapersByTitleSortByTitle               = "RetrieveAcceptedPapersByTitleSortByTitle"
	FuncRetrieveAcceptedPapersByTitleSortByPublishTime         = "RetrieveAcceptedPapersByTitleSortByPublishTime"
	FuncRetrieveAcceptedPapersByAuthorSortByTitle              = "RetrieveAcceptedPapersByAuthorSortByTitle"
	FuncRetrieveAcceptedPapersByAuthorSortByPublishTime        = "RetrieveAcceptedPapersByAuthorSortByPublishTime"
	FuncRetrieveAcceptedPapersByKeywordSortByTitle             = "RetrieveAcceptedPapersByKeywordSortByTitle"
	FuncRetrieveAcceptedPapersByKeywordSortByPublishTime       = "RetrieveAcceptedPapersByKeywordSortByPublishTime"
	FuncRetrieveAcceptedPapersByUploaderSortByTitle            = "RetrieveAcceptedPapersByUploaderSortByTitle"
	FuncRetrieveAcceptedPapersByUploaderSortByUploadTime       = "RetrieveAcceptedPapersByUploaderSortByUploadTime"
	FuncRetrieveAcceptedPapersByUploaderSortByPublishTime      = "RetrieveAcceptedPapersByUploaderSortByPublishTime"
	FuncRetrieveRejectedPapersByUploaderSortByTitle            = "RetrieveRejectedPapersByUploaderSortByTitle"
	FuncRetrieveRejectedPapersByUploaderSortByUploadTime       = "RetrieveRejectedPapersByUploaderSortByUploadTime"
	FuncRetrieveRejectedPapersByUploaderSortByPublishTime      = "RetrieveRejectedPapersByUploaderSortByPublishTime"
	FuncRetrieveReviewingPapersByUploaderSortByTitle           = "RetrieveReviewingPapersByUploaderSortByTitle"
	FuncRetrieveReviewingPapersByUploaderSortByUploadTime      = "RetrieveReviewingPapersByUploaderSortByUploadTime"
	FuncRetrievePaperById                                      = "RetrievePaperById"
	FuncRetrievePeerReviewsByQuery                             = "RetrievePeerReviewsByQuery"
	FuncRetrievePeerReviewsByReviewerSortByCreateTime          = "RetrievePeerReviewsByReviewerSortByCreateTime"
	FuncRetrieveAcceptedPeerReviewsByReviewerSortByCreateTime  = "RetrieveAcceptedPeerReviewsByReviewerSortByCreateTime"
	FuncRetrieveRejectedPeerReviewsByReviewerSortByCreateTime  = "RetrieveRejectedPeerReviewsByReviewerSortByCreateTime"
	FuncRetrieveReviewingPeerReviewsByReviewerSortByCreateTime = "RetrieveReviewingPeerReviewsByReviewerSortByCreateTime"
)
