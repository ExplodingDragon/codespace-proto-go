package codespacev1_test

import (
	"testing"

	codespacev1 "gitea.dev/codespace-proto-go/codespace/v1"
	"gitea.dev/codespace-proto-go/codespace/v1/codespacev1connect"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestDevContainerFeatureOptionsAreTyped(t *testing.T) {
	feature := &codespacev1.DevContainerFeature{
		Reference: "ghcr.io/example/features/tool:1",
		Options: []*codespacev1.DevContainerFeatureOption{
			{Name: "version", Value: &codespacev1.DevContainerFeatureOption_StringValue{StringValue: "1.2.3"}},
			{Name: "enabled", Value: &codespacev1.DevContainerFeatureOption_BoolValue{BoolValue: true}},
		},
	}

	if feature.Options[0].GetStringValue() != "1.2.3" {
		t.Fatalf("string option = %q", feature.Options[0].GetStringValue())
	}
	if !feature.Options[1].GetBoolValue() {
		t.Fatal("boolean option is false")
	}
}

func TestManagerServiceName(t *testing.T) {
	if codespacev1connect.ManagerServiceName != "codespace.v1.ManagerService" {
		t.Fatalf("service name = %q", codespacev1connect.ManagerServiceName)
	}
}

func TestManagerServiceRequestsCarryProtocolVersionAsFieldOne(t *testing.T) {
	requests := []proto.Message{
		&codespacev1.RegisterManagerRequest{},
		&codespacev1.DeclareManagerRequest{},
		&codespacev1.FetchOperationsRequest{},
		&codespacev1.FinalizeOperationRequest{},
		&codespacev1.UpdateLogRequest{},
		&codespacev1.ReportRuntimeMetadataRequest{},
		&codespacev1.RequestRuntimeAccessRequest{},
		&codespacev1.RequestIdleStopRequest{},
		&codespacev1.ValidateOpenTokenRequest{},
		&codespacev1.ValidatePublicEndpointRequest{},
		&codespacev1.VerifySSHPublicKeyRequest{},
		&codespacev1.ReportInstancesRequest{},
		&codespacev1.ReportRuntimeTransitionRequest{},
		&codespacev1.RevalidateGatewaySessionRequest{},
	}

	for _, request := range requests {
		descriptor := request.ProtoReflect().Descriptor()
		field := descriptor.Fields().ByName("protocol_version")
		if field == nil {
			t.Fatalf("%s does not define protocol_version", descriptor.FullName())
		}
		if field.Number() != protoreflect.FieldNumber(1) {
			t.Fatalf("%s protocol_version field number = %d", descriptor.FullName(), field.Number())
		}
	}
}
