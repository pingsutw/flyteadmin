package implementations

import (
	"context"
	"errors"
	"fmt"

	cloudevents "github.com/cloudevents/sdk-go/v2"

	"k8s.io/apimachinery/pkg/util/sets"

	"github.com/flyteorg/flyteadmin/pkg/async/notifications/interfaces"

	"github.com/flyteorg/flytestdlib/logger"
	"github.com/flyteorg/flytestdlib/promutils"
	"github.com/golang/protobuf/proto"
)

type CloudEventPublisher struct {
	pub           cloudevents.Client
	systemMetrics eventPublisherSystemMetrics
	events        sets.String
}

func (p *CloudEventPublisher) Publish(ctx context.Context, notificationType string, msg proto.Message) error {
	p.systemMetrics.PublishTotal.Inc()

	if !p.shouldPublishEvent(notificationType) {
		return nil
	}
	logger.Debugf(ctx, "Publishing the following message [%+v]", msg)

	event := cloudevents.NewEvent()
	// CloudEvent specification: https://github.com/cloudevents/spec/blob/v1.0/spec.md#required-attributes
	event.SetType("com.flyte.workflow")
	event.SetSource("https://github.com/flyteorg/flyteadmin")

	if err := event.SetData(cloudevents.ApplicationJSON, &msg); err != nil {
		p.systemMetrics.PublishError.Inc()
		return errors.New("failed to encode data")
	}

	if result := p.pub.Send(context.Background(), event); cloudevents.IsUndelivered(result) {
		p.systemMetrics.PublishError.Inc()
		return errors.New(fmt.Sprintf("failed to send: %v", result))
	} else {
		p.systemMetrics.PublishSuccess.Inc()
		logger.Infof(ctx, "sent, accepted: %t", cloudevents.IsACK(result))
	}

	return nil
}

func (p *CloudEventPublisher) shouldPublishEvent(notificationType string) bool {
	return p.events.Has(notificationType)
}

func NewCloudEventsPublisher(client cloudevents.Client, scope promutils.Scope, eventTypes []string) interfaces.Publisher {
	eventSet := sets.NewString()

	for _, event := range eventTypes {
		if event == AllTypes || event == AllTypesShort {
			for _, e := range supportedEvents {
				eventSet = eventSet.Insert(e)
			}
			break
		}
		if e, found := supportedEvents[event]; found {
			eventSet = eventSet.Insert(e)
		} else {
			logger.Errorf(context.Background(), "Unsupported event type [%s] in the config")
		}
	}

	return &CloudEventPublisher{
		pub:           client,
		systemMetrics: newEventPublisherSystemMetrics(scope.NewSubScope("cloudevents_publisher")),
		events:        eventSet,
	}
}
