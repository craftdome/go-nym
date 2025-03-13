package mixnet_test

import (
	"github.com/craftdome/go-nym/contracts/mixnet"
	"github.com/craftdome/go-nym/contracts/mixnet/query/delegations"
	"github.com/craftdome/go-nym/contracts/mixnet/query/intervals"
	"github.com/craftdome/go-nym/contracts/mixnet/query/nodes"
	"github.com/craftdome/go-nym/contracts/mixnet/query/rewards"
	"github.com/craftdome/go-nym/contracts/mixnet/shared/models"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"testing"
)

const (
	endpoint = "nym-grpc.polkachu.com:15390"
	contract = "n17srjznxl9dvzdkpwpw24gg668wc73val88a6m5ajg6ankwvz9wtst0cznr"
)

var (
	query *mixnet.QueryClient
)

func init() {
	conn, err := grpc.NewClient(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	query = mixnet.NewQueryClient(conn, contract)
}

func TestQueryClient_Contract_GetAdmin(t *testing.T) {
	resp, err := query.Contract.GetAdmin(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", resp)

	if resp.Admin == "" {
		t.Error("admin address is empty")
	}
}

func TestQueryClient_Contract_GetVersion(t *testing.T) {
	resp, err := query.Contract.GetVersion(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", resp)

	if resp.ContractName != "nym-mixnet-contract" {
		t.Errorf("want ContractName: nym-mixnet-contract, got %s", resp.ContractName)
	}

	if resp.BuildTimestamp == "" {
		t.Error("BuildTimestamp is empty")
	}

	if resp.BuildVersion == "" {
		t.Error("BuildVersion is empty")
	}

	if resp.CommitSHA == "" {
		t.Error("CommitSHA is empty")
	}

	if resp.CommitTimestamp == "" {
		t.Error("CommitTimestamp is empty")
	}

	if resp.CommitBranch == "" {
		t.Error("CommitBranch is empty")
	}

	if resp.RustcVersion == "" {
		t.Error("RustcVersion is empty")
	}

	if resp.CargoDebug == "" {
		t.Error("CargoDebug is empty")
	}

	if resp.CargoOptLevel == "" {
		t.Error("CargoOptLevel is empty")
	}
}

func TestQueryClient_Contract_GetCW2Version(t *testing.T) {
	resp, err := query.Contract.GetCW2Version(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", resp)

	if resp.Contract != "crate:nym-mixnet-contract" {
		t.Errorf("want Contract: nym-mixnet-contract, got %s", resp.Contract)
	}

	if resp.Version == "" {
		t.Error("Version is empty")
	}
}

func TestQueryClient_Contract_GetStateParams(t *testing.T) {
	resp, err := query.Contract.GetStateParams(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", resp)

	if resp.OperatorsParams.MinimumPledge.Denom == "" {
		t.Error("OperatorsParams.MinimumPledge.Denom is empty")
	}
}

func TestQueryClient_Contract_GetState(t *testing.T) {
	resp, err := query.Contract.GetState(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", resp)

	if resp.VestingContractAddress == "" {
		t.Error("VestingContractAddress is empty")
	}

	if resp.RewardingDenom != "unym" {
		t.Errorf("want RewardingDenom: unym, got %s", resp.RewardingDenom)
	}

	if resp.RewardingValidatorAddress == "" {
		t.Error("RewardingValidatorAddress is empty")
	}

	if resp.Owner == "" {
		t.Error("Owner is empty")
	}
}

func TestQueryClient_Contract_GetCurrentNodeVersion(t *testing.T) {
	resp, err := query.Contract.GetCurrentNodeVersion(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", resp)

	if resp.ID == 0 {
		t.Errorf("want ID == 0, got %d", resp.ID)
	}

	if resp.Info.Semver == "" {
		t.Errorf("want Info.Semver != '', got %s", resp.Info.Semver)
	}

	if resp.Info.IntroducedAtHeight == 0 {
		t.Errorf("want Info.IntroducedAtHeight != 0, got %d", resp.Info.IntroducedAtHeight)
	}
}

func TestQueryClient_Contract_GetNodeVersionHistory(t *testing.T) {
	resp, err := query.Contract.GetNodeVersionHistory(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", resp)

	if resp.StartNextAfter == 0 {
		t.Errorf("want StartNextAfter == 0, got %d", resp.StartNextAfter)
	}

}

func TestClient_Contract_GetParams(t *testing.T) {
	resp, err := query.Contract.GetRewardingParams(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", resp)

	if f, _ := resp.Interval.SybilResistance.Float32(); f == 0 {
		t.Errorf("want Interval.SybilResistance.Float32 != 0, got %f", f)
	}
}

func TestClient_Contract_GetEpochStatus(t *testing.T) {
	resp, err := query.Contract.GetEpochStatus(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", resp)

	if resp.BeingAdvancedBy == "" {
		t.Error("BeingAdvancedBy is empty")
	}

	if resp.State.String() == "" {
		t.Error("State is empty")
	}
}

func TestClient_Contract_GetIntervalStatus(t *testing.T) {
	resp, err := query.Contract.GetIntervalStatus(t.Context())
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("%+v\n", resp)

	if resp.CurrentBlocktime == 0 {
		t.Errorf("want CurrentBlocktime != 0, got %d", resp.CurrentBlocktime)
	}

	if resp.Interval.ID < 28 {
		t.Errorf("want Interval.ID >= 28, got %d", resp.Interval.ID)
	}

	if resp.Interval.EpochsInInterval != 720 {
		t.Errorf("want Interval.EpochsInInterval == 720, got %d", resp.Interval.EpochsInInterval)
	}

	if resp.Interval.EpochLength.Secs != 3600 {
		t.Errorf("want Interval.EpochLength.Secs == 3600, got %d", resp.Interval.EpochLength.Secs)
	}

	if resp.Interval.CurrentEpochStart.IsZero() {
		t.Errorf("want Interval.CurrentEpochStart != 0, got %v", resp.Interval.CurrentEpochStart)
	}
}

func TestClient_Nodes_GetAllBonded(t *testing.T) {
	params := nodes.GetAllBondedParams{}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Nodes.GetAllBonded(t.Context(), params)
		if err != nil {
			t.Errorf("GetAllBonded() error = %v", err)
		}
		t.Logf("len: %d, start_next_after: %d\n", len(got.Nodes), *got.StartNextAfter)

		params.StartAfter = *got.StartNextAfter
	})

	t.Run("Valid with start_next_after", func(t *testing.T) {
		got, err := query.Nodes.GetAllBonded(t.Context(), params)
		if err != nil {
			t.Errorf("GetAllBonded() error = %v", err)
			return
		}
		t.Logf("len: %d, start_next_after: %d\n", len(got.Nodes), *got.StartNextAfter)

		if len(got.Nodes) > 0 && got.Nodes[0].NodeID <= params.StartAfter {
			t.Errorf("got.Nodes[0].NodeID = %v, want gt %v (start_next_after)", got.Nodes[0].NodeID, params.StartAfter)
		}
	})
}

func TestClient_Nodes_GetAllDetailed(t *testing.T) {
	params := nodes.GetAllDetailedParams{
		StartAfter: 0,
		Limit:      10,
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Nodes.GetAllDetailed(t.Context(), params)
		if err != nil {
			t.Fatalf("GetAllDetailed() error = %v", err)
		}
		t.Logf("len: %d, start_next_after: %d\n", len(got.Nodes), *got.StartNextAfter)

		if len(got.Nodes) != int(params.Limit) {
			t.Errorf("len(got.Nodes) = %d, want %d", len(got.Nodes), params.Limit)
		}

		params.StartAfter = *got.StartNextAfter
	})

	t.Run("Valid with start_next_after", func(t *testing.T) {
		got, err := query.Nodes.GetAllDetailed(t.Context(), params)
		if err != nil {
			t.Fatalf("GetAllDetailed() error = %v", err)
		}
		t.Logf("len: %d, start_next_after: %d\n", len(got.Nodes), *got.StartNextAfter)

		if len(got.Nodes) > 0 && got.Nodes[0].BondInformation.NodeID <= params.StartAfter {
			t.Errorf("got.Nodes[0].NodeID = %v, want gt %v (start_next_after)", got.Nodes[0].BondInformation.NodeID, params.StartAfter)
		}
	})
}

func TestClient_Nodes_GetUnbonded(t *testing.T) {
	params := nodes.GetUnbondedParams{
		NodeID: 119,
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Nodes.GetUnbonded(t.Context(), params)
		if err != nil {
			t.Fatalf("GetUnbonded() error = %v", err)
		}
		t.Logf("%+v\n", got)

		if got.NodeID != params.NodeID {
			t.Errorf("got.NodeID = %d, want %d", got.NodeID, params.NodeID)
		}
	})
}

func TestClient_Nodes_GetAllUnbonded(t *testing.T) {
	params := nodes.GetAllUnbondedParams{
		Limit: 10,
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Nodes.GetAllUnbonded(t.Context(), params)
		if err != nil {
			t.Fatalf("GetAllUnbonded() error = %v", err)
		}
		t.Logf("len: %d, start_next_after: %d\n", len(got.Nodes), *got.StartNextAfter)

		params.StartAfter = *got.StartNextAfter
	})

	t.Run("Valid with start_next_after", func(t *testing.T) {
		got, err := query.Nodes.GetAllUnbonded(t.Context(), params)
		if err != nil {
			t.Fatalf("GetAllUnbonded() error = %v", err)
		}
		t.Logf("len: %d, start_next_after: %d\n", len(got.Nodes), *got.StartNextAfter)

		// Doesn't work in outside
		//if len(got.Nodes) > 0 && got.Nodes[0].NodeID <= params.StartAfter {
		//	t.Errorf("got.Nodes[0].NodeID = %v, want gt %v (start_next_after)", got.Nodes[0].NodeID, params.StartAfter)
		//}
	})
}

func TestClient_Nodes_GetUnbondedByOperator(t *testing.T) {
	params := nodes.GetUnbondedByOwnerParams{
		Owner: "n1vs7aplr9n208artds98j2h7mdpyhlrq4kj6k67",
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Nodes.GetUnbondedByOwner(t.Context(), params)
		if err != nil {
			t.Fatalf("GetUnbondedByOwner() error = %v", err)
		}
		t.Logf("len: %d, start_next_after: %d\n", len(got.Nodes), *got.StartNextAfter)

		params.StartAfter = *got.StartNextAfter
	})

	t.Run("Valid with start_next_after", func(t *testing.T) {
		got, err := query.Nodes.GetUnbondedByOwner(t.Context(), params)
		if err != nil {
			t.Fatalf("GetUnbondedByOwner() error = %v", err)
		}
		t.Logf("len: %d, start_next_after: %d\n", len(got.Nodes), *got.StartNextAfter)

		// Doesn't work in outside
		//if len(got.Nodes) > 0 && got.Nodes[0].NodeID <= params.StartAfter {
		//	t.Errorf("got.Nodes[0].NodeID = %v, want gt %v (start_next_after)", got.Nodes[0].NodeID, params.StartAfter)
		//}
	})
}

func TestClient_Nodes_GetUnbondedByIdentityKey(t *testing.T) {
	params := nodes.GetUnbondedByIdentityKeyParams{
		IdentityKey: "5qVwSJZSArAJibfGaS4G3oV7hidemit3WJFg62YMB2s6",
		StartAfter:  0,
		Limit:       10,
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Nodes.GetUnbondedByIdentityKey(t.Context(), params)
		if err != nil {
			t.Fatalf("GetUnbondedByIdentityKey() error = %v", err)
		}
		t.Logf("len: %d, start_next_after: %d\n", len(got.Nodes), *got.StartNextAfter)
	})

	t.Run("Valid with start_next_after", func(t *testing.T) {
		got, err := query.Nodes.GetUnbondedByIdentityKey(t.Context(), params)
		if err != nil {
			t.Fatalf("GetUnbondedByIdentityKey() error = %v", err)
		}
		t.Logf("len: %d, start_next_after: %d\n", len(got.Nodes), *got.StartNextAfter)

		if len(got.Nodes) > 0 && got.Nodes[0].NodeID <= params.StartAfter {
			t.Errorf("got.Nodes[0].NodeID = %v, want gt %v (start_next_after)", got.Nodes[0].NodeID, params.StartAfter)
		}
	})
}

func TestClient_Nodes_GetDetailedByOwner(t *testing.T) {
	params := nodes.GetDetailedByOwnerParams{
		Owner: "n1xwme0cstucs4ddycz7yvmwdnawv0zz5dam9596",
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Nodes.GetDetailedByOwner(t.Context(), params)
		if err != nil {
			t.Fatalf("GetDetailedByOwner() error = %v", err)
		}
		t.Logf("%+v", got)

		if f, _ := got.RewardingDetails.CostParams.ProfitMarginPercent.Float32(); f < 0.2 || f > 0.5 {
			t.Errorf("want RewardingDetails.CostParams.ProfitMarginPercent = [0.2, 0.5], got %v", f)
		}
	})
}

func TestClient_Nodes_GetDetailed(t *testing.T) {
	params := nodes.GetDetailedParams{
		NodeID: 895,
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Nodes.GetDetailed(t.Context(), params)
		if err != nil {
			t.Fatalf("GetDetailed() error = %v", err)
		}
		t.Logf("%+v", got)

		if f, _ := got.RewardingDetails.CostParams.ProfitMarginPercent.Float32(); f < 0.2 || f > 0.5 {
			t.Errorf("want RewardingDetails.CostParams.ProfitMarginPercent = [0.2, 0.5], got %v", f)
		}
	})
}

func TestClient_Nodes_GetDetailedByIdentityKey(t *testing.T) {
	params := nodes.GetDetailedByIdentityKeyParams{
		IdentityKey: "J3Wfpxca9mwnbipScWTkKCrbNnZ9J5M4YmjKCFTnXwWN",
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Nodes.GetDetailedByIdentityKey(t.Context(), params)
		if err != nil {
			t.Fatalf("GetDetailedByIdentityKey() error = %v", err)
		}
		t.Logf("%+v", got)

		if f, _ := got.RewardingDetails.CostParams.ProfitMarginPercent.Float32(); f < 0.2 || f > 0.5 {
			t.Errorf("want RewardingDetails.CostParams.ProfitMarginPercent = [0.2, 0.5], got %v", f)
		}
	})
}

