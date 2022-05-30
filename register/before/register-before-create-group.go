package before

import (
	"context"
	"database/sql"

	"github.com/heroiclabs/nakama-common/api"
	"github.com/heroiclabs/nakama-common/runtime"
	u "github.com/smhmayboudi/nakama-modules-go/util"
	"go.opentelemetry.io/contrib/propagators/b3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

func RegisterBeforeCreateGroup(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, in *api.CreateGroupRequest) (*api.CreateGroupRequest, error) {
	logger.WithFields(u.Inject(ctx, b3.B3MultipleHeader)).WithFields(map[string]interface{}{"name": "RegisterBeforeCreateGroup", "in": in}).Debug("")
	ctx = u.Extract(ctx, b3.B3SingleHeader)
	_, span := otel.Tracer(u.LoadConfig(logger).InstrumentationName).Start(
		ctx,
		"RegisterBeforeCreateGroup",
		trace.WithSpanKind(trace.SpanKindInternal))
	defer span.End()

	// if err := u.Redpanda(ctx, logger, map[string]interface{}{"name": "RegisterBeforeCreateGroup", "in": in}); err != nil {
	// 	u.HandleError(ctx, logger, span, err, "Error calling redpanda")
	// 	return in, err
	// }
	return in, nil
}
