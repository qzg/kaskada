package api

import (
	"context"
	"reflect"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/descriptorpb"

	apiv1alpha "github.com/kaskada-ai/kaskada/gen/proto/go/kaskada/kaskada/v1alpha"
)

type apiClient struct {
	query           QueryClient
	materialization MaterializationClient
	table           TableClient
	view            ViewClient
}

type ApiClient interface {
	LoadFile(name string, fileInput *apiv1alpha.FileInput) error
	Create(item protoreflect.ProtoMessage) error
	Delete(item protoreflect.ProtoMessage) error
	Get(item protoreflect.ProtoMessage) (protoreflect.ProtoMessage, error)
	List(item protoreflect.ProtoMessage) ([]protoreflect.ProtoMessage, error)
	Query(*apiv1alpha.CreateQueryRequest) (*apiv1alpha.CreateQueryResponse, error)
}

func NewApiClient() ApiClient {
	ctx, conn := getContextAndConnection()
	return &apiClient{
		query:           NewQueryServiceClient(ctx, conn),
		materialization: NewMaterializationServiceClient(ctx, conn),
		table:           NewTableServiceClient(ctx, conn),
		view:            NewViewServiceClient(ctx, conn),
	}
}

func (c apiClient) LoadFile(name string, fileInput *apiv1alpha.FileInput) error {
	return c.table.LoadFile(name, fileInput)
}

func (c apiClient) Query(req *apiv1alpha.CreateQueryRequest) (*apiv1alpha.CreateQueryResponse, error) {
	return c.query.Query(req)
}

func (c apiClient) Create(item protoreflect.ProtoMessage) error {
	kind := reflect.TypeOf(item).String()
	switch t := item.(type) {
	case *apiv1alpha.Materialization:
		return c.materialization.Create(t)
	case *apiv1alpha.Table:
		return c.table.Create(t)
	case *apiv1alpha.View:
		return c.view.Create(t)
	default:
		log.Fatal().Str("kind", kind).Msg("unknown item kind for create")
		return nil
	}
}

func (c apiClient) Delete(item protoreflect.ProtoMessage) error {
	kind := reflect.TypeOf(item).String()
	switch t := item.(type) {
	case *apiv1alpha.Materialization:
		return c.materialization.Delete(t.MaterializationName)
	case *apiv1alpha.Table:
		return c.table.Delete(t.TableName)
	case *apiv1alpha.View:
		return c.view.Delete(t.ViewName)
	default:
		log.Fatal().Str("kind", kind).Msg("unknown item kind for delete")
		return nil
	}
}

func (c apiClient) Get(item protoreflect.ProtoMessage) (protoreflect.ProtoMessage, error) {
	kind := reflect.TypeOf(item).String()
	switch t := item.(type) {
	case *apiv1alpha.Materialization:
		return c.materialization.Get(t.MaterializationName)
	case *apiv1alpha.Table:
		return c.table.Get(t.TableName)
	case *apiv1alpha.View:
		return c.view.Get(t.ViewName)
	default:
		log.Fatal().Str("kind", kind).Msg("unknown item kind for get")
		return nil, nil
	}
}

func (c apiClient) List(item protoreflect.ProtoMessage) ([]protoreflect.ProtoMessage, error) {
	kind := reflect.TypeOf(item).String()
	results := make([]protoreflect.ProtoMessage, 0)
	switch item.(type) {
	case *apiv1alpha.ListMaterializationsRequest:
		materializations, err := c.materialization.List()
		if err != nil {
			return nil, err
		}
		for _, m := range materializations {
			results = append(results, m)
		}
	case *apiv1alpha.ListTablesRequest:
		tables, err := c.table.List()
		if err != nil {
			return nil, err
		}
		for _, t := range tables {
			results = append(results, t)
		}
	case *apiv1alpha.ListViewsRequest:
		views, err := c.view.List()
		if err != nil {
			return nil, err
		}
		for _, v := range views {
			results = append(results, v)
		}
	default:
		log.Fatal().Str("kind", kind).Msg("unknown item kind for list")
	}
	return results, nil
}

func GetName(item protoreflect.ProtoMessage) string {
	kind := reflect.TypeOf(item).String()
	switch t := item.(type) {
	case *apiv1alpha.Materialization:
		return t.MaterializationName
	case *apiv1alpha.Table:
		return t.TableName
	case *apiv1alpha.View:
		return t.ViewName
	default:
		log.Fatal().Str("kind", kind).Msg("unknown item kind for getName")
		return ""
	}
}

func clearOutputOnly[M protoreflect.ProtoMessage](message M) M {
	msg := message.ProtoReflect()
	// Iterate over each field in the message
	msg.Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {

		// Get the options associated with each field (options are represented as a proto message)
		options := fd.Options().(*descriptorpb.FieldOptions)

		// Iterate over each option
		options.ProtoReflect().Range(func(fd2 protoreflect.FieldDescriptor, v2 protoreflect.Value) bool {
			// If the field behavior option is set
			if fd2.FullName() == "google.api.field_behavior" {

				// Iterate over the assigned field behaviors
				behaviors := v2.List()
				for i := 0; i < behaviors.Len(); i++ {
					// ...and zero out the field if it's behavior is "output only"
					if behaviors.Get(i).Enum() == annotations.FieldBehavior_OUTPUT_ONLY.Number() {
						msg.Clear(fd)
						return false
					}
				}
			}
			return true
		})
		return true
	})
	return msg.Interface().(M)
}

func clearOutputOnlyList[M protoreflect.ProtoMessage](messages []M) []M {
	output := make([]M, 0, len(messages))
	for _, m := range messages {
		output = append(output, clearOutputOnly(m))
	}
	return output
}

func getClientID() string {
	clientId := viper.GetString("kaskada-client-id")
	if clientId == "" {
		log.Debug().Msg("no client-id found, initiating request without passing client-id header.")
	}
	return clientId
}

func getContextAndConnection() (context.Context, *grpc.ClientConn) {
	ctx := metadata.AppendToOutgoingContext(context.Background(), "client-id", getClientID())

	opts := []grpc.DialOption{}
	if viper.GetBool("use-tls") {
		creds := credentials.NewTLS(nil)
		opts = append(opts, grpc.WithTransportCredentials(creds))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	var err error
	serverAddr := viper.GetString("kaskada-api-server")
	conn, err := grpc.Dial(serverAddr, opts...)
	if err != nil {
		log.Fatal().Err(err).Msg("failed to dial the API")
		return nil, nil
	}
	return ctx, conn
}