func TestClient_Nodes_GetRewardingDetails(t *testing.T) {
	params := nodes.GetRewardingDetailsParams{
		NodeID: 895,
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Nodes.GetRewardingDetails(t.Context(), params)
		if err != nil {
			t.Fatalf("GetRewardingDetails() error = %v", err)
		}
		t.Logf("%+v", got)

		if f, _ := got.CostParams.ProfitMarginPercent.Float32(); f < 0.2 || f > 0.5 {
			t.Errorf("want CostParams.ProfitMarginPercent = [0.2, 0.5], got %v", f)
		}
	})
}

func TestClient_Nodes_GetStakeSaturation(t *testing.T) {
	params := nodes.GetStakeSaturationParams{
		NodeID: 895,
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Nodes.GetStakeSaturation(t.Context(), params)
		if err != nil {
			t.Fatalf("GetStakeSaturation() error = %v", err)
		}
		t.Logf("%+v", got)

		if got.NodeID != params.NodeID {
			t.Errorf("got node_id %d, want %d", got.NodeID, params.NodeID)
		}
	})
}

func TestClient_Nodes_GetEpochAssignmentByRole(t *testing.T) {
	t.Run("Valid Layer1", func(t *testing.T) {
		params := nodes.GetEpochAssignmentByRoleParams{
			Role: models.Layer1,
		}

		got, err := query.Nodes.GetEpochAssignmentByRole(t.Context(), params)
		if err != nil {
			t.Fatalf("GetEpochAssignmentByRole() error = %v", err)
		}
		t.Logf("%+v", got)

		if got.EpochID == 0 {
			t.Errorf("got epoch_id %d, want != %d", got.EpochID, 0)
		}

		if len(got.Nodes) == 0 {
			t.Errorf("len(nodes) = %d, want > 0", len(got.Nodes))
		}
	})

	t.Run("Not valid", func(t *testing.T) {
		params := nodes.GetEpochAssignmentByRoleParams{
			Role: "l22",
		}
		got, err := query.Nodes.GetEpochAssignmentByRole(t.Context(), params)
		if err == nil {
			t.Logf("%+v", got)
			t.Fatalf("GetEpochAssignmentByRole() got nil, want error")
		}
	})

	t.Run("Valid Layer2", func(t *testing.T) {
		params := nodes.GetEpochAssignmentByRoleParams{
			Role: models.Layer2,
		}

		got, err := query.Nodes.GetEpochAssignmentByRole(t.Context(), params)
		if err != nil {
			t.Fatalf("GetEpochAssignmentByRole() error = %v", err)
		}
		t.Logf("%+v", got)

		if got.EpochID == 0 {
			t.Errorf("got epoch_id %d, want != %d", got.EpochID, 0)
		}

		if len(got.Nodes) == 0 {
			t.Errorf("len(nodes) = %d, want > 0", len(got.Nodes))
		}
	})

	t.Run("Valid Layer3", func(t *testing.T) {
		params := nodes.GetEpochAssignmentByRoleParams{
			Role: models.Layer3,
		}

		got, err := query.Nodes.GetEpochAssignmentByRole(t.Context(), params)
		if err != nil {
			t.Fatalf("GetEpochAssignmentByRole() error = %v", err)
		}
		t.Logf("%+v", got)

		if got.EpochID == 0 {
			t.Errorf("got epoch_id %d, want != %d", got.EpochID, 0)
		}

		if len(got.Nodes) == 0 {
			t.Errorf("len(nodes) = %d, want > 0", len(got.Nodes))
		}
	})

	t.Run("Valid EntryGateway", func(t *testing.T) {
		params := nodes.GetEpochAssignmentByRoleParams{
			Role: models.EntryGateway,
		}

		got, err := query.Nodes.GetEpochAssignmentByRole(t.Context(), params)
		if err != nil {
			t.Fatalf("GetEpochAssignmentByRole() error = %v", err)
		}
		t.Logf("%+v", got)

		if got.EpochID == 0 {
			t.Errorf("got epoch_id %d, want != %d", got.EpochID, 0)
		}

		if len(got.Nodes) == 0 {
			t.Errorf("len(nodes) = %d, want > 0", len(got.Nodes))
		}
	})

	t.Run("Valid ExitGateway", func(t *testing.T) {
		params := nodes.GetEpochAssignmentByRoleParams{
			Role: models.ExitGateway,
		}

		got, err := query.Nodes.GetEpochAssignmentByRole(t.Context(), params)
		if err != nil {
			t.Fatalf("GetEpochAssignmentByRole() error = %v", err)
		}
		t.Logf("%+v", got)

		if got.EpochID == 0 {
			t.Errorf("got epoch_id %d, want != %d", got.EpochID, 0)
		}

		if len(got.Nodes) == 0 {
			t.Errorf("len(nodes) = %d, want > 0", len(got.Nodes))
		}
	})
}

