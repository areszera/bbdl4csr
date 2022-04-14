package utils

import "chaincode/lib"

func GetStatus(s1 string, s2 string, s3 string) string {
	switch lib.StatusCode(s1) + lib.StatusCode(s2) + lib.StatusCode(s3) {
	case 3 * lib.StatusCodeAccepted, 2*lib.StatusCodeAccepted + lib.StatusCodeRejected:
		return lib.StatusCodeAccepted.String()
	case 3 * lib.StatusCodeRejected, 2*lib.StatusCodeRejected + lib.StatusCodeAccepted:
		return lib.StatusCodeRejected.String()
	default:
		return lib.StatusCodeReviewing.String()
	}
}