package domain

import (
	"context"

	"github.com/kyma-incubator/compass/components/director/internal/domain/api"
	"github.com/kyma-incubator/compass/components/director/internal/domain/application"
	"github.com/kyma-incubator/compass/components/director/internal/domain/document"
	"github.com/kyma-incubator/compass/components/director/internal/domain/eventapi"
	"github.com/kyma-incubator/compass/components/director/internal/domain/healthcheck"
	"github.com/kyma-incubator/compass/components/director/internal/domain/runtime"
	"github.com/kyma-incubator/compass/components/director/internal/graphql"
)

type RootResolver struct {
	app         *application.Resolver
	api         *api.Resolver
	eventAPI    *eventapi.Resolver
	runtime     *runtime.Resolver
	healthCheck *healthcheck.Resolver
}

func NewRootResolver() *RootResolver {
	appSvc := application.NewService()
	apiSvc := api.NewService()
	eventAPISvc := eventapi.NewService()
	docSvc := document.NewService()
	runtimeSvc := runtime.NewService()
	healthCheckSvc := healthcheck.NewService()

	return &RootResolver{
		app:         application.NewResolver(appSvc, apiSvc, eventAPISvc, docSvc),
		api:         api.NewResolver(apiSvc),
		eventAPI:    eventapi.NewResolver(eventAPISvc),
		runtime:     runtime.NewResolver(runtimeSvc),
		healthCheck: healthcheck.NewResolver(healthCheckSvc),
	}
}

func (r *RootResolver) Mutation() graphql.MutationResolver {
	return &mutationResolver{r}
}
func (r *RootResolver) Query() graphql.QueryResolver {
	return &queryResolver{r}
}

func (r *RootResolver) Application() graphql.ApplicationResolver {
	return &applicationResolver{r}
}

type queryResolver struct {
	*RootResolver
}

func (r *queryResolver) Applications(ctx context.Context, filter []*graphql.LabelFilter, first *int, after *graphql.PageCursor) (*graphql.ApplicationPage, error) {
	return r.app.Applications(ctx, filter, first, after)
}
func (r *queryResolver) Application(ctx context.Context, id string) (*graphql.Application, error) {
	return r.app.Application(ctx, id)
}
func (r *queryResolver) Runtimes(ctx context.Context, filter []*graphql.LabelFilter, first *int, after *graphql.PageCursor) (*graphql.RuntimePage, error) {
	return r.runtime.Runtimes(ctx, filter, first, after)
}
func (r *queryResolver) Runtime(ctx context.Context, id string) (*graphql.Runtime, error) {
	return r.runtime.Runtime(ctx, id)
}
func (r *queryResolver) HealthChecks(ctx context.Context, types []graphql.HealthCheckType, origin *string, first *int, after *graphql.PageCursor) (*graphql.HealthCheckPage, error) {
	return r.healthCheck.HealthChecks(ctx, types, origin, first, after)
}

type mutationResolver struct {
	*RootResolver
}