func TestClient_Nodes_GetEpochAssignmentMetadata(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		got, err := query.Nodes.GetEpochAssignmentMetadata(t.Context())
		if err != nil {
			t.Fatalf("GetEpochAssignmentMetadata() error = %v", err)
		}
		t.Logf("%+v", got)

		if got.Metadata.EpochID == 0 {
			t.Errorf("got epoch_id %d, want != %d", got.Metadata.EpochID, 0)
		}

		if got.Metadata.EntryGatewayMetadata.NumNodes == 0 {
			t.Errorf("got eg_metadata.NumNodes, want != %d", 0)
		}

		if got.Metadata.ExitGatewayMetadata.NumNodes == 0 {
			t.Errorf("got xg_metadata.NumNodes, want != %d", 0)
		}

		if got.Metadata.Layer1Metadata.NumNodes == 0 {
			t.Errorf("got l1_gateway_metadata.NumNodes, want != %d", 0)
		}

		if got.Metadata.Layer2Metadata.NumNodes == 0 {
			t.Errorf("got l2_gateway_metadata.NumNodes, want != %d", 0)
		}

		if got.Metadata.Layer3Metadata.NumNodes == 0 {
			t.Errorf("got l3_gateway_metadata.NumNodes, want != %d", 0)
		}

		if got.Metadata.StandbyMetadata.NumNodes != 0 {
			t.Errorf("got stb_gateway_metadata.NumNodes, want != %d", 0)
		}
	})
}

