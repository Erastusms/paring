import EventPublisher from "./events/publisher";
import EventSubscriber from "./events/subscriber";

// Interface untuk payload event (bisa di-extend sesuai kebutuhan)
interface OrderCreatedEvent {
  orderId: string;
  userId: string;
  total: number;
  createdAt: string;
}

const publisher = new EventPublisher();
const subscriber = new EventSubscriber();

// Subscribe ke event order.created
subscriber.subscribe("order.created", (data: OrderCreatedEvent) => {
  console.log("📥 Handling order.created event:", data);
});

// Simulasi publish event order.created setelah 2 detik
setTimeout(() => {
  publisher.publish("order.created", {
    orderId: "ORD-12345",
    userId: "USER-67890",
    total: 150000,
    createdAt: new Date().toISOString(),
  });
}, 2000);