func (r *mutationResolver) CreateApplication(ctx context.Context, in graphql.ApplicationInput) (*graphql.Application, error) {
	return r.app.CreateApplication(ctx, in)
}
func (r *mutationResolver) UpdateApplication(ctx context.Context, id string, in graphql.ApplicationInput) (*graphql.Application, error) {
	return r.app.UpdateApplication(ctx, id, in)
}
func (r *mutationResolver) DeleteApplication(ctx context.Context, id string) (*graphql.Application, error) {
	return r.app.DeleteApplication(ctx, id)
}
func (r *mutationResolver) AddApplicationLabel(ctx context.Context, applicationID string, label string, values []string) ([]string, error) {
	return r.app.AddApplicationLabel(ctx, applicationID, label, values)
}
func (r *mutationResolver) DeleteApplicationLabel(ctx context.Context, applicationID string, label string, values []string) ([]string, error) {
	return r.app.DeleteApplicationLabel(ctx, applicationID, label, values)
}
func (r *mutationResolver) AddApplicationAnnotation(ctx context.Context, applicationID string, annotation string, value string) (string, error) {
	return r.app.AddApplicationAnnotation(ctx, applicationID, annotation, value)
}
func (r *mutationResolver) DeleteApplicationAnnotation(ctx context.Context, applicationID string, annotation string) (*string, error) {
	return r.app.DeleteApplicationAnnotation(ctx, applicationID, annotation)
}
func (r *mutationResolver) AddApplicationWebhook(ctx context.Context, applicationID string, in graphql.ApplicationWebhookInput) (*graphql.ApplicationWebhook, error) {
	return r.app.AddApplicationWebhook(ctx, applicationID, in)
}
func (r *mutationResolver) UpdateApplicationWebhook(ctx context.Context, webhookID string, in graphql.ApplicationWebhookInput) (*graphql.ApplicationWebhook, error) {
	return r.app.UpdateApplicationWebhook(ctx, webhookID, in)
}
func (r *mutationResolver) DeleteApplicationWebhook(ctx context.Context, webhookID string) (*graphql.ApplicationWebhook, error) {
	return r.app.DeleteApplicationWebhook(ctx, webhookID)
}
func (r *mutationResolver) AddAPI(ctx context.Context, applicationID string, in graphql.APIDefinitionInput) (*graphql.APIDefinition, error) {
	return r.api.AddAPI(ctx, applicationID, in)
}
func (r *mutationResolver) UpdateAPI(ctx context.Context, id string, in graphql.APIDefinitionInput) (*graphql.APIDefinition, error) {
	return r.api.UpdateAPI(ctx, id, in)
}
func (r *mutationResolver) DeleteAPI(ctx context.Context, id string) (*graphql.APIDefinition, error) {
	return r.api.DeleteAPI(ctx, id)
}
func (r *mutationResolver) RefetchAPISpec(ctx context.Context, apiID string) (*graphql.APISpec, error) {
	return r.api.RefetchAPISpec(ctx, apiID)
}
func (r *mutationResolver) SetAPIAuth(ctx context.Context, apiID string, runtimeID string, in graphql.AuthInput) (*graphql.RuntimeAuth, error) {
	return r.api.SetAPIAuth(ctx, apiID, runtimeID, in)
}
func (r *mutationResolver) DeleteAPIAuth(ctx context.Context, apiID string, runtimeID string) (*graphql.RuntimeAuth, error) {
	return r.api.DeleteAPIAuth(ctx, apiID, runtimeID)
}
func (r *mutationResolver) AddEventAPI(ctx context.Context, applicationID string, in graphql.EventAPIDefinitionInput) (*graphql.EventAPIDefinition, error) {
	return r.eventAPI.AddEventAPI(ctx, applicationID, in)
}
func (r *mutationResolver) UpdateEventAPI(ctx context.Context, id string, in graphql.EventAPIDefinitionInput) (*graphql.EventAPIDefinition, error) {
	return r.eventAPI.UpdateEventAPI(ctx, id, in)
}
func (r *mutationResolver) DeleteEventAPI(ctx context.Context, id string) (*graphql.EventAPIDefinition, error) {
	return r.eventAPI.DeleteEventAPI(ctx, id)
}
func (r *mutationResolver) RefetchEventAPISpec(ctx context.Context, eventID string) (*graphql.EventAPISpec, error) {
	return r.eventAPI.RefetchEventAPISpec(ctx, eventID)
}
func (r *mutationResolver) CreateRuntime(ctx context.Context, in graphql.RuntimeInput) (*graphql.Runtime, error) {
	return r.runtime.CreateRuntime(ctx, in)
}
func (r *mutationResolver) UpdateRuntime(ctx context.Context, id string, in graphql.RuntimeInput) (*graphql.Runtime, error) {
	return r.runtime.UpdateRuntime(ctx, id, in)
}
func (r *mutationResolver) DeleteRuntime(ctx context.Context, id string) (*graphql.Runtime, error) {
	return r.runtime.DeleteRuntime(ctx, id)
}
func (r *mutationResolver) AddRuntimeLabel(ctx context.Context, runtimeID string, key string, values []string) ([]string, error) {
	return r.runtime.AddRuntimeLabel(ctx, runtimeID, key, values)
}
func (r *mutationResolver) DeleteRuntimeLabel(ctx context.Context, id string, key string, values []string) ([]string, error) {
	return r.runtime.DeleteRuntimeLabel(ctx, id, key, values)
}
func (r *mutationResolver) AddRuntimeAnnotation(ctx context.Context, runtimeID string, key string, value string) (string, error) {
	return r.runtime.AddRuntimeAnnotation(ctx, runtimeID, key, value)
}
func (r *mutationResolver) DeleteRuntimeAnnotation(ctx context.Context, id string, key string) (*string, error) {
	return r.runtime.DeleteRuntimeAnnotation(ctx, id, key)
}

type applicationResolver struct {
	*RootResolver
}

func (r *applicationResolver) Apis(ctx context.Context, obj *graphql.Application, group *string, first *int, after *graphql.PageCursor) (*graphql.APIDefinitionPage, error) {
	return r.app.Apis(ctx, obj, group, first, after)
}
func (r *applicationResolver) EventAPIs(ctx context.Context, obj *graphql.Application, group *string, first *int, after *graphql.PageCursor) (*graphql.EventAPIDefinitionPage, error) {
	return r.app.EventAPIs(ctx, obj, group, first, after)
}
func (r *applicationResolver) Documents(ctx context.Context, obj *graphql.Application, first *int, after *graphql.PageCursor) (*graphql.DocumentPage, error) {
	return r.app.Documents(ctx, obj, first, after)
}