func TestClient_Delegations_GetAll(t *testing.T) {
	params := delegations.GetDelegationsParams{
		//StartAfter: models.StorageKey{V1: 1, V2: "n18ds90dz0ezjy9a059k8zamrj7q2u98c9yq9th4"},
		Limit: 1,
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Delegations.GetAll(t.Context(), params)
		if err != nil {
			t.Fatalf("GetAll() error = %v", err)
		}
		t.Logf("%+v", got)

		if len(got.Delegations) > 0 {
			if got.Delegations[0].Owner == "" {
				t.Errorf("got owner %s, want not empty", got.Delegations[0].Owner)
			}

			if got.Delegations[0].Height == 0 {
				t.Errorf("got height %d, want > 0", got.Delegations[0].Height)
			}

			if got.Delegations[0].Amount.IsZero() {
				t.Errorf("got zero amount, want > 0")
			}
		}

		if got.StartNextAfter != nil {
			params.StartAfter = *got.StartNextAfter
		}
	})

	t.Run("Valid2", func(t *testing.T) {
		got, err := query.Delegations.GetAll(t.Context(), params)
		if err != nil {
			t.Fatalf("GetAll() error = %v", err)
		}
		t.Logf("%+v", got)

		if len(got.Delegations) > 0 {
			if got.Delegations[0].Amount.IsZero() {
				t.Errorf("got zero amount, want > 0")
			}

			if got.Delegations[0].Height == 0 {
				t.Errorf("got height %d, want > 0", got.Delegations[0].Height)
			}
		}

		if got.StartNextAfter != nil {
			params.Limit = models.DelegationPageMaxRetrievalLimit
			params.StartAfter = *got.StartNextAfter

			t.Run("Valid with StartNextAfter", func(t *testing.T) {
				var i int
			loop:
				got, err := query.Delegations.GetAll(t.Context(), params)
				if err != nil {
					t.Fatalf("GetAll() error = %v", err)
				}

				i += len(got.Delegations)
				if got.StartNextAfter != nil {
					params.StartAfter = *got.StartNextAfter
					goto loop
				}

				t.Logf("last %+v of %d", got, i)
			})
		}
	})
}

