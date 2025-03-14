// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql"
	"realm.pub/tavern/internal/ent/beacon"
	"realm.pub/tavern/internal/ent/file"
	"realm.pub/tavern/internal/ent/host"
	"realm.pub/tavern/internal/ent/quest"
	"realm.pub/tavern/internal/ent/tag"
	"realm.pub/tavern/internal/ent/task"
	"realm.pub/tavern/internal/ent/tome"
	"realm.pub/tavern/internal/ent/user"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (b *BeaconQuery) CollectFields(ctx context.Context, satisfies ...string) (*BeaconQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return b, nil
	}
	if err := b.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return b, nil
}

func (b *BeaconQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(beacon.Columns))
		selectedFields = []string{beacon.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "host":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&HostClient{config: b.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			b.withHost = query
		case "tasks":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TaskClient{config: b.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			b.WithNamedTasks(alias, func(wq *TaskQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[beacon.FieldName]; !ok {
				selectedFields = append(selectedFields, beacon.FieldName)
				fieldSeen[beacon.FieldName] = struct{}{}
			}
		case "principal":
			if _, ok := fieldSeen[beacon.FieldPrincipal]; !ok {
				selectedFields = append(selectedFields, beacon.FieldPrincipal)
				fieldSeen[beacon.FieldPrincipal] = struct{}{}
			}
		case "identifier":
			if _, ok := fieldSeen[beacon.FieldIdentifier]; !ok {
				selectedFields = append(selectedFields, beacon.FieldIdentifier)
				fieldSeen[beacon.FieldIdentifier] = struct{}{}
			}
		case "agentIdentifier":
			if _, ok := fieldSeen[beacon.FieldAgentIdentifier]; !ok {
				selectedFields = append(selectedFields, beacon.FieldAgentIdentifier)
				fieldSeen[beacon.FieldAgentIdentifier] = struct{}{}
			}
		case "lastSeenAt":
			if _, ok := fieldSeen[beacon.FieldLastSeenAt]; !ok {
				selectedFields = append(selectedFields, beacon.FieldLastSeenAt)
				fieldSeen[beacon.FieldLastSeenAt] = struct{}{}
			}
		case "interval":
			if _, ok := fieldSeen[beacon.FieldInterval]; !ok {
				selectedFields = append(selectedFields, beacon.FieldInterval)
				fieldSeen[beacon.FieldInterval] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		b.Select(selectedFields...)
	}
	return nil
}

type beaconPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []BeaconPaginateOption
}

func newBeaconPaginateArgs(rv map[string]any) *beaconPaginateArgs {
	args := &beaconPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &BeaconOrder{Field: &BeaconOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithBeaconOrder(order))
			}
		case *BeaconOrder:
			if v != nil {
				args.opts = append(args.opts, WithBeaconOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*BeaconWhereInput); ok {
		args.opts = append(args.opts, WithBeaconFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (f *FileQuery) CollectFields(ctx context.Context, satisfies ...string) (*FileQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return f, nil
	}
	if err := f.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return f, nil
}

func (f *FileQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(file.Columns))
		selectedFields = []string{file.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "createdAt":
			if _, ok := fieldSeen[file.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, file.FieldCreatedAt)
				fieldSeen[file.FieldCreatedAt] = struct{}{}
			}
		case "lastModifiedAt":
			if _, ok := fieldSeen[file.FieldLastModifiedAt]; !ok {
				selectedFields = append(selectedFields, file.FieldLastModifiedAt)
				fieldSeen[file.FieldLastModifiedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[file.FieldName]; !ok {
				selectedFields = append(selectedFields, file.FieldName)
				fieldSeen[file.FieldName] = struct{}{}
			}
		case "size":
			if _, ok := fieldSeen[file.FieldSize]; !ok {
				selectedFields = append(selectedFields, file.FieldSize)
				fieldSeen[file.FieldSize] = struct{}{}
			}
		case "hash":
			if _, ok := fieldSeen[file.FieldHash]; !ok {
				selectedFields = append(selectedFields, file.FieldHash)
				fieldSeen[file.FieldHash] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		f.Select(selectedFields...)
	}
	return nil
}

type filePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []FilePaginateOption
}

func newFilePaginateArgs(rv map[string]any) *filePaginateArgs {
	args := &filePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &FileOrder{Field: &FileOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithFileOrder(order))
			}
		case *FileOrder:
			if v != nil {
				args.opts = append(args.opts, WithFileOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*FileWhereInput); ok {
		args.opts = append(args.opts, WithFileFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (h *HostQuery) CollectFields(ctx context.Context, satisfies ...string) (*HostQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return h, nil
	}
	if err := h.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return h, nil
}

func (h *HostQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(host.Columns))
		selectedFields = []string{host.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "tags":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TagClient{config: h.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			h.WithNamedTags(alias, func(wq *TagQuery) {
				*wq = *query
			})
		case "beacons":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&BeaconClient{config: h.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			h.WithNamedBeacons(alias, func(wq *BeaconQuery) {
				*wq = *query
			})
		case "identifier":
			if _, ok := fieldSeen[host.FieldIdentifier]; !ok {
				selectedFields = append(selectedFields, host.FieldIdentifier)
				fieldSeen[host.FieldIdentifier] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[host.FieldName]; !ok {
				selectedFields = append(selectedFields, host.FieldName)
				fieldSeen[host.FieldName] = struct{}{}
			}
		case "primaryIP":
			if _, ok := fieldSeen[host.FieldPrimaryIP]; !ok {
				selectedFields = append(selectedFields, host.FieldPrimaryIP)
				fieldSeen[host.FieldPrimaryIP] = struct{}{}
			}
		case "platform":
			if _, ok := fieldSeen[host.FieldPlatform]; !ok {
				selectedFields = append(selectedFields, host.FieldPlatform)
				fieldSeen[host.FieldPlatform] = struct{}{}
			}
		case "lastSeenAt":
			if _, ok := fieldSeen[host.FieldLastSeenAt]; !ok {
				selectedFields = append(selectedFields, host.FieldLastSeenAt)
				fieldSeen[host.FieldLastSeenAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		h.Select(selectedFields...)
	}
	return nil
}

type hostPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []HostPaginateOption
}

func newHostPaginateArgs(rv map[string]any) *hostPaginateArgs {
	args := &hostPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &HostOrder{Field: &HostOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithHostOrder(order))
			}
		case *HostOrder:
			if v != nil {
				args.opts = append(args.opts, WithHostOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*HostWhereInput); ok {
		args.opts = append(args.opts, WithHostFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (q *QuestQuery) CollectFields(ctx context.Context, satisfies ...string) (*QuestQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return q, nil
	}
	if err := q.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return q, nil
}

func (q *QuestQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(quest.Columns))
		selectedFields = []string{quest.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "tome":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TomeClient{config: q.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			q.withTome = query
		case "bundle":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&FileClient{config: q.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			q.withBundle = query
		case "tasks":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&TaskClient{config: q.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			q.WithNamedTasks(alias, func(wq *TaskQuery) {
				*wq = *query
			})
		case "creator":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: q.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			q.withCreator = query
		case "createdAt":
			if _, ok := fieldSeen[quest.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, quest.FieldCreatedAt)
				fieldSeen[quest.FieldCreatedAt] = struct{}{}
			}
		case "lastModifiedAt":
			if _, ok := fieldSeen[quest.FieldLastModifiedAt]; !ok {
				selectedFields = append(selectedFields, quest.FieldLastModifiedAt)
				fieldSeen[quest.FieldLastModifiedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[quest.FieldName]; !ok {
				selectedFields = append(selectedFields, quest.FieldName)
				fieldSeen[quest.FieldName] = struct{}{}
			}
		case "parameters":
			if _, ok := fieldSeen[quest.FieldParameters]; !ok {
				selectedFields = append(selectedFields, quest.FieldParameters)
				fieldSeen[quest.FieldParameters] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		q.Select(selectedFields...)
	}
	return nil
}

type questPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []QuestPaginateOption
}

func newQuestPaginateArgs(rv map[string]any) *questPaginateArgs {
	args := &questPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &QuestOrder{Field: &QuestOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithQuestOrder(order))
			}
		case *QuestOrder:
			if v != nil {
				args.opts = append(args.opts, WithQuestOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*QuestWhereInput); ok {
		args.opts = append(args.opts, WithQuestFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TagQuery) CollectFields(ctx context.Context, satisfies ...string) (*TagQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TagQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(tag.Columns))
		selectedFields = []string{tag.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "hosts":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&HostClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			t.WithNamedHosts(alias, func(wq *HostQuery) {
				*wq = *query
			})
		case "name":
			if _, ok := fieldSeen[tag.FieldName]; !ok {
				selectedFields = append(selectedFields, tag.FieldName)
				fieldSeen[tag.FieldName] = struct{}{}
			}
		case "kind":
			if _, ok := fieldSeen[tag.FieldKind]; !ok {
				selectedFields = append(selectedFields, tag.FieldKind)
				fieldSeen[tag.FieldKind] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		t.Select(selectedFields...)
	}
	return nil
}

type tagPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TagPaginateOption
}

func newTagPaginateArgs(rv map[string]any) *tagPaginateArgs {
	args := &tagPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &TagOrder{Field: &TagOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithTagOrder(order))
			}
		case *TagOrder:
			if v != nil {
				args.opts = append(args.opts, WithTagOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*TagWhereInput); ok {
		args.opts = append(args.opts, WithTagFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TaskQuery) CollectFields(ctx context.Context, satisfies ...string) (*TaskQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TaskQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(task.Columns))
		selectedFields = []string{task.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "quest":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&QuestClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			t.withQuest = query
		case "beacon":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&BeaconClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			t.withBeacon = query
		case "createdAt":
			if _, ok := fieldSeen[task.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, task.FieldCreatedAt)
				fieldSeen[task.FieldCreatedAt] = struct{}{}
			}
		case "lastModifiedAt":
			if _, ok := fieldSeen[task.FieldLastModifiedAt]; !ok {
				selectedFields = append(selectedFields, task.FieldLastModifiedAt)
				fieldSeen[task.FieldLastModifiedAt] = struct{}{}
			}
		case "claimedAt":
			if _, ok := fieldSeen[task.FieldClaimedAt]; !ok {
				selectedFields = append(selectedFields, task.FieldClaimedAt)
				fieldSeen[task.FieldClaimedAt] = struct{}{}
			}
		case "execStartedAt":
			if _, ok := fieldSeen[task.FieldExecStartedAt]; !ok {
				selectedFields = append(selectedFields, task.FieldExecStartedAt)
				fieldSeen[task.FieldExecStartedAt] = struct{}{}
			}
		case "execFinishedAt":
			if _, ok := fieldSeen[task.FieldExecFinishedAt]; !ok {
				selectedFields = append(selectedFields, task.FieldExecFinishedAt)
				fieldSeen[task.FieldExecFinishedAt] = struct{}{}
			}
		case "output":
			if _, ok := fieldSeen[task.FieldOutput]; !ok {
				selectedFields = append(selectedFields, task.FieldOutput)
				fieldSeen[task.FieldOutput] = struct{}{}
			}
		case "error":
			if _, ok := fieldSeen[task.FieldError]; !ok {
				selectedFields = append(selectedFields, task.FieldError)
				fieldSeen[task.FieldError] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		t.Select(selectedFields...)
	}
	return nil
}

type taskPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TaskPaginateOption
}

func newTaskPaginateArgs(rv map[string]any) *taskPaginateArgs {
	args := &taskPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case []*TaskOrder:
			args.opts = append(args.opts, WithTaskOrder(v))
		case []any:
			var orders []*TaskOrder
			for i := range v {
				mv, ok := v[i].(map[string]any)
				if !ok {
					continue
				}
				var (
					err1, err2 error
					order      = &TaskOrder{Field: &TaskOrderField{}, Direction: entgql.OrderDirectionAsc}
				)
				if d, ok := mv[directionField]; ok {
					err1 = order.Direction.UnmarshalGQL(d)
				}
				if f, ok := mv[fieldField]; ok {
					err2 = order.Field.UnmarshalGQL(f)
				}
				if err1 == nil && err2 == nil {
					orders = append(orders, order)
				}
			}
			args.opts = append(args.opts, WithTaskOrder(orders))
		}
	}
	if v, ok := rv[whereField].(*TaskWhereInput); ok {
		args.opts = append(args.opts, WithTaskFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (t *TomeQuery) CollectFields(ctx context.Context, satisfies ...string) (*TomeQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return t, nil
	}
	if err := t.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return t, nil
}

func (t *TomeQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(tome.Columns))
		selectedFields = []string{tome.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "files":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&FileClient{config: t.config}).Query()
			)
			if err := query.collectField(ctx, opCtx, field, path, satisfies...); err != nil {
				return err
			}
			t.WithNamedFiles(alias, func(wq *FileQuery) {
				*wq = *query
			})
		case "createdAt":
			if _, ok := fieldSeen[tome.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, tome.FieldCreatedAt)
				fieldSeen[tome.FieldCreatedAt] = struct{}{}
			}
		case "lastModifiedAt":
			if _, ok := fieldSeen[tome.FieldLastModifiedAt]; !ok {
				selectedFields = append(selectedFields, tome.FieldLastModifiedAt)
				fieldSeen[tome.FieldLastModifiedAt] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[tome.FieldName]; !ok {
				selectedFields = append(selectedFields, tome.FieldName)
				fieldSeen[tome.FieldName] = struct{}{}
			}
		case "description":
			if _, ok := fieldSeen[tome.FieldDescription]; !ok {
				selectedFields = append(selectedFields, tome.FieldDescription)
				fieldSeen[tome.FieldDescription] = struct{}{}
			}
		case "paramDefs":
			if _, ok := fieldSeen[tome.FieldParamDefs]; !ok {
				selectedFields = append(selectedFields, tome.FieldParamDefs)
				fieldSeen[tome.FieldParamDefs] = struct{}{}
			}
		case "eldritch":
			if _, ok := fieldSeen[tome.FieldEldritch]; !ok {
				selectedFields = append(selectedFields, tome.FieldEldritch)
				fieldSeen[tome.FieldEldritch] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		t.Select(selectedFields...)
	}
	return nil
}

type tomePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []TomePaginateOption
}

func newTomePaginateArgs(rv map[string]any) *tomePaginateArgs {
	args := &tomePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &TomeOrder{Field: &TomeOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithTomeOrder(order))
			}
		case *TomeOrder:
			if v != nil {
				args.opts = append(args.opts, WithTomeOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*TomeWhereInput); ok {
		args.opts = append(args.opts, WithTomeFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "name":
			if _, ok := fieldSeen[user.FieldName]; !ok {
				selectedFields = append(selectedFields, user.FieldName)
				fieldSeen[user.FieldName] = struct{}{}
			}
		case "photoURL":
			if _, ok := fieldSeen[user.FieldPhotoURL]; !ok {
				selectedFields = append(selectedFields, user.FieldPhotoURL)
				fieldSeen[user.FieldPhotoURL] = struct{}{}
			}
		case "isActivated":
			if _, ok := fieldSeen[user.FieldIsActivated]; !ok {
				selectedFields = append(selectedFields, user.FieldIsActivated)
				fieldSeen[user.FieldIsActivated] = struct{}{}
			}
		case "isAdmin":
			if _, ok := fieldSeen[user.FieldIsAdmin]; !ok {
				selectedFields = append(selectedFields, user.FieldIsAdmin)
				fieldSeen[user.FieldIsAdmin] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

func limitRows(partitionBy string, limit int, orderBy ...sql.Querier) func(s *sql.Selector) {
	return func(s *sql.Selector) {
		d := sql.Dialect(s.Dialect())
		s.SetDistinct(false)
		with := d.With("src_query").
			As(s.Clone()).
			With("limited_query").
			As(
				d.Select("*").
					AppendSelectExprAs(
						sql.RowNumber().PartitionBy(partitionBy).OrderExpr(orderBy...),
						"row_number",
					).
					From(d.Table("src_query")),
			)
		t := d.Table("limited_query").As(s.TableName())
		*s = *d.Select(s.UnqualifiedColumns()...).
			From(t).
			Where(sql.LTE(t.C("row_number"), limit)).
			Prefix(with)
	}
}

// mayAddCondition appends another type condition to the satisfies list
// if condition is enabled (Node/Nodes) and it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond string) []string {
	if len(satisfies) == 0 {
		return satisfies
	}
	for _, s := range satisfies {
		if typeCond == s {
			return satisfies
		}
	}
	return append(satisfies, typeCond)
}
