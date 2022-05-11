package credentials

import (
	"github.com/iden3/go-circuits"
	"github.com/iden3/go-iden3-auth/types"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestVerifyCredentialFetchRequest(t *testing.T) {

	var message types.CredentialFetchRequest
	message.Type = CredentialFetchRequestMessageType
	message.Data = types.CredentialFetchRequestMessageData{ClaimID: "992fc184-c902-4f9a-af62-b383cc5e1eb4", Schema: "KYCAgeCredential"}

	zkpProof := types.ZeroKnowledgeProof{
		Type:      types.ZeroKnowledgeProofType,
		CircuitID: circuits.AuthCircuitID,
	}
	zkpProof.ProofData = &types.ProofData{
		A: []string{
			"6807142976568489254129987481389970790048870221943660648833750801722749769662",
			"13811182779758948993435669124001052501939669904238445458453308627013829993881",
			"1",
		},
		B: [][]string{
			{
				"1100658387420856656999617260396587549490320987275888589013664343574809180330",
				"6271619554100652532302412650545865559102683218896584596952129504406572338279",
			},
			{
				"14732910796480272245291363689840710264816417845998668210234805961967222411399",
				"697511497805383174761860295477525070010524578030535203059896030784240207952",
			},
			{
				"1",
				"0",
			}},
		C: []string{
			"3322888400314063147927477851922827359406772099015587732727269650428166130415",
			"11791447421105500417246293414158106577578665220990150855390594651727173683574",
			"1",
		},
	}
	zkpProof.PubSignals = []string{
		"1",
		"18656147546666944484453899241916469544090258810192803949522794490493271005313",
		"379949150130214723420589610911161895495647789006649785264738141299135414272",
	}
	message.Data.Scope = []interface{}{zkpProof}

	err := VerifyCredentialFetchRequest(&message)
	assert.Nil(t, err)
}

func TestExtractDataFromCredentialFetchRequest(t *testing.T) {

	var message types.CredentialFetchRequest
	message.Type = CredentialFetchRequestMessageType
	message.Data = types.CredentialFetchRequestMessageData{ClaimID: "992fc184-c902-4f9a-af62-b383cc5e1eb4", Schema: "KYCAgeCredential"}

	zkpProof := types.ZeroKnowledgeProof{
		Type:      types.ZeroKnowledgeProofType,
		CircuitID: circuits.AuthCircuitID,
	}

	zkpProof.ProofData = &types.ProofData{
		A: []string{
			"6807142976568489254129987481389970790048870221943660648833750801722749769662",
			"13811182779758948993435669124001052501939669904238445458453308627013829993881",
			"1",
		},
		B: [][]string{
			{
				"1100658387420856656999617260396587549490320987275888589013664343574809180330",
				"6271619554100652532302412650545865559102683218896584596952129504406572338279",
			},
			{
				"14732910796480272245291363689840710264816417845998668210234805961967222411399",
				"697511497805383174761860295477525070010524578030535203059896030784240207952",
			},
			{
				"1",
				"0",
			}},
		C: []string{
			"3322888400314063147927477851922827359406772099015587732727269650428166130415",
			"11791447421105500417246293414158106577578665220990150855390594651727173683574",
			"1",
		},
	}
	zkpProof.PubSignals = []string{
		"1",
		"18656147546666944484453899241916469544090258810192803949522794490493271005313",
		"379949150130214723420589610911161895495647789006649785264738141299135414272",
	}
	message.Data.Scope = []interface{}{zkpProof}

	token, err := ExtractMetadataFromCredentialFetchRequest(&message)
	assert.Nil(t, err)
	assert.Equal(t, "1", token.Challenge)
	assert.Equal(t, "992fc184-c902-4f9a-af62-b383cc5e1eb4", token.ClaimID)
	assert.Equal(t, "KYCAgeCredential", token.ClaimSchema)

}