func TestClient_Delegations_GetByNode(t *testing.T) {
	params := delegations.GetNodeDelegationsParams{
		NodeID: 895,
		Limit:  10,
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Delegations.GetByNode(t.Context(), params)
		if err != nil {
			t.Fatalf("GetByNode() error = %v", err)
		}
		t.Logf("%+v", got)

		if len(got.Delegations) > 0 {
			if got.Delegations[0].NodeID != params.NodeID {
				t.Errorf("got node_id %d, want %d", got.Delegations[0].NodeID, params.NodeID)
			}

			if got.Delegations[0].Amount.IsZero() {
				t.Errorf("got zero amount, want > 0")
			}

			if got.Delegations[0].Owner == "" {
				t.Errorf("got owner %s, want not empty", got.Delegations[0].Owner)
			}

			if got.Delegations[0].Height == 0 {
				t.Errorf("got height %d, want > 0", got.Delegations[0].Height)
			}
		}
	})
}

func TestClient_Delegations_GetByDelegator(t *testing.T) {
	params := delegations.GetDelegatorDelegationsParams{
		Delegator: "n1rnxpdpx3kldygsklfft0gech7fhfcux4zst5lw",
		//StartAfter: models.StorageKey{V1: 21, V2: "n1rnxpdpx3kldygsklfft0gech7fhfcux4zst5lw"},
		Limit: 1,
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Delegations.GetByDelegator(t.Context(), params)
		if err != nil {
			t.Fatalf("GetByDelegator() error = %v", err)
		}
		t.Logf("%+v", got)

		if len(got.Delegations) > 0 {
			if got.Delegations[0].Owner != params.Delegator {
				t.Errorf("got owner %s, want %s", got.Delegations[0].Owner, params.Delegator)
			}

			if got.Delegations[0].Amount.IsZero() {
				t.Errorf("got zero amount, want > 0")
			}

			if got.Delegations[0].Height == 0 {
				t.Errorf("got height %d, want > 0", got.Delegations[0].Height)
			}
		}

		if got.StartNextAfter != nil {
			params.StartAfter = *got.StartNextAfter
		}
	})

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Delegations.GetByDelegator(t.Context(), params)
		if err != nil {
			t.Fatalf("GetByDelegator() error = %v", err)
		}
		t.Logf("%+v", got)

		if len(got.Delegations) > 0 {
			if got.Delegations[0].Owner != params.Delegator {
				t.Errorf("got owner %s, want %s", got.Delegations[0].Owner, params.Delegator)
			}

			if got.Delegations[0].Amount.IsZero() {
				t.Errorf("got zero amount, want > 0")
			}

			if got.Delegations[0].Height == 0 {
				t.Errorf("got height %d, want > 0", got.Delegations[0].Height)
			}
		}

		if got.StartNextAfter != nil {
			params.Limit = models.DelegationPageDefaultRetrievalLimit
			params.StartAfter = *got.StartNextAfter

			t.Run("Valid with StartNextAfter", func(t *testing.T) {
				var i int
			loop:
				got, err := query.Delegations.GetByDelegator(t.Context(), params)
				if err != nil {
					t.Fatalf("GetByDelegator() error = %v", err)
				}

				i += len(got.Delegations)
				if got.StartNextAfter != nil {
					params.StartAfter = *got.StartNextAfter
					goto loop
				}

				t.Logf("last %+v of %d", got, i)
			})
		}
	})

}

