package route

import (
	"chaincode/lib"
	"chaincode/util"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
	"time"
)

func CreatePaper(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 5 {
		return shim.Error("function CreatePaper requires 5 arguments")
	}
	if args[0] == "" || args[1] == "" || args[2] == "" || args[3] == "" || args[4] == "" {
		return shim.Error("arguments should be nonempty")
	}
	email := args[0]
	title := args[1]
	abstract := args[2]
	authorsJSON := args[3]
	keywordsJSON := args[4]
	var authors []string
	err := json.Unmarshal([]byte(authorsJSON), &authors)
	if err != nil {
		return shim.Error(fmt.Sprintf("authors not in JSON: %s", err.Error()))
	}
	var keywords []string
	err = json.Unmarshal([]byte(keywordsJSON), &keywords)
	if err != nil {
		return shim.Error(fmt.Sprintf("keywords not in JSON: %s", err.Error()))
	}
	userResp := RetrieveUserByEmail(stub, []string{email})
	if userResp.Status != shim.OK {
		return shim.Error(fmt.Sprintf("cannot verify email: %s", userResp.Message))
	}
	if userResp.Payload == nil {
		return shim.Error(fmt.Sprintf("uploader %s not exists", email))
	}
	reviewerResp := retrieveAllSortedReviewers(stub, nil)
	if reviewerResp.Status != shim.OK {
		return shim.Error(fmt.Sprintf("failed to distribute reviewers: %s", reviewerResp.Message))
	}
	var reviewers []lib.User
	_ = json.Unmarshal(reviewerResp.Payload, &reviewers)
	if len(reviewers) < 3 {
		return shim.Error("not enough reviewers to distribute")
	}
	var reviewerEmails [3]string
	for i := 0; i < 3; i++ {
		reviewerEmails[i] = reviewers[i].Email
		_ = updateUserReviewingAdd(stub, []string{reviewers[i].Email})
	}
	now := time.Now().UnixNano()
	id := fmt.Sprintf("%x", sha256.Sum256([]byte(fmt.Sprintf("%s%d", email, now))))
	paper := lib.Paper{
		ID:          id,
		Uploader:    email,
		UploadTime:  now,
		Title:       title,
		Abstract:    abstract,
		Authors:     authors,
		Keywords:    keywords,
		PeerReviews: reviewerEmails,
		Status:      lib.StatusReviewing,
	}
	payload, err := util.PutLedger(stub, paper)
	if err != nil {
		return shim.Error(fmt.Sprintf("failed to put paper to ledger: %s", err.Error()))
	}
	return shim.Success(payload)
}