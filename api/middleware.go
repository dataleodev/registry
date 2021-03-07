package api

import (
	"context"
	"github.com/dataleodev/registry"

	log "github.com/go-kit/kit/log"
)

// Middleware describes a service middleware.
type Middleware func(registry.Service) registry.Service

type loggingMiddleware struct {
	logger log.Logger
	next   registry.Service
}

func (l loggingMiddleware) AuthThing(ctx context.Context, uuid string, authToken string) (node registry.Node, err error) {
	panic("implement me")
}

// LoggingMiddleware takes a logger as a dependency
// and returns a Service Middleware.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next registry.Service) registry.Service {
		return &loggingMiddleware{logger, next}
	}

}

func (l loggingMiddleware) Register(ctx context.Context, name string, email string, password, region string) (uuid string, err error) {
	defer func() {
		l.logger.Log("method", "Register", "name", name, "email", email, "password", password, "uuid", uuid, "err", err)
	}()
	return l.next.Register(ctx, name, email, password, region)
}
func (l loggingMiddleware) Login(ctx context.Context, uuid string, password string) (token string, err error) {
	defer func() {
		l.logger.Log("method", "Login", "uuid", uuid, "password", password, "token", token, "err", err)
	}()
	return l.next.Login(ctx, uuid, password)
}
func (l loggingMiddleware) ViewUser(ctx context.Context, token string, id string) (user registry.User, err error) {
	defer func() {
		l.logger.Log("method", "ViewUser", "token", token, "id", id, "user", user, "err", err)
	}()
	return l.next.ViewUser(ctx, token, id)
}
func (l loggingMiddleware) ListUsers(ctx context.Context, token string, args map[string]string) (users []registry.User, err error) {
	defer func() {
		l.logger.Log("method", "ListUsers", "token", token, "args", args, "users", users, "err", err)
	}()
	return l.next.ListUsers(ctx, token, args)
}
func (l loggingMiddleware) UpdateUser(ctx context.Context, token string, user registry.User) (err error) {
	defer func() {
		l.logger.Log("method", "UpdateUser", "token", token, "user", user, "err", err)
	}()
	return l.next.UpdateUser(ctx, token, user)
}
func (l loggingMiddleware) ChangePassword(ctx context.Context, authToken string, password string, oldPassword string) (err error) {
	defer func() {
		l.logger.Log("method", "ChangePassword", "authToken", authToken, "password", password, "oldPassword", oldPassword, "err", err)
	}()
	return l.next.ChangePassword(ctx, authToken, password, oldPassword)
}
func (l loggingMiddleware) AddNode(ctx context.Context, token string, node registry.Node) (err error) {
	defer func() {
		l.logger.Log("method", "AddNode", "token", token, "node", node, "err", err)
	}()
	return l.next.AddNode(ctx, token, node)
}
func (l loggingMiddleware) GetNode(ctx context.Context, token string, id string) (node registry.Node, err error) {
	defer func() {
		l.logger.Log("method", "GetNode", "token", token, "id", id, "node", node, "err", err)
	}()
	return l.next.GetNode(ctx, token, id)
}
func (l loggingMiddleware) ListNodes(ctx context.Context, token string, region string) (nodes []registry.Node, err error) {
	defer func() {
		l.logger.Log("method", "ListNodes", "token", token, "region", region, "nodes", nodes, "err", err)
	}()
	return l.next.ListNodes(ctx, token, region)
}
func (l loggingMiddleware) DeleteNode(ctx context.Context, token string, id string) (err error) {
	defer func() {
		l.logger.Log("method", "DeleteNode", "token", token, "id", id, "err", err)
	}()
	return l.next.DeleteNode(ctx, token, id)
}
func (l loggingMiddleware) UpdateNode(ctx context.Context, token string, id string, node registry.Node) (err error) {
	defer func() {
		l.logger.Log("method", "UpdateNode", "token", token, "id", id, "node", node, "err", err)
	}()
	return l.next.UpdateNode(ctx, token, id, node)
}
func (l loggingMiddleware) AddRegion(ctx context.Context, token string, region registry.Region) (err error) {
	defer func() {
		l.logger.Log("method", "AddRegion", "token", token, "region", region, "err", err)
	}()
	return l.next.AddRegion(ctx, token, region)
}
func (l loggingMiddleware) ListRegions(ctx context.Context, token string) (regions []registry.Region, err error) {
	defer func() {
		l.logger.Log("method", "ListRegions", "token", token, "regions", regions, "err", err)
	}()
	return l.next.ListRegions(ctx, token)
}

type eventsMiddleware struct {
	next registry.Service
}

func (e eventsMiddleware) AuthThing(ctx context.Context, uuid string, authToken string) (node registry.Node, err error) {
	panic("implement me")
}

// EventsMiddleware returns a Service Middleware.
func EventsMiddleware() Middleware {
	return func(next registry.Service) registry.Service {
		return &eventsMiddleware{next}
	}

}
func (e eventsMiddleware) Register(ctx context.Context, name string, email string, password, region string) (uuid string, err error) {
	// Implement your middleware logic here

	return e.next.Register(ctx, name, email, password, region)
}
func (e eventsMiddleware) Login(ctx context.Context, uuid string, password string) (token string, err error) {
	// Implement your middleware logic here

	return e.next.Login(ctx, uuid, password)
}
func (e eventsMiddleware) ViewUser(ctx context.Context, token string, id string) (user registry.User, err error) {
	// Implement your middleware logic here

	return e.next.ViewUser(ctx, token, id)
}
func (e eventsMiddleware) ListUsers(ctx context.Context, token string, args map[string]string) (users []registry.User, err error) {
	// Implement your middleware logic here

	return e.next.ListUsers(ctx, token, args)
}
func (e eventsMiddleware) UpdateUser(ctx context.Context, token string, user registry.User) (err error) {
	// Implement your middleware logic here

	return e.next.UpdateUser(ctx, token, user)
}
func (e eventsMiddleware) ChangePassword(ctx context.Context, authToken string, password string, oldPassword string) (err error) {
	// Implement your middleware logic here

	return e.next.ChangePassword(ctx, authToken, password, oldPassword)
}
func (e eventsMiddleware) AddNode(ctx context.Context, token string, node registry.Node) (err error) {
	// Implement your middleware logic here

	return e.next.AddNode(ctx, token, node)
}
func (e eventsMiddleware) GetNode(ctx context.Context, token string, id string) (node registry.Node, err error) {
	// Implement your middleware logic here

	return e.next.GetNode(ctx, token, id)
}
func (e eventsMiddleware) ListNodes(ctx context.Context, token string, region string) (nodes []registry.Node, err error) {
	// Implement your middleware logic here

	return e.next.ListNodes(ctx, token, region)
}
func (e eventsMiddleware) DeleteNode(ctx context.Context, token string, id string) (err error) {
	// Implement your middleware logic here

	return e.next.DeleteNode(ctx, token, id)
}
func (e eventsMiddleware) UpdateNode(ctx context.Context, token string, id string, node registry.Node) (err error) {
	// Implement your middleware logic here

	return e.next.UpdateNode(ctx, token, id, node)
}
func (e eventsMiddleware) AddRegion(ctx context.Context, token string, region registry.Region) (err error) {
	// Implement your middleware logic here

	return e.next.AddRegion(ctx, token, region)
}
func (e eventsMiddleware) ListRegions(ctx context.Context, token string) (regions []registry.Region, err error) {
	// Implement your middleware logic here

	return e.next.ListRegions(ctx, token)
}