func TestClient_Delegations_GetByNodeAndDelegator(t *testing.T) {
	params := delegations.GetNodeDelegationParams{
		NodeID:    1227,
		Delegator: "n1rnxpdpx3kldygsklfft0gech7fhfcux4zst5lw",
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Delegations.GetByNodeAndDelegator(t.Context(), params)
		if err != nil {
			t.Fatalf("GetByNodeAndDelegator() error = %v", err)
		}
		t.Logf("%+v", got)

		if got.Delegation != nil {
			if got.Delegation.NodeID != params.NodeID {
				t.Errorf("got node_id %d, want %d", got.Delegation.NodeID, params.NodeID)
			}

			if got.Delegation.Amount.IsZero() {
				t.Errorf("got zero amount, want > 0")
			}

			if got.Delegation.Height == 0 {
				t.Errorf("got height %d, want > 0", got.Delegation.Height)
			}

			if got.Delegation.Owner != params.Delegator {
				t.Errorf("got owner %s, want %s", got.Delegation.Owner, params.Delegator)
			}
		}
	})
}

func TestClient_Rewards_GetPendingByOwner(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		params := rewards.GetPendingByOwnerParams{
			Owner: "n1xwme0cstucs4ddycz7yvmwdnawv0zz5dam9596",
		}

		got, err := query.Rewards.GetPendingByOwner(t.Context(), params)
		if err != nil {
			t.Fatalf("GetPendingByOwner() error = %v", err)
		}
		t.Logf("%+v", got)

		if got.AmountEarned.IsZero() {
			t.Errorf("got zero amount, want > 0")
		}
	})
}

func TestClient_Rewards_GetPendingByNode(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		params := rewards.GetPendingByNodeParams{
			NodeID: 895,
		}

		got, err := query.Rewards.GetPendingByNode(t.Context(), params)
		if err != nil {
			t.Fatalf("GetPendingByNode() error = %v", err)
		}
		t.Logf("%+v", got)

		if got.AmountEarned.IsZero() {
			t.Errorf("got zero amount, want > 0")
		}
	})
}

func TestClient_Rewards_GetPendingByNodeAndDelegator(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		params := rewards.GetPendingByNodeAndDelegatorParams{
			Delegator: "n1rnxpdpx3kldygsklfft0gech7fhfcux4zst5lw",
			NodeID:    895,
		}

		got, err := query.Rewards.GetPendingByNodeAndDelegator(t.Context(), params)
		if err != nil {
			t.Fatalf("GetPendingByNodeAndDelegator() error = %v", err)
		}
		t.Logf("%+v", got)

		if got.AmountEarned.IsZero() {
			t.Errorf("got zero amount, want > 0")
		}
	})
}

func TestClient_Rewards_EstimateOperatorReward(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		params := rewards.EstimateOperatorRewardParams{
			NodeID:               895,
			EstimatedPerformance: 0.99,
			EstimatedWork:        0.98,
		}

		got, err := query.Rewards.EstimateOperatorReward(t.Context(), params)
		if err != nil {
			t.Fatalf("EstimateOperatorReward() error = %v", err)
		}
		t.Logf("%+v", got)
	})
}

