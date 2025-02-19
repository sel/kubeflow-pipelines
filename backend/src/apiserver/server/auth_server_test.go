package server

import (
	"context"
	"testing"

	api "github.com/kubeflow/pipelines/backend/api/v1beta1/go_client"
	kfpauth "github.com/kubeflow/pipelines/backend/src/apiserver/auth"
	"github.com/kubeflow/pipelines/backend/src/apiserver/client"
	"github.com/kubeflow/pipelines/backend/src/apiserver/common"
	"github.com/kubeflow/pipelines/backend/src/common/util"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"google.golang.org/grpc/metadata"
	authorizationv1 "k8s.io/api/authorization/v1"
)

func TestAuthorizeRequest_SingleUserMode(t *testing.T) {
	clients, manager, _ := initWithExperiment(t)
	defer clients.Close()
	authServer := AuthServer{resourceManager: manager}
	clients.SubjectAccessReviewClientFake = client.NewFakeSubjectAccessReviewClientUnauthorized()

	md := metadata.New(map[string]string{})
	ctx := metadata.NewIncomingContext(context.Background(), md)

	request := &api.AuthorizeRequest{
		Namespace: "ns1",
		Resources: api.AuthorizeRequest_VIEWERS,
		Verb:      api.AuthorizeRequest_GET,
	}

	_, err := authServer.Authorize(ctx, request)
	// Authz is completely skipped without checking anything.
	assert.Nil(t, err)
}

func TestAuthorizeRequest_InvalidRequest(t *testing.T) {
	viper.Set(common.MultiUserMode, "true")
	defer viper.Set(common.MultiUserMode, "false")

	clients, manager, _ := initWithExperiment(t)
	defer clients.Close()
	authServer := AuthServer{resourceManager: manager}

	md := metadata.New(map[string]string{})
	ctx := metadata.NewIncomingContext(context.Background(), md)

	request := &api.AuthorizeRequest{
		Namespace: "",
		Resources: api.AuthorizeRequest_UNASSIGNED_RESOURCES,
		Verb:      api.AuthorizeRequest_UNASSIGNED_VERB,
	}

	_, err := authServer.Authorize(ctx, request)
	assert.Error(t, err)
	assert.EqualError(t, err, "Authorize request is not valid: Invalid input error: Namespace is empty. Please specify a valid namespace.")
}

func TestAuthorizeRequest_Authorized(t *testing.T) {
	viper.Set(common.MultiUserMode, "true")
	defer viper.Set(common.MultiUserMode, "false")

	clients, manager, _ := initWithExperiment(t)
	defer clients.Close()
	authServer := AuthServer{resourceManager: manager}

	md := metadata.New(map[string]string{common.GoogleIAPUserIdentityHeader: "accounts.google.com:user@google.com"})
	ctx := metadata.NewIncomingContext(context.Background(), md)

	request := &api.AuthorizeRequest{
		Namespace: "ns1",
		Resources: api.AuthorizeRequest_VIEWERS,
		Verb:      api.AuthorizeRequest_GET,
	}

	_, err := authServer.Authorize(ctx, request)
	assert.Nil(t, err)
}

func TestAuthorizeRequest_Unauthorized(t *testing.T) {
	viper.Set(common.MultiUserMode, "true")
	defer viper.Set(common.MultiUserMode, "false")

	clients, manager, _ := initWithExperiment_SubjectAccessReview_Unauthorized(t)
	defer clients.Close()
	authServer := AuthServer{resourceManager: manager}

	userIdentity := "user@google.com"
	md := metadata.New(map[string]string{common.GoogleIAPUserIdentityHeader: common.GoogleIAPUserIdentityPrefix + userIdentity})
	ctx := metadata.NewIncomingContext(context.Background(), md)

	request := &api.AuthorizeRequest{
		Namespace: "ns1",
		Resources: api.AuthorizeRequest_VIEWERS,
		Verb:      api.AuthorizeRequest_GET,
	}

	_, err := authServer.Authorize(ctx, request)
	assert.Error(t, err)

	resourceAttributes := &authorizationv1.ResourceAttributes{
		Namespace: "ns1",
		Verb:      common.RbacResourceVerbGet,
		Group:     common.RbacKubeflowGroup,
		Version:   common.RbacPipelinesVersion,
		Resource:  common.RbacResourceTypeViewers,
	}
	assert.EqualError(t, err, wrapFailedAuthzRequestError(getPermissionDeniedError(userIdentity, resourceAttributes)).Error())
}

func TestAuthorizeRequest_EmptyUserIdPrefix(t *testing.T) {
	viper.Set(common.MultiUserMode, "true")
	defer viper.Set(common.MultiUserMode, "false")
	viper.Set(common.KubeflowUserIDPrefix, "")
	defer viper.Set(common.KubeflowUserIDPrefix, common.GoogleIAPUserIdentityPrefix)

	clients, manager, _ := initWithExperiment(t)
	defer clients.Close()
	authServer := AuthServer{resourceManager: manager}

	md := metadata.New(map[string]string{common.GoogleIAPUserIdentityHeader: "user@google.com"})
	ctx := metadata.NewIncomingContext(context.Background(), md)

	request := &api.AuthorizeRequest{
		Namespace: "ns1",
		Resources: api.AuthorizeRequest_VIEWERS,
		Verb:      api.AuthorizeRequest_GET,
	}

	_, err := authServer.Authorize(ctx, request)
	assert.Nil(t, err)
}

func TestAuthorizeRequest_Unauthenticated(t *testing.T) {
	viper.Set(common.MultiUserMode, "true")
	defer viper.Set(common.MultiUserMode, "false")

	clients, manager, _ := initWithExperiment(t)
	defer clients.Close()
	authServer := AuthServer{resourceManager: manager}

	md := metadata.New(map[string]string{"no-identity-header": "user"})
	ctx := metadata.NewIncomingContext(context.Background(), md)

	request := &api.AuthorizeRequest{
		Namespace: "ns1",
		Resources: api.AuthorizeRequest_VIEWERS,
		Verb:      api.AuthorizeRequest_GET,
	}

	_, err := authServer.Authorize(ctx, request)
	assert.NotNil(t, err)
	assert.EqualError(
		t,
		err,
		util.Wrap(kfpauth.IdentityHeaderMissingError, "Failed to authorize the request").Error(),
	)
}
