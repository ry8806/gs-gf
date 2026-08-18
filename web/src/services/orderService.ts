export type CreateOrderRequest = {
  productId: number;
  quantity: number;
};

export type PackAllocation = {
  packSize: number;
  quantity: number;
  // Decimal values are serialized as strings by the Go API.
  packCost: string;
};

export type OrderSummary = {
  orderId: number;
  productId: number;
  desiredQuantity: number;
  actualQuantity: number;
  // Decimal values are serialized as strings by the Go API.
  cost: string;
  packAllocation: PackAllocation[];
};

export async function createOrder(request: CreateOrderRequest): Promise<OrderSummary> {
  const response = await fetch("/api/orders", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(request),
  });

  if (!response.ok) {
    throw new Error(`Unable to create order (${response.status})`);
  }

  return response.json() as Promise<OrderSummary>;
}