func TestClient_Rewards_EstimateDelegatorReward(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		params := rewards.EstimateDelegatorRewardParams{
			Delegator:            "n1rnxpdpx3kldygsklfft0gech7fhfcux4zst5lw",
			NodeID:               895,
			EstimatedPerformance: 0.99,
			EstimatedWork:        0.98,
		}

		got, err := query.Rewards.EstimateDelegatorReward(t.Context(), params)
		if err != nil {
			t.Errorf("EstimateDelegatorReward() error = %v", err)
		}
		t.Logf("%+v", got)
	})
}

func TestClient_Intervals_GetPendingEpochEvents(t *testing.T) {
	params := intervals.GetPendingEpochEventsParams{
		Limit: 2,
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Intervals.GetPendingEpochEvents(t.Context(), params)
		if err != nil {
			t.Fatalf("GetPendingEpochEvents() error = %v", err)
		}
		t.Logf("%+v", got)

		if len(got.Events) == 0 {
			return
		}

		if got.Events[0].ID == 0 {
			t.Errorf("got zero ID, want > 0")
		}

		if got.Events[0].Event.CreatedAt == 0 {
			t.Errorf("got zero block height, want > 0")
		}

		kind := got.Events[0].Event.Kind

		switch {
		case kind.Delegate != nil:
			if kind.Delegate.NodeID == 0 {
				t.Errorf("got zero node_id, want > 0")
			}

			if kind.Delegate.Owner == "" {
				t.Errorf("got zero owner, want not empty")
			}

			if kind.Delegate.Amount.IsZero() {
				t.Errorf("got zero amount, want > 0")
			}
		case kind.Undelegate != nil:
			if kind.Undelegate.NodeID == 0 {
				t.Errorf("got zero node_id, want > 0")
			}
		case kind.UnbondNode != nil:
			if kind.UnbondNode.NodeID == 0 {
				t.Errorf("got zero node_id, want > 0")
			}
		case kind.NodeIncreasePledge != nil:
			if kind.NodeIncreasePledge.NodeID == 0 {
				t.Errorf("got zero node_id, want > 0")
			}

			if kind.NodeIncreasePledge.IncreaseBy.IsZero() {
				t.Errorf("got zero amount, want > 0")
			}
		case kind.NodeDecreasePledge != nil:
			if kind.NodeDecreasePledge.NodeID == 0 {
				t.Errorf("got zero node_id, want > 0")
			}

			if kind.NodeDecreasePledge.DecreaseBy.IsZero() {
				t.Errorf("got zero decrease_by, want > 0")
			}
		case kind.UpdateActiveSet != nil:
			if kind.UpdateActiveSet.Update.Mixnodes == 0 {
				t.Errorf("got zero mixnodes, want > 0")
			}

			if kind.UpdateActiveSet.Update.ExitGateways == 0 {
				t.Errorf("got zero exit_gateways, want > 0")
			}

			if kind.UpdateActiveSet.Update.EntryGateways == 0 {
				t.Errorf("got zero entry_gateways, want > 0")
			}
		default:
			t.Errorf("unknown kind: %+v", kind)
		}

		if got.StartNextAfter.IsZero() {
			return
		}

		params.StartAfter = got.StartNextAfter

		t.Run("Valid", func(t *testing.T) {
			params.Limit = models.EpochEventsMaxRetrievalLimit
			var i int
		loop:
			got, err := query.Intervals.GetPendingEpochEvents(t.Context(), params)
			if err != nil {
				t.Fatalf("GetPendingEpochEvents() error = %v", err)
			}
			t.Logf("%+v", got)

			i += len(got.Events)
			if !got.StartNextAfter.IsZero() {
				params.StartAfter = got.StartNextAfter
				goto loop
			}

			t.Logf("last: %+v of %d", got, i)
		})
	})
}

