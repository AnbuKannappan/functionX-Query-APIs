package handlers

import (
	"encoding/json"
	utils "functionX-Query-APIs/utils"
	"github.com/gin-gonic/gin"
	"math/big"
	"net/http"
	"os/exec"
	"strings"
)

type Master struct {
	Proposals []Proposals `json:"proposals"`
}

type Proposals struct {
	Content            Content            `json:"content"`
	Deposit_end_time   string             `json:"deposit_end_time"`
	Final_tally_result Final_tally_result `json:"final_tally_result"`
	Proposal_id        string             `json:"proposal_id"`
	Status             string             `json:"status"`
	Submit_time        string             `json:"submit_time"`
	Total_deposit      []Amount           `json:"total_deposit"`
	Voting_end_time    string             `json:"voting_end_time"`
	Voting_start_time  string             `json:"voting_start_time"`
}

type Content struct {
	Type        string      `json:"@type"`
	Amount      []Amount    `json:"amount"`
	Recipient   string      `json:recipient`
	Chain_name  string      `json:"chain_name"`
	Description string      `json:"description"`
	Params      interface{} `json:"params"`
	Title       string      `json:"title"`
}

type Final_tally_result struct {
	Abstain      string `json:"abstain"`
	No           string `json:"no"`
	No_with_veto string `json:"no_with_veto"`
	Yes          string `json:"yes"`
}

type Amount struct {
	Amount string `json:"amount"`
	Denom  string `json:"denom"`
}

type Result struct {
	Community_Spend_Pool    []Receiver_Details
	Sum_Of_Withdrawn_Amount *big.Int
}

type Receiver_Details struct {
	Amount_Withdrawn *big.Int
	Denom            string
	Timestamp        string
	Receiver_Address string
}

const COMMUNITY_POOL_SPEND_PROPOSAL = "CommunityPoolSpendProposal"
const PROPOSAL_STATUS_PASSED = "PROPOSAL_STATUS_PASSED"

func CommunityPollOutstanding(c *gin.Context) {

	cmd := exec.Command("fxcored", "query", "distribution", "community-pool", "--node", "https://fx-json.functionx.io:26657")

	stdout, err := cmd.Output()

	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	res := utils.YamlParser(string(stdout))

	c.JSON(http.StatusOK, res)

}

func CommunityPollDeductions(c *gin.Context) {

	cmd := exec.Command("fxcored", "q", "gov", "proposals", "-o", "json", "--node", "https://fx-json.functionx.io:26657")
	stdout, err := cmd.Output()

	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	output := stdout

	var res = Master{}
	var Receiver_Details = Receiver_Details{}
	var Result = Result{}

	if err := json.Unmarshal([]byte(output), &res); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	proposals := res.Proposals
	totalAmount := big.NewInt(0)

	for _, proposal := range proposals {
		//fmt.Print64f("ProposalId: %s, Type: %s, Status: %s, Amount: %s\n", proposal.Proposal_id, proposal.Content.Type, proposal.Status, proposal.Content.Amount)
		singleProposalAmount := big.NewInt(0)

		if strings.Contains(proposal.Content.Type, COMMUNITY_POOL_SPEND_PROPOSAL) && proposal.Status == PROPOSAL_STATUS_PASSED {
			for _, amount := range proposal.Content.Amount {
				currentAmount := new(big.Int)
				currentAmount, err := currentAmount.SetString(amount.Amount, 10)
				if !err {
					c.AbortWithStatus(http.StatusBadRequest)
					return
				}
				singleProposalAmount.Add(singleProposalAmount, currentAmount)
			}

			totalAmount.Add(totalAmount, singleProposalAmount)
			Receiver_Details.Amount_Withdrawn = singleProposalAmount
			Receiver_Details.Denom = "FX"
			Receiver_Details.Receiver_Address = proposal.Content.Recipient
			Receiver_Details.Timestamp = proposal.Submit_time

			Result.Community_Spend_Pool = append(Result.Community_Spend_Pool, Receiver_Details)
		}
	}

	Result.Sum_Of_Withdrawn_Amount = totalAmount

	c.JSON(http.StatusOK, Result)

}
