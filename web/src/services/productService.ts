export type PackSize = {
  size: number;
  // Decimal values are serialized as strings by the Go API.
  cost: string;
};

export type Product = {
  id: number;
  name: string;
  packSizes: PackSize[];
};

export async function getProduct(id: number): Promise<Product> {
  const response = await fetch(`/api/products/${id}`);

  if (!response.ok) {
    throw new Error(`Unable to load product (${response.status})`);
  }

  return response.json() as Promise<Product>;
}