func TestClient_Intervals_GetPendingIntervalEvents(t *testing.T) {
	params := intervals.GetPendingIntervalEventsParams{
		Limit: 2,
	}

	t.Run("Valid", func(t *testing.T) {
		got, err := query.Intervals.GetPendingIntervalEvents(t.Context(), params)
		if err != nil {
			t.Fatalf("GetPendingIntervalEvents() error = %v", err)
		}
		t.Logf("%+v", got)

		if len(got.Events) == 0 {
			return
		}

		if got.Events[0].ID == 0 {
			t.Errorf("got zero ID, want > 0")
		}

		if got.Events[0].Event.CreatedAt == 0 {
			t.Errorf("got zero block height, want > 0")
		}

		kind := got.Events[0].Event.Kind
		switch {
		case kind.ChangeNodeCostParams != nil:
			if kind.ChangeNodeCostParams.NodeID == 0 {
				t.Errorf("got zero node_id, want > 0")
			}

			if kind.ChangeNodeCostParams.NewCosts.IntervalOperatingCost.IsZero() &&
				kind.ChangeNodeCostParams.NewCosts.ProfitMarginPercent.IsZero() {
				t.Errorf("got zero interval_operating_cost && profit_margin_percent both, want > 0 at once")
			}
		case kind.UpdateIntervalConfig != nil:
			if kind.UpdateIntervalConfig.EpochsInInterval != 720 {
				t.Errorf("got %d epochs in interval, want 720", kind.UpdateIntervalConfig.EpochsInInterval)
			}

			if kind.UpdateIntervalConfig.EpochDurationSecs != 3600 {
				t.Errorf("got %d epoch duration, want 3600", kind.UpdateIntervalConfig.EpochDurationSecs)
			}
		case kind.UpdateRewardingParams != nil:
			if kind.UpdateRewardingParams.Update.SybilResistancePercent.IsZero() {
				t.Errorf("got zero update_sybil_resistance_percent, want > 0")
			}

			if kind.UpdateRewardingParams.Update.RewardPool.IsZero() {
				t.Errorf("got zero reward_pool, want > 0")
			}

			if kind.UpdateRewardingParams.Update.RewardedSetParams.EntryGateways == 0 {
				t.Errorf("got zero rewarded_set_entry_gateways, want > 0")
			}

			if kind.UpdateRewardingParams.Update.RewardedSetParams.ExitGateways == 0 {
				t.Errorf("got zero rewarded_set_exit_gateways, want > 0")
			}

			if kind.UpdateRewardingParams.Update.RewardedSetParams.Mixnodes == 0 {
				t.Errorf("got zero rewarded_set_mixnodes, want > 0")
			}

			if kind.UpdateRewardingParams.Update.IntervalPoolEmission.IsZero() {
				t.Errorf("got zero interval_pool_emission, want > 0")
			}

			if kind.UpdateRewardingParams.Update.ActiveSetWorkFactor.IsZero() {
				t.Errorf("got zero active_set_work_factor, want > 0")
			}

			if kind.UpdateRewardingParams.Update.StakingSupplyScaleFactor.IsZero() {
				t.Errorf("got zero staking_supply_scale_factor, want > 0")
			}

			if kind.UpdateRewardingParams.Update.StakingSupply.IsZero() {
				t.Errorf("got zero staking_supply, want > 0")
			}
		default:
			t.Errorf("unknown kind: %+v", kind)
		}

		if got.StartNextAfter.IsZero() {
			return
		}

		params.StartAfter = got.StartNextAfter

		t.Run("Valid", func(t *testing.T) {
			params.Limit = models.IntervalEventsMaxRetrievalLimit
			var i int
		loop:
			got, err := query.Intervals.GetPendingIntervalEvents(t.Context(), params)
			if err != nil {
				t.Fatalf("GetPendingIntervalEvents() error = %v", err)
			}
			t.Logf("%+v", got)

			i += len(got.Events)
			if !got.StartNextAfter.IsZero() {
				params.StartAfter = got.StartNextAfter
				goto loop
			}

			t.Logf("last: %+v of %d", got, i)
		})
	})

}

func TestClient_Intervals_GetPendingEpochEvent(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		events, err := query.Intervals.GetPendingEpochEvents(t.Context(), intervals.GetPendingEpochEventsParams{Limit: 1})
		if err != nil {
			t.Fatalf("GetPendingEpochEvents() error = %v", err)
		}

		if len(events.Events) == 0 {
			return
		}

		params := intervals.GetPendingEpochEventParams{
			EventID: events.Events[0].ID,
		}

		got, err := query.Intervals.GetPendingEpochEvent(t.Context(), params)
		if err != nil {
			t.Fatalf("GetPendingEpochEvent() error = %v", err)
		}
		t.Logf("%+v", got)

		if got.EventID != params.EventID {
			t.Errorf("got event_id = %v, want %v", got.EventID, params.EventID)
		}

		if got.Event.CreatedAt == 0 {
			t.Errorf("got zero block height, want > 0")
		}
	})
}

func TestClient_Intervals_GetPendingIntervalEvent(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		events, err := query.Intervals.GetPendingIntervalEvents(t.Context(), intervals.GetPendingIntervalEventsParams{Limit: 1})
		if err != nil {
			t.Fatalf("GetPendingIntervalEvents() error = %v", err)
		}

		if len(events.Events) == 0 {
			return
		}

		params := intervals.GetPendingIntervalEventParams{
			EventID: events.Events[0].ID,
		}

		got, err := query.Intervals.GetPendingIntervalEvent(t.Context(), params)
		if err != nil {
			t.Fatalf("GetPendingIntervalEvent() error = %v", err)
		}
		t.Logf("%+v", got)

		if got.EventID != params.EventID {
			t.Errorf("got event_id = %v, want %v", got.EventID, params.EventID)
		}

		if got.Event.CreatedAt == 0 {
			t.Errorf("got zero block height, want > 0")
		}
	})
}

func TestClient_Intervals_GetNumberOfPendingEvents(t *testing.T) {
	t.Run("Valid", func(t *testing.T) {
		got, err := query.Intervals.GetNumberOfPendingEvents(t.Context())
		if err != nil {
			t.Fatalf("GetNumberOfPendingEvents() error = %v", err)
		}
		t.Logf("%+v", got)
	})
}
