package componentv1_test

import (
	"testing"

	componentv1 "gitea.dev/codespace-proto-go/component/v1"
	"gitea.dev/codespace-proto-go/component/v1/componentv1connect"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

func TestComponentServiceContract(t *testing.T) {
	if componentv1connect.ComponentServiceName != "component.v1.ComponentService" {
		t.Fatalf("service name = %q", componentv1connect.ComponentServiceName)
	}

	requests := []proto.Message{
		&componentv1.GatewayControlRequest{},
		&componentv1.AuthorizeGatewayRequest{},
		&componentv1.IssueAgentAccessRequest{},
		&componentv1.CacheControlRequest{},
	}
	for _, request := range requests {
		descriptor := request.ProtoReflect().Descriptor()
		field := descriptor.Fields().ByName("protocol_version")
		if field == nil || field.Number() != protoreflect.FieldNumber(1) {
			t.Fatalf("%s protocol_version is not field 1", descriptor.FullName())
		}
	}

	open := (&componentv1.OpenCodeAuthorization{}).ProtoReflect().Descriptor().Fields()
	if open.ByName("runtime_uuid").Number() != 2 || open.ByName("endpoint_id").Number() != 3 {
		t.Fatal("OpenCodeAuthorization target fields changed")
	}
}
